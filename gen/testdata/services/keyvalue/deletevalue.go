// Code generated by thriftrw

package keyvalue

import (
	"fmt"
	"github.com/thriftrw/thriftrw-go/gen/testdata/exceptions"
	"github.com/thriftrw/thriftrw-go/gen/testdata/services"
	"github.com/thriftrw/thriftrw-go/wire"
	"strings"
)

type DeleteValueArgs struct{ Key *services.Key }

func (v *DeleteValueArgs) ToWire() wire.Value {
	var fields [1]wire.Field
	i := 0
	if v.Key != nil {
		fields[i] = wire.Field{ID: 1, Value: v.Key.ToWire()}
		i++
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]})
}
func _Key_Read(w wire.Value) (services.Key, error) {
	var x services.Key
	err := x.FromWire(w)
	return x, err
}
func (v *DeleteValueArgs) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				var x services.Key
				x, err = _Key_Read(field.Value)
				v.Key = &x
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
func (v *DeleteValueArgs) String() string {
	var fields [1]string
	i := 0
	if v.Key != nil {
		fields[i] = fmt.Sprintf("Key: %v", *(v.Key))
		i++
	}
	return fmt.Sprintf("DeleteValueArgs{%v}", strings.Join(fields[:i], ", "))
}

type DeleteValueResult struct {
	DoesNotExist  *exceptions.DoesNotExistException
	InternalError *services.InternalError
}

func (v *DeleteValueResult) ToWire() wire.Value {
	var fields [2]wire.Field
	i := 0
	if v.DoesNotExist != nil {
		fields[i] = wire.Field{ID: 1, Value: v.DoesNotExist.ToWire()}
		i++
	}
	if v.InternalError != nil {
		fields[i] = wire.Field{ID: 2, Value: v.InternalError.ToWire()}
		i++
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]})
}
func _DoesNotExistException_Read(w wire.Value) (*exceptions.DoesNotExistException, error) {
	var v exceptions.DoesNotExistException
	err := v.FromWire(w)
	return &v, err
}
func _InternalError_Read(w wire.Value) (*services.InternalError, error) {
	var v services.InternalError
	err := v.FromWire(w)
	return &v, err
}
func (v *DeleteValueResult) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.DoesNotExist, err = _DoesNotExistException_Read(field.Value)
				if err != nil {
					return err
				}
			}
		case 2:
			if field.Value.Type() == wire.TStruct {
				v.InternalError, err = _InternalError_Read(field.Value)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
func (v *DeleteValueResult) String() string {
	var fields [2]string
	i := 0
	if v.DoesNotExist != nil {
		fields[i] = fmt.Sprintf("DoesNotExist: %v", v.DoesNotExist)
		i++
	}
	if v.InternalError != nil {
		fields[i] = fmt.Sprintf("InternalError: %v", v.InternalError)
		i++
	}
	return fmt.Sprintf("DeleteValueResult{%v}", strings.Join(fields[:i], ", "))
}

var DeleteValue = struct {
	IsException    func(error) bool
	Args           func(key *services.Key) *DeleteValueArgs
	WrapResponse   func(error) (*DeleteValueResult, error)
	UnwrapResponse func(*DeleteValueResult) error
}{}

func init() {
	DeleteValue.IsException = func(err error) bool {
		switch err.(type) {
		case *exceptions.DoesNotExistException:
			return true
		case *services.InternalError:
			return true
		default:
			return false
		}
	}
	DeleteValue.Args = func(key *services.Key) *DeleteValueArgs {
		return &DeleteValueArgs{Key: key}
	}
	DeleteValue.WrapResponse = func(err error) (*DeleteValueResult, error) {
		if err == nil {
			return &DeleteValueResult{}, nil
		}
		switch e := err.(type) {
		case *exceptions.DoesNotExistException:
			return &DeleteValueResult{DoesNotExist: e}, nil
		case *services.InternalError:
			return &DeleteValueResult{InternalError: e}, nil
		}
		return nil, err
	}
	DeleteValue.UnwrapResponse = func(result *DeleteValueResult) (err error) {
		if result.DoesNotExist != nil {
			err = result.DoesNotExist
			return
		}
		if result.InternalError != nil {
			err = result.InternalError
			return
		}
		return
	}
}
