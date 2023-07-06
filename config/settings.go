package settings

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Rojo   bool `yaml:"setup-rojo"`
	Aftman bool `yaml:"setup-aftman"`
	Wally  bool `yaml:"setup-wally"`
	Selene bool `yaml:"setup-selene"`

	AftmanTools []string `yaml:"aftman-tools"`
}

var config Config
var yamlPath string

func Init(path string) {

	yamlPath = path

	data, err := os.ReadFile(yamlPath)

	if err != nil {
		doesNotExist := os.IsNotExist(err)

		if doesNotExist {
			config = Config{}

			SaveConfig()
		} else {
			panic(err)
		}
	}

	if err := yaml.Unmarshal(data, &config); err != nil {
		panic(err)
	}
}

func SaveConfig() {
	data, err := yaml.Marshal(&config)

	if err != nil {
		panic(err)
	}

	err = os.WriteFile(yamlPath, data, 0644)

	if err != nil {
		panic(err)
	}

}

func GetConfig() *Config {
	return &config
}

func SetConfig(newConfig Config) {
	config = newConfig
	SaveConfig()
}
