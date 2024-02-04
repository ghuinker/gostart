package core

import (
	"os"
)

func IsDebug() bool {
	value := os.Getenv("DEBUG")
	return value == "true"
}
