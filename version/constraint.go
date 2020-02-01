package version

import (
	"github.com/Masterminds/semver/v3"
)

type Constraints interface {
	Check(v Version) bool
	String() string
}

type constraint struct {
	constraint *semver.Constraints
}

func NewConstraint(v string) (Constraints, error) {
	cs, err := semver.NewConstraint(v)
	if err != nil {
		return nil, err
	}
	return &constraint{
		constraint: cs,
	}, nil
}

func (cs constraint) Check(v Version) bool {
	ver, _ := semver.NewVersion(v.Original())
	return cs.constraint.Check(ver)
}

func (cs constraint) String() string {
	return cs.constraint.String()
}
