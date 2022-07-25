package utils

import (
	"os"
	"strconv"
)

func GetIntEnv(key string) int64 {
	env, _ := strconv.ParseInt(os.Getenv(key), 10, 64)

	return env
}
