package configuration

import (
	"fmt"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

/*
 * API_SECRET - (required) authentication token for management endpoints
 * DATABASE_URL - (required) URL for database to store API Keys
 * CACHE_TYPE - (optional, default = "memory") Type of cache to use. Allowed values = no, memory, redis
 * CACHE_MEMORY_LIMIT - (optional) Maximum number of keys to cache
 * CACHE_REDIS_ADDR - (optional) URL of Redis cache
 * CACHE_REDIS_PASSWORD - (optional) Redis cache password
 * FAILS_TO_BAN - (optional) Number per minute of failed tries to authenticate before ban per IP (default = 10)
 * BAN_DURATION - (optional) Ban duration (in seconds, default = 1800)
 */

type DatabaseConfiguration struct {
	Url string
}

type CacheConfiguration struct {
	Type          string
	MemoryLimit   int
	RedisAddr     string
	RedisPassword string
	HasRedis      bool
}

type SecurityConfiguration struct {
	FailsToBan  int
	BanDuration int
}

type Configuration struct {
	ApiSecret string
	Database  DatabaseConfiguration
	Cache     CacheConfiguration
	Security  SecurityConfiguration
}

func New() *Configuration {
	return &Configuration{
		ApiSecret: getEnv("API_SECRET", "", true),
		Database: DatabaseConfiguration{
			Url: getEnv("DATABASE_URL", "", true),
		},
		Cache: CacheConfiguration{
			Type:          getEnv("CACHE_TYPE", "memory", false),
			MemoryLimit:   getEnvAsNumber("CACHE_MEMORY_LIMIT", 100000, false),
			RedisAddr:     getEnv("CACHE_REDIS_ADDR", "", false),
			RedisPassword: getEnv("CACHE_REDIS_PASSWORD", "", false),
			HasRedis:      getEnv("CACHE_REDIS_ADDR", "", false) != "",
		},
		Security: SecurityConfiguration{
			FailsToBan:  getEnvAsNumber("FAILS_TO_BAN", 10, false),
			BanDuration: getEnvAsNumber("BAN_DURATION", 1800, false),
		},
	}
}

func getEnv(env string, defaultValue string, required bool) string {
	value := os.Getenv(env)
	if value == "" {
		if required {
			panic(fmt.Errorf("missing required environment variable %s", env))
		}
		return defaultValue
	}

	return value
}

func getEnvAsNumber(env string, defaultValue int, required bool) int {
	value := os.Getenv(env)
	if value == "" {
		if required {
			panic(fmt.Errorf("missing required environment variable %s", env))
		}
		return defaultValue
	}

	result, err := strconv.Atoi(env)
	if err != nil {
		panic(fmt.Errorf("invalid value for %s: %s", env, value))
	}

	return result
}
