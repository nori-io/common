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

package interfaces_test

import (
	"testing"

	"github.com/nori-io/nori-common/interfaces"
	"github.com/stretchr/testify/assert"
)

func TestSessionVerification_String(t *testing.T) {
	a := assert.New(t)

	a.Equal("NoVerify", interfaces.NoVerify.String())
	a.Equal("WhiteList", interfaces.WhiteList.String())
	a.Equal("BlackList", interfaces.BlackList.String())

	none1 := interfaces.SessionVerification(-1)
	a.Equal("", none1.String())

	none2 := interfaces.SessionVerification(1000)
	a.Equal("", none2.String())
}
