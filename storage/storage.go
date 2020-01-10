package storage

import "fmt"

type (
	Storage interface {
		// CreateBucket creates a new bucket with the given name and returns it.
		CreateBucket(name string) (Bucket, error)
		// CreateBucketIfNotExists creates a new bucket if it doesn't already exist and returns it.
		CreateBucketIfNotExists(name string) (Bucket, error)
		// Bucket returns a bucket by name.
		Bucket(name string) (Bucket, error)
		// DeleteBucket deletes a bucket with the given name.
		DeleteBucket(name string) error
	}

	// Bucket represents a collection of key/value pairs
	Bucket interface {
		KeyValue
		Cursor() Cursor
	}

	KeyValue interface {
		// Get retrieves the value for a key.
		Get(key string) ([]byte, error)
		// Set sets the value for a key.
		Set(key string, value []byte) error
		// Delete removes a key
		Delete(key string) error
		// ForEach executes a function for each key/value pair
		ForEach(fn func(k string, v []byte) error) error
	}

	Cursor interface {
		First() (key string, value []byte)
		Last() (key string, value []byte)
		Next() (key string, value []byte)
		Prev() (key string, value []byte)
		// Seek moves the cursor to a seek key and returns it,
		// If the key does not exist then then next key is used.
		// If there are no keys, a nil key is returned
		Seek(seek string) (key string, value []byte)
		// Delete removes current key-value
		Delete() error
		// HasNext returns true if next element exists
		HasNext() bool
	}

	NotFound struct {
		Entity string
	}

	AlreadyExists struct {
		Entity string
	}
)

func (e NotFound) Error() string {
	return fmt.Sprintf("Entity %s not found", e.Entity)
}

func (e AlreadyExists) Error() string {
	return fmt.Sprintf("Entity %s already exists", e.Entity)
}
