package app

import "os"

func ParseEnvOrDefaultValue(env, defaultValue string) string {
	parsedEnv := os.Getenv(env)
	if parsedEnv == "" {
		return defaultValue
	}
	return parsedEnv
}
