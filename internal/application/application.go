package application

import (
	"crypto/ed25519"
	"encoding/hex"
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func decodeToPublicKey(applicationPublicKey string) ed25519.PublicKey {
	rawKey := []byte(applicationPublicKey)
	applicaitonPublicKey := make([]byte, hex.DecodedLen(len(rawKey)))
	hex.Decode(applicaitonPublicKey, rawKey)
	return ed25519.PublicKey(applicaitonPublicKey)
}

func Run(publicKeyFromDiscord string) error {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/api/interactions", func(c *gin.Context) {
		if !discordgo.VerifyInteraction(c.Request, decodeToPublicKey(publicKeyFromDiscord)) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		var rootMessage discordgo.Message
		err := c.BindJSON(&rootMessage)
		logrus.WithError(err).Infof("%+v", rootMessage)
		if err == nil {
			if rootMessage.Type == 1 {
				c.JSON(http.StatusOK, gin.H{
					"type": 1,
				})
				return
			}
		}
		c.Abort()
	})
	return r.Run()
}
