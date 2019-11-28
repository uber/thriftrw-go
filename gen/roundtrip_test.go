// Copyright (c) 2015 Uber Technologies, Inc.
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

package gen

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"

	"go.uber.org/thriftrw/protocol"
	"go.uber.org/thriftrw/wire"

	"github.com/stretchr/testify/assert"
)

// assertRoundTrip checks if x.ToWire() results in the given Value and whether
// x.FromWire() with the given value results in the original x.
func assertRoundTrip(t *testing.T, x thriftType, v wire.Value, msg string, args ...interface{}) bool {
	message := fmt.Sprintf(msg, args...)

	b, v := assertSerialization(t, x, v, message)
	if !b {
		return false
	}

	xType := reflect.TypeOf(x)
	if xType.Kind() == reflect.Ptr {
		xType = xType.Elem()
	}

	gotX := reflect.New(xType).Interface().(thriftType)
	if assert.NoError(t, gotX.FromWire(v), "FromWire: %v", message) {
		return assert.Equal(t, x, gotX, "FromWire: %v", message)
	}

	return false
}

// assertSerialization checks if x.ToWire() results in the provided wire.Value.
func assertSerialization(t *testing.T, x thriftType, v wire.Value, message string) (bool, wire.Value) {
	if w, err := x.ToWire(); assert.NoError(t, err, "failed to serialize: %v", x) {
		if !assert.True(
			t, wire.ValuesAreEqual(v, w), "%v: %v.ToWire() != %v", message, x, v) {
			return false, v
		}

		var buff bytes.Buffer
		if !assert.NoError(t, protocol.Binary.Encode(w, &buff), "%v: failed to serialize", message) {
			return false, v
		}

		// Flip v to deserialize(serialize(x.ToWire())) to ensure full round trip.
		newV, err := protocol.Binary.Decode(bytes.NewReader(buff.Bytes()), v.Type())
		if !assert.NoError(t, err, "%v: failed to deserialize", message) {
			return false, newV
		}

		if !assert.True(
			t, wire.ValuesAreEqual(newV, v), "%v: deserialize(serialize(%v.ToWire())) != %v", message, x, v) {
			return false, newV
		}

		v = newV
	}

	return true, v
}