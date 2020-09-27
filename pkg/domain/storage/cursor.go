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

//go:generate mockgen -destination=../mocks/storage/cursor.go -package=mocks github.com/nori-io/common/v3/pkg/domain/storage Cursor
type Cursor interface {
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
