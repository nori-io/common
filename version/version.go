package version

import (
	"github.com/hashicorp/go-version"
)

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
	semVer *version.Version
}

func NewVersion(v string) (Version, error) {
	semVer, err := version.NewSemver(v)
	if err != nil {
		return nil, err
	}
	return &ver{
		semVer: semVer,
	}, nil
}

func (v *ver) Compare(other Version) int {
	o, _ := version.NewSemver(other.String())
	return v.semVer.Compare(o)
}

// Equal tests if two versions are equal.
func (v *ver) Equal(other Version) bool {
	o, _ := version.NewSemver(other.String())
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
	return v.semVer.Segments()
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
