// Code generated by thriftrw v1.30.0. DO NOT EDIT.
// @generated

package uuid_conflict

import (
	errors "errors"
	fmt "fmt"
	multierr "go.uber.org/multierr"
	typedefs "go.uber.org/thriftrw/gen/internal/tests/typedefs"
	stream "go.uber.org/thriftrw/protocol/stream"
	thriftreflect "go.uber.org/thriftrw/thriftreflect"
	wire "go.uber.org/thriftrw/wire"
	zapcore "go.uber.org/zap/zapcore"
	strings "strings"
)

type UUID string

// UUIDPtr returns a pointer to a UUID
func (v UUID) Ptr() *UUID {
	return &v
}

// ToWire translates UUID into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
func (v UUID) ToWire() (wire.Value, error) {
	x := (string)(v)
	return wire.NewValueString(x), error(nil)
}

// String returns a readable string representation of UUID.
func (v UUID) String() string {
	x := (string)(v)
	return (string)(x)
}

func (v UUID) Encode(sw stream.Writer) error {
	x := (string)(v)
	return sw.WriteString(x)
}

// FromWire deserializes UUID from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
func (v *UUID) FromWire(w wire.Value) error {
	x, err := w.GetString(), error(nil)
	*v = (UUID)(x)
	return err
}

// Decode deserializes UUID directly off the wire.
func (v *UUID) Decode(sr stream.Reader) error {
	x, err := sr.ReadString()
	*v = (UUID)(x)
	return err
}

// Equals returns true if this UUID is equal to the provided
// UUID.
func (lhs UUID) Equals(rhs UUID) bool {
	return ((string)(lhs) == (string)(rhs))
}

type UUIDConflict struct {
	LocalUUID    UUID           `json:"localUUID,required"`
	ImportedUUID *typedefs.UUID `json:"importedUUID,required"`
}

// ToWire translates a UUIDConflict struct into a Thrift-level intermediate
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
func (v *UUIDConflict) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	w, err = v.LocalUUID.ToWire()
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	if v.ImportedUUID == nil {
		return w, errors.New("field ImportedUUID of UUIDConflict is required")
	}
	w, err = v.ImportedUUID.ToWire()
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 2, Value: w}
	i++

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _UUID_Read(w wire.Value) (UUID, error) {
	var x UUID
	err := x.FromWire(w)
	return x, err
}

func _UUID_1_Read(w wire.Value) (*typedefs.UUID, error) {
	var x typedefs.UUID
	err := x.FromWire(w)
	return &x, err
}

// FromWire deserializes a UUIDConflict struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a UUIDConflict struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v UUIDConflict
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *UUIDConflict) FromWire(w wire.Value) error {
	var err error

	localUUIDIsSet := false
	importedUUIDIsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.LocalUUID, err = _UUID_Read(field.Value)
				if err != nil {
					return err
				}
				localUUIDIsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TStruct {
				v.ImportedUUID, err = _UUID_1_Read(field.Value)
				if err != nil {
					return err
				}
				importedUUIDIsSet = true
			}
		}
	}

	if !localUUIDIsSet {
		return errors.New("field LocalUUID of UUIDConflict is required")
	}

	if !importedUUIDIsSet {
		return errors.New("field ImportedUUID of UUIDConflict is required")
	}

	return nil
}

// Encode serializes a UUIDConflict struct directly into bytes, without going
// through an intermediary type.
//
// An error is returned if a UUIDConflict struct could not be encoded.
func (v *UUIDConflict) Encode(sw stream.Writer) error {
	if err := sw.WriteStructBegin(); err != nil {
		return err
	}

	if err := sw.WriteFieldBegin(stream.FieldHeader{ID: 1, Type: wire.TBinary}); err != nil {
		return err
	}
	if err := v.LocalUUID.Encode(sw); err != nil {
		return err
	}
	if err := sw.WriteFieldEnd(); err != nil {
		return err
	}

	if v.ImportedUUID == nil {
		return errors.New("field ImportedUUID of UUIDConflict is required")
	}
	if err := sw.WriteFieldBegin(stream.FieldHeader{ID: 2, Type: wire.TStruct}); err != nil {
		return err
	}
	if err := v.ImportedUUID.Encode(sw); err != nil {
		return err
	}
	if err := sw.WriteFieldEnd(); err != nil {
		return err
	}

	return sw.WriteStructEnd()
}

func _UUID_Decode(sr stream.Reader) (UUID, error) {
	var x UUID
	err := x.Decode(sr)
	return x, err
}

func _UUID_1_Decode(sr stream.Reader) (*typedefs.UUID, error) {
	var x typedefs.UUID
	err := x.Decode(sr)
	return &x, err
}

// Decode deserializes a UUIDConflict struct directly from its Thrift-level
// representation, without going through an intemediary type.
//
// An error is returned if a UUIDConflict struct could not be generated from the wire
// representation.
func (v *UUIDConflict) Decode(sr stream.Reader) error {

	localUUIDIsSet := false
	importedUUIDIsSet := false

	if err := sr.ReadStructBegin(); err != nil {
		return err
	}

	fh, ok, err := sr.ReadFieldBegin()
	if err != nil {
		return err
	}

	for ok {
		switch {
		case fh.ID == 1 && fh.Type == wire.TBinary:
			v.LocalUUID, err = _UUID_Decode(sr)
			if err != nil {
				return err
			}
			localUUIDIsSet = true
		case fh.ID == 2 && fh.Type == wire.TStruct:
			v.ImportedUUID, err = _UUID_1_Decode(sr)
			if err != nil {
				return err
			}
			importedUUIDIsSet = true
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

	if !localUUIDIsSet {
		return errors.New("field LocalUUID of UUIDConflict is required")
	}

	if !importedUUIDIsSet {
		return errors.New("field ImportedUUID of UUIDConflict is required")
	}

	return nil
}

// String returns a readable string representation of a UUIDConflict
// struct.
func (v *UUIDConflict) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [2]string
	i := 0
	fields[i] = fmt.Sprintf("LocalUUID: %v", v.LocalUUID)
	i++
	fields[i] = fmt.Sprintf("ImportedUUID: %v", v.ImportedUUID)
	i++

	return fmt.Sprintf("UUIDConflict{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this UUIDConflict match the
// provided UUIDConflict.
//
// This function performs a deep comparison.
func (v *UUIDConflict) Equals(rhs *UUIDConflict) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !(v.LocalUUID == rhs.LocalUUID) {
		return false
	}
	if !v.ImportedUUID.Equals(rhs.ImportedUUID) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of UUIDConflict.
func (v *UUIDConflict) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	enc.AddString("localUUID", (string)(v.LocalUUID))
	err = multierr.Append(err, enc.AddObject("importedUUID", v.ImportedUUID))
	return err
}

// GetLocalUUID returns the value of LocalUUID if it is set or its
// zero value if it is unset.
func (v *UUIDConflict) GetLocalUUID() (o UUID) {
	if v != nil {
		o = v.LocalUUID
	}
	return
}

// GetImportedUUID returns the value of ImportedUUID if it is set or its
// zero value if it is unset.
func (v *UUIDConflict) GetImportedUUID() (o *typedefs.UUID) {
	if v != nil {
		o = v.ImportedUUID
	}
	return
}

// IsSetImportedUUID returns true if ImportedUUID is not nil.
func (v *UUIDConflict) IsSetImportedUUID() bool {
	return v != nil && v.ImportedUUID != nil
}

// ThriftModule represents the IDL file used to generate this package.
var ThriftModule = &thriftreflect.ThriftModule{
	Name:     "uuid_conflict",
	Package:  "go.uber.org/thriftrw/gen/internal/tests/uuid_conflict",
	FilePath: "uuid_conflict.thrift",
	SHA1:     "c7ab8450f4c3a548cde8938fe7e150cf1b8f9493",
	Includes: []*thriftreflect.ThriftModule{
		typedefs.ThriftModule,
	},
	Raw: rawIDL,
}

const rawIDL = "include \"./typedefs.thrift\"\n\ntypedef string UUID\n\nstruct UUIDConflict {\n    1: required UUID localUUID\n    2: required typedefs.UUID importedUUID\n}\n"
