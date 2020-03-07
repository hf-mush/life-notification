package common

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const envVarEnvFile = "ENV_FILE"
const defaultEnvFile = ".env"

// LoadEnv apply .env to ENVIRONMENT VARIABLE.
func LoadEnv() {
	if _, found := os.LookupEnv(envVarEnvFile); !found {
		os.Setenv(envVarEnvFile, defaultEnvFile)
		log.Println(fmt.Sprintf("no env file specified. try to load default %s.", defaultEnvFile))
	}

	envFile := os.Getenv(envVarEnvFile)
	if err := godotenv.Load(envFile); err != nil {
		log.Println(fmt.Sprintf("no env file loaded %#v", err))
	} else {
		log.Println(fmt.Sprintf("env file loaded: %s", envFile))
	}
}

// GetSlackAPIURL return slack api url.
func GetSlackAPIURL() string {
	return os.Getenv("SLACK_API_URL")
}

// GetSlackAPIToken return slack api token.
func GetSlackAPIToken() string {
	return os.Getenv("SLACK_API_TOKEN")
}

// GetSlackChannel return slack channel.
func GetSlackChannel() string {
	return os.Getenv("SLACK_CHANNEL")
}
