package application

import (
	"crypto/ed25519"
	"encoding/hex"
	"github.com/kkweon/discord-ping-pong/internal/common"
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func decodeToPublicKey(applicationPublicKey string) ed25519.PublicKey {
	rawKey := []byte(applicationPublicKey)
	byteKey := make([]byte, hex.DecodedLen(len(rawKey)))
	_, _ = hex.Decode(byteKey, rawKey)
	return byteKey
}

func GetRouter(pubKey ed25519.PublicKey) *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/api/interactions", func(c *gin.Context) {
		if !discordgo.VerifyInteraction(c.Request, pubKey) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		var rootMessage common.DiscordInteraction
		err := c.BindJSON(&rootMessage)
		logrus.WithError(err).Infof("%+v", rootMessage)
		if err == nil {
			if rootMessage.Type == common.DiscordInteractionTypePing {
				c.JSON(http.StatusOK, gin.H{
					"type": common.DiscordInteractionCallbackTypePong,
				})
				return
			} else if rootMessage.Type == common.DiscordInteractionTypeApplicationCommand {
				response := common.DiscordInteractionResponse{
					Type: common.DiscordInteractionCallbackTypeChannelMessageWithSource,
					Data: common.DiscordInteractionApplicationCommandCallbackData{
						Content: "Pong!",
					},
				}
				c.JSON(http.StatusOK, response)
				return
			}
		}
		c.Abort()
	})

	r.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})

	return r
}

func Run(publicKeyFromDiscord string) error {
	r := GetRouter(decodeToPublicKey(publicKeyFromDiscord))
	return r.Run()
}
