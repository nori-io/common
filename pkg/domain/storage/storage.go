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
	"github.com/nori-io/common/v3/pkg/domain/meta"
)

const StorageInterface meta.Interface = "core/Storage@1.0.0"

//go:generate mockgen -destination=../mocks/storage/storage.go -package=mocks github.com/nori-io/common/v3/pkg/domain/storage Storage
type Storage interface {
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
