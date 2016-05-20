// Code generated by thriftrw

package keyvalue

import (
	"errors"
	"fmt"
	"github.com/thriftrw/thriftrw-go/gen/testdata/exceptions"
	"github.com/thriftrw/thriftrw-go/gen/testdata/services"
	"github.com/thriftrw/thriftrw-go/gen/testdata/unions"
	"github.com/thriftrw/thriftrw-go/wire"
	"strings"
)

type GetManyValuesArgs struct {
	Range []services.Key `json:"range"`
}

type _List_Key_ValueList []services.Key

func (v _List_Key_ValueList) ForEach(f func(wire.Value) error) error {
	for _, x := range v {
		w, err := x.ToWire()
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

func (v _List_Key_ValueList) Close() {
}

func (v *GetManyValuesArgs) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Range != nil {
		w, err = wire.NewValueList(wire.List{ValueType: wire.TBinary, Size: len(v.Range), Items: _List_Key_ValueList(v.Range)}), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _List_Key_Read(l wire.List) ([]services.Key, error) {
	if l.ValueType != wire.TBinary {
		return nil, nil
	}
	o := make([]services.Key, 0, l.Size)
	err := l.Items.ForEach(func(x wire.Value) error {
		i, err := _Key_Read(x)
		if err != nil {
			return err
		}
		o = append(o, i)
		return nil
	})
	l.Items.Close()
	return o, err
}

func (v *GetManyValuesArgs) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TList {
				v.Range, err = _List_Key_Read(field.Value.GetList())
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (v *GetManyValuesArgs) String() string {
	var fields [1]string
	i := 0
	if v.Range != nil {
		fields[i] = fmt.Sprintf("Range: %v", v.Range)
		i++
	}
	return fmt.Sprintf("GetManyValuesArgs{%v}", strings.Join(fields[:i], ", "))
}

func (v *GetManyValuesArgs) MethodName() string {
	return "getManyValues"
}

func (v *GetManyValuesArgs) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

type GetManyValuesResult struct {
	Success      []*unions.ArbitraryValue          `json:"success"`
	DoesNotExist *exceptions.DoesNotExistException `json:"doesNotExist,omitempty"`
}

type _List_ArbitraryValue_ValueList []*unions.ArbitraryValue

func (v _List_ArbitraryValue_ValueList) ForEach(f func(wire.Value) error) error {
	for _, x := range v {
		w, err := x.ToWire()
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

func (v _List_ArbitraryValue_ValueList) Close() {
}

func (v *GetManyValuesResult) ToWire() (wire.Value, error) {
	count := 0
	if v.Success != nil {
		count++
	}
	if v.DoesNotExist != nil {
		count++
	}
	if count != 1 {
		return wire.Value{}, fmt.Errorf("GetManyValuesResult should receive exactly one field value: received %v values", count)
	}
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Success != nil {
		w, err = wire.NewValueList(wire.List{ValueType: wire.TStruct, Size: len(v.Success), Items: _List_ArbitraryValue_ValueList(v.Success)}), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}
	if v.DoesNotExist != nil {
		w, err = v.DoesNotExist.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _ArbitraryValue_Read(w wire.Value) (*unions.ArbitraryValue, error) {
	var v unions.ArbitraryValue
	err := v.FromWire(w)
	return &v, err
}

func _List_ArbitraryValue_Read(l wire.List) ([]*unions.ArbitraryValue, error) {
	if l.ValueType != wire.TStruct {
		return nil, nil
	}
	o := make([]*unions.ArbitraryValue, 0, l.Size)
	err := l.Items.ForEach(func(x wire.Value) error {
		i, err := _ArbitraryValue_Read(x)
		if err != nil {
			return err
		}
		o = append(o, i)
		return nil
	})
	l.Items.Close()
	return o, err
}

func (v *GetManyValuesResult) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TList {
				v.Success, err = _List_ArbitraryValue_Read(field.Value.GetList())
				if err != nil {
					return err
				}
			}
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.DoesNotExist, err = _DoesNotExistException_Read(field.Value)
				if err != nil {
					return err
				}
			}
		}
	}
	count := 0
	if v.Success != nil {
		count++
	}
	if v.DoesNotExist != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("GetManyValuesResult should receive exactly one field value: received %v values", count)
	}
	return nil
}

func (v *GetManyValuesResult) String() string {
	var fields [2]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}
	if v.DoesNotExist != nil {
		fields[i] = fmt.Sprintf("DoesNotExist: %v", v.DoesNotExist)
		i++
	}
	return fmt.Sprintf("GetManyValuesResult{%v}", strings.Join(fields[:i], ", "))
}

func (v *GetManyValuesResult) MethodName() string {
	return "getManyValues"
}

func (v *GetManyValuesResult) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}

var GetManyValuesHelper = struct {
	IsException    func(error) bool
	Args           func(range2 []services.Key) *GetManyValuesArgs
	WrapResponse   func([]*unions.ArbitraryValue, error) (*GetManyValuesResult, error)
	UnwrapResponse func(*GetManyValuesResult) ([]*unions.ArbitraryValue, error)
}{}

func init() {
	GetManyValuesHelper.IsException = func(err error) bool {
		switch err.(type) {
		case *exceptions.DoesNotExistException:
			return true
		default:
			return false
		}
	}
	GetManyValuesHelper.Args = func(range2 []services.Key) *GetManyValuesArgs {
		return &GetManyValuesArgs{Range: range2}
	}
	GetManyValuesHelper.WrapResponse = func(success []*unions.ArbitraryValue, err error) (*GetManyValuesResult, error) {
		if err == nil {
			return &GetManyValuesResult{Success: success}, nil
		}
		switch e := err.(type) {
		case *exceptions.DoesNotExistException:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for GetManyValuesResult.DoesNotExist")
			}
			return &GetManyValuesResult{DoesNotExist: e}, nil
		}
		return nil, err
	}
	GetManyValuesHelper.UnwrapResponse = func(result *GetManyValuesResult) (success []*unions.ArbitraryValue, err error) {
		if result.DoesNotExist != nil {
			err = result.DoesNotExist
			return
		}
		if result.Success != nil {
			success = result.Success
			return
		}
		err = errors.New("expected a non-void result")
		return
	}
}
