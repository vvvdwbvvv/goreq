package main

import (
	"github.com/fatih/color"
	"github.com/spf13/viper"
)

// Config holds the application configuration loaded from .greqrc
type Config struct {
	DefaultHeaders map[string]string `mapstructure:"default_headers"`
	BaseURL        string            `mapstructure:"base_url"`
	Timeout        int               `mapstructure:"timeout"`
}

// LoadConfig reads the configuration from .greqrc file
// If the file is not found, it returns default configuration
func LoadConfig() Config {
	viper.SetConfigName(".greqrc")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME")
	viper.AddConfigPath(".")

	// Set defaults
	viper.SetDefault("timeout", 30)
	viper.SetDefault("default_headers", map[string]string{
		"Accept":     "application/json",
		"User-Agent": "greq/1.0",
	})

	var config Config
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			color.New(color.FgYellow).Printf("Warning: %v\n", err)
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		color.New(color.FgYellow).Printf("Warning: Failed to parse config: %v\n", err)
		// Return default config if unmarshal fails
		return Config{
			DefaultHeaders: viper.GetStringMapString("default_headers"),
			Timeout:        viper.GetInt("timeout"),
		}
	}

	return config
}
