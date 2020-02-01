package version_test

import (
	"testing"

	"github.com/nori-io/nori-common/v2/version"
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

func TestVer_Metadata(t *testing.T) {
	a := assert.New(t)

	v, err := version.NewVersion("1.0.0+build.05")
	a.NoError(err)
	a.Equal("build.05", v.Metadata())

	v2, err := version.NewVersion("1.0.0")
	a.NoError(err)
	a.Equal("", v2.Metadata())
}

func TestVer_Prerelease(t *testing.T) {
	a := assert.New(t)

	v, err := version.NewVersion("1.0.0-rc1")
	a.NoError(err)
	a.Equal("rc1", v.Prerelease())

	v2, err := version.NewVersion("1.0.0")
	a.NoError(err)
	a.Equal("", v2.Prerelease())
}

func TestVer_Segments(t *testing.T) {
	a := assert.New(t)

	v, err := version.NewVersion("1.0.5")
	a.NoError(err)

	a.Equal([]int{1, 0, 5}, v.Segments())
	a.NotEqual([]int{0, 1, 5}, v.Segments())
}
