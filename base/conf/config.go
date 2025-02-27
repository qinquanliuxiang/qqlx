package conf

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func LoadConfig(configPath string) error {
	_, err := os.Stat(configPath)
	if os.IsNotExist(err) {
		return fmt.Errorf("configuration file %s does not exist", configPath)
	}
	if err != nil {
		return fmt.Errorf("stat configuration file %s faild. err: %w", configPath, err)
	}

	log.Printf("loading configuration file: %s", configPath)
	configDir := filepath.Dir(configPath)
	configBase := filepath.Base(configPath)
	viper.SetConfigName(configBase)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configDir)
	err = viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("reading configuration files %s faild. err: %w", configPath, err)
	}
	log.Println("configuration file read successfully")
	return nil
}
