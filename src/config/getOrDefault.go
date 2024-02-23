package config

import (
	"os"
	"strconv"
)

func getEnvOrDefault(key string, defaultValue string) string {
	env := os.Getenv(key)

	if env == "" {
		return defaultValue
	}

	return env
}

func getEnvOrDefaultUint(key string, defaultValue uint) uint {
	env := os.Getenv(key)

	if env == "" {
		return defaultValue
	}

	value, err := strconv.ParseUint(env, 10, 64)
	if err != nil {
		return defaultValue
	}

	return uint(value)
}

func getEnvOrDefaultInt(key string, defaultValue int) int {
	env := os.Getenv(key)

	if env == "" {
		return defaultValue
	}

	value, err := strconv.ParseInt(env, 10, 64)
	if err != nil {
		return defaultValue
	}

	return int(value)
}

func getEnvOrDefaultBool(key string, defaultValue bool) bool {
	env := os.Getenv(key)

	if env == "" {
		return defaultValue
	}

	value, err := strconv.ParseBool(env)
	if err != nil {
		return defaultValue
	}

	return value
}
