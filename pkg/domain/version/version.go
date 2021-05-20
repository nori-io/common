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

//go:generate mockgen -destination=../mocks/version/version.go -package=mocks github.com/nori-io/common/v5/pkg/domain/version Version
//go:generate mockgen -destination=../mocks/version/constraints.go -package=mocks github.com/nori-io/common/v5/pkg/domain/version Constraints

type (
	Version interface {
		// Compare compares this version to another version. This
		// returns -1, 0, or 1 if this version is smaller, equal,
		// or larger than the other version, respectively.
		Compare(other Version) int

		// Equal tests if two versions are equal.
		Equal(other Version) bool

		// Metadata returns any metadata that was part of the version string.
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

		// String returns the full version string included pre-release and metadata information.
		String() string

		// Original returns the original parsed version as-is
		Original() string
	}

	Constraints interface {
		Check(v Version) bool
		String() string
	}
)
