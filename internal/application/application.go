package application

import (
	"bytes"
	"crypto/ed25519"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/kkweon/discord-ping-pong/internal/common"
	"github.com/kkweon/discord-ping-pong/internal/searcher"
	"google.golang.org/api/customsearch/v1"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Application struct {
	ApplicationPublicKey string
	ApplicationID        string
	HttpClient           httpClient
	Searcher             searcher.Searcher
}

func New(applicationPublicKey, applicationID string, svc searcher.Searcher) Application {
	return Application{
		ApplicationPublicKey: applicationPublicKey,
		ApplicationID:        applicationID,
		HttpClient: &http.Client{
			Timeout: time.Minute,
		},
		Searcher: svc,
	}
}

func decodeToPublicKey(applicationPublicKey string) ed25519.PublicKey {
	rawKey := []byte(applicationPublicKey)
	byteKey := make([]byte, hex.DecodedLen(len(rawKey)))
	_, _ = hex.Decode(byteKey, rawKey)
	return byteKey
}

func GetRouter(pubKey ed25519.PublicKey, handleTermSearch func(term string, token string, useEmbeds bool)) *gin.Engine {
	r := gin.Default()

	r.Use(requestLogger())

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
		if err == nil {
			if rootMessage.Type == common.DiscordInteractionTypePing {
				c.JSON(http.StatusOK, gin.H{
					"type": common.DiscordInteractionCallbackTypePong,
				})
				return
			} else if rootMessage.Type == common.DiscordInteractionTypeApplicationCommand {
				switch rootMessage.Data.Name {
				case "ping":
					response := common.DiscordInteractionResponse{
						Type: common.DiscordInteractionCallbackTypeChannelMessageWithSource,
						Data: common.DiscordInteractionApplicationCommandCallbackData{
							Content: "Pong!",
						},
					}
					c.JSON(http.StatusOK, response)
					return
				case "define":
					response := common.DiscordInteractionResponse{
						Type: common.DiscordInteractionCallbackTypeDeferredChannelMessageWithSource,
					}
					c.JSON(http.StatusOK, response)

					var term string
					var useEmbeds bool

					for _, option := range rootMessage.Data.Options {
						if option.Name == "term" && option.Value.StringValue != nil {
							// option.Value contains the search term
							// PATCH /webhooks/{application.id}/{interaction.token}/messages/@original
							term = *option.Value.StringValue

						} else if option.Name == "use_embeds" && option.Value.BoolValue != nil {
							useEmbeds = *option.Value.BoolValue
						}
					}

					if term != "" {
						go handleTermSearch(term, rootMessage.Token, useEmbeds)
					}

					return
				}
			}
		}
		log.WithError(err).Warn("did not return any value")
		c.Abort()
	})

	r.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})

	return r
}

func requestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		buf, _ := ioutil.ReadAll(c.Request.Body)
		log.WithFields(log.Fields{
			"method": c.Request.Method,
			"URL":    c.Request.URL,
		}).Infof("%s", string(buf))
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(buf))
		c.Next()
	}
}

func (a *Application) Run() error {
	r := GetRouter(decodeToPublicKey(a.ApplicationPublicKey), a.searchTermAndEditDiscordMessage)
	return r.Run()
}

func (a *Application) searchTermAndEditDiscordMessage(term string, token string, useEmbeds bool) {
	searchResult, err := a.Searcher.Search(term)
	if err != nil {
		log.WithError(err).WithField("term", term).Warn("Search failed")
		return
	}

	var body *common.DiscordEditWebhookMessage

	if useEmbeds {
		body = convertToEmbeds(searchResult)
	} else {
		body = &common.DiscordEditWebhookMessage{
			Content: searcher.SearchResultToString(searchResult),
		}
	}

	bodyBS, err := json.Marshal(body)
	req, err := http.NewRequest(http.MethodPatch, fmt.Sprintf("https://%s/webhooks/%s/%s/messages/@original", common.DiscordAPIv8URL, a.ApplicationID, token), bytes.NewReader(bodyBS))
	if err != nil {
		log.WithError(err).WithFields(log.Fields{
			"token":         token,
			"applicationID": a.ApplicationID,
			"body":          body,
		}).Warnf("failed to build a new request")
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := a.HttpClient.Do(req)
	respBs, _ := ioutil.ReadAll(resp.Body)

	log.WithError(err).WithFields(log.Fields{
		"response":    string(respBs),
		"status code": resp.Status,
		"URL":         req.URL,
		"body":        string(bodyBS),
	}).Info("end of term handler")
}

func convertToEmbeds(searchResult *customsearch.Search) *common.DiscordEditWebhookMessage {
	var embeds []*common.DiscordEmbed

	for _, item := range searchResult.Items {
		embeds = append(embeds, &common.DiscordEmbed{
			Title:       item.Title,
			Type:        "rich",
			Description: item.Snippet,
			URL:         item.Link,
		})
	}

	return &common.DiscordEditWebhookMessage{
		Embeds: embeds,
	}
}
