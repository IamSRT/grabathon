package config

import (
	"os"
)

/**
 * Created by Sai Ravi Teja K on 22, Jul 2019
 */

var (
	DBConnectionString = "root:@tcp(127.0.0.1:3306)/grabathon?charset=utf8mb4&parseTime=True"
	ServerPort         = "4000"
	LogLevel           = GetEnv("LOG_LEVEL", "info")
)

// get environment variable
func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}