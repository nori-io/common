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
	"net/http"
)

// Server executes BeforeFuncs prior to invoking the endpoint.
// Client executes BeforeFuncs after creating the request
// but prior to invoking the HTTP client.
type BeforeFunc func(context.Context, *http.Request) context.Context

type ServerAfterFunc func(context.Context, http.ResponseWriter) context.Context

type ClientAfterFunc func(context.Context, *http.Response) context.Context

func ServerBefore(before ...BeforeFunc) ServerOption {
	return func(s *Server) { s.before = append(s.before, before...) }
}

func ServerAfter(after ...ServerAfterFunc) ServerOption {
	return func(s *Server) { s.after = append(s.after, after...) }
}

func ClientBefore(before ...BeforeFunc) ClientOption {
	return func(c *Client) { c.before = append(c.before, before...) }
}

func ClientAfter(after ...ClientAfterFunc) ClientOption {
	return func(c *Client) { c.after = append(c.after, after...) }
}
