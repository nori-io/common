package version

import (
	"strings"

	"github.com/hashicorp/go-version"
)

type Constraints interface {
	Check(v Version) bool
	String() string
}

type constraint struct {
	constraint *version.Constraint
}

type constraints []*constraint

func NewConstraint(v string) (Constraints, error) {
	cs, err := version.NewConstraint(v)
	if err != nil {
		return nil, err
	}
	csList := constraints{}
	for _, c := range cs {
		csList = append(csList, &constraint{
			constraint: c,
		})
	}
	return csList, nil
}

func (cs constraint) Check(v Version) bool {
	ver, _ := version.NewSemver(v.Original())
	return cs.constraint.Check(ver)
}

func (cs constraint) String() string {
	return cs.constraint.String()
}

func (cs constraints) Check(v Version) bool {
	for _, c := range cs {
		if !c.Check(v) {
			return false
		}
	}

	return true
}

func (cs constraints) String() string {
	items := make([]string, len(cs))
	for i, c := range cs {
		items[i] = c.String()
	}
	return strings.Join(items, ",")
}
