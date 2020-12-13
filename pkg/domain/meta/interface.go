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

package meta

import (
	"fmt"
	"strings"
)

type Interface string

func NewInterface(name, version string) Interface {
	if len(name) == 0 {
		panic("interface name cannot be empty")
	}
	if len(version) == 0 {
		version = "0.0.0"
	}
	return Interface(fmt.Sprintf("%s@%s", name, version))
}

func (i Interface) Constraint() string {
	return "^" + i.Version()
}

func (i Interface) Equal(other Interface) bool {
	return i == other
}

func (i Interface) IsUndefined() bool {
	return len(i) == 0
}

func (i Interface) Name() string {
	parts := strings.Split(string(i), "@")

	if len(i) == 0 || len(parts[0]) == 0 {
		panic("Interface name cannot be empty")
	}

	return parts[0]
}

func (i Interface) String() string {
	return string(i)
}

func (i Interface) Version() string {
	parts := strings.Split(string(i), "@")
	if len(parts) != 2 {
		return "0.0.0"
	}
	return parts[1]
}
