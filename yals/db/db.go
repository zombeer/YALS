package db

import (
	"time"
	"yals/models"

	"github.com/go-redis/redis"
)

// Redis Database instance
var rdb = GetRDB()
var ttl = time.Minute * 1

// GetRDB :: Database instance getter
func GetRDB() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

// GetOrCreateUsingURL :: check if URL already exist.
func GetOrCreateUsingURL(url string) (*models.Item, error) {
	shortcut, err := rdb.Get(url).Result()
	if err != nil {
		return Create(url)
	}
	return GetByShortcut(shortcut)
}

// GetByShortcut :: Gets URL from database using shortcut
func GetByShortcut(shortcut string) (*models.Item, error) {
	i := models.Item{Shortcut: shortcut}
	url, err := rdb.Get(shortcut).Result()
	if err != nil {
		return nil, err
	}
	i.URL = url
	UpdateTTL(&i)
	return &i, err
}

// Create :: writes items pair to database.
func Create(url string) (*models.Item, error) {
	i := models.Item{}
	i.URL = url
	i.GenerateHASH()
	// Primary item
	rdb.Set(i.Shortcut, i.URL, ttl)
	// Helper item
	rdb.Set(i.URL, i.Shortcut, ttl)
	return &i, nil
}

// UpdateTTL :: Updates ttl
func UpdateTTL(i *models.Item) error {
	rdb.Expire(i.Shortcut, ttl)
	rdb.Expire(i.URL, ttl)
	return nil
}
