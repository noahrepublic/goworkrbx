/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	settings "github.com/noahrepublic/goworkrbx/config"
	toolname "github.com/noahrepublic/goworkrbx/tools"
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

			aftmanTool := toolname.New(strings.Join(args, " "))

			if !aftmanTool.IsValid {
				panic(aftmanTool.Err)
			}

			config.AftmanTools = append(config.AftmanTools, aftmanTool.Inner)

			settings.SaveConfig()

			fmt.Printf("Added %s\n", aftmanTool.Inner)
		},
	}

	aftmanRemove := &cobra.Command{
		Use:   "remove",
		Short: "Remove an aftman tool",

		Run: func(cmd *cobra.Command, args []string) {
			search := strings.Join(args, " ")

			aftmanTool := toolname.New(search)

			if aftmanTool.IsValid {
				search = aftmanTool.Inner
			}

			config := settings.GetConfig()

			hasFound := false

			for i, tool := range config.AftmanTools {
				aftmanTool = toolname.New(tool)

				if strings.Contains(aftmanTool.Name, search) { // could have multiple by same dev
					config.AftmanTools = append(config.AftmanTools[:i], config.AftmanTools[i+1:]...)

					settings.SaveConfig()

					fmt.Printf("Removed %s\n", tool)

					hasFound = true

					break
				}
			}

			if !hasFound {
				fmt.Printf("Could not find %s\n", search)
				return
			}

		},
	}

	aftmanCmd.AddCommand(aftmanAdd)
	aftmanCmd.AddCommand(aftmanRemove)

	aftmanEnabled = aftmanCmd.Flags().BoolP("enable", "e", true, "Enable aftman init")
}
