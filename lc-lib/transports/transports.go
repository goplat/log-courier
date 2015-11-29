/*
 * Copyright 2014-2015 Jason Woods.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package transports

import (
	"github.com/driskell/log-courier/lc-lib/addresspool"
	"github.com/driskell/log-courier/lc-lib/config"
	"github.com/driskell/log-courier/lc-lib/core"
)

// Response is the generic interface implemented by all response structures
// Endpoints create responses with the necessary metadata for sending
type Response interface {
	Endpoint() Endpoint
}

// Endpoint is the interface implemented by the consumer of a transport,
// to allow the transport to communicate back
type Endpoint interface {
	Pool() *addresspool.Pool
	Ready()
	Fail()
	Recover()
	Finished()
	ResponseChan() chan<- Response
}

// Transport is the generic interface that all transports implement
type Transport interface {
	ReloadConfig(*config.Network) bool
	Write(string, []*core.EventDescriptor) error
	Ping() error
	Fail()
	Shutdown()
}

// transportFactory is the interface that all transport factories implement. The
// transport factory should store the transport's configuration and, when
// NewTransport is called, return an instance of the transport that obeys that
// configuration
type transportFactory interface {
	NewTransport(Endpoint) Transport
}

// NewTransport returns a Transport interface initialised from the given Factory
func NewTransport(factory interface{}, endpoint Endpoint) Transport {
	return factory.(transportFactory).NewTransport(endpoint)
}