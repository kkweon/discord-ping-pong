package main

import (
	"crypto/ed25519"
	"encoding/hex"
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var rawKey = []byte("APPLICATION_PUBLIC_KEY")

var applicaitonPublicKey []byte

func init() {
	applicaitonPublicKey = make([]byte, hex.EncodedLen(len(rawKey)))
	hex.Encode(applicaitonPublicKey, rawKey)
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/api/interactions", func(c *gin.Context) {
		if !discordgo.VerifyInteraction(c.Request, ed25519.PublicKey(applicaitonPublicKey)) {
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
	r.Run()
}
