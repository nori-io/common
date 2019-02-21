/*
Copyright 2019 The Nori Authors.
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

	"github.com/nori-io/nori-common/meta"
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
			ID:         meta.PluginID("nori/a"),
			Constraint: ">=1.0.0",
		},
		meta.Dependency{
			Interface: meta.Auth,
		},
	}

	description := meta.Description{
		Name:        "Nori Common",
		Description: "Nori common interfaces and data structures",
	}

	core := meta.Core{
		VersionConstraint: "1.0.0",
	}

	Interface := meta.Custom

	license := meta.License{
		Title: "Apache 2.0 License",
		Type:  "Apache 2.0",
		URI:   "http://www.apache.org/licenses/LICENSE-2.0",
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
		Interface:    Interface,
		License:      license,
		Links:        links,
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
	a.Equal(tags, data.GetTags(), "incorrect tags")

}
