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
package meta_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/nori-io/nori-common/v2/meta"
)

func TestInterface_Dependency(t *testing.T) {
	a := assert.New(t)

	var i = meta.Interface("nori/Auth@1.0.0")
	dep := i.Dependency()

	a.Equal(meta.Dependency{
		Constraint: "^1.0.0",
		Interface:  i,
	}, dep)
}

func TestInterface_Equal(t *testing.T) {
	a := assert.New(t)

	i := meta.Interface("nori/Any@1.0.0")
	same := meta.Interface("nori/Any@1.0.0")
	sameV2 := meta.Interface("nori/Any@2.0.0")
	other := meta.Interface("nori/Other@1.0.0")
	noVersion := meta.Interface("nori/Any")

	a.True(i.Equal(same))
	a.False(i.Equal(sameV2))
	a.False(i.Equal(other))
	a.False(i.Equal(noVersion))
}

func TestInterface_IsUndefined(t *testing.T) {
	a := assert.New(t)

	empty := meta.Interface("")
	auth := meta.Interface("nori/auth@1.0.0")

	a.True(empty.IsUndefined())
	a.False(auth.IsUndefined())
}

func TestInterface_NameAndVersion(t *testing.T) {
	a := assert.New(t)

	noVersion := "0.0.0"
	name := "nori/Auth"
	ver := "1.2.3"

	empty := meta.Interface("")
	authNoVersion := meta.NewInterface(name, "")
	authWithVersion := meta.NewInterface(name, ver)

	a.Equal("", empty.Name())
	a.Equal(noVersion, empty.Version())

	a.Equal(name, authNoVersion.Name())
	a.Equal(noVersion, authNoVersion.Version())

	a.Equal(name, authWithVersion.Name())
	a.Equal(ver, authWithVersion.Version())
}
