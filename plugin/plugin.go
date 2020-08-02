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

package plugin

import (
	"context"

	"github.com/nori-io/common/v3/config"
	"github.com/nori-io/common/v3/logger"
	"github.com/nori-io/common/v3/meta"
)

//go:generate mockgen -destination=../mocks/plugin_plugin.go -package=mocks github.com/nori-io/common/v3/plugin Plugin
type Plugin interface {
	Meta() meta.Meta
	Instance() interface{}
	Init(ctx context.Context, config config.Config, log logger.FieldLogger) error
	Start(ctx context.Context, registry Registry) error
	Stop(ctx context.Context, registry Registry) error
}

//go:generate mockgen -destination=../mocks/plugin_installable.go -package=mocks github.com/nori-io/common/v3/plugin Installable
type Installable interface {
	Install(ctx context.Context, registry Registry) error
	UnInstall(ctx context.Context, registry Registry) error
}
