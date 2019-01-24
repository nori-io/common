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

package interfaces

import (
	"context"
	"errors"

	"github.com/secure2work/nori/core/endpoint"
)

const (
	AuthTokenContextKey = "nori.auth.token"
	AuthDataContextKey  = "nori.auth.data"
)

var (
	ErrorOptionNotDefined = errors.New("option not defined")
)

type AccessTokenOption func(interface{}) interface{}

type Auth interface {
	AccessToken(ops AccessTokenOption) (string, error)
	Authenticated() endpoint.Middleware
	IsAuthenticated(context.Context) bool
}
