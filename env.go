package main

import (
	"os"
	"strconv"
)

func GreetingEnabled() bool { return parseBool("DISPLAY_GREETING", false) }
func MaxItems() int         { return parseInt("MAX_ITEMS", 100) }
func RetryCount() int       { return parseInt("RETRY_COUNT", 3) }

func parseBool(k string, d bool) bool {
	if v, ok := os.LookupEnv(k); ok {
		if b, err := strconv.ParseBool(v); err == nil {
			return b
		}
	}
	return d
}

func parseInt(k string, d int) int {
	if v, ok := os.LookupEnv(k); ok {
		if i, err := strconv.Atoi(v); err == nil {
			return i
		}
	}
	return d
}
