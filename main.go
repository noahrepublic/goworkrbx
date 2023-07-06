/*
Copyright © 2023 Noah Williams noahdevelopsthings@gmail.com
*/
package main

import (
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/noahrepublic/goworkrbx/cmd"
	settings "github.com/noahrepublic/goworkrbx/config"
)

func main() {
	home, err := homedir.Dir()

	if err != nil {
		panic(err)
	}

	yamlPath := filepath.Join(home, ".goworkrbx.yaml")
	settings.Init(yamlPath)

	cmd.Execute()
}
