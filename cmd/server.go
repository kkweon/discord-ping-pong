package cmd

import (
	"errors"
	"github.com/kkweon/discord-ping-pong/internal/searcher"
	"os"

	"github.com/kkweon/discord-ping-pong/internal/application"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the server",
	RunE: func(cmd *cobra.Command, args []string) error {
		pKey := os.Getenv("APPLICATION_PUBLIC_KEY")

		if pKey == "" {
			return errors.New("env APPLICATION_PUBLIC_KEY is required. Get your application key from https://discord.com/developers/applications")
		}

		applicationID = os.Getenv("APPLICATION_ID")
		if applicationID == "" {
			return errors.New("env APPLICATION_ID is required. Get your application ID from https://discord.com/developers/applications")
		}

		googleAPIKey := os.Getenv("GOOGLE_API_KEY")
		if googleAPIKey == "" {
			return errors.New("env GOOGLE_API_KEY is required. Get your API KEY from https://console.cloud.google.com/apis/credentials")
		}

		app := application.New(pKey, applicationID, searcher.New(googleAPIKey))
		return app.Run()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
