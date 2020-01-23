package version_test

import (
	"testing"

	"github.com/nori-io/nori-common/v2/version"
	"github.com/stretchr/testify/assert"
)

func TestConstraint_Check(t *testing.T) {
	a := assert.New(t)

	c, e := version.NewConstraint(">= 2.5.0")
	a.NoError(e)
	a.Equal(">= 2.5.0", c.String())

	v1, e := version.NewVersion("2.5.0")
	v2, e := version.NewVersion("2.5.9")
	v3, e := version.NewVersion("2.4.0")

	a.True(c.Check(v1))
	a.True(c.Check(v2))
	a.False(c.Check(v3))
}
