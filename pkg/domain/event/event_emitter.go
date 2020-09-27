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

package event

//go:generate mockgen -destination=../mocks/event/event_emitter.go -package=mocks github.com/nori-io/common/v3/pkg/domain/event EventEmitter
type EventEmitter interface {
	Use(pattern string, middleware ...Middleware)
	On(name string, middleware ...Middleware) (evt <-chan Event, unsubscribe func())
	Emit(event Event)
}

type Middleware func(event *Event)

type Event struct {
	Name   string
	Params map[string]interface{}
}
