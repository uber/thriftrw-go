// Code generated by thriftrw v1.3.0
// @generated

package collision

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"go.uber.org/thriftrw/wire"
	"math"
	"strconv"
	"strings"
)

type LittlePotatoe int64

func (v LittlePotatoe) ToWire() (wire.Value, error) {
	x := (int64)(v)
	return wire.NewValueI64(x), error(nil)
}

func (v LittlePotatoe) String() string {
	x := (int64)(v)
	return fmt.Sprint(x)
}

func (v *LittlePotatoe) FromWire(w wire.Value) error {
	x, err := w.GetI64(), error(nil)
	*v = (LittlePotatoe)(x)
	return err
}

func (lhs LittlePotatoe) Equals(rhs LittlePotatoe) bool {
	return (lhs == rhs)
}

type MyEnum int32

const (
	MyEnumX       MyEnum = 123
	MyEnumY       MyEnum = 456
	MyEnumZ       MyEnum = 789
	MyEnumFooBar  MyEnum = 790
	MyEnumFooBar2 MyEnum = 791
)

func MyEnum_Values() []MyEnum {
	return []MyEnum{MyEnumX, MyEnumY, MyEnumZ, MyEnumFooBar, MyEnumFooBar2}
}

func (v *MyEnum) UnmarshalText(value []byte) error {
	switch string(value) {
	case "X":
		*v = MyEnumX
		return nil
	case "Y":
		*v = MyEnumY
		return nil
	case "Z":
		*v = MyEnumZ
		return nil
	case "FooBar":
		*v = MyEnumFooBar
		return nil
	case "foo_bar":
		*v = MyEnumFooBar2
		return nil
	default:
		return fmt.Errorf("unknown enum value %q for %q", value, "MyEnum")
	}
}

func (v MyEnum) ToWire() (wire.Value, error) {
	return wire.NewValueI32(int32(v)), nil
}

func (v *MyEnum) FromWire(w wire.Value) error {
	*v = (MyEnum)(w.GetI32())
	return nil
}

func (v MyEnum) String() string {
	w := int32(v)
	switch w {
	case 123:
		return "X"
	case 456:
		return "Y"
	case 789:
		return "Z"
	case 790:
		return "FooBar"
	case 791:
		return "foo_bar"
	}
	return fmt.Sprintf("MyEnum(%d)", w)
}

func (v MyEnum) Equals(rhs MyEnum) bool {
	return v == rhs
}

func (v MyEnum) MarshalJSON() ([]byte, error) {
	switch int32(v) {
	case 123:
		return ([]byte)("\"X\""), nil
	case 456:
		return ([]byte)("\"Y\""), nil
	case 789:
		return ([]byte)("\"Z\""), nil
	case 790:
		return ([]byte)("\"FooBar\""), nil
	case 791:
		return ([]byte)("\"foo_bar\""), nil
	}
	return ([]byte)(strconv.FormatInt(int64(v), 10)), nil
}

func (v *MyEnum) UnmarshalJSON(text []byte) error {
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
			return fmt.Errorf("enum overflow from JSON %q for %q", text, "MyEnum")
		}
		if x < math.MinInt32 {
			return fmt.Errorf("enum underflow from JSON %q for %q", text, "MyEnum")
		}
		*v = (MyEnum)(x)
		return nil
	case string:
		return v.UnmarshalText([]byte(w))
	default:
		return fmt.Errorf("invalid JSON value %q (%T) to unmarshal into %q", t, t, "MyEnum")
	}
}

type PrimitiveContainers struct {
	A []string            `json:"ListOrSetOrMap"`
	B map[string]struct{} `json:"List_Or_SetOrMap"`
	C map[string]string   `json:"ListOrSet_Or_Map"`
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

func (_List_String_ValueList) Close() {
}

type _Set_String_ValueList map[string]struct{}

func (v _Set_String_ValueList) ForEach(f func(wire.Value) error) error {
	for x := range v {
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

func (v _Set_String_ValueList) Size() int {
	return len(v)
}

func (_Set_String_ValueList) ValueType() wire.Type {
	return wire.TBinary
}

func (_Set_String_ValueList) Close() {
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

func (_Map_String_String_MapItemList) Close() {
}

func (v *PrimitiveContainers) ToWire() (wire.Value, error) {
	var (
		fields [3]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.A != nil {
		w, err = wire.NewValueList(_List_String_ValueList(v.A)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	if v.B != nil {
		w, err = wire.NewValueSet(_Set_String_ValueList(v.B)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 3, Value: w}
		i++
	}
	if v.C != nil {
		w, err = wire.NewValueMap(_Map_String_String_MapItemList(v.C)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 5, Value: w}
		i++
	}
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

func _Set_String_Read(s wire.ValueList) (map[string]struct{}, error) {
	if s.ValueType() != wire.TBinary {
		return nil, nil
	}
	o := make(map[string]struct{}, s.Size())
	err := s.ForEach(func(x wire.Value) error {
		i, err := x.GetString(), error(nil)
		if err != nil {
			return err
		}
		o[i] = struct{}{}
		return nil
	})
	s.Close()
	return o, err
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

func (v *PrimitiveContainers) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TList {
				v.A, err = _List_String_Read(field.Value.GetList())
				if err != nil {
					return err
				}
			}
		case 3:
			if field.Value.Type() == wire.TSet {
				v.B, err = _Set_String_Read(field.Value.GetSet())
				if err != nil {
					return err
				}
			}
		case 5:
			if field.Value.Type() == wire.TMap {
				v.C, err = _Map_String_String_Read(field.Value.GetMap())
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (v *PrimitiveContainers) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [3]string
	i := 0
	if v.A != nil {
		fields[i] = fmt.Sprintf("A: %v", v.A)
		i++
	}
	if v.B != nil {
		fields[i] = fmt.Sprintf("B: %v", v.B)
		i++
	}
	if v.C != nil {
		fields[i] = fmt.Sprintf("C: %v", v.C)
		i++
	}
	return fmt.Sprintf("PrimitiveContainers{%v}", strings.Join(fields[:i], ", "))
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

func _Set_String_Equals(lhs, rhs map[string]struct{}) bool {
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

func (v *PrimitiveContainers) Equals(rhs *PrimitiveContainers) bool {
	if !((v.A == nil && rhs.A == nil) || (v.A != nil && rhs.A != nil && _List_String_Equals(v.A, rhs.A))) {
		return false
	}
	if !((v.B == nil && rhs.B == nil) || (v.B != nil && rhs.B != nil && _Set_String_Equals(v.B, rhs.B))) {
		return false
	}
	if !((v.C == nil && rhs.C == nil) || (v.C != nil && rhs.C != nil && _Map_String_String_Equals(v.C, rhs.C))) {
		return false
	}
	return true
}

type StructCollision struct {
	CollisionField  bool   `json:"collisionField"`
	CollisionField2 string `json:"collision_field"`
}

func (v *StructCollision) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	w, err = wire.NewValueBool(v.CollisionField), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	w, err = wire.NewValueString(v.CollisionField2), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 2, Value: w}
	i++
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *StructCollision) FromWire(w wire.Value) error {
	var err error
	collisionFieldIsSet := false
	collision_fieldIsSet := false
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBool {
				v.CollisionField, err = field.Value.GetBool(), error(nil)
				if err != nil {
					return err
				}
				collisionFieldIsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TBinary {
				v.CollisionField2, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				collision_fieldIsSet = true
			}
		}
	}
	if !collisionFieldIsSet {
		return errors.New("field CollisionField of StructCollision is required")
	}
	if !collision_fieldIsSet {
		return errors.New("field CollisionField2 of StructCollision is required")
	}
	return nil
}

func (v *StructCollision) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [2]string
	i := 0
	fields[i] = fmt.Sprintf("CollisionField: %v", v.CollisionField)
	i++
	fields[i] = fmt.Sprintf("CollisionField2: %v", v.CollisionField2)
	i++
	return fmt.Sprintf("StructCollision{%v}", strings.Join(fields[:i], ", "))
}

func (v *StructCollision) Equals(rhs *StructCollision) bool {
	if !(v.CollisionField == rhs.CollisionField) {
		return false
	}
	if !(v.CollisionField2 == rhs.CollisionField2) {
		return false
	}
	return true
}

type UnionCollision struct {
	CollisionField  *bool   `json:"collisionField,omitempty"`
	CollisionField2 *string `json:"collision_field,omitempty"`
}

func (v *UnionCollision) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.CollisionField != nil {
		w, err = wire.NewValueBool(*(v.CollisionField)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	if v.CollisionField2 != nil {
		w, err = wire.NewValueString(*(v.CollisionField2)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}
	if i != 1 {
		return wire.Value{}, fmt.Errorf("UnionCollision should have exactly one field: got %v fields", i)
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *UnionCollision) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBool {
				var x bool
				x, err = field.Value.GetBool(), error(nil)
				v.CollisionField = &x
				if err != nil {
					return err
				}
			}
		case 2:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.CollisionField2 = &x
				if err != nil {
					return err
				}
			}
		}
	}
	count := 0
	if v.CollisionField != nil {
		count++
	}
	if v.CollisionField2 != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("UnionCollision should have exactly one field: got %v fields", count)
	}
	return nil
}

func (v *UnionCollision) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [2]string
	i := 0
	if v.CollisionField != nil {
		fields[i] = fmt.Sprintf("CollisionField: %v", *(v.CollisionField))
		i++
	}
	if v.CollisionField2 != nil {
		fields[i] = fmt.Sprintf("CollisionField2: %v", *(v.CollisionField2))
		i++
	}
	return fmt.Sprintf("UnionCollision{%v}", strings.Join(fields[:i], ", "))
}

func _Bool_EqualsPtr(lhs, rhs *bool) bool {
	if lhs != nil && rhs != nil {
		x := *lhs
		y := *rhs
		return (x == y)
	}
	return lhs == nil && rhs == nil
}

func _String_EqualsPtr(lhs, rhs *string) bool {
	if lhs != nil && rhs != nil {
		x := *lhs
		y := *rhs
		return (x == y)
	}
	return lhs == nil && rhs == nil
}

func (v *UnionCollision) Equals(rhs *UnionCollision) bool {
	if !_Bool_EqualsPtr(v.CollisionField, rhs.CollisionField) {
		return false
	}
	if !_String_EqualsPtr(v.CollisionField2, rhs.CollisionField2) {
		return false
	}
	return true
}

type WithDefault struct {
	Pouet *StructCollision2 `json:"pouet,omitempty"`
}

func (v *WithDefault) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Pouet == nil {
		v.Pouet = &StructCollision2{CollisionField: false, CollisionField2: "false indeed"}
	}
	{
		w, err = v.Pouet.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _StructCollision_Read(w wire.Value) (*StructCollision2, error) {
	var v StructCollision2
	err := v.FromWire(w)
	return &v, err
}

func (v *WithDefault) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.Pouet, err = _StructCollision_Read(field.Value)
				if err != nil {
					return err
				}
			}
		}
	}
	if v.Pouet == nil {
		v.Pouet = &StructCollision2{CollisionField: false, CollisionField2: "false indeed"}
	}
	return nil
}

func (v *WithDefault) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [1]string
	i := 0
	if v.Pouet != nil {
		fields[i] = fmt.Sprintf("Pouet: %v", v.Pouet)
		i++
	}
	return fmt.Sprintf("WithDefault{%v}", strings.Join(fields[:i], ", "))
}

func (v *WithDefault) Equals(rhs *WithDefault) bool {
	if !((v.Pouet == nil && rhs.Pouet == nil) || (v.Pouet != nil && rhs.Pouet != nil && v.Pouet.Equals(rhs.Pouet))) {
		return false
	}
	return true
}

type LittlePotatoe2 float64

func (v LittlePotatoe2) ToWire() (wire.Value, error) {
	x := (float64)(v)
	return wire.NewValueDouble(x), error(nil)
}

func (v LittlePotatoe2) String() string {
	x := (float64)(v)
	return fmt.Sprint(x)
}

func (v *LittlePotatoe2) FromWire(w wire.Value) error {
	x, err := w.GetDouble(), error(nil)
	*v = (LittlePotatoe2)(x)
	return err
}

func (lhs LittlePotatoe2) Equals(rhs LittlePotatoe2) bool {
	return (lhs == rhs)
}

type MyEnum2 int32

const (
	MyEnum2X MyEnum2 = 12
	MyEnum2Y MyEnum2 = 34
	MyEnum2Z MyEnum2 = 56
)

func MyEnum2_Values() []MyEnum2 {
	return []MyEnum2{MyEnum2X, MyEnum2Y, MyEnum2Z}
}

func (v *MyEnum2) UnmarshalText(value []byte) error {
	switch string(value) {
	case "X":
		*v = MyEnum2X
		return nil
	case "Y":
		*v = MyEnum2Y
		return nil
	case "Z":
		*v = MyEnum2Z
		return nil
	default:
		return fmt.Errorf("unknown enum value %q for %q", value, "MyEnum2")
	}
}

func (v MyEnum2) ToWire() (wire.Value, error) {
	return wire.NewValueI32(int32(v)), nil
}

func (v *MyEnum2) FromWire(w wire.Value) error {
	*v = (MyEnum2)(w.GetI32())
	return nil
}

func (v MyEnum2) String() string {
	w := int32(v)
	switch w {
	case 12:
		return "X"
	case 34:
		return "Y"
	case 56:
		return "Z"
	}
	return fmt.Sprintf("MyEnum2(%d)", w)
}

func (v MyEnum2) Equals(rhs MyEnum2) bool {
	return v == rhs
}

func (v MyEnum2) MarshalJSON() ([]byte, error) {
	switch int32(v) {
	case 12:
		return ([]byte)("\"X\""), nil
	case 34:
		return ([]byte)("\"Y\""), nil
	case 56:
		return ([]byte)("\"Z\""), nil
	}
	return ([]byte)(strconv.FormatInt(int64(v), 10)), nil
}

func (v *MyEnum2) UnmarshalJSON(text []byte) error {
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
			return fmt.Errorf("enum overflow from JSON %q for %q", text, "MyEnum2")
		}
		if x < math.MinInt32 {
			return fmt.Errorf("enum underflow from JSON %q for %q", text, "MyEnum2")
		}
		*v = (MyEnum2)(x)
		return nil
	case string:
		return v.UnmarshalText([]byte(w))
	default:
		return fmt.Errorf("invalid JSON value %q (%T) to unmarshal into %q", t, t, "MyEnum2")
	}
}

type StructCollision2 struct {
	CollisionField  bool   `json:"collisionField"`
	CollisionField2 string `json:"collision_field"`
}

func (v *StructCollision2) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	w, err = wire.NewValueBool(v.CollisionField), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	w, err = wire.NewValueString(v.CollisionField2), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 2, Value: w}
	i++
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *StructCollision2) FromWire(w wire.Value) error {
	var err error
	collisionFieldIsSet := false
	collision_fieldIsSet := false
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBool {
				v.CollisionField, err = field.Value.GetBool(), error(nil)
				if err != nil {
					return err
				}
				collisionFieldIsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TBinary {
				v.CollisionField2, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				collision_fieldIsSet = true
			}
		}
	}
	if !collisionFieldIsSet {
		return errors.New("field CollisionField of StructCollision2 is required")
	}
	if !collision_fieldIsSet {
		return errors.New("field CollisionField2 of StructCollision2 is required")
	}
	return nil
}

func (v *StructCollision2) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [2]string
	i := 0
	fields[i] = fmt.Sprintf("CollisionField: %v", v.CollisionField)
	i++
	fields[i] = fmt.Sprintf("CollisionField2: %v", v.CollisionField2)
	i++
	return fmt.Sprintf("StructCollision2{%v}", strings.Join(fields[:i], ", "))
}

func (v *StructCollision2) Equals(rhs *StructCollision2) bool {
	if !(v.CollisionField == rhs.CollisionField) {
		return false
	}
	if !(v.CollisionField2 == rhs.CollisionField2) {
		return false
	}
	return true
}

type UnionCollision2 struct {
	CollisionField  *bool   `json:"collisionField,omitempty"`
	CollisionField2 *string `json:"collision_field,omitempty"`
}

func (v *UnionCollision2) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.CollisionField != nil {
		w, err = wire.NewValueBool(*(v.CollisionField)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	if v.CollisionField2 != nil {
		w, err = wire.NewValueString(*(v.CollisionField2)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}
	if i != 1 {
		return wire.Value{}, fmt.Errorf("UnionCollision2 should have exactly one field: got %v fields", i)
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *UnionCollision2) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBool {
				var x bool
				x, err = field.Value.GetBool(), error(nil)
				v.CollisionField = &x
				if err != nil {
					return err
				}
			}
		case 2:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.CollisionField2 = &x
				if err != nil {
					return err
				}
			}
		}
	}
	count := 0
	if v.CollisionField != nil {
		count++
	}
	if v.CollisionField2 != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("UnionCollision2 should have exactly one field: got %v fields", count)
	}
	return nil
}

func (v *UnionCollision2) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [2]string
	i := 0
	if v.CollisionField != nil {
		fields[i] = fmt.Sprintf("CollisionField: %v", *(v.CollisionField))
		i++
	}
	if v.CollisionField2 != nil {
		fields[i] = fmt.Sprintf("CollisionField2: %v", *(v.CollisionField2))
		i++
	}
	return fmt.Sprintf("UnionCollision2{%v}", strings.Join(fields[:i], ", "))
}

func (v *UnionCollision2) Equals(rhs *UnionCollision2) bool {
	if !_Bool_EqualsPtr(v.CollisionField, rhs.CollisionField) {
		return false
	}
	if !_String_EqualsPtr(v.CollisionField2, rhs.CollisionField2) {
		return false
	}
	return true
}
