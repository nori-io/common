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
	enum "github.com/nori-io/common/v4/pkg/domain/enum/meta"
)

type (
	//go:generate mockgen -destination=../mocks/meta/meta/meta.go -package=mocks github.com/nori-io/common/v4/pkg/domain/meta Meta
	Meta interface {
		GetID() ID
		GetAuthor() Author
		GetDependencies() []Dependency
		GetDescription() Description
		GetInterface() Interface
		GetLicense() []License
		GetLinks() []Link
		GetRepository() Repository
		GetTags() []string
	}

	PluginID string

	//go:generate mockgen -destination=../mocks/meta/meta/id.go -package=mocks github.com/nori-io/common/v4/pkg/domain/meta ID
	ID interface {
		GetID() PluginID
		GetVersion() string
		String() string
	}

	//go:generate mockgen -destination=../mocks/meta/meta/author.go -package=mocks github.com/nori-io/common/v4/pkg/domain/meta Author
	Author interface {
		GetName() string
		GetURL() string
	}

	Dependency interface {
		Constraint() string
		Name() string
		String() string
		Version() string
	}

	//go:generate mockgen -destination=../mocks/meta/meta/description.go -package=mocks github.com/nori-io/common/v4/pkg/domain/meta Description
	Description interface {
		GetTitle() string
		GetDescription() string
	}

	//go:generate mockgen -destination=../mocks/meta/meta/license.go -package=mocks github.com/nori-io/common/v4/pkg/domain/meta License
	License interface {
		GetTitle() string
		GetType() enum.LicenseType
		GetURL() string
	}

	//go:generate mockgen -destination=../mocks/meta/meta/link.go -package=mocks github.com/nori-io/common/v4/pkg/domain/meta Link
	Link interface {
		GetTitle() string
		GetURL() string
	}

	//go:generate mockgen -destination=../mocks/meta/meta/repository.go -package=mocks github.com/nori-io/common/v4/pkg/domain/meta Repository
	Repository interface {
		GetType() enum.RepositoryType
		GetURL() string
	}
)
