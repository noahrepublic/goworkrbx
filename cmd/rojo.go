/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	settings "github.com/noahrepublic/goworkrbx/config"
	"github.com/spf13/cobra"
)

var rojoEnabled *bool

// rojoCmd represents the rojo command
var rojoCmd = &cobra.Command{
	Use:   "rojo",
	Short: "Enable or disable automatically creating a rojo project",

	Run: func(cmd *cobra.Command, args []string) {
		config := settings.GetConfig()

		config.Rojo = *rojoEnabled

		settings.SaveConfig()
	},
}

func init() {
	rootCmd.AddCommand(rojoCmd)

	rojoEnabled = rojoCmd.Flags().BoolP("enable", "e", true, "Enable rojo")

}
