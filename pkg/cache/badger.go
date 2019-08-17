package cache

import (
	"fmt"
	"time"

	badger "github.com/dgraph-io/badger"
	zip "stillgrove.com/goexp/pkg/gzip"
)

type BadgerCache struct {
	db *badger.DB
}

func NewBadgerCache(dbName string) (c Cache, err error) {
	db, err := badger.Open(badger.DefaultOptions(dbName + ".db"))
	if err != nil {
		return c, err
	}
	return BadgerCache{
		db: db,
	}, nil
}

func (b BadgerCache) LoadRecord(key string) (payload []byte, err error) {
	var zipped []byte
	err = b.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return fmt.Errorf("Item not found - %v", err)
		}

		err = item.Value(func(val []byte) error {
			return nil
		})
		if err != nil {
			return fmt.Errorf("Couldn't fetch value - %v", err)
		}

		zipped, err = item.ValueCopy(nil)
		if err != nil {
			return fmt.Errorf("Couldn't copy value - %v", err)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	payload, err = zip.Unzip(zipped)
	if err != nil {
		return nil, err
	}

	return payload, err
}

func (b BadgerCache) SaveRecord(key string, record []byte, expiration time.Duration) (err error) {

	txn := b.db.NewTransaction(true)

	e := badger.NewEntry([]byte(key), zip.Zip(record)).WithTTL(expiration)
	if err := txn.SetEntry(e); err == badger.ErrTxnTooBig {
		_ = txn.Commit()
		txn = b.db.NewTransaction(true)
		_ = txn.SetEntry(e)
	}

	err = txn.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (b BadgerCache) SaveRecords(updates map[string][]byte, expiration time.Duration) (err error) {
	txn := b.db.NewTransaction(true)
	for k, v := range updates {
		e := badger.NewEntry([]byte(k), zip.Zip(v)).WithTTL(expiration)
		if err := txn.SetEntry(e); err == badger.ErrTxnTooBig {
			_ = txn.Commit()
			txn = b.db.NewTransaction(true)
			_ = txn.SetEntry(e)
		}
	}
	err = txn.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (b BadgerCache) LoadRecords() (outputs map[string][]byte, err error) {
	outputs = make(map[string][]byte)
	err = b.db.View(func(txn *badger.Txn) error {
		var k []byte
		var item *badger.Item

		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10

		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item = it.Item()
			k = item.Key()
			err := item.Value(func(v []byte) error {
				v, _ = zip.Unzip(v)
				outputs[string(k)] = v
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})

	return outputs, err
}

func (b BadgerCache) Close() {
	b.db.Close()
}
