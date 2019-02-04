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
)

type Marshaller func(interface{}) ([]byte, error)
type Unmarshaller func(source []byte, destination interface{}) error

type PubSub interface {
	NewPublisher(Marshaller, ...func(interface{})) (Publisher, error)
	NewSubscriber(Unmarshaller, ...func(interface{})) (Subscriber, error)
	Publisher() Publisher
	Subscriber() Subscriber
}

type Publisher interface {
	Publish(context.Context, string, interface{}) error
	Errors() <-chan error
}

type Subscriber interface {
	Start() <-chan Message
	Errors() <-chan error
	Stop() error
}

type Message interface {
	UnMarshal(interface{}) error
	Done() error
}
