// Code generated by thriftrw
// @generated

package cache

import (
	"fmt"
	"go.uber.org/thriftrw/wire"
	"strings"
)

type ClearArgs struct{}

func (v *ClearArgs) ToWire() (wire.Value, error) {
	var (
		fields [0]wire.Field
		i      int = 0
	)
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *ClearArgs) FromWire(w wire.Value) error {
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		}
	}
	return nil
}

func (v *ClearArgs) String() string {
	var fields [0]string
	i := 0
	return fmt.Sprintf("ClearArgs{%v}", strings.Join(fields[:i], ", "))
}

func (v *ClearArgs) MethodName() string {
	return "clear"
}

func (v *ClearArgs) EnvelopeType() wire.EnvelopeType {
	return wire.OneWay
}

var ClearHelper = struct{ Args func() *ClearArgs }{}

func init() {
	ClearHelper.Args = func() *ClearArgs {
		return &ClearArgs{}
	}
}
