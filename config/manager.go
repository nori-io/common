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

import "github.com/secure2work/nori-common/meta"

type Manager interface {
	Register(meta.Meta) Config
	PluginVariables(id meta.ID) []Variable
}

type Config interface {
	Bool(key string, desc string) Bool
	Float(key string, desc string) Float
	Int(key string, desc string) Int
	UInt(key string, desc string) UInt
	Slice(key, delimiter string, desc string) Slice
	String(key string, desc string) String
	StringMap(key string, desc string) StringMap
}

type (
	Bool      func() bool
	Float     func() float64
	Int       func() int
	UInt      func() uint
	Slice     func() []interface{}
	String    func() string
	StringMap func() map[string]interface{}
)

type Variable struct {
	Name        string
	Description string
}
