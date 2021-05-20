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

//go:generate mockgen -destination=../mocks/storage/bucket.go -package=mocks github.com/nori-io/common/v4/pkg/domain/storage Bucket

// Bucket represents a collection of key/value pairs
type Bucket interface {
	KeyValue
	Cursor() Cursor
}

//go:generate mockgen -destination=../mocks/storage/key_value.go -package=mocks github.com/nori-io/common/v4/pkg/domain/storage KeyValue

type KeyValue interface {
	// Get retrieves the value for a key.
	Get(key string) ([]byte, error)
	// Set sets the value for a key.
	Set(key string, value []byte) error
	// Delete removes a key
	Delete(key string) error
	// ForEach executes a function for each key/value pair
	ForEach(fn func(k string, v []byte) error) error
}
