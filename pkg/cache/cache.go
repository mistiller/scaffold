package cache

import(
	"time"
)

type Cache interface{
	SaveRecord(key string, record []byte, expiration time.Duration) error
	LoadRecord(key string) (b []byte, err error)
	Close()
}