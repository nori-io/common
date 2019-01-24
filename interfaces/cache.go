/*
Copyright 2019 The Nori Authors.

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

package interfaces

import (
	"time"

	"errors"
)

type Cache interface {
	Clear() error
	Delete(key []byte) error
	Get(key []byte) ([]byte, error)
	Set(key []byte, value []byte, ttl time.Duration) error
}

var CacheKeyNotFound = errors.New("key not found")
