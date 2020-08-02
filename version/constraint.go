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

package version

import (
	"github.com/Masterminds/semver/v3"
)

//go:generate mockgen -destination=../mocks/version_constraints.go -package=mocks github.com/nori-io/common/v3/version Constraints
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
