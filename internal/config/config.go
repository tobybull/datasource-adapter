package config

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"strings"

	"github.com/spf13/viper"
)

// Config holds the application's configuration.
type Config struct {
	LogLevel       string `mapstructure:"log_level"`
	IsProduction   bool   `mapstructure:"is_production"`
	ServicePort    int    `mapstructure:"service_port"`
	ExternalAPIURL string `mapstructure:"external_api_url"`
	DatabaseURL    string `mapstructure:"database_url"`
}

// LoadConfig loads the application's configuration using the Viper library.
// It supports loading from a file and/or environment variables.
func LoadConfig() (*Config, error) {
	viper.SetConfigName("application") // Name of the configuration file
	viper.SetConfigType("yaml")        // Configuration file type (e.g., yaml, json)
	viper.AddConfigPath(".")           // Path where to search for the configuration file

	// Attempt to read the configuration file.
	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if !errors.As(err, &configFileNotFoundError) {
			// Config file was found but another error occurred.
			return nil, fmt.Errorf("error reading config file: %w", err)
		}
		// Config file not found; carry on.  We will try to load from ENV
		fmt.Println("No application.yaml file found, using environment variables")
	}

	// Load .env file if it exists.  This will not override existing environment variables.
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file; not a problem if not present")
	}

	// Set environment variable prefixes and bind them to Viper.
	viper.SetEnvPrefix("simple")                           // e.g., SIMPLE_LOG_LEVEL
	viper.AutomaticEnv()                                   // Read in environment variables that match.
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // Replace . with _ in env vars

	//check if all required ENVs are set
	requiredVars := []string{"log_level", "service_port", "external_api_url", "database_url"} // in viper
	var missingVars []string
	for _, key := range requiredVars {
		if !viper.IsSet(key) {
			missingVars = append(missingVars, key)
		}
	}
	if len(missingVars) > 0 {
		return nil, fmt.Errorf("missing required configuration variables: %s", strings.Join(missingVars, ", "))
	}

	var config Config
	// Use viper to get each value.  Viper will check env vars first, then the config file.
	config.LogLevel = viper.GetString("log_level")
	config.IsProduction = viper.GetBool("is_production")
	config.ServicePort = viper.GetInt("service_port")
	config.ExternalAPIURL = viper.GetString("external_api_url")
	config.DatabaseURL = viper.GetString("database_url")

	return &config, nil
}
