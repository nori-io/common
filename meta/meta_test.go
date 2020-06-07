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

package meta_test

import (
	"testing"

	"github.com/nori-io/nori-common/v2/meta"
	"github.com/stretchr/testify/assert"
)

func TestData_MetaInterface(t *testing.T) {
	a := assert.New(t)

	id := meta.ID{
		ID:      meta.PluginID("nori/common"),
		Version: "1.0.0",
	}

	author := meta.Author{
		Name: "Nori Developers",
		URI:  "https://nori.io",
	}

	dependencies := []meta.Dependency{
		meta.Dependency{
			Interface: meta.Interface("Auth@1.0.0"),
		},
	}

	description := meta.Description{
		Name:        "Nori Common",
		Description: "Nori common interfaces and data structures",
	}

	core := meta.Core{
		VersionConstraint: "1.0.0",
	}

	Interface := meta.Interface("Custom@1.0.0")

	license := []meta.License{
		meta.License{
			Title: "Apache 2.0 License",
			Type:  "Apache 2.0",
			URI:   "http://www.apache.org/licenses/LICENSE-2.0",
		},
	}

	repo := meta.Repository{
		Type: "git",
		URI:  "github.com/nori-io/nori-common",
	}

	links := []meta.Link{
		meta.Link{
			Title: "github",
			URL:   "https://github.com/nori-io/nori-common",
		},
	}

	tags := []string{"nori", "common", "test"}

	data := meta.Data{
		ID:           id,
		Author:       author,
		Dependencies: dependencies,
		Description:  description,
		Core:         core,
		Interface:    meta.Interface(Interface),
		License:      license,
		Links:        links,
		Repository:   repo,
		Tags:         tags,
	}

	a.Equal(id, data.Id(), "incorrect ID")
	a.Equal(author, data.GetAuthor(), "incorrect Author")
	a.Equal(dependencies, data.GetDependencies(), "incorrect Dependencies")
	a.Equal(description, data.GetDescription(), "incorrect Description")
	a.Equal(core, data.GetCore(), "incorrect Core")
	a.Equal(Interface, data.GetInterface(), "incorrect Interface")
	a.Equal(license, data.GetLicense(), "incorrect License")
	a.Equal(links, data.GetLinks(), "incorrect links")
	a.Equal(repo, data.GetRepository(), "incorrect repo")
	a.Equal(tags, data.GetTags(), "incorrect tags")

}

func TestDependency_GetConstraint(t *testing.T) {
	a := assert.New(t)

	dep1 := meta.Dependency{
		Interface: "nori/Test@1.2.0",
	}
	_, err := dep1.GetConstraint()
	a.NoError(err)

	dep2 := meta.Dependency{
		Interface: "nori/Test",
	}
	_, err = dep2.GetConstraint()
	a.NoError(err)

	dep3 := meta.Dependency{
		Interface: "nori/Test@word",
	}
	_, err = dep3.GetConstraint()
	a.Error(err)
}

func TestDependency_String(t *testing.T) {
	a := assert.New(t)

	a.Equal("dependency(interface: nori/Test@1.2.0, constraint: ^1.2.0)", meta.Dependency{
		Interface: "nori/Test@1.2.0",
	}.String())

	a.Equal("dependency(interface: nori/Test, constraint: ^0.0.0)", meta.Dependency{
		Interface: "nori/Test",
	}.String())
}

func TestCore_GetConstraint(t *testing.T) {
	a := assert.New(t)

	core1 := meta.Core{
		VersionConstraint: "^1.2.0",
	}
	_, err := core1.GetConstraint()
	a.NoError(err)

	core2 := meta.Core{
		VersionConstraint: "",
	}
	_, err = core2.GetConstraint()
	a.Error(err)

	core3 := meta.Core{
		VersionConstraint: "word",
	}
	_, err = core3.GetConstraint()
	a.Error(err)
}

func TestID_String(t *testing.T) {
	a := assert.New(t)

	a.Equal("nori/Test:1.2.0", meta.ID{
		ID:      "nori/Test",
		Version: "1.2.0",
	}.String())

	a.Equal("nori/Test:", meta.ID{
		ID:      "nori/Test",
		Version: "",
	}.String())

	a.Equal("nori/Test:word", meta.ID{
		ID:      "nori/Test",
		Version: "word",
	}.String())
}

func TestID_GetVersion(t *testing.T) {
	a := assert.New(t)

	v, err := meta.ID{
		ID:      "nori/Test",
		Version: "1.2.0",
	}.GetVersion()

	a.Equal("1.2.0", v.String())
	a.NoError(err)

	_, err = meta.ID{
		ID:      "nori/Test",
		Version: "",
	}.GetVersion()
	a.Error(err)

	_, err = meta.ID{
		ID:      "nori/Test",
		Version: "word",
	}.GetVersion()
	a.Error(err)
}
