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

package errors

import (
	"fmt"

	"github.com/nori-io/common/v5/pkg/domain/meta"
)

type InterfaceAssertError struct {
	Interface meta.Interface
}

func (e InterfaceAssertError) Error() string {
	return fmt.Sprintf("can't assert interface to %s", e.Interface)
}

type EntityNotFound struct {
	Entity string
}

func (e EntityNotFound) Error() string {
	return fmt.Sprintf("Entity %s not found", e.Entity)
}

type EntityAlreadyExists struct {
	Entity string
}

func (e EntityAlreadyExists) Error() string {
	return fmt.Sprintf("Entity %s already exists", e.Entity)
}
