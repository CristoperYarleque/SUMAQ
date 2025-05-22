package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

type Config struct {
	GeneralDB      DatabaseConfig
	MetaDatabaseDB DatabaseConfig
}

type DatabaseConfig struct {
	Host      string
	Username  string
	Password  string
	Port      string
	PageLimit string
	Adapter   string
	Database  string
	URI       string
}

func GetConfigPath(env string) string {

	configPaths := map[string]string{
		"development": "./config/config-development.yml",
		"production":  "./config/config-production.yml",
	}

	path, exists := configPaths[env]
	if !exists {
		path = configPaths["development"]
	}

	return path
}

func GetConfig(configPath string) (*Config, error) {
	cfgFile, err := LoadConfig(configPath)
	if err != nil {
		return nil, fmt.Errorf("error loading config: %w", err)
	}
	cfg, err := ParseConfig(cfgFile)
	if err != nil {
		return nil, fmt.Errorf("error parsing config: %w", err)
	}
	return cfg, nil
}

func LoadConfig(filename string) (*viper.Viper, error) {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		return nil, fmt.Errorf("could not get current file path")
	}

	rootDir := filepath.Join(file, "../..")
	configCleanPath := filepath.Clean(filename)
	configPath := filepath.Join(rootDir, configCleanPath)

	v := viper.New()
	v.SetConfigFile(configPath)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	err := v.ReadInConfig()

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("config file not found: %w", err)
	}

	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	return v, nil
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config
	err := v.Unmarshal(&c)
	if err != nil {
		log.Printf("Unable to decode into struct, %v", err.Error())
		return nil, err
	}
	return &c, nil
}
