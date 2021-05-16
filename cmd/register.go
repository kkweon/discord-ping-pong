package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/kkweon/discord-ping-pong/internal/common"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
)

var (
	applicationID string
	guildID       string
	dryRun        bool
	botToken      string
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register slash commands",
	RunE: func(cmd *cobra.Command, args []string) error {
		commandList := []*common.DiscordApplicationCommand{
			{
				Name:        "ping",
				Description: "Healthcheck the server status",
			},
			{
				Name:        "define",
				Description: "Google Search this term",
				Options: []common.DiscordApplicationCommandOption{
					{
						Type:        common.DiscordApplicationCommandOptionTypeString,
						Name:        "term",
						Description: "Term to search",
						Required:    true,
					},
				},
			},
		}
		commandListSerialized, err := json.Marshal(commandList)
		if err != nil {
			return err
		}

		if dryRun {

			commandListSerializedIndented, _ := json.MarshalIndent(commandList, "", "  ")

			fmt.Printf(`Sending a request to %s with the following body:
%s`, getURL(), string(commandListSerializedIndented))

			return nil
		}
		url := getURL()
		logrus.WithField("URL", url).Info("Sending a request")
		request, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(commandListSerialized))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Authorization", fmt.Sprintf("Bot %s", botToken))
		if err != nil {
			return err
		}

		client := http.Client{}
		response, err := client.Do(request)
		if err != nil {
			return err
		}

		if response.StatusCode != http.StatusOK {
			body, _ := ioutil.ReadAll(response.Body)
			return fmt.Errorf("expected HTTP Status but got = %v with the response = %s", response.StatusCode, body)
		}

		return nil
	},
}

func getURL() string {
	if guildID != "" {
		return fmt.Sprintf("https://%s/applications/%s/guilds/%s/commands", common.DiscordAPIv8URL, applicationID, guildID)
	}
	return fmt.Sprintf("https://%s/applications/%s/commands", common.DiscordAPIv8URL, applicationID)
}

func init() {
	rootCmd.AddCommand(registerCmd)
	registerCmd.Flags().StringVar(&botToken, "bot-token", "", "Bot Token (required)")
	must(registerCmd.MarkFlagRequired("bot-token"))
	registerCmd.Flags().StringVar(&applicationID, "application-id", "", "Application ID (required)")
	must(registerCmd.MarkFlagRequired("application-id"))

	registerCmd.Flags().StringVar(&guildID, "guild-id", "", "Guild ID is given then it will create a Guild Application Command.")
	registerCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Do not send a request")
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
