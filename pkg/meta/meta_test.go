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

	enum "github.com/nori-io/common/v3/pkg/domain/enum/meta"
	dommeta "github.com/nori-io/common/v3/pkg/domain/meta"
	"github.com/nori-io/common/v3/pkg/meta"
	"github.com/stretchr/testify/assert"
)

func TestMeta_Interface(t *testing.T) {
	a := assert.New(t)

	id := meta.ID{
		ID:      dommeta.PluginID("nori/common"),
		Version: "1.0.0",
	}

	author := meta.Author{
		Name: "Nori Developers",
		URL:  "https://nori.io",
	}

	dependencies := []dommeta.Dependency{
		dommeta.NewInterface("nori/Any", "0.1.0"),
	}

	description := meta.Description{
		Title:       "Nori Common",
		Description: "Nori common interfaces and data structures",
	}

	Interface := dommeta.Interface("core/Dummy@0.1.0")

	license := []dommeta.License{
		meta.License{
			Title: "Apache 2.0 License",
			Type:  enum.Apache2_0,
			URL:   "http://www.apache.org/licenses/LICENSE-2.0",
		},
	}

	repo := meta.Repository{
		Type: enum.Git,
		URL:  "github.com/nori-io/common",
	}

	links := []dommeta.Link{
		meta.Link{
			Title: "github",
			URL:   "https://github.com/nori-io/common",
		},
	}

	tags := []string{"nori", "common", "test"}

	data := meta.Meta{
		ID:           id,
		Author:       author,
		Dependencies: dependencies,
		Description:  description,
		Interface:    Interface,
		License:      license,
		Links:        links,
		Repository:   repo,
		Tags:         tags,
	}

	a.Equal(id, data.GetID(), "incorrect ID")
	a.Equal(author, data.GetAuthor(), "incorrect Author")
	a.Equal(dependencies, data.GetDependencies(), "incorrect Dependencies")
	a.Equal(description, data.GetDescription(), "incorrect Description")
	a.Equal(Interface, data.GetInterface(), "incorrect Interface")
	a.Equal(license, data.GetLicense(), "incorrect License")
	a.Equal(links, data.GetLinks(), "incorrect links")
	a.Equal(repo, data.GetRepository(), "incorrect repo")
	a.Equal(tags, data.GetTags(), "incorrect tags")

}

func TestID_String(t *testing.T) {
	a := assert.New(t)

	a.Equal("nori/any:1.2.0", meta.ID{
		ID:      "nori/any",
		Version: "1.2.0",
	}.String())

	a.Equal("nori/any:", meta.ID{
		ID:      "nori/any",
		Version: "",
	}.String())

	a.Equal("nori/any:word", meta.ID{
		ID:      "nori/any",
		Version: "word",
	}.String())
}

func TestID_GetIDAndVersion(t *testing.T) {
	a := assert.New(t)

	const (
		id  = dommeta.PluginID("nori/any")
		ver = "1.2.0"
	)

	v := meta.ID{
		ID:      id,
		Version: ver,
	}

	a.Equal(id, v.GetID())
	a.Equal(ver, v.GetVersion())
}

func TestID_GetEmptyVersion(t *testing.T) {
	a := assert.New(t)

	const (
		id = dommeta.PluginID("nori/any")
	)

	v := meta.ID{
		ID:      id,
		Version: "",
	}

	a.Equal(id, v.GetID())
	a.Equal("", v.GetVersion())
}

func TestAuthor(t *testing.T) {
	a := assert.New(t)

	const (
		name = "nori authors"
		url  = "https://nori.io"
	)

	author := meta.Author{
		Name: name,
		URL:  url,
	}

	a.Equal(name, author.GetName())
	a.Equal(url, author.GetURL())
}

func TestDescription(t *testing.T) {
	a := assert.New(t)

	const (
		title       = "lorem ipsum"
		description = "lorem ipsum dolor sit amet"
	)

	desc := meta.Description{
		Title:       title,
		Description: description,
	}

	a.Equal(title, desc.GetTitle())
	a.Equal(description, desc.GetDescription())
}

func TestLicense(t *testing.T) {
	a := assert.New(t)

	const (
		title = "license"
		typ   = enum.Apache2_0
		url   = "https://example.com/license"
	)

	desc := meta.License{
		Title: title,
		Type:  typ,
		URL:   url,
	}

	a.Equal(title, desc.GetTitle())
	a.Equal(typ, desc.GetType())
	a.Equal(url, desc.GetURL())
}

func TestLink(t *testing.T) {
	a := assert.New(t)

	const (
		title = "nori official website"
		url   = "https://nori.io"
	)

	author := meta.Link{
		Title: title,
		URL:   url,
	}

	a.Equal(title, author.GetTitle())
	a.Equal(url, author.GetURL())
}

func TestRepository(t *testing.T) {
	a := assert.New(t)

	const (
		typ = enum.Git
		url = "https://nori.io"
	)

	author := meta.Repository{
		Type: typ,
		URL:  url,
	}

	a.Equal(typ, author.GetType())
	a.Equal(url, author.GetURL())
}
