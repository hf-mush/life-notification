package common

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const varDotEnv = "DOT_ENV_FILE"
const defaultDotEnv = ".env"

// LoadDotEnv apply .env to VARIABLE.
func LoadDotEnv() {
	if _, found := os.LookupEnv(varDotEnv); !found {
		os.Setenv(varDotEnv, defaultDotEnv)
		log.Println(fmt.Sprintf("no env file specified. try to load default %s.", defaultDotEnv))
	}

	dotEnvFile := os.Getenv(varDotEnv)
	if err := godotenv.Load(dotEnvFile); err != nil {
		log.Println(fmt.Sprintf("no env file loaded %#v", err))
	} else {
		log.Println(fmt.Sprintf("env file loaded: %s", dotEnvFile))
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
