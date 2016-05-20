// Code generated by thriftrw

package keyvalue

import (
	"errors"
	"fmt"
	"github.com/thriftrw/thriftrw-go/wire"
	"strings"
)

type SizeArgs struct{}

func (v *SizeArgs) ToWire() (wire.Value, error) {
	var (
		fields [0]wire.Field
		i      int = 0
	)
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *SizeArgs) FromWire(w wire.Value) error {
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		}
	}
	return nil
}

func (v *SizeArgs) String() string {
	var fields [0]string
	i := 0
	return fmt.Sprintf("SizeArgs{%v}", strings.Join(fields[:i], ", "))
}

func (v *SizeArgs) MethodName() string {
	return "size"
}

func (v *SizeArgs) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

type SizeResult struct {
	Success *int64 `json:"success,omitempty"`
}

func (v *SizeResult) ToWire() (wire.Value, error) {
	count := 0
	if v.Success != nil {
		count++
	}
	if count != 1 {
		return wire.Value{}, fmt.Errorf("SizeResult should receive exactly one field value: received %v values", count)
	}
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Success != nil {
		w, err = wire.NewValueI64(*(v.Success)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *SizeResult) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TI64 {
				var x int64
				x, err = field.Value.GetI64(), error(nil)
				v.Success = &x
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
	if count != 1 {
		return fmt.Errorf("SizeResult should receive exactly one field value: received %v values", count)
	}
	return nil
}

func (v *SizeResult) String() string {
	var fields [1]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", *(v.Success))
		i++
	}
	return fmt.Sprintf("SizeResult{%v}", strings.Join(fields[:i], ", "))
}

func (v *SizeResult) MethodName() string {
	return "size"
}

func (v *SizeResult) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}

var SizeHelper = struct {
	IsException    func(error) bool
	Args           func() *SizeArgs
	WrapResponse   func(int64, error) (*SizeResult, error)
	UnwrapResponse func(*SizeResult) (int64, error)
}{}

func init() {
	SizeHelper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}
	SizeHelper.Args = func() *SizeArgs {
		return &SizeArgs{}
	}
	SizeHelper.WrapResponse = func(success int64, err error) (*SizeResult, error) {
		if err == nil {
			return &SizeResult{Success: &success}, nil
		}
		return nil, err
	}
	SizeHelper.UnwrapResponse = func(result *SizeResult) (success int64, err error) {
		if result.Success != nil {
			success = *result.Success
			return
		}
		err = errors.New("expected a non-void result")
		return
	}
}
