/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"strings"

	settings "github.com/noahrepublic/goworkrbx/config"
	"github.com/spf13/cobra"
)

var aftmanEnabled *bool

// aftmanCmd represents the aftman command
var aftmanCmd = &cobra.Command{
	Use:   "aftman",
	Short: "Enable or disable aftman init",

	Run: func(cmd *cobra.Command, args []string) {
		config := settings.GetConfig()

		config.Aftman = *aftmanEnabled

		settings.SaveConfig()
	},
}

func init() {
	rootCmd.AddCommand(aftmanCmd)

	aftmanAdd := &cobra.Command{
		Use:   "add",
		Short: "Add a new aftman tool",

		Run: func(cmd *cobra.Command, args []string) {
			config := settings.GetConfig()

			aftmanTool := strings.Join(args, " ")

			config.AftmanTools = append(config.AftmanTools, aftmanTool)

			settings.SaveConfig()
		},
	}

	aftmanCmd.AddCommand(aftmanAdd)

	aftmanEnabled = aftmanCmd.Flags().BoolP("enable", "e", true, "Enable aftman init")
}
