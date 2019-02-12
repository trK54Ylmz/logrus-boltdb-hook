package logrusbolt

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/sirupsen/logrus"
	bolt "go.etcd.io/bbolt"
)

const (
	format = "2006-01-02 15:04:05.000000000"
)

type BoltHook struct {
	DBLoc     string
	Bucket    string
	Formatter logrus.Formatter

	db *bolt.DB
}

// Creates a hook for instance of logrus logger
func NewHook(b BoltHook) (*BoltHook, error) {
	boltDB, err := bolt.Open(b.DBLoc, 0600, nil)

	if err != nil {
		return nil, err
	}

	return &BoltHook{
		DBLoc:     b.DBLoc,
		Bucket:    b.Bucket,
		Formatter: b.Formatter,
		db:        boltDB,
	}, nil
}

// Formats boltdb key
func (b *BoltHook) now() string {
	return time.Now().Format(format) + "." + fmt.Sprint(rand.Uint32())
}

// Calls Fire method when event is fired
func (b *BoltHook) Fire(e *logrus.Entry) error {
	return b.db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(b.Bucket))

		if err != nil {
			return err
		}

		bytes, err := b.Formatter.Format(e)

		if err != nil {
			return err
		}

		bucket.Put([]byte(b.now()), bytes)

		return nil
	})
}

// Returns the available logging levels in logrus
func (b *BoltHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
