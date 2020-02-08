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

package version_test

import (
	"testing"

	"github.com/nori-io/nori-common/v2/version"
	"github.com/stretchr/testify/assert"
)

func TestConstraint_NewConstraint(t *testing.T) {
	a := assert.New(t)

	_, e := version.NewConstraint(">=2.5.0")
	a.NoError(e)

	_, e = version.NewConstraint("^2.5.0")
	a.NoError(e)

	_, e = version.NewConstraint("")
	a.Error(e)

	_, e = version.NewConstraint("word")
	a.Error(e)
}

func TestConstraint_Check(t *testing.T) {
	a := assert.New(t)

	c, e := version.NewConstraint(">= 2.5.0")
	a.NoError(e)
	a.Equal(">=2.5.0", c.String())

	v1, e := version.NewVersion("2.5.0")
	v2, e := version.NewVersion("2.5.9")
	v3, e := version.NewVersion("2.4.0")

	a.True(c.Check(v1))
	a.True(c.Check(v2))
	a.False(c.Check(v3))
}

func TestConstraint_CheckV2(t *testing.T) {
	a := assert.New(t)

	c, e := version.NewConstraint("^2.5.0")
	a.NoError(e)
	a.Equal("^2.5.0", c.String())

	v1, e := version.NewVersion("2.4.0")
	v2, e := version.NewVersion("2.5.0")
	v3, e := version.NewVersion("2.5.9")
	v4, e := version.NewVersion("2.6.0")
	v5, e := version.NewVersion("3.0.0")
	v6, e := version.NewVersion("3.0.1")
	v7, e := version.NewVersion("3.1.0")

	a.False(c.Check(v1))
	a.True(c.Check(v2))
	a.True(c.Check(v3))
	a.True(c.Check(v4))
	a.False(c.Check(v5))
	a.False(c.Check(v6))
	a.False(c.Check(v7))
}
