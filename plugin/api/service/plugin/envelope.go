// Code generated by thriftrw-plugin-envelope
// @generated

// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package plugin

import (
	"go.uber.org/thriftrw/internal/envelope"
	"go.uber.org/thriftrw/wire"
	"go.uber.org/thriftrw/plugin/api"
)

// Client implements a Plugin client.
type client struct {
	client envelope.Client
}

// NewClient builds a new Plugin client.
func NewClient(c envelope.Client) api.Plugin {
	return &client{
		client: c,
	}
}

func (c *client) Goodbye() (err error) {
	args := GoodbyeHelper.Args()

	var body wire.Value
	body, err = args.ToWire()
	if err != nil {
		return
	}

	body, err = c.client.Send("goodbye", body)
	if err != nil {
		return
	}

	var result GoodbyeResult
	if err = result.FromWire(body); err != nil {
		return
	}

	err = GoodbyeHelper.UnwrapResponse(&result)
	return
}

func (c *client) Handshake(
	_Request *api.HandshakeRequest,
) (success *api.HandshakeResponse, err error) {
	args := HandshakeHelper.Args(_Request)

	var body wire.Value
	body, err = args.ToWire()
	if err != nil {
		return
	}

	body, err = c.client.Send("handshake", body)
	if err != nil {
		return
	}

	var result HandshakeResult
	if err = result.FromWire(body); err != nil {
		return
	}

	success, err = HandshakeHelper.UnwrapResponse(&result)
	return
}

// Handler serves an implementation of the Plugin service.
type Handler struct {
	impl api.Plugin
}

// NewHandler builds a new Plugin handler.
func NewHandler(service api.Plugin) Handler {
	return Handler{
		impl: service,
	}
}

// Handle receives and handles a request for the Plugin service.
func (h Handler) Handle(name string, reqValue wire.Value) (wire.Value, error) {
	switch name {

	case "goodbye":
		var args GoodbyeArgs
		if err := args.FromWire(reqValue); err != nil {
			return wire.Value{}, err
		}

		result, err := GoodbyeHelper.WrapResponse(
			h.impl.Goodbye(),
		)
		if err != nil {
			return wire.Value{}, err
		}

		return result.ToWire()

	case "handshake":
		var args HandshakeArgs
		if err := args.FromWire(reqValue); err != nil {
			return wire.Value{}, err
		}

		result, err := HandshakeHelper.WrapResponse(
			h.impl.Handshake(args.Request),
		)
		if err != nil {
			return wire.Value{}, err
		}

		return result.ToWire()

	default:

		return wire.Value{}, envelope.ErrUnknownMethod(name)

	}
}
