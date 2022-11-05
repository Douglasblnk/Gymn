package utils

import "os"

func IsProductionEnv() bool {
	return os.Getenv("APP_ENV") == "prod"
}
