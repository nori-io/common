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

package config

import (
	"github.com/nori-io/nori-common/v2/config/types"
	"github.com/nori-io/nori-common/v2/meta"
)

type Manager interface {
	Register(meta.Meta) Config
	PluginVariables(id meta.ID) []Variable
}

type Config interface {
	Bool(key string, desc string) Bool
	Float(key string, desc string) Float
	Int(key string, desc string) Int
	Int8(key string, desc string) Int8
	Int32(key string, desc string) Int32
	Int64(key string, desc string) Int64
	UInt(key string, desc string) UInt
	UInt32(key string, desc string) UInt32
	UInt64(key string, desc string) UInt64
	Slice(key string, desc string) Slice
	SliceInt(key string, desc string) SliceInt
	SliceString(key string, desc string) SliceString
	String(key string, desc string) String
	StringMap(key string, desc string) StringMap
	StringMapInt(key string, desc string) StringMapInt
	StringMapSliceString(key string, desc string) StringMapSliceString
	StringMapString(key string, desc string) StringMapString
}

type (
	Bool                 func() bool
	Float                func() float64
	Int                  func() int
	Int8                 func() int8
	Int32                func() int32
	Int64                func() int64
	UInt                 func() uint
	UInt32               func() uint32
	UInt64               func() uint64
	Slice                func() []interface{}
	SliceInt             func() []int
	SliceString          func() []string
	String               func() string
	StringMap            func() map[string]interface{}
	StringMapInt         func() map[string]int
	StringMapSliceString func() map[string][]string
	StringMapString      func() map[string]string
)

type Variable struct {
	Name        string
	Description string
	Type        types.Kind
}
