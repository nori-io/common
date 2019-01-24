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
	"time"

	"github.com/secure2work/nori/core/endpoint"
)

const (
	SessionIdContextKey = "nori.session.id"
)

type SessionVerification int

type Session interface {
	Get(key []byte, data interface{}) error
	Save(key []byte, data interface{}, exp time.Duration) error
	Delete(key []byte) error
	SessionId(ctx context.Context) []byte
	Verify() endpoint.Middleware
}

const (
	NoVerify SessionVerification = iota
	WhiteList
	BlackList
)

var sessionVerificationNames = [...]string{
	"NoVerify",
	"WhiteList",
	"BlackList",
}

func (sv SessionVerification) String() string {
	if sv < 0 || int(sv) >= len(sessionVerificationNames) {
		return ""
	}
	return sessionVerificationNames[sv]
}

type State string

const (
	SessionActive  State = "active"
	SessionClosed  State = "closed"
	SessionLocked  State = "locked"
	SessionBlocked State = "blocked"
	SessionExpired State = "expired"
	SessionError   State = "error"
)
