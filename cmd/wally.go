/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	settings "github.com/noahrepublic/goworkrbx/config"
	"github.com/spf13/cobra"
)

var wallyEnabled *bool

// wallyCmd represents the wally command
var wallyCmd = &cobra.Command{
	Use:   "wally",
	Short: "Enable or disable default wally init setup",

	Run: func(cmd *cobra.Command, args []string) {
		config := settings.GetConfig()

		config.Wally = *wallyEnabled

		settings.SaveConfig()
	},
}

func init() {
	rootCmd.AddCommand(wallyCmd)

	wallyEnabled = wallyCmd.Flags().BoolP("enable", "e", false, "Enable wally init")
}
