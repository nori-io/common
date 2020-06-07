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

package meta

import (
	"fmt"

	"github.com/nori-io/nori-common/v2/version"
)

type PluginID string

type ID struct {
	ID      PluginID
	Version string
}

func (id ID) GetVersion() (version.Version, error) {
	v, err := version.NewVersion(id.Version)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (id ID) String() string {
	return fmt.Sprintf("%s:%s", id.ID, id.Version)
}

type Author struct {
	Name string
	URI  string
}

type Core struct {
	VersionConstraint string
}

func (d Core) GetConstraint() (version.Constraints, error) {
	constraints, err := version.NewConstraint(d.VersionConstraint)
	if err != nil {
		return nil, err
	}
	return constraints, nil
}

type Dependency struct {
	Interface Interface
}

func (d Dependency) GetConstraint() (version.Constraints, error) {
	constraints, err := version.NewConstraint(d.Interface.Constraint())
	if err != nil {
		return nil, err
	}
	return constraints, nil
}

func (d Dependency) String() string {
	return fmt.Sprintf("dependency(interface: %s, constraint: %s)", d.Interface, d.Interface.Constraint())
}

type Description struct {
	Name        string
	Description string
}

type License struct {
	Title string
	Type  string
	URI   string
}

type Link struct {
	Title string
	URL   string
}

type Repository struct {
	Type string
	URI  string
}

//go:generate mockgen -destination=../mocks/meta_meta.go -package=mocks github.com/nori-io/nori-common/v2/meta Meta
type Meta interface {
	Id() ID
	GetAuthor() Author
	GetDependencies() []Dependency
	GetDescription() Description
	GetCore() Core
	GetInterface() Interface
	GetLicense() []License
	GetLinks() []Link
	GetRepository() Repository
	GetTags() []string
}

type Data struct {
	ID           ID
	Author       Author
	Dependencies []Dependency
	Description  Description
	Core         Core
	Interface    Interface
	License      []License
	Links        []Link
	Repository   Repository
	Tags         []string
}

func (m Data) Id() ID {
	return m.ID
}

func (m Data) GetAuthor() Author {
	return m.Author
}

func (m Data) GetDependencies() []Dependency {
	return m.Dependencies
}

func (m Data) GetDescription() Description {
	return m.Description
}

func (m Data) GetCore() Core {
	return m.Core
}

func (m Data) GetInterface() Interface {
	return m.Interface
}

func (m Data) GetLicense() []License {
	return m.License
}

func (m Data) GetLinks() []Link {
	return m.Links
}

func (m Data) GetRepository() Repository {
	return m.Repository
}

func (m Data) GetTags() []string {
	return m.Tags
}
