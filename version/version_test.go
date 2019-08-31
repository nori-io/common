package version_test

import (
	"testing"

	"github.com/nori-io/nori-common/version"
	"github.com/stretchr/testify/assert"
)

func TestNewVersion(t *testing.T) {
	a := assert.New(t)

	_, e := version.NewVersion("1.0.0")
	a.NoError(e)

	_, e = version.NewVersion("ver")
	a.Error(e)
}

func TestVer_Equal(t *testing.T) {
	a := assert.New(t)

	v1, _ := version.NewVersion("1.0.5")
	v2, _ := version.NewVersion("1.0.5")
	v3, _ := version.NewVersion("2.0.0")

	a.True(v1.Equal(v2))
	a.False(v1.Equal(v3))
}

func TestVer_Compare(t *testing.T) {
	a := assert.New(t)

	v1, _ := version.NewVersion("1.0.5")
	v2, _ := version.NewVersion("1.0.5")
	v3, _ := version.NewVersion("1.0.1")
	v4, _ := version.NewVersion("1.0.9")

	a.Equal(0, v1.Compare(v2))
	a.Equal(1, v1.Compare(v3))
	a.Equal(-1, v1.Compare(v4))
}
