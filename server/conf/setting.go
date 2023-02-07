package conf

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type SystemSetting struct {
	Get struct {
		On  bool   `yaml:"on"`
		Key string `yaml:"key"`
	} `yaml:"get"`
	Save struct {
		On  bool   `yaml:"on"`
		Key string `yaml:"key"`
	} `yaml:"save"`
	Edit struct {
		On  bool   `yaml:"on"`
		Key string `yaml:"key"`
	} `yaml:"edit"`
	Message struct {
		MaxLength int `yaml:"maxLength"`
	} `yaml:"message"`
}

func GetSystemSetting() SystemSetting {
	log.SetPrefix("[conf]")
	data, err := os.ReadFile("./setting.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	var t SystemSetting

	err = yaml.Unmarshal(data, &t)

	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return t
}
