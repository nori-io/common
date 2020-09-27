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

package meta_test

import (
	"testing"

	"github.com/nori-io/common/v3/pkg/domain/meta"
	"github.com/stretchr/testify/assert"
)

func TestNewInterface(t *testing.T) {
	a := assert.New(t)

	const (
		name = "nori/Dummy"
		ver  = "1.0.3"
	)

	var i = meta.NewInterface(name, ver)

	a.Equal(name, i.Name())
	a.Equal(ver, i.Version())
}

func TestNewInterface_EmptyVersion(t *testing.T) {
	a := assert.New(t)

	const (
		name = "nori/Dummy"
		ver  = "0.0.0"
	)

	var i = meta.NewInterface(name, "")

	a.Equal(name, i.Name())
	a.Equal(ver, i.Version())
}

func TestNewInterface_Panic(t *testing.T) {
	a := assert.New(t)

	a.Panics(func() {
		_ = meta.NewInterface("", "1.0.0")
	}, "meta.NewInterface function did not panic on empty Interface name")
}

func TestInterface(t *testing.T) {
	a := assert.New(t)

	const (
		name = "nori/Dummy"
		ver  = "1.0.3"
	)

	var i = meta.Interface(name + "@" + ver)

	a.Equal(name, i.Name())
	a.Equal(ver, i.Version())
}

func TestInterface_Constraint(t *testing.T) {
	a := assert.New(t)

	a.Equal("^1.0.0", meta.Interface("nori/Dummy@1.0.0").Constraint())
}

func TestInterface_EqualSuccess(t *testing.T) {
	a := assert.New(t)

	x := meta.Interface("nori/Any@1.0.0")
	y := meta.Interface("nori/Any@1.0.0")

	a.True(x.Equal(y))
}

func TestInterface_EqualFailure_DifferentVersions(t *testing.T) {
	a := assert.New(t)

	x := meta.Interface("nori/Any@1.0.0")
	y := meta.Interface("nori/Any@1.0.1")

	a.False(x.Equal(y))
}

func TestInterface_EqualFailure_DifferentNames(t *testing.T) {
	a := assert.New(t)

	x := meta.Interface("nori/Any@1.0.0")
	y := meta.Interface("nori/Some@1.0.0")

	a.False(x.Equal(y))
}

func TestInterface_IsUndefined(t *testing.T) {
	a := assert.New(t)

	empty := meta.Interface("")
	auth := meta.Interface("nori/auth@1.0.0")

	a.True(empty.IsUndefined())
	a.False(auth.IsUndefined())
}

func TestInterface_String(t *testing.T) {
	a := assert.New(t)

	x := "nori/any@1.0.0"
	y := meta.Interface(x)
	a.Equal(x, y.String())

	m := "nori/any"
	n := meta.Interface(m)
	a.Equal(m, n.String())
}

func TestInterface_EmptyName(t *testing.T) {
	a := assert.New(t)

	x := meta.Interface("")
	a.Panics(func() {
		x.Name()
	}, "Interface.Name() function did not panic on empty Interface name")
}

func TestInterface_EmptyNameWithVersion(t *testing.T) {
	a := assert.New(t)

	x := meta.Interface("@1.0.0")
	a.Panics(func() {
		x.Name()
	}, "Interface.Name() function did not panic on empty Interface name")
}

func TestInterface_NameWithEmptyVersion(t *testing.T) {
	a := assert.New(t)

	x := meta.Interface("nori/any")
	a.Equal("0.0.0", x.Version())
}
