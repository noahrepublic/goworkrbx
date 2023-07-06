/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os/exec"

	settings "github.com/noahrepublic/goworkrbx/config"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Automatically sets up a new project using the set config",

	Run: func(cmd *cobra.Command, args []string) {
		config := settings.GetConfig()

		exec.Command("ls")

		if config.Aftman {
			fmt.Println("Setting up Aftman...")

			_, err := exec.Command("aftman.exe", "init").Output()

			if err != nil {
				fmt.Println("could not run command: ", err)
			}

			for _, tool := range config.AftmanTools {
				fmt.Printf("Adding %v to Aftman...\n", tool)
				fmt.Printf("aftman.exe add %v\n", tool)
				_, err := exec.Command("aftman.exe", "add", tool).Output()

				if err != nil {
					fmt.Println(err)
				}
			}

			_, err = exec.Command("aftman.exe", "install").Output()

			if err != nil {
				panic(err)
			}
		}

		if config.Wally {
			fmt.Println("Setting up Wally...")
			_, err := exec.Command("wally.exe", "init").Output()

			if err != nil {
				panic(err)
			}
		}

		if config.Selene {
			fmt.Println("Setting up Selene...")
			_, err := exec.Command("selene.exe", "generate-roblox-std").Output()

			if err != nil {
				panic(err)
			}
		}

		if config.Rojo {
			fmt.Println("Setting up Rojo...")

			_, err := exec.Command("rojo.exe", "init").Output()

			if err != nil {
				panic(err)
			}

			_, err = exec.Command("rojo.exe", "build", "-o", "game.rbxlx").Output()

			if err != nil {
				panic(err)
			}

		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
