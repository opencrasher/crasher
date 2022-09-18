package app

import "os"

func parseEnvOrDefaultValue(env, defaultValue string) string {
	parsedEnv := os.Getenv(env)
	if parsedEnv == "" {
		return defaultValue
	}
	return parsedEnv
}
