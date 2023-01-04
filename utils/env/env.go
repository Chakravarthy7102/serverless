package env

import "os"

func GetEnv(env, defaultValue string) string {
	envinorment := os.Getenv(env)

	if envinorment == "" {
		return defaultValue
	}

	return envinorment
}
