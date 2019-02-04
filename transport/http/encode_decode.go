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

package http

import (
	"context"
	"encoding/json"
	"net/http"
)

type DecodeRequest func(context.Context, *http.Request) (request interface{}, err error)

type EncodeResponse func(context.Context, http.ResponseWriter, interface{}) error

type StatusCoder interface {
	StatusCode() int
}

type Headerer interface {
	Headers() http.Header
}

// This function can be called before main
func EncodeWrapper(
	ctx context.Context,
	w http.ResponseWriter,
	response interface{},
	encode EncodeResponse,
) error {
	// Headerer interface
	if headerer, ok := response.(Headerer); ok {
		for k := range headerer.Headers() {
			w.Header().Set(k, headerer.Headers().Get(k))
		}
	}

	// StatusCoder interface
	code := http.StatusOK
	if sc, ok := response.(StatusCoder); ok {
		code = sc.StatusCode()
		w.WriteHeader(code)
	}
	if code == http.StatusNoContent {
		return nil
	}

	return encode(ctx, w, response)
}

func EncodeJSONResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	return EncodeWrapper(ctx, w, response, func(context.Context, http.ResponseWriter, interface{}) error {
		return json.NewEncoder(w).Encode(response)
	})
}
