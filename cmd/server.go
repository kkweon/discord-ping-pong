package cmd

import (
	"errors"
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
			return errors.New("env APPLICATION_PUBLIC_KEY is required. Get your application key from https://discord.com/developers/applications.")
		}

		return application.Run(pKey)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
