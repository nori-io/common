package meta

import (
	"fmt"

	enum "github.com/nori-io/common/v3/pkg/domain/enum/meta"
	"github.com/nori-io/common/v3/pkg/domain/meta"
)

type (
	Meta struct {
		ID           ID
		Author       Author
		Dependencies []meta.Dependency
		Description  Description
		Interface    meta.Interface
		License      []meta.License
		Links        []meta.Link
		Repository   Repository
		Tags         []string
	}

	ID struct {
		ID      meta.PluginID
		Version string
	}

	Author struct {
		Name string
		URL  string
	}

	Description struct {
		Title       string
		Description string
	}

	License struct {
		Title string
		Type  enum.LicenseType
		URL   string
	}

	Link struct {
		Title string
		URL   string
	}

	Repository struct {
		Type enum.RepositoryType
		URL  string
	}
)

// Meta
func (m Meta) GetID() meta.ID {
	return m.ID
}

func (m Meta) GetAuthor() meta.Author {
	return m.Author
}

func (m Meta) GetDependencies() []meta.Dependency {
	return m.Dependencies
}

func (m Meta) GetDescription() meta.Description {
	return m.Description
}

func (m Meta) GetInterface() meta.Interface {
	return m.Interface
}

func (m Meta) GetLicense() []meta.License {
	return m.License
}

func (m Meta) GetLinks() []meta.Link {
	return m.Links
}

func (m Meta) GetRepository() meta.Repository {
	return m.Repository
}

func (m Meta) GetTags() []string {
	return m.Tags
}

// ID
func (id ID) GetID() meta.PluginID {
	return id.ID
}

func (id ID) GetVersion() string {
	return id.Version
}

func (id ID) String() string {
	return fmt.Sprintf("%s:%s", id.ID, id.Version)
}

// Author
func (a Author) GetName() string {
	return a.Name
}

func (a Author) GetURL() string {
	return a.URL
}

// Description
func (d Description) GetTitle() string {
	return d.Title
}

func (d Description) GetDescription() string {
	return d.Description
}

// License
func (l License) GetTitle() string {
	return l.Title
}

func (l License) GetType() enum.LicenseType {
	return l.Type
}

func (l License) GetURL() string {
	return l.URL
}

// Link
func (l Link) GetTitle() string {
	return l.Title
}

func (l Link) GetURL() string {
	return l.URL
}

// Repository
func (r Repository) GetType() enum.RepositoryType {
	return r.Type
}

func (r Repository) GetURL() string {
	return r.URL
}
