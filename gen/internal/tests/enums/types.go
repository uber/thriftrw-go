// Code generated by thriftrw v1.13.0. DO NOT EDIT.
// @generated

package enums

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go.uber.org/thriftrw/wire"
	"go.uber.org/zap/zapcore"
	"math"
	"strconv"
	"strings"
)

type EmptyEnum int32

// UnmarshalText tries to decode EmptyEnum from a byte slice
// containing its name.
func (v *EmptyEnum) UnmarshalText(value []byte) error {
	switch string(value) {
	default:
		return fmt.Errorf("unknown enum value %q for %q", value, "EmptyEnum")
	}
}

// MarshalText encodes EmptyEnum to text.
//
// If the enum value is recognized, its name is returned. Otherwise,
// its integer value is returned.
//
// This implements the TextMarshaler interface.
func (v EmptyEnum) MarshalText() ([]byte, error) {
	return []byte(strconv.FormatInt(int64(v), 10)), nil
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of EmptyEnum.
func (v EmptyEnum) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddInt32("value", int32(v))
	return nil
}

// Ptr returns a pointer to this enum value.
func (v EmptyEnum) Ptr() *EmptyEnum {
	return &v
}

// ToWire translates EmptyEnum into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// Enums are represented as 32-bit integers over the wire.
func (v EmptyEnum) ToWire() (wire.Value, error) {
	return wire.NewValueI32(int32(v)), nil
}

// FromWire deserializes EmptyEnum from its Thrift-level
// representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TI32)
//   if err != nil {
//     return EmptyEnum(0), err
//   }
//
//   var v EmptyEnum
//   if err := v.FromWire(x); err != nil {
//     return EmptyEnum(0), err
//   }
//   return v, nil
func (v *EmptyEnum) FromWire(w wire.Value) error {
	*v = (EmptyEnum)(w.GetI32())
	return nil
}

// String returns a readable string representation of EmptyEnum.
func (v EmptyEnum) String() string {
	w := int32(v)
	return fmt.Sprintf("EmptyEnum(%d)", w)
}

// Equals returns true if this EmptyEnum value matches the provided
// value.
func (v EmptyEnum) Equals(rhs EmptyEnum) bool {
	return v == rhs
}

// MarshalJSON serializes EmptyEnum into JSON.
//
// If the enum value is recognized, its name is returned. Otherwise,
// its integer value is returned.
//
// This implements json.Marshaler.
func (v EmptyEnum) MarshalJSON() ([]byte, error) {
	return ([]byte)(strconv.FormatInt(int64(v), 10)), nil
}

// UnmarshalJSON attempts to decode EmptyEnum from its JSON
// representation.
//
// This implementation supports both, numeric and string inputs. If a
// string is provided, it must be a known enum name.
//
// This implements json.Unmarshaler.
func (v *EmptyEnum) UnmarshalJSON(text []byte) error {
	d := json.NewDecoder(bytes.NewReader(text))
	d.UseNumber()
	t, err := d.Token()
	if err != nil {
		return err
	}

	switch w := t.(type) {
	case json.Number:
		x, err := w.Int64()
		if err != nil {
			return err
		}
		if x > math.MaxInt32 {
			return fmt.Errorf("enum overflow from JSON %q for %q", text, "EmptyEnum")
		}
		if x < math.MinInt32 {
			return fmt.Errorf("enum underflow from JSON %q for %q", text, "EmptyEnum")
		}
		*v = (EmptyEnum)(x)
		return nil
	case string:
		return v.UnmarshalText([]byte(w))
	default:
		return fmt.Errorf("invalid JSON value %q (%T) to unmarshal into %q", t, t, "EmptyEnum")
	}
}

type EnumDefault int32

const (
	EnumDefaultFoo EnumDefault = 0
	EnumDefaultBar EnumDefault = 1
	EnumDefaultBaz EnumDefault = 2
)

// EnumDefault_Values returns all recognized values of EnumDefault.
func EnumDefault_Values() []EnumDefault {
	return []EnumDefault{
		EnumDefaultFoo,
		EnumDefaultBar,
		EnumDefaultBaz,
	}
}

// UnmarshalText tries to decode EnumDefault from a byte slice
// containing its name.
//
//   var v EnumDefault
//   err := v.UnmarshalText([]byte("Foo"))
func (v *EnumDefault) UnmarshalText(value []byte) error {
	switch string(value) {
	case "Foo":
		*v = EnumDefaultFoo
		return nil
	case "Bar":
		*v = EnumDefaultBar
		return nil
	case "Baz":
		*v = EnumDefaultBaz
		return nil
	default:
		return fmt.Errorf("unknown enum value %q for %q", value, "EnumDefault")
	}
}

// MarshalText encodes EnumDefault to text.
//
// If the enum value is recognized, its name is returned. Otherwise,
// its integer value is returned.
//
// This implements the TextMarshaler interface.
func (v EnumDefault) MarshalText() ([]byte, error) {
	switch int32(v) {
	case 0:
		return []byte("Foo"), nil
	case 1:
		return []byte("Bar"), nil
	case 2:
		return []byte("Baz"), nil
	}
	return []byte(strconv.FormatInt(int64(v), 10)), nil
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of EnumDefault.
func (v EnumDefault) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddInt32("value", int32(v))
	switch int32(v) {
	case 0:
		enc.AddString("name", "Foo")
	case 1:
		enc.AddString("name", "Bar")
	case 2:
		enc.AddString("name", "Baz")
	}
	return nil
}

// Ptr returns a pointer to this enum value.
func (v EnumDefault) Ptr() *EnumDefault {
	return &v
}

// ToWire translates EnumDefault into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// Enums are represented as 32-bit integers over the wire.
func (v EnumDefault) ToWire() (wire.Value, error) {
	return wire.NewValueI32(int32(v)), nil
}

// FromWire deserializes EnumDefault from its Thrift-level
// representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TI32)
//   if err != nil {
//     return EnumDefault(0), err
//   }
//
//   var v EnumDefault
//   if err := v.FromWire(x); err != nil {
//     return EnumDefault(0), err
//   }
//   return v, nil
func (v *EnumDefault) FromWire(w wire.Value) error {
	*v = (EnumDefault)(w.GetI32())
	return nil
}

// String returns a readable string representation of EnumDefault.
func (v EnumDefault) String() string {
	w := int32(v)
	switch w {
	case 0:
		return "Foo"
	case 1:
		return "Bar"
	case 2:
		return "Baz"
	}
	return fmt.Sprintf("EnumDefault(%d)", w)
}

// Equals returns true if this EnumDefault value matches the provided
// value.
func (v EnumDefault) Equals(rhs EnumDefault) bool {
	return v == rhs
}

// MarshalJSON serializes EnumDefault into JSON.
//
// If the enum value is recognized, its name is returned. Otherwise,
// its integer value is returned.
//
// This implements json.Marshaler.
func (v EnumDefault) MarshalJSON() ([]byte, error) {
	switch int32(v) {
	case 0:
		return ([]byte)("\"Foo\""), nil
	case 1:
		return ([]byte)("\"Bar\""), nil
	case 2:
		return ([]byte)("\"Baz\""), nil
	}
	return ([]byte)(strconv.FormatInt(int64(v), 10)), nil
}

// UnmarshalJSON attempts to decode EnumDefault from its JSON
// representation.
//
// This implementation supports both, numeric and string inputs. If a
// string is provided, it must be a known enum name.
//
// This implements json.Unmarshaler.
func (v *EnumDefault) UnmarshalJSON(text []byte) error {
	d := json.NewDecoder(bytes.NewReader(text))
	d.UseNumber()
	t, err := d.Token()
	if err != nil {
		return err
	}

	switch w := t.(type) {
	case json.Number:
		x, err := w.Int64()
		if err != nil {
			return err
		}
		if x > math.MaxInt32 {
			return fmt.Errorf("enum overflow from JSON %q for %q", text, "EnumDefault")
		}
		if x < math.MinInt32 {
			return fmt.Errorf("enum underflow from JSON %q for %q", text, "EnumDefault")
		}
		*v = (EnumDefault)(x)
		return nil
	case string:
		return v.UnmarshalText([]byte(w))
	default:
		return fmt.Errorf("invalid JSON value %q (%T) to unmarshal into %q", t, t, "EnumDefault")
	}
}

type EnumWithDuplicateName int32

const (
	EnumWithDuplicateNameA EnumWithDuplicateName = 0
	EnumWithDuplicateNameB EnumWithDuplicateName = 1
	EnumWithDuplicateNameC EnumWithDuplicateName = 2
	EnumWithDuplicateNameP EnumWithDuplicateName = 3
	EnumWithDuplicateNameQ EnumWithDuplicateName = 4
	EnumWithDuplicateNameR EnumWithDuplicateName = 5
	EnumWithDuplicateNameX EnumWithDuplicateName = 6
	EnumWithDuplicateNameY EnumWithDuplicateName = 7
	EnumWithDuplicateNameZ EnumWithDuplicateName = 8
)

// EnumWithDuplicateName_Values returns all recognized values of EnumWithDuplicateName.
func EnumWithDuplicateName_Values() []EnumWithDuplicateName {
	return []EnumWithDuplicateName{
		EnumWithDuplicateNameA,
		EnumWithDuplicateNameB,
		EnumWithDuplicateNameC,
		EnumWithDuplicateNameP,
		EnumWithDuplicateNameQ,
		EnumWithDuplicateNameR,
		EnumWithDuplicateNameX,
		EnumWithDuplicateNameY,
		EnumWithDuplicateNameZ,
	}
}

// UnmarshalText tries to decode EnumWithDuplicateName from a byte slice
// containing its name.
//
//   var v EnumWithDuplicateName
//   err := v.UnmarshalText([]byte("A"))
func (v *EnumWithDuplicateName) UnmarshalText(value []byte) error {
	switch string(value) {
	case "A":
		*v = EnumWithDuplicateNameA
		return nil
	case "B":
		*v = EnumWithDuplicateNameB
		return nil
	case "C":
		*v = EnumWithDuplicateNameC
		return nil
	case "P":
		*v = EnumWithDuplicateNameP
		return nil
	case "Q":
		*v = EnumWithDuplicateNameQ
		return nil
	case "R":
		*v = EnumWithDuplicateNameR
		return nil
	case "X":
		*v = EnumWithDuplicateNameX
		return nil
	case "Y":
		*v = EnumWithDuplicateNameY
		return nil
	case "Z":
		*v = EnumWithDuplicateNameZ
		return nil
	default:
		return fmt.Errorf("unknown enum value %q for %q", value, "EnumWithDuplicateName")
	}
}

// MarshalText encodes EnumWithDuplicateName to text.
//
// If the enum value is recognized, its name is returned. Otherwise,
// its integer value is returned.
//
// This implements the TextMarshaler interface.
func (v EnumWithDuplicateName) MarshalText() ([]byte, error) {
	switch int32(v) {
	case 0:
		return []byte("A"), nil
	case 1:
		return []byte("B"), nil
	case 2:
		return []byte("C"), nil
	case 3:
		return []byte("P"), nil
	case 4:
		return []byte("Q"), nil
	case 5:
		return []byte("R"), nil
	case 6:
		return []byte("X"), nil
	case 7:
		return []byte("Y"), nil
	case 8:
		return []byte("Z"), nil
	}
	return []byte(strconv.FormatInt(int64(v), 10)), nil
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of EnumWithDuplicateName.
func (v EnumWithDuplicateName) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddInt32("value", int32(v))
	switch int32(v) {
	case 0:
		enc.AddString("name", "A")
	case 1:
		enc.AddString("name", "B")
	case 2:
		enc.AddString("name", "C")
	case 3:
		enc.AddString("name", "P")
	case 4:
		enc.AddString("name", "Q")
	case 5:
		enc.AddString("name", "R")
	case 6:
		enc.AddString("name", "X")
	case 7:
		enc.AddString("name", "Y")
	case 8:
		enc.AddString("name", "Z")
	}
	return nil
}

// Ptr returns a pointer to this enum value.
func (v EnumWithDuplicateName) Ptr() *EnumWithDuplicateName {
	return &v
}

// ToWire translates EnumWithDuplicateName into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// Enums are represented as 32-bit integers over the wire.
func (v EnumWithDuplicateName) ToWire() (wire.Value, error) {
	return wire.NewValueI32(int32(v)), nil
}

// FromWire deserializes EnumWithDuplicateName from its Thrift-level
// representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TI32)
//   if err != nil {
//     return EnumWithDuplicateName(0), err
//   }
//
//   var v EnumWithDuplicateName
//   if err := v.FromWire(x); err != nil {
//     return EnumWithDuplicateName(0), err
//   }
//   return v, nil
func (v *EnumWithDuplicateName) FromWire(w wire.Value) error {
	*v = (EnumWithDuplicateName)(w.GetI32())
	return nil
}

// String returns a readable string representation of EnumWithDuplicateName.
func (v EnumWithDuplicateName) String() string {
	w := int32(v)
	switch w {
	case 0:
		return "A"
	case 1:
		return "B"
	case 2:
		return "C"
	case 3:
		return "P"
	case 4:
		return "Q"
	case 5:
		return "R"
	case 6:
		return "X"
	case 7:
		return "Y"
	case 8:
		return "Z"
	}
	return fmt.Sprintf("EnumWithDuplicateName(%d)", w)
}

// Equals returns true if this EnumWithDuplicateName value matches the provided
// value.
func (v EnumWithDuplicateName) Equals(rhs EnumWithDuplicateName) bool {
	return v == rhs
}

// MarshalJSON serializes EnumWithDuplicateName into JSON.
//
// If the enum value is recognized, its name is returned. Otherwise,
// its integer value is returned.
//
// This implements json.Marshaler.
func (v EnumWithDuplicateName) MarshalJSON() ([]byte, error) {
	switch int32(v) {
	case 0:
		return ([]byte)("\"A\""), nil
	case 1:
		return ([]byte)("\"B\""), nil
	case 2:
		return ([]byte)("\"C\""), nil
	case 3:
		return ([]byte)("\"P\""), nil
	case 4:
		return ([]byte)("\"Q\""), nil
	case 5:
		return ([]byte)("\"R\""), nil
	case 6:
		return ([]byte)("\"X\""), nil
	case 7:
		return ([]byte)("\"Y\""), nil
	case 8:
		return ([]byte)("\"Z\""), nil
	}
	return ([]byte)(strconv.FormatInt(int64(v), 10)), nil
}

// UnmarshalJSON attempts to decode EnumWithDuplicateName from its JSON
// representation.
//
// This implementation supports both, numeric and string inputs. If a
// string is provided, it must be a known enum name.
//
// This implements json.Unmarshaler.
func (v *EnumWithDuplicateName) UnmarshalJSON(text []byte) error {
	d := json.NewDecoder(bytes.NewReader(text))
	d.UseNumber()
	t, err := d.Token()
	if err != nil {
		return err
	}

	switch w := t.(type) {
	case json.Number:
		x, err := w.Int64()
		if err != nil {
			return err
		}
		if x > math.MaxInt32 {
			return fmt.Errorf("enum overflow from JSON %q for %q", text, "EnumWithDuplicateName")
		}
		if x < math.MinInt32 {
			return fmt.Errorf("enum underflow from JSON %q for %q", text, "EnumWithDuplicateName")
		}
		*v = (EnumWithDuplicateName)(x)
		return nil
	case string:
		return v.UnmarshalText([]byte(w))
	default:
		return fmt.Errorf("invalid JSON value %q (%T) to unmarshal into %q", t, t, "EnumWithDuplicateName")
	}
}

type EnumWithDuplicateValues int32

const (
	EnumWithDuplicateValuesP EnumWithDuplicateValues = 0
	EnumWithDuplicateValuesQ EnumWithDuplicateValues = -1
	EnumWithDuplicateValuesR EnumWithDuplicateValues = 0
)

// EnumWithDuplicateValues_Values returns all recognized values of EnumWithDuplicateValues.
func EnumWithDuplicateValues_Values() []EnumWithDuplicateValues {
	return []EnumWithDuplicateValues{
		EnumWithDuplicateValuesP,
		EnumWithDuplicateValuesQ,
		EnumWithDuplicateValuesR,
	}
}

// UnmarshalText tries to decode EnumWithDuplicateValues from a byte slice
// containing its name.
//
//   var v EnumWithDuplicateValues
//   err := v.UnmarshalText([]byte("P"))
func (v *EnumWithDuplicateValues) UnmarshalText(value []byte) error {
	switch string(value) {
	case "P":
		*v = EnumWithDuplicateValuesP
		return nil
	case "Q":
		*v = EnumWithDuplicateValuesQ
		return nil
	case "R":
		*v = EnumWithDuplicateValuesR
		return nil
	default:
		return fmt.Errorf("unknown enum value %q for %q", value, "EnumWithDuplicateValues")
	}
}

// MarshalText encodes EnumWithDuplicateValues to text.
//
// If the enum value is recognized, its name is returned. Otherwise,
// its integer value is returned.
//
// This implements the TextMarshaler interface.
func (v EnumWithDuplicateValues) MarshalText() ([]byte, error) {
	switch int32(v) {
	case 0:
		return []byte("P"), nil
	case -1:
		return []byte("Q"), nil
	}
	return []byte(strconv.FormatInt(int64(v), 10)), nil
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of EnumWithDuplicateValues.
func (v EnumWithDuplicateValues) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddInt32("value", int32(v))
	switch int32(v) {
	case 0:
		enc.AddString("name", "P")
	case -1:
		enc.AddString("name", "Q")
	}
	return nil
}

// Ptr returns a pointer to this enum value.
func (v EnumWithDuplicateValues) Ptr() *EnumWithDuplicateValues {
	return &v
}

// ToWire translates EnumWithDuplicateValues into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// Enums are represented as 32-bit integers over the wire.
func (v EnumWithDuplicateValues) ToWire() (wire.Value, error) {
	return wire.NewValueI32(int32(v)), nil
}

// FromWire deserializes EnumWithDuplicateValues from its Thrift-level
// representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TI32)
//   if err != nil {
//     return EnumWithDuplicateValues(0), err
//   }
//
//   var v EnumWithDuplicateValues
//   if err := v.FromWire(x); err != nil {
//     return EnumWithDuplicateValues(0), err
//   }
//   return v, nil
func (v *EnumWithDuplicateValues) FromWire(w wire.Value) error {
	*v = (EnumWithDuplicateValues)(w.GetI32())
	return nil
}

// String returns a readable string representation of EnumWithDuplicateValues.
func (v EnumWithDuplicateValues) String() string {
	w := int32(v)
	switch w {
	case 0:
		return "P"
	case -1:
		return "Q"
	}
	return fmt.Sprintf("EnumWithDuplicateValues(%d)", w)
}

// Equals returns true if this EnumWithDuplicateValues value matches the provided
// value.
func (v EnumWithDuplicateValues) Equals(rhs EnumWithDuplicateValues) bool {
	return v == rhs
}

// MarshalJSON serializes EnumWithDuplicateValues into JSON.
//
// If the enum value is recognized, its name is returned. Otherwise,
// its integer value is returned.
//
// This implements json.Marshaler.
func (v EnumWithDuplicateValues) MarshalJSON() ([]byte, error) {
	switch int32(v) {
	case 0:
		return ([]byte)("\"P\""), nil
	case -1:
		return ([]byte)("\"Q\""), nil
	}
	return ([]byte)(strconv.FormatInt(int64(v), 10)), nil
}

// UnmarshalJSON attempts to decode EnumWithDuplicateValues from its JSON
// representation.
//
// This implementation supports both, numeric and string inputs. If a
// string is provided, it must be a known enum name.
//
// This implements json.Unmarshaler.
func (v *EnumWithDuplicateValues) UnmarshalJSON(text []byte) error {
	d := json.NewDecoder(bytes.NewReader(text))
	d.UseNumber()
	t, err := d.Token()
	if err != nil {
		return err
	}

	switch w := t.(type) {
	case json.Number:
		x, err := w.Int64()
		if err != nil {
			return err
		}
		if x > math.MaxInt32 {
			return fmt.Errorf("enum overflow from JSON %q for %q", text, "EnumWithDuplicateValues")
		}
		if x < math.MinInt32 {
			return fmt.Errorf("enum underflow from JSON %q for %q", text, "EnumWithDuplicateValues")
		}
		*v = (EnumWithDuplicateValues)(x)
		return nil
	case string:
		return v.UnmarshalText([]byte(w))
	default:
		return fmt.Errorf("invalid JSON value %q (%T) to unmarshal into %q", t, t, "EnumWithDuplicateValues")
	}
}

type EnumWithValues int32

const (
	EnumWithValuesX EnumWithValues = 123
	EnumWithValuesY EnumWithValues = 456
	EnumWithValuesZ EnumWithValues = 789
)

// EnumWithValues_Values returns all recognized values of EnumWithValues.
func EnumWithValues_Values() []EnumWithValues {
	return []EnumWithValues{
		EnumWithValuesX,
		EnumWithValuesY,
		EnumWithValuesZ,
	}
}

// UnmarshalText tries to decode EnumWithValues from a byte slice
// containing its name.
//
//   var v EnumWithValues
//   err := v.UnmarshalText([]byte("X"))
func (v *EnumWithValues) UnmarshalText(value []byte) error {
	switch string(value) {
	case "X":
		*v = EnumWithValuesX
		return nil
	case "Y":
		*v = EnumWithValuesY
		return nil
	case "Z":
		*v = EnumWithValuesZ
		return nil
	default:
		return fmt.Errorf("unknown enum value %q for %q", value, "EnumWithValues")
	}
}

// MarshalText encodes EnumWithValues to text.
//
// If the enum value is recognized, its name is returned. Otherwise,
// its integer value is returned.
//
// This implements the TextMarshaler interface.
func (v EnumWithValues) MarshalText() ([]byte, error) {
	switch int32(v) {
	case 123:
		return []byte("X"), nil
	case 456:
		return []byte("Y"), nil
	case 789:
		return []byte("Z"), nil
	}
	return []byte(strconv.FormatInt(int64(v), 10)), nil
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of EnumWithValues.
func (v EnumWithValues) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddInt32("value", int32(v))
	switch int32(v) {
	case 123:
		enc.AddString("name", "X")
	case 456:
		enc.AddString("name", "Y")
	case 789:
		enc.AddString("name", "Z")
	}
	return nil
}

// Ptr returns a pointer to this enum value.
func (v EnumWithValues) Ptr() *EnumWithValues {
	return &v
}

// ToWire translates EnumWithValues into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// Enums are represented as 32-bit integers over the wire.
func (v EnumWithValues) ToWire() (wire.Value, error) {
	return wire.NewValueI32(int32(v)), nil
}

// FromWire deserializes EnumWithValues from its Thrift-level
// representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TI32)
//   if err != nil {
//     return EnumWithValues(0), err
//   }
//
//   var v EnumWithValues
//   if err := v.FromWire(x); err != nil {
//     return EnumWithValues(0), err
//   }
//   return v, nil
func (v *EnumWithValues) FromWire(w wire.Value) error {
	*v = (EnumWithValues)(w.GetI32())
	return nil
}

// String returns a readable string representation of EnumWithValues.
func (v EnumWithValues) String() string {
	w := int32(v)
	switch w {
	case 123:
		return "X"
	case 456:
		return "Y"
	case 789:
		return "Z"
	}
	return fmt.Sprintf("EnumWithValues(%d)", w)
}

// Equals returns true if this EnumWithValues value matches the provided
// value.
func (v EnumWithValues) Equals(rhs EnumWithValues) bool {
	return v == rhs
}

// MarshalJSON serializes EnumWithValues into JSON.
//
// If the enum value is recognized, its name is returned. Otherwise,
// its integer value is returned.
//
// This implements json.Marshaler.
func (v EnumWithValues) MarshalJSON() ([]byte, error) {
	switch int32(v) {
	case 123:
		return ([]byte)("\"X\""), nil
	case 456:
		return ([]byte)("\"Y\""), nil
	case 789:
		return ([]byte)("\"Z\""), nil
	}
	return ([]byte)(strconv.FormatInt(int64(v), 10)), nil
}

// UnmarshalJSON attempts to decode EnumWithValues from its JSON
// representation.
//
// This implementation supports both, numeric and string inputs. If a
// string is provided, it must be a known enum name.
//
// This implements json.Unmarshaler.
func (v *EnumWithValues) UnmarshalJSON(text []byte) error {
	d := json.NewDecoder(bytes.NewReader(text))
	d.UseNumber()
	t, err := d.Token()
	if err != nil {
		return err
	}

	switch w := t.(type) {
	case json.Number:
		x, err := w.Int64()
		if err != nil {
			return err
		}
		if x > math.MaxInt32 {
			return fmt.Errorf("enum overflow from JSON %q for %q", text, "EnumWithValues")
		}
		if x < math.MinInt32 {
			return fmt.Errorf("enum underflow from JSON %q for %q", text, "EnumWithValues")
		}
		*v = (EnumWithValues)(x)
		return nil
	case string:
		return v.UnmarshalText([]byte(w))
	default:
		return fmt.Errorf("invalid JSON value %q (%T) to unmarshal into %q", t, t, "EnumWithValues")
	}
}

// Kinds of records stored in the database.
type RecordType int32

const (
	// Name of the user.
	RecordTypeName RecordType = 0
	// Home address of the user.
	//
	// This record is always present.
	RecordTypeHomeAddress RecordType = 1
	// Home address of the user.
	//
	// This record may not be present.
	RecordTypeWorkAddress RecordType = 2
)

// RecordType_Values returns all recognized values of RecordType.
func RecordType_Values() []RecordType {
	return []RecordType{
		RecordTypeName,
		RecordTypeHomeAddress,
		RecordTypeWorkAddress,
	}
}

// UnmarshalText tries to decode RecordType from a byte slice
// containing its name.
//
//   var v RecordType
//   err := v.UnmarshalText([]byte("NAME"))
func (v *RecordType) UnmarshalText(value []byte) error {
	switch string(value) {
	case "NAME":
		*v = RecordTypeName
		return nil
	case "HOME_ADDRESS":
		*v = RecordTypeHomeAddress
		return nil
	case "WORK_ADDRESS":
		*v = RecordTypeWorkAddress
		return nil
	default:
		return fmt.Errorf("unknown enum value %q for %q", value, "RecordType")
	}
}

// MarshalText encodes RecordType to text.
//
// If the enum value is recognized, its name is returned. Otherwise,
// its integer value is returned.
//
// This implements the TextMarshaler interface.
func (v RecordType) MarshalText() ([]byte, error) {
	switch int32(v) {
	case 0:
		return []byte("NAME"), nil
	case 1:
		return []byte("HOME_ADDRESS"), nil
	case 2:
		return []byte("WORK_ADDRESS"), nil
	}
	return []byte(strconv.FormatInt(int64(v), 10)), nil
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of RecordType.
func (v RecordType) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddInt32("value", int32(v))
	switch int32(v) {
	case 0:
		enc.AddString("name", "NAME")
	case 1:
		enc.AddString("name", "HOME_ADDRESS")
	case 2:
		enc.AddString("name", "WORK_ADDRESS")
	}
	return nil
}

// Ptr returns a pointer to this enum value.
func (v RecordType) Ptr() *RecordType {
	return &v
}

// ToWire translates RecordType into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// Enums are represented as 32-bit integers over the wire.
func (v RecordType) ToWire() (wire.Value, error) {
	return wire.NewValueI32(int32(v)), nil
}

// FromWire deserializes RecordType from its Thrift-level
// representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TI32)
//   if err != nil {
//     return RecordType(0), err
//   }
//
//   var v RecordType
//   if err := v.FromWire(x); err != nil {
//     return RecordType(0), err
//   }
//   return v, nil
func (v *RecordType) FromWire(w wire.Value) error {
	*v = (RecordType)(w.GetI32())
	return nil
}

// String returns a readable string representation of RecordType.
func (v RecordType) String() string {
	w := int32(v)
	switch w {
	case 0:
		return "NAME"
	case 1:
		return "HOME_ADDRESS"
	case 2:
		return "WORK_ADDRESS"
	}
	return fmt.Sprintf("RecordType(%d)", w)
}

// Equals returns true if this RecordType value matches the provided
// value.
func (v RecordType) Equals(rhs RecordType) bool {
	return v == rhs
}

// MarshalJSON serializes RecordType into JSON.
//
// If the enum value is recognized, its name is returned. Otherwise,
// its integer value is returned.
//
// This implements json.Marshaler.
func (v RecordType) MarshalJSON() ([]byte, error) {
	switch int32(v) {
	case 0:
		return ([]byte)("\"NAME\""), nil
	case 1:
		return ([]byte)("\"HOME_ADDRESS\""), nil
	case 2:
		return ([]byte)("\"WORK_ADDRESS\""), nil
	}
	return ([]byte)(strconv.FormatInt(int64(v), 10)), nil
}

// UnmarshalJSON attempts to decode RecordType from its JSON
// representation.
//
// This implementation supports both, numeric and string inputs. If a
// string is provided, it must be a known enum name.
//
// This implements json.Unmarshaler.
func (v *RecordType) UnmarshalJSON(text []byte) error {
	d := json.NewDecoder(bytes.NewReader(text))
	d.UseNumber()
	t, err := d.Token()
	if err != nil {
		return err
	}

	switch w := t.(type) {
	case json.Number:
		x, err := w.Int64()
		if err != nil {
			return err
		}
		if x > math.MaxInt32 {
			return fmt.Errorf("enum overflow from JSON %q for %q", text, "RecordType")
		}
		if x < math.MinInt32 {
			return fmt.Errorf("enum underflow from JSON %q for %q", text, "RecordType")
		}
		*v = (RecordType)(x)
		return nil
	case string:
		return v.UnmarshalText([]byte(w))
	default:
		return fmt.Errorf("invalid JSON value %q (%T) to unmarshal into %q", t, t, "RecordType")
	}
}

type RecordTypeValues int32

const (
	RecordTypeValuesFoo RecordTypeValues = 0
	RecordTypeValuesBar RecordTypeValues = 1
)

// RecordTypeValues_Values returns all recognized values of RecordTypeValues.
func RecordTypeValues_Values() []RecordTypeValues {
	return []RecordTypeValues{
		RecordTypeValuesFoo,
		RecordTypeValuesBar,
	}
}

// UnmarshalText tries to decode RecordTypeValues from a byte slice
// containing its name.
//
//   var v RecordTypeValues
//   err := v.UnmarshalText([]byte("FOO"))
func (v *RecordTypeValues) UnmarshalText(value []byte) error {
	switch string(value) {
	case "FOO":
		*v = RecordTypeValuesFoo
		return nil
	case "BAR":
		*v = RecordTypeValuesBar
		return nil
	default:
		return fmt.Errorf("unknown enum value %q for %q", value, "RecordTypeValues")
	}
}

// MarshalText encodes RecordTypeValues to text.
//
// If the enum value is recognized, its name is returned. Otherwise,
// its integer value is returned.
//
// This implements the TextMarshaler interface.
func (v RecordTypeValues) MarshalText() ([]byte, error) {
	switch int32(v) {
	case 0:
		return []byte("FOO"), nil
	case 1:
		return []byte("BAR"), nil
	}
	return []byte(strconv.FormatInt(int64(v), 10)), nil
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of RecordTypeValues.
func (v RecordTypeValues) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddInt32("value", int32(v))
	switch int32(v) {
	case 0:
		enc.AddString("name", "FOO")
	case 1:
		enc.AddString("name", "BAR")
	}
	return nil
}

// Ptr returns a pointer to this enum value.
func (v RecordTypeValues) Ptr() *RecordTypeValues {
	return &v
}

// ToWire translates RecordTypeValues into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// Enums are represented as 32-bit integers over the wire.
func (v RecordTypeValues) ToWire() (wire.Value, error) {
	return wire.NewValueI32(int32(v)), nil
}

// FromWire deserializes RecordTypeValues from its Thrift-level
// representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TI32)
//   if err != nil {
//     return RecordTypeValues(0), err
//   }
//
//   var v RecordTypeValues
//   if err := v.FromWire(x); err != nil {
//     return RecordTypeValues(0), err
//   }
//   return v, nil
func (v *RecordTypeValues) FromWire(w wire.Value) error {
	*v = (RecordTypeValues)(w.GetI32())
	return nil
}

// String returns a readable string representation of RecordTypeValues.
func (v RecordTypeValues) String() string {
	w := int32(v)
	switch w {
	case 0:
		return "FOO"
	case 1:
		return "BAR"
	}
	return fmt.Sprintf("RecordTypeValues(%d)", w)
}

// Equals returns true if this RecordTypeValues value matches the provided
// value.
func (v RecordTypeValues) Equals(rhs RecordTypeValues) bool {
	return v == rhs
}

// MarshalJSON serializes RecordTypeValues into JSON.
//
// If the enum value is recognized, its name is returned. Otherwise,
// its integer value is returned.
//
// This implements json.Marshaler.
func (v RecordTypeValues) MarshalJSON() ([]byte, error) {
	switch int32(v) {
	case 0:
		return ([]byte)("\"FOO\""), nil
	case 1:
		return ([]byte)("\"BAR\""), nil
	}
	return ([]byte)(strconv.FormatInt(int64(v), 10)), nil
}

// UnmarshalJSON attempts to decode RecordTypeValues from its JSON
// representation.
//
// This implementation supports both, numeric and string inputs. If a
// string is provided, it must be a known enum name.
//
// This implements json.Unmarshaler.
func (v *RecordTypeValues) UnmarshalJSON(text []byte) error {
	d := json.NewDecoder(bytes.NewReader(text))
	d.UseNumber()
	t, err := d.Token()
	if err != nil {
		return err
	}

	switch w := t.(type) {
	case json.Number:
		x, err := w.Int64()
		if err != nil {
			return err
		}
		if x > math.MaxInt32 {
			return fmt.Errorf("enum overflow from JSON %q for %q", text, "RecordTypeValues")
		}
		if x < math.MinInt32 {
			return fmt.Errorf("enum underflow from JSON %q for %q", text, "RecordTypeValues")
		}
		*v = (RecordTypeValues)(x)
		return nil
	case string:
		return v.UnmarshalText([]byte(w))
	default:
		return fmt.Errorf("invalid JSON value %q (%T) to unmarshal into %q", t, t, "RecordTypeValues")
	}
}

type StructWithOptionalEnum struct {
	E *EnumDefault `json:"e,omitempty"`
}

// ToWire translates a StructWithOptionalEnum struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *StructWithOptionalEnum) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.E != nil {
		w, err = v.E.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _EnumDefault_Read(w wire.Value) (EnumDefault, error) {
	var v EnumDefault
	err := v.FromWire(w)
	return v, err
}

// FromWire deserializes a StructWithOptionalEnum struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a StructWithOptionalEnum struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v StructWithOptionalEnum
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *StructWithOptionalEnum) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TI32 {
				var x EnumDefault
				x, err = _EnumDefault_Read(field.Value)
				v.E = &x
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// String returns a readable string representation of a StructWithOptionalEnum
// struct.
func (v *StructWithOptionalEnum) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.E != nil {
		fields[i] = fmt.Sprintf("E: %v", *(v.E))
		i++
	}

	return fmt.Sprintf("StructWithOptionalEnum{%v}", strings.Join(fields[:i], ", "))
}

func _EnumDefault_EqualsPtr(lhs, rhs *EnumDefault) bool {
	if lhs != nil && rhs != nil {

		x := *lhs
		y := *rhs
		return x.Equals(y)
	}
	return lhs == nil && rhs == nil
}

// Equals returns true if all the fields of this StructWithOptionalEnum match the
// provided StructWithOptionalEnum.
//
// This function performs a deep comparison.
func (v *StructWithOptionalEnum) Equals(rhs *StructWithOptionalEnum) bool {
	if !_EnumDefault_EqualsPtr(v.E, rhs.E) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of StructWithOptionalEnum.
func (v *StructWithOptionalEnum) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	if v.E != nil {
		if err := enc.AddObject("e", *v.E); err != nil {
			return err
		}
	}
	return nil
}

// GetE returns the value of E if it is set or its
// zero value if it is unset.
func (v *StructWithOptionalEnum) GetE() (o EnumDefault) {
	if v.E != nil {
		return *v.E
	}

	return
}

type LowerCaseEnum int32

const (
	LowerCaseEnumContaining LowerCaseEnum = 0
	LowerCaseEnumLowerCase  LowerCaseEnum = 1
	LowerCaseEnumItems      LowerCaseEnum = 2
)

// LowerCaseEnum_Values returns all recognized values of LowerCaseEnum.
func LowerCaseEnum_Values() []LowerCaseEnum {
	return []LowerCaseEnum{
		LowerCaseEnumContaining,
		LowerCaseEnumLowerCase,
		LowerCaseEnumItems,
	}
}

// UnmarshalText tries to decode LowerCaseEnum from a byte slice
// containing its name.
//
//   var v LowerCaseEnum
//   err := v.UnmarshalText([]byte("containing"))
func (v *LowerCaseEnum) UnmarshalText(value []byte) error {
	switch string(value) {
	case "containing":
		*v = LowerCaseEnumContaining
		return nil
	case "lower_case":
		*v = LowerCaseEnumLowerCase
		return nil
	case "items":
		*v = LowerCaseEnumItems
		return nil
	default:
		return fmt.Errorf("unknown enum value %q for %q", value, "LowerCaseEnum")
	}
}

// MarshalText encodes LowerCaseEnum to text.
//
// If the enum value is recognized, its name is returned. Otherwise,
// its integer value is returned.
//
// This implements the TextMarshaler interface.
func (v LowerCaseEnum) MarshalText() ([]byte, error) {
	switch int32(v) {
	case 0:
		return []byte("containing"), nil
	case 1:
		return []byte("lower_case"), nil
	case 2:
		return []byte("items"), nil
	}
	return []byte(strconv.FormatInt(int64(v), 10)), nil
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of LowerCaseEnum.
func (v LowerCaseEnum) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddInt32("value", int32(v))
	switch int32(v) {
	case 0:
		enc.AddString("name", "containing")
	case 1:
		enc.AddString("name", "lower_case")
	case 2:
		enc.AddString("name", "items")
	}
	return nil
}

// Ptr returns a pointer to this enum value.
func (v LowerCaseEnum) Ptr() *LowerCaseEnum {
	return &v
}

// ToWire translates LowerCaseEnum into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// Enums are represented as 32-bit integers over the wire.
func (v LowerCaseEnum) ToWire() (wire.Value, error) {
	return wire.NewValueI32(int32(v)), nil
}

// FromWire deserializes LowerCaseEnum from its Thrift-level
// representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TI32)
//   if err != nil {
//     return LowerCaseEnum(0), err
//   }
//
//   var v LowerCaseEnum
//   if err := v.FromWire(x); err != nil {
//     return LowerCaseEnum(0), err
//   }
//   return v, nil
func (v *LowerCaseEnum) FromWire(w wire.Value) error {
	*v = (LowerCaseEnum)(w.GetI32())
	return nil
}

// String returns a readable string representation of LowerCaseEnum.
func (v LowerCaseEnum) String() string {
	w := int32(v)
	switch w {
	case 0:
		return "containing"
	case 1:
		return "lower_case"
	case 2:
		return "items"
	}
	return fmt.Sprintf("LowerCaseEnum(%d)", w)
}

// Equals returns true if this LowerCaseEnum value matches the provided
// value.
func (v LowerCaseEnum) Equals(rhs LowerCaseEnum) bool {
	return v == rhs
}

// MarshalJSON serializes LowerCaseEnum into JSON.
//
// If the enum value is recognized, its name is returned. Otherwise,
// its integer value is returned.
//
// This implements json.Marshaler.
func (v LowerCaseEnum) MarshalJSON() ([]byte, error) {
	switch int32(v) {
	case 0:
		return ([]byte)("\"containing\""), nil
	case 1:
		return ([]byte)("\"lower_case\""), nil
	case 2:
		return ([]byte)("\"items\""), nil
	}
	return ([]byte)(strconv.FormatInt(int64(v), 10)), nil
}

// UnmarshalJSON attempts to decode LowerCaseEnum from its JSON
// representation.
//
// This implementation supports both, numeric and string inputs. If a
// string is provided, it must be a known enum name.
//
// This implements json.Unmarshaler.
func (v *LowerCaseEnum) UnmarshalJSON(text []byte) error {
	d := json.NewDecoder(bytes.NewReader(text))
	d.UseNumber()
	t, err := d.Token()
	if err != nil {
		return err
	}

	switch w := t.(type) {
	case json.Number:
		x, err := w.Int64()
		if err != nil {
			return err
		}
		if x > math.MaxInt32 {
			return fmt.Errorf("enum overflow from JSON %q for %q", text, "LowerCaseEnum")
		}
		if x < math.MinInt32 {
			return fmt.Errorf("enum underflow from JSON %q for %q", text, "LowerCaseEnum")
		}
		*v = (LowerCaseEnum)(x)
		return nil
	case string:
		return v.UnmarshalText([]byte(w))
	default:
		return fmt.Errorf("invalid JSON value %q (%T) to unmarshal into %q", t, t, "LowerCaseEnum")
	}
}
