package meta_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/secure2work/nori-common/meta"
)

func TestInterface_Dependency(t *testing.T) {
	a := assert.New(t)

	i := meta.Auth
	dep := i.Dependency("1.0")

	a.Equal(meta.Dependency{
		ID:         "",
		Constraint: "~1.0",
		Interface:  meta.Auth,
	}, dep)
}
