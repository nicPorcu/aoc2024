package util

import (
	"log"
	"os"
	"strings"
)

func CheckWithFatalError(e error, msg string) {
	if e != nil {
		log.Fatalf(msg, e)
	}
}

func ReadFile(path string) []string {
	dat, err := os.ReadFile(path)
	CheckWithFatalError(err, "Failed to read file %v")
	return strings.Split(string(dat), "\n")
}

func Abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}
