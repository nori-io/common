/*
Copyright 2018-2020 The Nori Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package storage

import (
	"fmt"

	"github.com/nori-io/common/v2/meta"
)

type (
	//go:generate mockgen -destination=../mocks/storage_storage.go -package=mocks github.com/nori-io/common/storage Storage
	Storage interface {
		// CreateBucket creates a new bucket with the given name and returns it.
		CreateBucket(name string) (Bucket, error)
		// CreateBucketIfNotExists creates a new bucket if it doesn't already exist and returns it.
		CreateBucketIfNotExists(name string) (Bucket, error)
		// Bucket returns a bucket by name.
		Bucket(name string) (Bucket, error)
		// DeleteBucket deletes a bucket with the given name.
		DeleteBucket(name string) error
		// Close current storage: connection to db, file etc
		Close() error
	}

	// Bucket represents a collection of key/value pairs
	//go:generate mockgen -destination=../mocks/storage_bucket.go -package=mocks github.com/nori-io/common/storage Bucket
	Bucket interface {
		KeyValue
		Cursor() Cursor
	}

	//go:generate mockgen -destination=../mocks/storage_key_value.go -package=mocks github.com/nori-io/common/storage KeyValue
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

	//go:generate mockgen -destination=../mocks/storage_cursor.go -package=mocks github.com/nori-io/common/storage Cursor
	Cursor interface {
		First() (key string, value []byte)
		Last() (key string, value []byte)
		Next() (key string, value []byte)
		Prev() (key string, value []byte)
		// Seek moves the cursor to a seek key and returns it,
		// If the key does not exist then then next key is used.
		// If there are no keys, an empty key is returned
		Seek(seek string) (key string, value []byte)
		// Delete removes current key-value
		Delete() error
		// HasNext returns true if next element exists
		HasNext() bool
		// Free current cursor, close cursor transactions, free used memory etc
		Close() error
	}

	NotFound struct {
		Entity string
	}

	AlreadyExists struct {
		Entity string
	}
)

const StorageInterface meta.Interface = "core/Storage@1.0.0"

func (e NotFound) Error() string {
	return fmt.Sprintf("Entity %s not found", e.Entity)
}

func (e AlreadyExists) Error() string {
	return fmt.Sprintf("Entity %s already exists", e.Entity)
}
