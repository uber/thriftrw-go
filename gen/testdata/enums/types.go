// Code generated by thriftrw v0.4.0
// @generated

package enums

import (
	"fmt"
	"go.uber.org/thriftrw/wire"
	"strconv"
	"strings"
)

type EmptyEnum int32

func (v EmptyEnum) ToWire() (wire.Value, error) {
	return wire.NewValueI32(int32(v)), nil
}

func (v *EmptyEnum) FromWire(w wire.Value) error {
	*v = (EmptyEnum)(w.GetI32())
	return nil
}

func (v EmptyEnum) String() string {
	w := int32(v)
	return fmt.Sprintf("EmptyEnum(%d)", w)
}

func (v EmptyEnum) MarshalText() (text []byte, err error) {
	return ([]byte)(strconv.FormatInt(int64(v), 10)), nil
}

func (v EmptyEnum) UnmarshalText(text []byte) error {
	w, err := strconv.ParseInt(string(text), 10, 32)
	if err == nil {
		v = (EmptyEnum)(w)
		return err
	}
	e := err.(*strconv.NumError)
	if e.Err != strconv.ErrSyntax {
		return err
	}
	return fmt.Errorf("impossible to unmarshal %q from %q", "EmptyEnum", text)
}

type EnumDefault int32

const (
	EnumDefaultFoo EnumDefault = 0
	EnumDefaultBar EnumDefault = 1
	EnumDefaultBaz EnumDefault = 2
)

func (v EnumDefault) ToWire() (wire.Value, error) {
	return wire.NewValueI32(int32(v)), nil
}

func (v *EnumDefault) FromWire(w wire.Value) error {
	*v = (EnumDefault)(w.GetI32())
	return nil
}

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

func (v EnumDefault) MarshalText() (text []byte, err error) {
	w := int32(v)
	switch w {
	case 0:
		return ([]byte)("Foo"), nil
	case 1:
		return ([]byte)("Bar"), nil
	case 2:
		return ([]byte)("Baz"), nil
	}
	return ([]byte)(strconv.FormatInt(int64(v), 10)), nil
}

func (v EnumDefault) UnmarshalText(text []byte) error {
	w, err := strconv.ParseInt(string(text), 10, 32)
	if err == nil {
		v = (EnumDefault)(w)
		return err
	}
	e := err.(*strconv.NumError)
	if e.Err != strconv.ErrSyntax {
		return err
	}
	switch string(text) {
	case "Foo":
		v = (EnumDefault)(0)
		return nil
	case "Bar":
		v = (EnumDefault)(1)
		return nil
	case "Baz":
		v = (EnumDefault)(2)
		return nil
	}
	return fmt.Errorf("impossible to unmarshal %q from %q", "EnumDefault", text)
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

func (v EnumWithDuplicateName) ToWire() (wire.Value, error) {
	return wire.NewValueI32(int32(v)), nil
}

func (v *EnumWithDuplicateName) FromWire(w wire.Value) error {
	*v = (EnumWithDuplicateName)(w.GetI32())
	return nil
}

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

func (v EnumWithDuplicateName) MarshalText() (text []byte, err error) {
	w := int32(v)
	switch w {
	case 0:
		return ([]byte)("A"), nil
	case 1:
		return ([]byte)("B"), nil
	case 2:
		return ([]byte)("C"), nil
	case 3:
		return ([]byte)("P"), nil
	case 4:
		return ([]byte)("Q"), nil
	case 5:
		return ([]byte)("R"), nil
	case 6:
		return ([]byte)("X"), nil
	case 7:
		return ([]byte)("Y"), nil
	case 8:
		return ([]byte)("Z"), nil
	}
	return ([]byte)(strconv.FormatInt(int64(v), 10)), nil
}

func (v EnumWithDuplicateName) UnmarshalText(text []byte) error {
	w, err := strconv.ParseInt(string(text), 10, 32)
	if err == nil {
		v = (EnumWithDuplicateName)(w)
		return err
	}
	e := err.(*strconv.NumError)
	if e.Err != strconv.ErrSyntax {
		return err
	}
	switch string(text) {
	case "A":
		v = (EnumWithDuplicateName)(0)
		return nil
	case "B":
		v = (EnumWithDuplicateName)(1)
		return nil
	case "C":
		v = (EnumWithDuplicateName)(2)
		return nil
	case "P":
		v = (EnumWithDuplicateName)(3)
		return nil
	case "Q":
		v = (EnumWithDuplicateName)(4)
		return nil
	case "R":
		v = (EnumWithDuplicateName)(5)
		return nil
	case "X":
		v = (EnumWithDuplicateName)(6)
		return nil
	case "Y":
		v = (EnumWithDuplicateName)(7)
		return nil
	case "Z":
		v = (EnumWithDuplicateName)(8)
		return nil
	}
	return fmt.Errorf("impossible to unmarshal %q from %q", "EnumWithDuplicateName", text)
}

type EnumWithDuplicateValues int32

const (
	EnumWithDuplicateValuesP EnumWithDuplicateValues = 0
	EnumWithDuplicateValuesQ EnumWithDuplicateValues = -1
	EnumWithDuplicateValuesR EnumWithDuplicateValues = 0
)

func (v EnumWithDuplicateValues) ToWire() (wire.Value, error) {
	return wire.NewValueI32(int32(v)), nil
}

func (v *EnumWithDuplicateValues) FromWire(w wire.Value) error {
	*v = (EnumWithDuplicateValues)(w.GetI32())
	return nil
}

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

func (v EnumWithDuplicateValues) MarshalText() (text []byte, err error) {
	w := int32(v)
	switch w {
	case 0:
		return ([]byte)("P"), nil
	case -1:
		return ([]byte)("Q"), nil
	}
	return ([]byte)(strconv.FormatInt(int64(v), 10)), nil
}

func (v EnumWithDuplicateValues) UnmarshalText(text []byte) error {
	w, err := strconv.ParseInt(string(text), 10, 32)
	if err == nil {
		v = (EnumWithDuplicateValues)(w)
		return err
	}
	e := err.(*strconv.NumError)
	if e.Err != strconv.ErrSyntax {
		return err
	}
	switch string(text) {
	case "P":
		v = (EnumWithDuplicateValues)(0)
		return nil
	case "Q":
		v = (EnumWithDuplicateValues)(-1)
		return nil
	case "R":
		v = (EnumWithDuplicateValues)(0)
		return nil
	}
	return fmt.Errorf("impossible to unmarshal %q from %q", "EnumWithDuplicateValues", text)
}

type EnumWithValues int32

const (
	EnumWithValuesX EnumWithValues = 123
	EnumWithValuesY EnumWithValues = 456
	EnumWithValuesZ EnumWithValues = 789
)

func (v EnumWithValues) ToWire() (wire.Value, error) {
	return wire.NewValueI32(int32(v)), nil
}

func (v *EnumWithValues) FromWire(w wire.Value) error {
	*v = (EnumWithValues)(w.GetI32())
	return nil
}

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

func (v EnumWithValues) MarshalText() (text []byte, err error) {
	w := int32(v)
	switch w {
	case 123:
		return ([]byte)("X"), nil
	case 456:
		return ([]byte)("Y"), nil
	case 789:
		return ([]byte)("Z"), nil
	}
	return ([]byte)(strconv.FormatInt(int64(v), 10)), nil
}

func (v EnumWithValues) UnmarshalText(text []byte) error {
	w, err := strconv.ParseInt(string(text), 10, 32)
	if err == nil {
		v = (EnumWithValues)(w)
		return err
	}
	e := err.(*strconv.NumError)
	if e.Err != strconv.ErrSyntax {
		return err
	}
	switch string(text) {
	case "X":
		v = (EnumWithValues)(123)
		return nil
	case "Y":
		v = (EnumWithValues)(456)
		return nil
	case "Z":
		v = (EnumWithValues)(789)
		return nil
	}
	return fmt.Errorf("impossible to unmarshal %q from %q", "EnumWithValues", text)
}

type RecordType int32

const (
	RecordTypeName        RecordType = 0
	RecordTypeHomeAddress RecordType = 1
	RecordTypeWorkAddress RecordType = 2
)

func (v RecordType) ToWire() (wire.Value, error) {
	return wire.NewValueI32(int32(v)), nil
}

func (v *RecordType) FromWire(w wire.Value) error {
	*v = (RecordType)(w.GetI32())
	return nil
}

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

func (v RecordType) MarshalText() (text []byte, err error) {
	w := int32(v)
	switch w {
	case 0:
		return ([]byte)("NAME"), nil
	case 1:
		return ([]byte)("HOME_ADDRESS"), nil
	case 2:
		return ([]byte)("WORK_ADDRESS"), nil
	}
	return ([]byte)(strconv.FormatInt(int64(v), 10)), nil
}

func (v RecordType) UnmarshalText(text []byte) error {
	w, err := strconv.ParseInt(string(text), 10, 32)
	if err == nil {
		v = (RecordType)(w)
		return err
	}
	e := err.(*strconv.NumError)
	if e.Err != strconv.ErrSyntax {
		return err
	}
	switch string(text) {
	case "NAME":
		v = (RecordType)(0)
		return nil
	case "HOME_ADDRESS":
		v = (RecordType)(1)
		return nil
	case "WORK_ADDRESS":
		v = (RecordType)(2)
		return nil
	}
	return fmt.Errorf("impossible to unmarshal %q from %q", "RecordType", text)
}

type StructWithOptionalEnum struct {
	E *EnumDefault `json:"e,omitempty"`
}

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

func (v *StructWithOptionalEnum) String() string {
	var fields [1]string
	i := 0
	if v.E != nil {
		fields[i] = fmt.Sprintf("E: %v", *(v.E))
		i++
	}
	return fmt.Sprintf("StructWithOptionalEnum{%v}", strings.Join(fields[:i], ", "))
}

type LowerCaseEnum int32

const (
	LowerCaseEnumContaining LowerCaseEnum = 0
	LowerCaseEnumLowerCase  LowerCaseEnum = 1
	LowerCaseEnumItems      LowerCaseEnum = 2
)

func (v LowerCaseEnum) ToWire() (wire.Value, error) {
	return wire.NewValueI32(int32(v)), nil
}

func (v *LowerCaseEnum) FromWire(w wire.Value) error {
	*v = (LowerCaseEnum)(w.GetI32())
	return nil
}

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

func (v LowerCaseEnum) MarshalText() (text []byte, err error) {
	w := int32(v)
	switch w {
	case 0:
		return ([]byte)("containing"), nil
	case 1:
		return ([]byte)("lower_case"), nil
	case 2:
		return ([]byte)("items"), nil
	}
	return ([]byte)(strconv.FormatInt(int64(v), 10)), nil
}

func (v LowerCaseEnum) UnmarshalText(text []byte) error {
	w, err := strconv.ParseInt(string(text), 10, 32)
	if err == nil {
		v = (LowerCaseEnum)(w)
		return err
	}
	e := err.(*strconv.NumError)
	if e.Err != strconv.ErrSyntax {
		return err
	}
	switch string(text) {
	case "containing":
		v = (LowerCaseEnum)(0)
		return nil
	case "lower_case":
		v = (LowerCaseEnum)(1)
		return nil
	case "items":
		v = (LowerCaseEnum)(2)
		return nil
	}
	return fmt.Errorf("impossible to unmarshal %q from %q", "LowerCaseEnum", text)
}
