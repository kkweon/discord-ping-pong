package cmd

import (
	"github.com/spf13/cobra"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "discord-ping-pong",
	Short: "Discord bot to be used as a slash command",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
