package env

import (
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func WithDefaultString(name, defaultValue string) string {
	val := os.Getenv(name)
	if val == "" {
		val = defaultValue
	}
	return val
}

func WithDefaultInt(name string, defaultValue int) int {
	result := defaultValue
	val := os.Getenv(name)
	if val != "" {
		var err error
		if result, err = strconv.Atoi(val); err != nil {
			log.Printf("could not load environment variable, expected integer, got type: %T\n", val)
			result = defaultValue
		}
	}
	return result
}
