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
	"net/http"

	"github.com/nori-io/nori-common/endpoint"
	"github.com/nori-io/nori-common/logger"
	"github.com/nori-io/nori-common/transport"
)

type Server struct {
	e            endpoint.Endpoint
	decode       DecodeRequest
	encode       EncodeResponse
	before       []ServerBeforeFunc
	after        []ServerAfterFunc
	errorEncoder ErrorEncoder
	errorHandler transport.ErrorHandler
}

type ServerOption func(*Server)

func ServerErrorLogger(logger logger.Logger) ServerOption {
	return func(s *Server) { s.errorHandler = transport.NewLogErrorHandler(logger) }
}

func ServerErrorHandler(errorHandler transport.ErrorHandler) ServerOption {
	return func(s *Server) { s.errorHandler = errorHandler }
}

func NewServer(
	e endpoint.Endpoint,
	decode DecodeRequest,
	encode EncodeResponse,
	options ...ServerOption,
) *Server {
	s := &Server{
		e:            e,
		decode:       decode,
		encode:       encode,
		errorEncoder: DefaultErrorEncoder,
		//errorHandler: , // add nop error handler
	}
	for _, option := range options {
		option(s)
	}
	return s
}

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	for _, f := range s.before {
		ctx = f(ctx, r)
	}

	request, err := s.decode(ctx, r)
	if err != nil {
		if s.errorHandler != nil {
			s.errorHandler.Handle(ctx, err)
		}
		s.errorEncoder(ctx, err, w)
		return
	}

	response, err := s.e(ctx, request)

	if err != nil {
		if s.errorHandler != nil {
			s.errorHandler.Handle(ctx, err)
		}
		s.errorEncoder(ctx, err, w)
		return
	}

	for _, f := range s.after {
		ctx = f(ctx, w)
	}

	if err := s.encode(ctx, w, response); err != nil {
		if s.errorHandler != nil {
			s.errorHandler.Handle(ctx, err)
		}
		s.errorEncoder(ctx, err, w)
		return
	}
}
