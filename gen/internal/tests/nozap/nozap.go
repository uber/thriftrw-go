// Code generated by thriftrw v1.32.0. DO NOT EDIT.
// @generated

package nozap

import (
	bytes "bytes"
	json "encoding/json"
	errors "errors"
	fmt "fmt"
	stream "go.uber.org/thriftrw/protocol/stream"
	thriftreflect "go.uber.org/thriftrw/thriftreflect"
	wire "go.uber.org/thriftrw/wire"
	math "math"
	strconv "strconv"
	strings "strings"
)

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
//	var v EnumDefault
//	err := v.UnmarshalText([]byte("Foo"))
func (v *EnumDefault) UnmarshalText(value []byte) error {
	switch s := string(value); s {
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
		val, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			return fmt.Errorf("unknown enum value %q for %q: %v", s, "EnumDefault", err)
		}
		*v = EnumDefault(val)
		return nil
	}
}

// MarshalText encodes EnumDefault to text.
//
// If the enum value is recognized, its name is returned.
// Otherwise, its integer value is returned.
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

// Ptr returns a pointer to this enum value.
func (v EnumDefault) Ptr() *EnumDefault {
	return &v
}

// Encode encodes EnumDefault directly to bytes.
//
//	sWriter := BinaryStreamer.Writer(writer)
//
//	var v EnumDefault
//	return v.Encode(sWriter)
func (v EnumDefault) Encode(sw stream.Writer) error {
	return sw.WriteInt32(int32(v))
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
//	x, err := binaryProtocol.Decode(reader, wire.TI32)
//	if err != nil {
//	    return EnumDefault(0), err
//	}
//
//	var v EnumDefault
//	if err := v.FromWire(x); err != nil {
//	    return EnumDefault(0), err
//	}
//	return v, nil
func (v *EnumDefault) FromWire(w wire.Value) error {
	*v = (EnumDefault)(w.GetI32())
	return nil
}

// Decode reads off the encoded EnumDefault directly off of the wire.
//
//	sReader := BinaryStreamer.Reader(reader)
//
//	var v EnumDefault
//	if err := v.Decode(sReader); err != nil {
//	    return EnumDefault(0), err
//	}
//	return v, nil
func (v *EnumDefault) Decode(sr stream.Reader) error {
	i, err := sr.ReadInt32()
	if err != nil {
		return err
	}
	*v = (EnumDefault)(i)
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
// If the enum value is recognized, its name is returned.
// Otherwise, its integer value is returned.
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

type PrimitiveRequiredStruct struct {
	BoolField          bool               `json:"boolField,required"`
	ByteField          int8               `json:"byteField,required"`
	Int16Field         int16              `json:"int16Field,required"`
	Int32Field         int32              `json:"int32Field,required"`
	Int64Field         int64              `json:"int64Field,required"`
	DoubleField        float64            `json:"doubleField,required"`
	StringField        string             `json:"stringField,required"`
	BinaryField        []byte             `json:"binaryField,required"`
	ListOfStrings      []string           `json:"listOfStrings,required"`
	SetOfInts          map[int32]struct{} `json:"setOfInts,required"`
	MapOfIntsToDoubles map[int64]float64  `json:"mapOfIntsToDoubles,required"`
}

type _List_String_ValueList []string

func (v _List_String_ValueList) ForEach(f func(wire.Value) error) error {
	for _, x := range v {
		w, err := wire.NewValueString(x), error(nil)
		if err != nil {
			return err
		}
		err = f(w)
		if err != nil {
			return err
		}
	}
	return nil
}

func (v _List_String_ValueList) Size() int {
	return len(v)
}

func (_List_String_ValueList) ValueType() wire.Type {
	return wire.TBinary
}

func (_List_String_ValueList) Close() {}

type _Set_I32_mapType_ValueList map[int32]struct{}

func (v _Set_I32_mapType_ValueList) ForEach(f func(wire.Value) error) error {
	for x := range v {
		w, err := wire.NewValueI32(x), error(nil)
		if err != nil {
			return err
		}

		if err := f(w); err != nil {
			return err
		}
	}
	return nil
}

func (v _Set_I32_mapType_ValueList) Size() int {
	return len(v)
}

func (_Set_I32_mapType_ValueList) ValueType() wire.Type {
	return wire.TI32
}

func (_Set_I32_mapType_ValueList) Close() {}

type _Map_I64_Double_MapItemList map[int64]float64

func (m _Map_I64_Double_MapItemList) ForEach(f func(wire.MapItem) error) error {
	for k, v := range m {
		kw, err := wire.NewValueI64(k), error(nil)
		if err != nil {
			return err
		}

		vw, err := wire.NewValueDouble(v), error(nil)
		if err != nil {
			return err
		}
		err = f(wire.MapItem{Key: kw, Value: vw})
		if err != nil {
			return err
		}
	}
	return nil
}

func (m _Map_I64_Double_MapItemList) Size() int {
	return len(m)
}

func (_Map_I64_Double_MapItemList) KeyType() wire.Type {
	return wire.TI64
}

func (_Map_I64_Double_MapItemList) ValueType() wire.Type {
	return wire.TDouble
}

func (_Map_I64_Double_MapItemList) Close() {}

// ToWire translates a PrimitiveRequiredStruct struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//	x, err := v.ToWire()
//	if err != nil {
//		return err
//	}
//
//	if err := binaryProtocol.Encode(x, writer); err != nil {
//		return err
//	}
func (v *PrimitiveRequiredStruct) ToWire() (wire.Value, error) {
	var (
		fields [11]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	w, err = wire.NewValueBool(v.BoolField), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++

	w, err = wire.NewValueI8(v.ByteField), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 2, Value: w}
	i++

	w, err = wire.NewValueI16(v.Int16Field), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 3, Value: w}
	i++

	w, err = wire.NewValueI32(v.Int32Field), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 4, Value: w}
	i++

	w, err = wire.NewValueI64(v.Int64Field), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 5, Value: w}
	i++

	w, err = wire.NewValueDouble(v.DoubleField), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 6, Value: w}
	i++

	w, err = wire.NewValueString(v.StringField), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 7, Value: w}
	i++
	if v.BinaryField == nil {
		return w, errors.New("field BinaryField of PrimitiveRequiredStruct is required")
	}
	w, err = wire.NewValueBinary(v.BinaryField), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 8, Value: w}
	i++

	w, err = wire.NewValueList(_List_String_ValueList(v.ListOfStrings)), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 9, Value: w}
	i++
	if v.SetOfInts == nil {
		return w, errors.New("field SetOfInts of PrimitiveRequiredStruct is required")
	}
	w, err = wire.NewValueSet(_Set_I32_mapType_ValueList(v.SetOfInts)), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 10, Value: w}
	i++
	if v.MapOfIntsToDoubles == nil {
		return w, errors.New("field MapOfIntsToDoubles of PrimitiveRequiredStruct is required")
	}
	w, err = wire.NewValueMap(_Map_I64_Double_MapItemList(v.MapOfIntsToDoubles)), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 11, Value: w}
	i++

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _List_String_Read(l wire.ValueList) ([]string, error) {
	if l.ValueType() != wire.TBinary {
		return nil, nil
	}

	o := make([]string, 0, l.Size())
	err := l.ForEach(func(x wire.Value) error {
		i, err := x.GetString(), error(nil)
		if err != nil {
			return err
		}
		o = append(o, i)
		return nil
	})
	l.Close()
	return o, err
}

func _Set_I32_mapType_Read(s wire.ValueList) (map[int32]struct{}, error) {
	if s.ValueType() != wire.TI32 {
		return nil, nil
	}

	o := make(map[int32]struct{}, s.Size())
	err := s.ForEach(func(x wire.Value) error {
		i, err := x.GetI32(), error(nil)
		if err != nil {
			return err
		}

		o[i] = struct{}{}
		return nil
	})
	s.Close()
	return o, err
}

func _Map_I64_Double_Read(m wire.MapItemList) (map[int64]float64, error) {
	if m.KeyType() != wire.TI64 {
		return nil, nil
	}

	if m.ValueType() != wire.TDouble {
		return nil, nil
	}

	o := make(map[int64]float64, m.Size())
	err := m.ForEach(func(x wire.MapItem) error {
		k, err := x.Key.GetI64(), error(nil)
		if err != nil {
			return err
		}

		v, err := x.Value.GetDouble(), error(nil)
		if err != nil {
			return err
		}

		o[k] = v
		return nil
	})
	m.Close()
	return o, err
}

// FromWire deserializes a PrimitiveRequiredStruct struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a PrimitiveRequiredStruct struct
// from the provided intermediate representation.
//
//	x, err := binaryProtocol.Decode(reader, wire.TStruct)
//	if err != nil {
//		return nil, err
//	}
//
//	var v PrimitiveRequiredStruct
//	if err := v.FromWire(x); err != nil {
//		return nil, err
//	}
//	return &v, nil
func (v *PrimitiveRequiredStruct) FromWire(w wire.Value) error {
	var err error

	boolFieldIsSet := false
	byteFieldIsSet := false
	int16FieldIsSet := false
	int32FieldIsSet := false
	int64FieldIsSet := false
	doubleFieldIsSet := false
	stringFieldIsSet := false
	binaryFieldIsSet := false
	listOfStringsIsSet := false
	setOfIntsIsSet := false
	mapOfIntsToDoublesIsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBool {
				v.BoolField, err = field.Value.GetBool(), error(nil)
				if err != nil {
					return err
				}
				boolFieldIsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TI8 {
				v.ByteField, err = field.Value.GetI8(), error(nil)
				if err != nil {
					return err
				}
				byteFieldIsSet = true
			}
		case 3:
			if field.Value.Type() == wire.TI16 {
				v.Int16Field, err = field.Value.GetI16(), error(nil)
				if err != nil {
					return err
				}
				int16FieldIsSet = true
			}
		case 4:
			if field.Value.Type() == wire.TI32 {
				v.Int32Field, err = field.Value.GetI32(), error(nil)
				if err != nil {
					return err
				}
				int32FieldIsSet = true
			}
		case 5:
			if field.Value.Type() == wire.TI64 {
				v.Int64Field, err = field.Value.GetI64(), error(nil)
				if err != nil {
					return err
				}
				int64FieldIsSet = true
			}
		case 6:
			if field.Value.Type() == wire.TDouble {
				v.DoubleField, err = field.Value.GetDouble(), error(nil)
				if err != nil {
					return err
				}
				doubleFieldIsSet = true
			}
		case 7:
			if field.Value.Type() == wire.TBinary {
				v.StringField, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				stringFieldIsSet = true
			}
		case 8:
			if field.Value.Type() == wire.TBinary {
				v.BinaryField, err = field.Value.GetBinary(), error(nil)
				if err != nil {
					return err
				}
				binaryFieldIsSet = true
			}
		case 9:
			if field.Value.Type() == wire.TList {
				v.ListOfStrings, err = _List_String_Read(field.Value.GetList())
				if err != nil {
					return err
				}
				listOfStringsIsSet = true
			}
		case 10:
			if field.Value.Type() == wire.TSet {
				v.SetOfInts, err = _Set_I32_mapType_Read(field.Value.GetSet())
				if err != nil {
					return err
				}
				setOfIntsIsSet = true
			}
		case 11:
			if field.Value.Type() == wire.TMap {
				v.MapOfIntsToDoubles, err = _Map_I64_Double_Read(field.Value.GetMap())
				if err != nil {
					return err
				}
				mapOfIntsToDoublesIsSet = true
			}
		}
	}

	if !boolFieldIsSet {
		return errors.New("field BoolField of PrimitiveRequiredStruct is required")
	}

	if !byteFieldIsSet {
		return errors.New("field ByteField of PrimitiveRequiredStruct is required")
	}

	if !int16FieldIsSet {
		return errors.New("field Int16Field of PrimitiveRequiredStruct is required")
	}

	if !int32FieldIsSet {
		return errors.New("field Int32Field of PrimitiveRequiredStruct is required")
	}

	if !int64FieldIsSet {
		return errors.New("field Int64Field of PrimitiveRequiredStruct is required")
	}

	if !doubleFieldIsSet {
		return errors.New("field DoubleField of PrimitiveRequiredStruct is required")
	}

	if !stringFieldIsSet {
		return errors.New("field StringField of PrimitiveRequiredStruct is required")
	}

	if !binaryFieldIsSet {
		return errors.New("field BinaryField of PrimitiveRequiredStruct is required")
	}

	if !listOfStringsIsSet {
		return errors.New("field ListOfStrings of PrimitiveRequiredStruct is required")
	}

	if !setOfIntsIsSet {
		return errors.New("field SetOfInts of PrimitiveRequiredStruct is required")
	}

	if !mapOfIntsToDoublesIsSet {
		return errors.New("field MapOfIntsToDoubles of PrimitiveRequiredStruct is required")
	}

	return nil
}

func _List_String_Encode(val []string, sw stream.Writer) error {

	lh := stream.ListHeader{
		Type:   wire.TBinary,
		Length: len(val),
	}
	if err := sw.WriteListBegin(lh); err != nil {
		return err
	}

	for _, v := range val {
		if err := sw.WriteString(v); err != nil {
			return err
		}
	}
	return sw.WriteListEnd()
}

func _Set_I32_mapType_Encode(val map[int32]struct{}, sw stream.Writer) error {

	sh := stream.SetHeader{
		Type:   wire.TI32,
		Length: len(val),
	}

	if err := sw.WriteSetBegin(sh); err != nil {
		return err
	}

	for v, _ := range val {

		if err := sw.WriteInt32(v); err != nil {
			return err
		}
	}
	return sw.WriteSetEnd()
}

func _Map_I64_Double_Encode(val map[int64]float64, sw stream.Writer) error {

	mh := stream.MapHeader{
		KeyType:   wire.TI64,
		ValueType: wire.TDouble,
		Length:    len(val),
	}
	if err := sw.WriteMapBegin(mh); err != nil {
		return err
	}

	for k, v := range val {
		if err := sw.WriteInt64(k); err != nil {
			return err
		}
		if err := sw.WriteDouble(v); err != nil {
			return err
		}
	}

	return sw.WriteMapEnd()
}

// Encode serializes a PrimitiveRequiredStruct struct directly into bytes, without going
// through an intermediary type.
//
// An error is returned if a PrimitiveRequiredStruct struct could not be encoded.
func (v *PrimitiveRequiredStruct) Encode(sw stream.Writer) error {
	if err := sw.WriteStructBegin(); err != nil {
		return err
	}

	if err := sw.WriteFieldBegin(stream.FieldHeader{ID: 1, Type: wire.TBool}); err != nil {
		return err
	}
	if err := sw.WriteBool(v.BoolField); err != nil {
		return err
	}
	if err := sw.WriteFieldEnd(); err != nil {
		return err
	}

	if err := sw.WriteFieldBegin(stream.FieldHeader{ID: 2, Type: wire.TI8}); err != nil {
		return err
	}
	if err := sw.WriteInt8(v.ByteField); err != nil {
		return err
	}
	if err := sw.WriteFieldEnd(); err != nil {
		return err
	}

	if err := sw.WriteFieldBegin(stream.FieldHeader{ID: 3, Type: wire.TI16}); err != nil {
		return err
	}
	if err := sw.WriteInt16(v.Int16Field); err != nil {
		return err
	}
	if err := sw.WriteFieldEnd(); err != nil {
		return err
	}

	if err := sw.WriteFieldBegin(stream.FieldHeader{ID: 4, Type: wire.TI32}); err != nil {
		return err
	}
	if err := sw.WriteInt32(v.Int32Field); err != nil {
		return err
	}
	if err := sw.WriteFieldEnd(); err != nil {
		return err
	}

	if err := sw.WriteFieldBegin(stream.FieldHeader{ID: 5, Type: wire.TI64}); err != nil {
		return err
	}
	if err := sw.WriteInt64(v.Int64Field); err != nil {
		return err
	}
	if err := sw.WriteFieldEnd(); err != nil {
		return err
	}

	if err := sw.WriteFieldBegin(stream.FieldHeader{ID: 6, Type: wire.TDouble}); err != nil {
		return err
	}
	if err := sw.WriteDouble(v.DoubleField); err != nil {
		return err
	}
	if err := sw.WriteFieldEnd(); err != nil {
		return err
	}

	if err := sw.WriteFieldBegin(stream.FieldHeader{ID: 7, Type: wire.TBinary}); err != nil {
		return err
	}
	if err := sw.WriteString(v.StringField); err != nil {
		return err
	}
	if err := sw.WriteFieldEnd(); err != nil {
		return err
	}

	if v.BinaryField == nil {
		return errors.New("field BinaryField of PrimitiveRequiredStruct is required")
	}
	if err := sw.WriteFieldBegin(stream.FieldHeader{ID: 8, Type: wire.TBinary}); err != nil {
		return err
	}
	if err := sw.WriteBinary(v.BinaryField); err != nil {
		return err
	}
	if err := sw.WriteFieldEnd(); err != nil {
		return err
	}

	if err := sw.WriteFieldBegin(stream.FieldHeader{ID: 9, Type: wire.TList}); err != nil {
		return err
	}
	if err := _List_String_Encode(v.ListOfStrings, sw); err != nil {
		return err
	}
	if err := sw.WriteFieldEnd(); err != nil {
		return err
	}

	if v.SetOfInts == nil {
		return errors.New("field SetOfInts of PrimitiveRequiredStruct is required")
	}
	if err := sw.WriteFieldBegin(stream.FieldHeader{ID: 10, Type: wire.TSet}); err != nil {
		return err
	}
	if err := _Set_I32_mapType_Encode(v.SetOfInts, sw); err != nil {
		return err
	}
	if err := sw.WriteFieldEnd(); err != nil {
		return err
	}

	if v.MapOfIntsToDoubles == nil {
		return errors.New("field MapOfIntsToDoubles of PrimitiveRequiredStruct is required")
	}
	if err := sw.WriteFieldBegin(stream.FieldHeader{ID: 11, Type: wire.TMap}); err != nil {
		return err
	}
	if err := _Map_I64_Double_Encode(v.MapOfIntsToDoubles, sw); err != nil {
		return err
	}
	if err := sw.WriteFieldEnd(); err != nil {
		return err
	}

	return sw.WriteStructEnd()
}

func _List_String_Decode(sr stream.Reader) ([]string, error) {
	lh, err := sr.ReadListBegin()
	if err != nil {
		return nil, err
	}

	if lh.Type != wire.TBinary {
		for i := 0; i < lh.Length; i++ {
			if err := sr.Skip(lh.Type); err != nil {
				return nil, err
			}
		}
		return nil, sr.ReadListEnd()
	}

	o := make([]string, 0, lh.Length)
	for i := 0; i < lh.Length; i++ {
		v, err := sr.ReadString()
		if err != nil {
			return nil, err
		}
		o = append(o, v)
	}

	if err = sr.ReadListEnd(); err != nil {
		return nil, err
	}
	return o, err
}

func _Set_I32_mapType_Decode(sr stream.Reader) (map[int32]struct{}, error) {
	sh, err := sr.ReadSetBegin()
	if err != nil {
		return nil, err
	}

	if sh.Type != wire.TI32 {
		for i := 0; i < sh.Length; i++ {
			if err := sr.Skip(sh.Type); err != nil {
				return nil, err
			}
		}
		return nil, sr.ReadSetEnd()
	}

	o := make(map[int32]struct{}, sh.Length)
	for i := 0; i < sh.Length; i++ {
		v, err := sr.ReadInt32()
		if err != nil {
			return nil, err
		}

		o[v] = struct{}{}
	}

	if err = sr.ReadSetEnd(); err != nil {
		return nil, err
	}
	return o, err
}

func _Map_I64_Double_Decode(sr stream.Reader) (map[int64]float64, error) {
	mh, err := sr.ReadMapBegin()
	if err != nil {
		return nil, err
	}

	if mh.KeyType != wire.TI64 || mh.ValueType != wire.TDouble {
		for i := 0; i < mh.Length; i++ {
			if err := sr.Skip(mh.KeyType); err != nil {
				return nil, err
			}

			if err := sr.Skip(mh.ValueType); err != nil {
				return nil, err
			}
		}
		return nil, sr.ReadMapEnd()
	}

	o := make(map[int64]float64, mh.Length)
	for i := 0; i < mh.Length; i++ {
		k, err := sr.ReadInt64()
		if err != nil {
			return nil, err
		}

		v, err := sr.ReadDouble()
		if err != nil {
			return nil, err
		}

		o[k] = v
	}

	if err = sr.ReadMapEnd(); err != nil {
		return nil, err
	}
	return o, err
}

// Decode deserializes a PrimitiveRequiredStruct struct directly from its Thrift-level
// representation, without going through an intemediary type.
//
// An error is returned if a PrimitiveRequiredStruct struct could not be generated from the wire
// representation.
func (v *PrimitiveRequiredStruct) Decode(sr stream.Reader) error {

	boolFieldIsSet := false
	byteFieldIsSet := false
	int16FieldIsSet := false
	int32FieldIsSet := false
	int64FieldIsSet := false
	doubleFieldIsSet := false
	stringFieldIsSet := false
	binaryFieldIsSet := false
	listOfStringsIsSet := false
	setOfIntsIsSet := false
	mapOfIntsToDoublesIsSet := false

	if err := sr.ReadStructBegin(); err != nil {
		return err
	}

	fh, ok, err := sr.ReadFieldBegin()
	if err != nil {
		return err
	}

	for ok {
		switch {
		case fh.ID == 1 && fh.Type == wire.TBool:
			v.BoolField, err = sr.ReadBool()
			if err != nil {
				return err
			}
			boolFieldIsSet = true
		case fh.ID == 2 && fh.Type == wire.TI8:
			v.ByteField, err = sr.ReadInt8()
			if err != nil {
				return err
			}
			byteFieldIsSet = true
		case fh.ID == 3 && fh.Type == wire.TI16:
			v.Int16Field, err = sr.ReadInt16()
			if err != nil {
				return err
			}
			int16FieldIsSet = true
		case fh.ID == 4 && fh.Type == wire.TI32:
			v.Int32Field, err = sr.ReadInt32()
			if err != nil {
				return err
			}
			int32FieldIsSet = true
		case fh.ID == 5 && fh.Type == wire.TI64:
			v.Int64Field, err = sr.ReadInt64()
			if err != nil {
				return err
			}
			int64FieldIsSet = true
		case fh.ID == 6 && fh.Type == wire.TDouble:
			v.DoubleField, err = sr.ReadDouble()
			if err != nil {
				return err
			}
			doubleFieldIsSet = true
		case fh.ID == 7 && fh.Type == wire.TBinary:
			v.StringField, err = sr.ReadString()
			if err != nil {
				return err
			}
			stringFieldIsSet = true
		case fh.ID == 8 && fh.Type == wire.TBinary:
			v.BinaryField, err = sr.ReadBinary()
			if err != nil {
				return err
			}
			binaryFieldIsSet = true
		case fh.ID == 9 && fh.Type == wire.TList:
			v.ListOfStrings, err = _List_String_Decode(sr)
			if err != nil {
				return err
			}
			listOfStringsIsSet = true
		case fh.ID == 10 && fh.Type == wire.TSet:
			v.SetOfInts, err = _Set_I32_mapType_Decode(sr)
			if err != nil {
				return err
			}
			setOfIntsIsSet = true
		case fh.ID == 11 && fh.Type == wire.TMap:
			v.MapOfIntsToDoubles, err = _Map_I64_Double_Decode(sr)
			if err != nil {
				return err
			}
			mapOfIntsToDoublesIsSet = true
		default:
			if err := sr.Skip(fh.Type); err != nil {
				return err
			}
		}

		if err := sr.ReadFieldEnd(); err != nil {
			return err
		}

		if fh, ok, err = sr.ReadFieldBegin(); err != nil {
			return err
		}
	}

	if err := sr.ReadStructEnd(); err != nil {
		return err
	}

	if !boolFieldIsSet {
		return errors.New("field BoolField of PrimitiveRequiredStruct is required")
	}

	if !byteFieldIsSet {
		return errors.New("field ByteField of PrimitiveRequiredStruct is required")
	}

	if !int16FieldIsSet {
		return errors.New("field Int16Field of PrimitiveRequiredStruct is required")
	}

	if !int32FieldIsSet {
		return errors.New("field Int32Field of PrimitiveRequiredStruct is required")
	}

	if !int64FieldIsSet {
		return errors.New("field Int64Field of PrimitiveRequiredStruct is required")
	}

	if !doubleFieldIsSet {
		return errors.New("field DoubleField of PrimitiveRequiredStruct is required")
	}

	if !stringFieldIsSet {
		return errors.New("field StringField of PrimitiveRequiredStruct is required")
	}

	if !binaryFieldIsSet {
		return errors.New("field BinaryField of PrimitiveRequiredStruct is required")
	}

	if !listOfStringsIsSet {
		return errors.New("field ListOfStrings of PrimitiveRequiredStruct is required")
	}

	if !setOfIntsIsSet {
		return errors.New("field SetOfInts of PrimitiveRequiredStruct is required")
	}

	if !mapOfIntsToDoublesIsSet {
		return errors.New("field MapOfIntsToDoubles of PrimitiveRequiredStruct is required")
	}

	return nil
}

// String returns a readable string representation of a PrimitiveRequiredStruct
// struct.
func (v *PrimitiveRequiredStruct) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [11]string
	i := 0
	fields[i] = fmt.Sprintf("BoolField: %v", v.BoolField)
	i++
	fields[i] = fmt.Sprintf("ByteField: %v", v.ByteField)
	i++
	fields[i] = fmt.Sprintf("Int16Field: %v", v.Int16Field)
	i++
	fields[i] = fmt.Sprintf("Int32Field: %v", v.Int32Field)
	i++
	fields[i] = fmt.Sprintf("Int64Field: %v", v.Int64Field)
	i++
	fields[i] = fmt.Sprintf("DoubleField: %v", v.DoubleField)
	i++
	fields[i] = fmt.Sprintf("StringField: %v", v.StringField)
	i++
	fields[i] = fmt.Sprintf("BinaryField: %v", v.BinaryField)
	i++
	fields[i] = fmt.Sprintf("ListOfStrings: %v", v.ListOfStrings)
	i++
	fields[i] = fmt.Sprintf("SetOfInts: %v", v.SetOfInts)
	i++
	fields[i] = fmt.Sprintf("MapOfIntsToDoubles: %v", v.MapOfIntsToDoubles)
	i++

	return fmt.Sprintf("PrimitiveRequiredStruct{%v}", strings.Join(fields[:i], ", "))
}

func _List_String_Equals(lhs, rhs []string) bool {
	if len(lhs) != len(rhs) {
		return false
	}

	for i, lv := range lhs {
		rv := rhs[i]
		if !(lv == rv) {
			return false
		}
	}

	return true
}

func _Set_I32_mapType_Equals(lhs, rhs map[int32]struct{}) bool {
	if len(lhs) != len(rhs) {
		return false
	}

	for x := range rhs {
		if _, ok := lhs[x]; !ok {
			return false
		}
	}

	return true
}

func _Map_I64_Double_Equals(lhs, rhs map[int64]float64) bool {
	if len(lhs) != len(rhs) {
		return false
	}

	for lk, lv := range lhs {
		rv, ok := rhs[lk]
		if !ok {
			return false
		}
		if !(lv == rv) {
			return false
		}
	}
	return true
}

// Equals returns true if all the fields of this PrimitiveRequiredStruct match the
// provided PrimitiveRequiredStruct.
//
// This function performs a deep comparison.
func (v *PrimitiveRequiredStruct) Equals(rhs *PrimitiveRequiredStruct) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !(v.BoolField == rhs.BoolField) {
		return false
	}
	if !(v.ByteField == rhs.ByteField) {
		return false
	}
	if !(v.Int16Field == rhs.Int16Field) {
		return false
	}
	if !(v.Int32Field == rhs.Int32Field) {
		return false
	}
	if !(v.Int64Field == rhs.Int64Field) {
		return false
	}
	if !(v.DoubleField == rhs.DoubleField) {
		return false
	}
	if !(v.StringField == rhs.StringField) {
		return false
	}
	if !bytes.Equal(v.BinaryField, rhs.BinaryField) {
		return false
	}
	if !_List_String_Equals(v.ListOfStrings, rhs.ListOfStrings) {
		return false
	}
	if !_Set_I32_mapType_Equals(v.SetOfInts, rhs.SetOfInts) {
		return false
	}
	if !_Map_I64_Double_Equals(v.MapOfIntsToDoubles, rhs.MapOfIntsToDoubles) {
		return false
	}

	return true
}

// GetBoolField returns the value of BoolField if it is set or its
// zero value if it is unset.
func (v *PrimitiveRequiredStruct) GetBoolField() (o bool) {
	if v != nil {
		o = v.BoolField
	}
	return
}

// GetByteField returns the value of ByteField if it is set or its
// zero value if it is unset.
func (v *PrimitiveRequiredStruct) GetByteField() (o int8) {
	if v != nil {
		o = v.ByteField
	}
	return
}

// GetInt16Field returns the value of Int16Field if it is set or its
// zero value if it is unset.
func (v *PrimitiveRequiredStruct) GetInt16Field() (o int16) {
	if v != nil {
		o = v.Int16Field
	}
	return
}

// GetInt32Field returns the value of Int32Field if it is set or its
// zero value if it is unset.
func (v *PrimitiveRequiredStruct) GetInt32Field() (o int32) {
	if v != nil {
		o = v.Int32Field
	}
	return
}

// GetInt64Field returns the value of Int64Field if it is set or its
// zero value if it is unset.
func (v *PrimitiveRequiredStruct) GetInt64Field() (o int64) {
	if v != nil {
		o = v.Int64Field
	}
	return
}

// GetDoubleField returns the value of DoubleField if it is set or its
// zero value if it is unset.
func (v *PrimitiveRequiredStruct) GetDoubleField() (o float64) {
	if v != nil {
		o = v.DoubleField
	}
	return
}

// GetStringField returns the value of StringField if it is set or its
// zero value if it is unset.
func (v *PrimitiveRequiredStruct) GetStringField() (o string) {
	if v != nil {
		o = v.StringField
	}
	return
}

// GetBinaryField returns the value of BinaryField if it is set or its
// zero value if it is unset.
func (v *PrimitiveRequiredStruct) GetBinaryField() (o []byte) {
	if v != nil {
		o = v.BinaryField
	}
	return
}

// IsSetBinaryField returns true if BinaryField is not nil.
func (v *PrimitiveRequiredStruct) IsSetBinaryField() bool {
	return v != nil && v.BinaryField != nil
}

// GetListOfStrings returns the value of ListOfStrings if it is set or its
// zero value if it is unset.
func (v *PrimitiveRequiredStruct) GetListOfStrings() (o []string) {
	if v != nil {
		o = v.ListOfStrings
	}
	return
}

// IsSetListOfStrings returns true if ListOfStrings is not nil.
func (v *PrimitiveRequiredStruct) IsSetListOfStrings() bool {
	return v != nil && v.ListOfStrings != nil
}

// GetSetOfInts returns the value of SetOfInts if it is set or its
// zero value if it is unset.
func (v *PrimitiveRequiredStruct) GetSetOfInts() (o map[int32]struct{}) {
	if v != nil {
		o = v.SetOfInts
	}
	return
}

// IsSetSetOfInts returns true if SetOfInts is not nil.
func (v *PrimitiveRequiredStruct) IsSetSetOfInts() bool {
	return v != nil && v.SetOfInts != nil
}

// GetMapOfIntsToDoubles returns the value of MapOfIntsToDoubles if it is set or its
// zero value if it is unset.
func (v *PrimitiveRequiredStruct) GetMapOfIntsToDoubles() (o map[int64]float64) {
	if v != nil {
		o = v.MapOfIntsToDoubles
	}
	return
}

// IsSetMapOfIntsToDoubles returns true if MapOfIntsToDoubles is not nil.
func (v *PrimitiveRequiredStruct) IsSetMapOfIntsToDoubles() bool {
	return v != nil && v.MapOfIntsToDoubles != nil
}

type Primitives PrimitiveRequiredStruct

// ToWire translates Primitives into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
func (v *Primitives) ToWire() (wire.Value, error) {
	x := (*PrimitiveRequiredStruct)(v)
	return x.ToWire()
}

// String returns a readable string representation of Primitives.
func (v *Primitives) String() string {
	x := (*PrimitiveRequiredStruct)(v)

	return fmt.Sprint(x)
}

func (v *Primitives) Encode(sw stream.Writer) error {
	x := (*PrimitiveRequiredStruct)(v)
	return x.Encode(sw)
}

// FromWire deserializes Primitives from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
func (v *Primitives) FromWire(w wire.Value) error {
	return (*PrimitiveRequiredStruct)(v).FromWire(w)
}

// Decode deserializes Primitives directly off the wire.
func (v *Primitives) Decode(sr stream.Reader) error {
	return (*PrimitiveRequiredStruct)(v).Decode(sr)
}

// Equals returns true if this Primitives is equal to the provided
// Primitives.
func (lhs *Primitives) Equals(rhs *Primitives) bool {
	return (*PrimitiveRequiredStruct)(lhs).Equals((*PrimitiveRequiredStruct)(rhs))
}

type StringList []string

// ToWire translates StringList into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
func (v StringList) ToWire() (wire.Value, error) {
	x := ([]string)(v)
	return wire.NewValueList(_List_String_ValueList(x)), error(nil)
}

// String returns a readable string representation of StringList.
func (v StringList) String() string {
	x := ([]string)(v)

	return fmt.Sprint(x)
}

func (v StringList) Encode(sw stream.Writer) error {
	x := ([]string)(v)
	return _List_String_Encode(x, sw)
}

// FromWire deserializes StringList from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
func (v *StringList) FromWire(w wire.Value) error {
	x, err := _List_String_Read(w.GetList())
	*v = (StringList)(x)
	return err
}

// Decode deserializes StringList directly off the wire.
func (v *StringList) Decode(sr stream.Reader) error {
	x, err := _List_String_Decode(sr)
	*v = (StringList)(x)
	return err
}

// Equals returns true if this StringList is equal to the provided
// StringList.
func (lhs StringList) Equals(rhs StringList) bool {
	return _List_String_Equals(([]string)(lhs), ([]string)(rhs))
}

type _Map_String_String_MapItemList map[string]string

func (m _Map_String_String_MapItemList) ForEach(f func(wire.MapItem) error) error {
	for k, v := range m {
		kw, err := wire.NewValueString(k), error(nil)
		if err != nil {
			return err
		}

		vw, err := wire.NewValueString(v), error(nil)
		if err != nil {
			return err
		}
		err = f(wire.MapItem{Key: kw, Value: vw})
		if err != nil {
			return err
		}
	}
	return nil
}

func (m _Map_String_String_MapItemList) Size() int {
	return len(m)
}

func (_Map_String_String_MapItemList) KeyType() wire.Type {
	return wire.TBinary
}

func (_Map_String_String_MapItemList) ValueType() wire.Type {
	return wire.TBinary
}

func (_Map_String_String_MapItemList) Close() {}

func _Map_String_String_Encode(val map[string]string, sw stream.Writer) error {

	mh := stream.MapHeader{
		KeyType:   wire.TBinary,
		ValueType: wire.TBinary,
		Length:    len(val),
	}
	if err := sw.WriteMapBegin(mh); err != nil {
		return err
	}

	for k, v := range val {
		if err := sw.WriteString(k); err != nil {
			return err
		}
		if err := sw.WriteString(v); err != nil {
			return err
		}
	}

	return sw.WriteMapEnd()
}

func _Map_String_String_Read(m wire.MapItemList) (map[string]string, error) {
	if m.KeyType() != wire.TBinary {
		return nil, nil
	}

	if m.ValueType() != wire.TBinary {
		return nil, nil
	}

	o := make(map[string]string, m.Size())
	err := m.ForEach(func(x wire.MapItem) error {
		k, err := x.Key.GetString(), error(nil)
		if err != nil {
			return err
		}

		v, err := x.Value.GetString(), error(nil)
		if err != nil {
			return err
		}

		o[k] = v
		return nil
	})
	m.Close()
	return o, err
}

func _Map_String_String_Decode(sr stream.Reader) (map[string]string, error) {
	mh, err := sr.ReadMapBegin()
	if err != nil {
		return nil, err
	}

	if mh.KeyType != wire.TBinary || mh.ValueType != wire.TBinary {
		for i := 0; i < mh.Length; i++ {
			if err := sr.Skip(mh.KeyType); err != nil {
				return nil, err
			}

			if err := sr.Skip(mh.ValueType); err != nil {
				return nil, err
			}
		}
		return nil, sr.ReadMapEnd()
	}

	o := make(map[string]string, mh.Length)
	for i := 0; i < mh.Length; i++ {
		k, err := sr.ReadString()
		if err != nil {
			return nil, err
		}

		v, err := sr.ReadString()
		if err != nil {
			return nil, err
		}

		o[k] = v
	}

	if err = sr.ReadMapEnd(); err != nil {
		return nil, err
	}
	return o, err
}

func _Map_String_String_Equals(lhs, rhs map[string]string) bool {
	if len(lhs) != len(rhs) {
		return false
	}

	for lk, lv := range lhs {
		rv, ok := rhs[lk]
		if !ok {
			return false
		}
		if !(lv == rv) {
			return false
		}
	}
	return true
}

type StringMap map[string]string

// ToWire translates StringMap into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
func (v StringMap) ToWire() (wire.Value, error) {
	x := (map[string]string)(v)
	return wire.NewValueMap(_Map_String_String_MapItemList(x)), error(nil)
}

// String returns a readable string representation of StringMap.
func (v StringMap) String() string {
	x := (map[string]string)(v)

	return fmt.Sprint(x)
}

func (v StringMap) Encode(sw stream.Writer) error {
	x := (map[string]string)(v)
	return _Map_String_String_Encode(x, sw)
}

// FromWire deserializes StringMap from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
func (v *StringMap) FromWire(w wire.Value) error {
	x, err := _Map_String_String_Read(w.GetMap())
	*v = (StringMap)(x)
	return err
}

// Decode deserializes StringMap directly off the wire.
func (v *StringMap) Decode(sr stream.Reader) error {
	x, err := _Map_String_String_Decode(sr)
	*v = (StringMap)(x)
	return err
}

// Equals returns true if this StringMap is equal to the provided
// StringMap.
func (lhs StringMap) Equals(rhs StringMap) bool {
	return _Map_String_String_Equals((map[string]string)(lhs), (map[string]string)(rhs))
}

// ThriftModule represents the IDL file used to generate this package.
var ThriftModule = &thriftreflect.ThriftModule{
	Name:     "nozap",
	Package:  "go.uber.org/thriftrw/gen/internal/tests/nozap",
	FilePath: "nozap.thrift",
	SHA1:     "05f7228060eeb97fbd181a7a2660a6483799ac78",
	Raw:      rawIDL,
}

const rawIDL = "enum EnumDefault {\n    Foo, Bar, Baz\n}\n\nstruct PrimitiveRequiredStruct {\n    1: required bool boolField\n    2: required byte byteField\n    3: required i16 int16Field\n    4: required i32 int32Field\n    5: required i64 int64Field\n    6: required double doubleField\n    7: required string stringField\n    8: required binary binaryField\n    9: required list<string> listOfStrings\n    10: required set<i32> setOfInts\n    11: required map<i64, double> mapOfIntsToDoubles\n}\n\ntypedef map<string, string> StringMap\ntypedef PrimitiveRequiredStruct Primitives\ntypedef list<string> StringList\n"
