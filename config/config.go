package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var StorageDir string
var Editor string

func InitConfig() {

	viper.SetDefault("storageDir", "")
	viper.SetDefault("editor", "")

	viper.BindEnv("storageDir", "AZRB_STORAGE_DIR")
	viper.BindEnv("editor", "AZRB_EDITOR")

	home, err := os.UserHomeDir()
	if err != nil {
		os.Exit(1)
	}

	if viper.GetString("storageDir") == "" {
		viper.Set("storageDir", filepath.Join(home, ".azrb"))
	}
	StorageDir = viper.GetString("storageDir")

	if viper.GetString("editor") == "" {
		if os.Getenv("OS") == "Windows_NT" {
			viper.Set("editor", "notepad")
		} else {
			viper.Set("editor", "nvim")
		}
	}
	Editor = viper.GetString("editor")

	viper.SetConfigName(".azrb")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(home)
	_ = viper.ReadInConfig()
}
	