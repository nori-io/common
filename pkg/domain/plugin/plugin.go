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

	"github.com/nori-io/common/v4/pkg/domain/config"
	"github.com/nori-io/common/v4/pkg/domain/logger"
	"github.com/nori-io/common/v4/pkg/domain/meta"
	"github.com/nori-io/common/v4/pkg/domain/registry"
)

//go:generate mockgen -destination=../mocks/plugin/plugin.go -package=mocks github.com/nori-io/common/v4/pkg/domain/plugin Plugin

type Plugin interface {
	Meta() meta.Meta
	Instance() interface{}
	Init(ctx context.Context, config config.Config, log logger.FieldLogger) error
	Start(ctx context.Context, registry registry.Registry) error
	Stop(ctx context.Context, registry registry.Registry) error
}
