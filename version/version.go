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

//go:generate mockgen -destination=../mocks/version_version.go -package=mocks github.com/nori-io/nori-common/v2/version Version
type Version interface {
	// Compare compares this version to another version. This
	// returns -1, 0, or 1 if this version is smaller, equal,
	// or larger than the other version, respectively.
	Compare(other Version) int

	// Equal tests if two versions are equal.
	Equal(other Version) bool

	// Metadata returns any metadata that was part of the version
	// string.
	//
	// Metadata is anything that comes after the "+" in the version.
	// For example, with "1.2.3+beta", the metadata is "beta".
	Metadata() string

	// Prerelease returns any prerelease data that is part of the version,
	// or blank if there is no prerelease data.
	//
	// Prerelease information is anything that comes after the "-" in the
	// version (but before any metadata). For example, with "1.2.3-beta",
	// the prerelease information is "beta".
	Prerelease() string

	// Segments returns the numeric segments of the version as a slice of ints.
	Segments() []int

	// String returns the full version string included pre-release
	// and metadata information.
	String() string

	// Original returns the original parsed version as-is
	Original() string
}

type ver struct {
	semVer *semver.Version
}

func NewVersion(v string) (Version, error) {
	semVer, err := semver.NewVersion(v)
	if err != nil {
		return nil, err
	}
	return &ver{
		semVer: semVer,
	}, nil
}

func (v *ver) Compare(other Version) int {
	o, _ := semver.NewVersion(other.String())
	return v.semVer.Compare(o)
}

// Equal tests if two versions are equal.
func (v *ver) Equal(other Version) bool {
	o, _ := semver.NewVersion(other.String())
	return v.semVer.Equal(o)
}

// Metadata returns any metadata that was part of the version
// string.
func (v *ver) Metadata() string {
	return v.semVer.Metadata()
}

// Prerelease returns any prerelease data that is part of the version,
// or blank if there is no prerelease data.
func (v *ver) Prerelease() string {
	return v.semVer.Prerelease()
}

// Segments returns the numeric segments of the version as a slice of ints.
func (v *ver) Segments() []int {
	return []int{
		int(v.semVer.Major()),
		int(v.semVer.Minor()),
		int(v.semVer.Patch()),
	}
}

// String returns the full version string included pre-release
// and metadata information.
func (v *ver) String() string {
	return v.semVer.String()
}

// Original returns the original parsed version as-is
func (v *ver) Original() string {
	return v.semVer.Original()
}
