package cache

import (
	"fmt"
	"time"
	bolt "github.com/boltdb/bolt"

	z "stillgrove.com/goexp/pkg/gzip"
)

type BoltCache struct {
	client *bolt.DB
	bucketName string
}

func NewBoltCache(DBName, bucketName string)(cache Cache, err error){
	db, err := bolt.Open(DBName+".db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return cache, err
	}
	//defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte(bucketName))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	return BoltCache{
		client: db,
		bucketName: bucketName,
	}, nil
}

func (c BoltCache) SaveRecord(key string, record []byte, expiration time.Duration) (err error){
	//err = c.client.Set(key, z.Zip(record), expiration).Err()

	err = c.client.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(c.bucketName))
		err := b.Put([]byte(key), z.Zip(record))
		return err
	})

	return err
}

func (c BoltCache) LoadRecord(key string) (by []byte, err error){
	/*rec, err := c.client.Get(key).Result()
	if err != nil {
		return b, fmt.Errorf("Fetch record: %v", err)
	}
	b0 := []byte(rec)
	b, err = z.Unzip(b0)
	if err != nil {
		return b, fmt.Errorf("Unzipping record: %v", err)
	}*/

	var rec []byte
	c.client.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(c.bucketName))
		v := b.Get([]byte(key))

		rec, _ = z.Unzip([]byte(v))
		
		return nil
	})

	return rec, nil
}

func (c BoltCache) Close() {
	c.client.Close()
}