package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/go-redis/redis/v8"
	"github.com/mentatxx/traefik-api-key-forward-auth/internal/configuration"
	"github.com/mentatxx/traefik-api-key-forward-auth/models"
)

const (
	redisLockKey = "db_migration_lock"
)

func migrateWithLock(configuration *configuration.Configuration, database *gorm.DB) error {
	if configuration.Cache.HasRedis {

		// Initialize Redis client for locking
		redisClient := redis.NewClient(&redis.Options{
			Addr:     configuration.Cache.RedisAddr,
			Password: configuration.Cache.RedisPassword, // Provide password if required
			DB:       0,
		})

		// Attempt to acquire the lock
		if acquired := acquireLock(redisClient); !acquired {
			log.Println("Failed to acquire lock. Another instance may be performing migrations.")
			os.Exit(1)
		}
		defer func() {
			if err := releaseLock(redisClient); err != nil {
				log.Println("Error releasing lock:", err)
			}
		}()
	}

	// Run auto migrations
	modelsToMigrate := []interface{}{
		&models.Key{},
	}
	database.AutoMigrate(modelsToMigrate...)

	return nil
}

func Connect(config *configuration.Configuration) (*gorm.DB, error) {
	// url := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(config.Database.Url), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = migrateWithLock(config, db)
	if err != nil {
		return nil, err
	}
	return db, err
}
