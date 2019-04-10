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

package plugin

import (
	"github.com/nori-io/nori-common/config"
	"github.com/nori-io/nori-common/interfaces"
	"github.com/nori-io/nori-common/meta"
)

type Registry interface {
	Resolve(dep meta.Dependency) (interface{}, error)

	Auth() (interfaces.Auth, error)
	Authorize() (interfaces.Authorize, error)
	Cache() (interfaces.Cache, error)
	Config() config.Manager
	Http() (interfaces.Http, error)
	HTTPTransport() (interfaces.HTTPTransport, error)
	Logger(meta meta.Meta) interfaces.Logger
	Mail() (interfaces.Mail, error)
	PubSub() (interfaces.PubSub, error)
	Session() (interfaces.Session, error)
	Sql() (interfaces.SQL, error)
	Templates() (interfaces.Templates, error)
}
