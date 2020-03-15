package conf

import (
	"log"
	"os"
)

func GetLogLevel() string {
	level, ok := os.LookupEnv("Log_Level")
	if !ok {
		log.Println("log level not set, use debug")
		return "debug"
	}

	return level
}
