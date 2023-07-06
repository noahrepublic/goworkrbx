/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	settings "github.com/noahrepublic/goworkrbx/config"
	"github.com/spf13/cobra"
)

var seleneEnabled *bool

// seleneCmd represents the selene command
var seleneCmd = &cobra.Command{
	Use:   "selene",
	Short: "Automatically creates a selene.toml file",

	Run: func(cmd *cobra.Command, args []string) {
		config := settings.GetConfig()

		config.Selene = *seleneEnabled

		settings.SaveConfig()
	},
}

func init() {
	rootCmd.AddCommand(seleneCmd)

	seleneEnabled = seleneCmd.Flags().BoolP("enable", "e", true, "Enable selene.toml")
}
