package cache

import (
	"fmt"
	"time"
	"github.com/go-redis/redis"

	z "stillgrove.com/goexp/pkg/gzip"
)

type Cache struct {
	client *redis.Client
	compressed bool
}

func BuildCache(host string, port int, password string, db int, compressed bool)(cache Cache, err error){
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: password,
		DB:       db,  // 0 = default DB
	})

	_, err = client.Ping().Result()
	if err != nil {
		return cache, err
	}

	if err != nil {
		return cache, err
	}

	return Cache{
		client: client,
		compressed: compressed,
	}, nil
}

func (c *Cache) SaveRecord(key string, record []byte, expiration time.Duration) (err error){
	err = c.client.Set(key, z.Zip(record), expiration).Err()
	return err
}

func (c *Cache) LoadRecord(key string) (b []byte, err error){
	rec, err := c.client.Get(key).Result()
	if err != nil {
		return b, fmt.Errorf("Fetch record: %v", err)
	}
	b0 := []byte(rec)
	b, err = z.Unzip(b0)
	if err != nil {
		return b, fmt.Errorf("Unzipping record: %v", err)
	}

	return b, nil
}