package main

import (
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/api/interactions", func(c *gin.Context) {
		var rootMessage discordgo.Message
		err := c.BindJSON(&rootMessage)
		logrus.WithError(err).Infof("%+v", rootMessage)
		if err == nil {
			if rootMessage.Type == 1 {
				c.JSON(http.StatusOK, rootMessage)
				return
			}
		}
		c.Abort()
	})
	r.Run()
}
