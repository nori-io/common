package interfaces_test

import (
	"testing"

	"github.com/secure2work/nori-common/interfaces"
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
