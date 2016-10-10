// Code generated by thriftrw
// @generated

// Copyright (c) 2016 Uber Technologies, Inc.
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

package api

import (
	"errors"
	"fmt"
	"go.uber.org/thriftrw/wire"
	"strings"
)

type Argument struct {
	Name string `json:"name"`
	Type *Type  `json:"type"`
}

func (v *Argument) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	w, err = wire.NewValueString(v.Name), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	if v.Type == nil {
		return w, errors.New("field Type of Argument is required")
	}
	w, err = v.Type.ToWire()
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 2, Value: w}
	i++
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _Type_Read(w wire.Value) (*Type, error) {
	var v Type
	err := v.FromWire(w)
	return &v, err
}

func (v *Argument) FromWire(w wire.Value) error {
	var err error
	nameIsSet := false
	typeIsSet := false
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.Name, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				nameIsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TStruct {
				v.Type, err = _Type_Read(field.Value)
				if err != nil {
					return err
				}
				typeIsSet = true
			}
		}
	}
	if !nameIsSet {
		return errors.New("field Name of Argument is required")
	}
	if !typeIsSet {
		return errors.New("field Type of Argument is required")
	}
	return nil
}

func (v *Argument) String() string {
	var fields [2]string
	i := 0
	fields[i] = fmt.Sprintf("Name: %v", v.Name)
	i++
	fields[i] = fmt.Sprintf("Type: %v", v.Type)
	i++
	return fmt.Sprintf("Argument{%v}", strings.Join(fields[:i], ", "))
}

type Feature int32

const (
	FeatureServiceGenerator Feature = 1
)

func (v Feature) ToWire() (wire.Value, error) {
	return wire.NewValueI32(int32(v)), nil
}

func (v *Feature) FromWire(w wire.Value) error {
	*v = (Feature)(w.GetI32())
	return nil
}

func (v Feature) String() string {
	w := int32(v)
	switch w {
	case 1:
		return "SERVICE_GENERATOR"
	}
	return fmt.Sprintf("Feature(%d)", w)
}

type Function struct {
	Name       string      `json:"name"`
	ThriftName string      `json:"thriftName"`
	Arguments  []*Argument `json:"arguments"`
	ReturnType *Type       `json:"returnType,omitempty"`
	Exceptions []*Argument `json:"exceptions"`
	OneWay     *bool       `json:"oneWay,omitempty"`
}

type _List_Argument_ValueList []*Argument

func (v _List_Argument_ValueList) ForEach(f func(wire.Value) error) error {
	for i, x := range v {
		if x == nil {
			return fmt.Errorf("invalid [%v]: value is nil", i)
		}
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

func (v _List_Argument_ValueList) Size() int {
	return len(v)
}

func (_List_Argument_ValueList) ValueType() wire.Type {
	return wire.TStruct
}

func (_List_Argument_ValueList) Close() {
}

func (v *Function) ToWire() (wire.Value, error) {
	var (
		fields [6]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	w, err = wire.NewValueString(v.Name), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	w, err = wire.NewValueString(v.ThriftName), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 2, Value: w}
	i++
	if v.Arguments == nil {
		return w, errors.New("field Arguments of Function is required")
	}
	w, err = wire.NewValueList(_List_Argument_ValueList(v.Arguments)), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 3, Value: w}
	i++
	if v.ReturnType != nil {
		w, err = v.ReturnType.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 4, Value: w}
		i++
	}
	if v.Exceptions != nil {
		w, err = wire.NewValueList(_List_Argument_ValueList(v.Exceptions)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 5, Value: w}
		i++
	}
	if v.OneWay != nil {
		w, err = wire.NewValueBool(*(v.OneWay)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 6, Value: w}
		i++
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _Argument_Read(w wire.Value) (*Argument, error) {
	var v Argument
	err := v.FromWire(w)
	return &v, err
}

func _List_Argument_Read(l wire.ValueList) ([]*Argument, error) {
	if l.ValueType() != wire.TStruct {
		return nil, nil
	}
	o := make([]*Argument, 0, l.Size())
	err := l.ForEach(func(x wire.Value) error {
		i, err := _Argument_Read(x)
		if err != nil {
			return err
		}
		o = append(o, i)
		return nil
	})
	l.Close()
	return o, err
}

func (v *Function) FromWire(w wire.Value) error {
	var err error
	nameIsSet := false
	thriftNameIsSet := false
	argumentsIsSet := false
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.Name, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				nameIsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TBinary {
				v.ThriftName, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				thriftNameIsSet = true
			}
		case 3:
			if field.Value.Type() == wire.TList {
				v.Arguments, err = _List_Argument_Read(field.Value.GetList())
				if err != nil {
					return err
				}
				argumentsIsSet = true
			}
		case 4:
			if field.Value.Type() == wire.TStruct {
				v.ReturnType, err = _Type_Read(field.Value)
				if err != nil {
					return err
				}
			}
		case 5:
			if field.Value.Type() == wire.TList {
				v.Exceptions, err = _List_Argument_Read(field.Value.GetList())
				if err != nil {
					return err
				}
			}
		case 6:
			if field.Value.Type() == wire.TBool {
				var x bool
				x, err = field.Value.GetBool(), error(nil)
				v.OneWay = &x
				if err != nil {
					return err
				}
			}
		}
	}
	if !nameIsSet {
		return errors.New("field Name of Function is required")
	}
	if !thriftNameIsSet {
		return errors.New("field ThriftName of Function is required")
	}
	if !argumentsIsSet {
		return errors.New("field Arguments of Function is required")
	}
	return nil
}

func (v *Function) String() string {
	var fields [6]string
	i := 0
	fields[i] = fmt.Sprintf("Name: %v", v.Name)
	i++
	fields[i] = fmt.Sprintf("ThriftName: %v", v.ThriftName)
	i++
	fields[i] = fmt.Sprintf("Arguments: %v", v.Arguments)
	i++
	if v.ReturnType != nil {
		fields[i] = fmt.Sprintf("ReturnType: %v", v.ReturnType)
		i++
	}
	if v.Exceptions != nil {
		fields[i] = fmt.Sprintf("Exceptions: %v", v.Exceptions)
		i++
	}
	if v.OneWay != nil {
		fields[i] = fmt.Sprintf("OneWay: %v", *(v.OneWay))
		i++
	}
	return fmt.Sprintf("Function{%v}", strings.Join(fields[:i], ", "))
}

type GenerateServiceRequest struct {
	RootServices []ServiceID            `json:"rootServices"`
	Services     map[ServiceID]*Service `json:"services"`
	Modules      map[ModuleID]*Module   `json:"modules"`
}

type _List_ServiceID_ValueList []ServiceID

func (v _List_ServiceID_ValueList) ForEach(f func(wire.Value) error) error {
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

func (v _List_ServiceID_ValueList) Size() int {
	return len(v)
}

func (_List_ServiceID_ValueList) ValueType() wire.Type {
	return wire.TI32
}

func (_List_ServiceID_ValueList) Close() {
}

type _Map_ServiceID_Service_MapItemList map[ServiceID]*Service

func (m _Map_ServiceID_Service_MapItemList) ForEach(f func(wire.MapItem) error) error {
	for k, v := range m {
		if v == nil {
			return fmt.Errorf("invalid [%v]: value is nil", k)
		}
		kw, err := k.ToWire()
		if err != nil {
			return err
		}
		vw, err := v.ToWire()
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

func (m _Map_ServiceID_Service_MapItemList) Size() int {
	return len(m)
}

func (_Map_ServiceID_Service_MapItemList) KeyType() wire.Type {
	return wire.TI32
}

func (_Map_ServiceID_Service_MapItemList) ValueType() wire.Type {
	return wire.TStruct
}

func (_Map_ServiceID_Service_MapItemList) Close() {
}

type _Map_ModuleID_Module_MapItemList map[ModuleID]*Module

func (m _Map_ModuleID_Module_MapItemList) ForEach(f func(wire.MapItem) error) error {
	for k, v := range m {
		if v == nil {
			return fmt.Errorf("invalid [%v]: value is nil", k)
		}
		kw, err := k.ToWire()
		if err != nil {
			return err
		}
		vw, err := v.ToWire()
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

func (m _Map_ModuleID_Module_MapItemList) Size() int {
	return len(m)
}

func (_Map_ModuleID_Module_MapItemList) KeyType() wire.Type {
	return wire.TI32
}

func (_Map_ModuleID_Module_MapItemList) ValueType() wire.Type {
	return wire.TStruct
}

func (_Map_ModuleID_Module_MapItemList) Close() {
}

func (v *GenerateServiceRequest) ToWire() (wire.Value, error) {
	var (
		fields [3]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.RootServices == nil {
		return w, errors.New("field RootServices of GenerateServiceRequest is required")
	}
	w, err = wire.NewValueList(_List_ServiceID_ValueList(v.RootServices)), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	if v.Services == nil {
		return w, errors.New("field Services of GenerateServiceRequest is required")
	}
	w, err = wire.NewValueMap(_Map_ServiceID_Service_MapItemList(v.Services)), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 2, Value: w}
	i++
	if v.Modules == nil {
		return w, errors.New("field Modules of GenerateServiceRequest is required")
	}
	w, err = wire.NewValueMap(_Map_ModuleID_Module_MapItemList(v.Modules)), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 3, Value: w}
	i++
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _ServiceID_Read(w wire.Value) (ServiceID, error) {
	var x ServiceID
	err := x.FromWire(w)
	return x, err
}

func _List_ServiceID_Read(l wire.ValueList) ([]ServiceID, error) {
	if l.ValueType() != wire.TI32 {
		return nil, nil
	}
	o := make([]ServiceID, 0, l.Size())
	err := l.ForEach(func(x wire.Value) error {
		i, err := _ServiceID_Read(x)
		if err != nil {
			return err
		}
		o = append(o, i)
		return nil
	})
	l.Close()
	return o, err
}

func _Service_Read(w wire.Value) (*Service, error) {
	var v Service
	err := v.FromWire(w)
	return &v, err
}

func _Map_ServiceID_Service_Read(m wire.MapItemList) (map[ServiceID]*Service, error) {
	if m.KeyType() != wire.TI32 {
		return nil, nil
	}
	if m.ValueType() != wire.TStruct {
		return nil, nil
	}
	o := make(map[ServiceID]*Service, m.Size())
	err := m.ForEach(func(x wire.MapItem) error {
		k, err := _ServiceID_Read(x.Key)
		if err != nil {
			return err
		}
		v, err := _Service_Read(x.Value)
		if err != nil {
			return err
		}
		o[k] = v
		return nil
	})
	m.Close()
	return o, err
}

func _ModuleID_Read(w wire.Value) (ModuleID, error) {
	var x ModuleID
	err := x.FromWire(w)
	return x, err
}

func _Module_Read(w wire.Value) (*Module, error) {
	var v Module
	err := v.FromWire(w)
	return &v, err
}

func _Map_ModuleID_Module_Read(m wire.MapItemList) (map[ModuleID]*Module, error) {
	if m.KeyType() != wire.TI32 {
		return nil, nil
	}
	if m.ValueType() != wire.TStruct {
		return nil, nil
	}
	o := make(map[ModuleID]*Module, m.Size())
	err := m.ForEach(func(x wire.MapItem) error {
		k, err := _ModuleID_Read(x.Key)
		if err != nil {
			return err
		}
		v, err := _Module_Read(x.Value)
		if err != nil {
			return err
		}
		o[k] = v
		return nil
	})
	m.Close()
	return o, err
}

func (v *GenerateServiceRequest) FromWire(w wire.Value) error {
	var err error
	rootServicesIsSet := false
	servicesIsSet := false
	modulesIsSet := false
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TList {
				v.RootServices, err = _List_ServiceID_Read(field.Value.GetList())
				if err != nil {
					return err
				}
				rootServicesIsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TMap {
				v.Services, err = _Map_ServiceID_Service_Read(field.Value.GetMap())
				if err != nil {
					return err
				}
				servicesIsSet = true
			}
		case 3:
			if field.Value.Type() == wire.TMap {
				v.Modules, err = _Map_ModuleID_Module_Read(field.Value.GetMap())
				if err != nil {
					return err
				}
				modulesIsSet = true
			}
		}
	}
	if !rootServicesIsSet {
		return errors.New("field RootServices of GenerateServiceRequest is required")
	}
	if !servicesIsSet {
		return errors.New("field Services of GenerateServiceRequest is required")
	}
	if !modulesIsSet {
		return errors.New("field Modules of GenerateServiceRequest is required")
	}
	return nil
}

func (v *GenerateServiceRequest) String() string {
	var fields [3]string
	i := 0
	fields[i] = fmt.Sprintf("RootServices: %v", v.RootServices)
	i++
	fields[i] = fmt.Sprintf("Services: %v", v.Services)
	i++
	fields[i] = fmt.Sprintf("Modules: %v", v.Modules)
	i++
	return fmt.Sprintf("GenerateServiceRequest{%v}", strings.Join(fields[:i], ", "))
}

type GenerateServiceResponse struct {
	Files map[string][]byte `json:"files"`
}

type _Map_String_Binary_MapItemList map[string][]byte

func (m _Map_String_Binary_MapItemList) ForEach(f func(wire.MapItem) error) error {
	for k, v := range m {
		if v == nil {
			return fmt.Errorf("invalid [%v]: value is nil", k)
		}
		kw, err := wire.NewValueString(k), error(nil)
		if err != nil {
			return err
		}
		vw, err := wire.NewValueBinary(v), error(nil)
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

func (m _Map_String_Binary_MapItemList) Size() int {
	return len(m)
}

func (_Map_String_Binary_MapItemList) KeyType() wire.Type {
	return wire.TBinary
}

func (_Map_String_Binary_MapItemList) ValueType() wire.Type {
	return wire.TBinary
}

func (_Map_String_Binary_MapItemList) Close() {
}

func (v *GenerateServiceResponse) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Files != nil {
		w, err = wire.NewValueMap(_Map_String_Binary_MapItemList(v.Files)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _Map_String_Binary_Read(m wire.MapItemList) (map[string][]byte, error) {
	if m.KeyType() != wire.TBinary {
		return nil, nil
	}
	if m.ValueType() != wire.TBinary {
		return nil, nil
	}
	o := make(map[string][]byte, m.Size())
	err := m.ForEach(func(x wire.MapItem) error {
		k, err := x.Key.GetString(), error(nil)
		if err != nil {
			return err
		}
		v, err := x.Value.GetBinary(), error(nil)
		if err != nil {
			return err
		}
		o[k] = v
		return nil
	})
	m.Close()
	return o, err
}

func (v *GenerateServiceResponse) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TMap {
				v.Files, err = _Map_String_Binary_Read(field.Value.GetMap())
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (v *GenerateServiceResponse) String() string {
	var fields [1]string
	i := 0
	if v.Files != nil {
		fields[i] = fmt.Sprintf("Files: %v", v.Files)
		i++
	}
	return fmt.Sprintf("GenerateServiceResponse{%v}", strings.Join(fields[:i], ", "))
}

type HandshakeRequest struct{}

func (v *HandshakeRequest) ToWire() (wire.Value, error) {
	var (
		fields [0]wire.Field
		i      int = 0
	)
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *HandshakeRequest) FromWire(w wire.Value) error {
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		}
	}
	return nil
}

func (v *HandshakeRequest) String() string {
	var fields [0]string
	i := 0
	return fmt.Sprintf("HandshakeRequest{%v}", strings.Join(fields[:i], ", "))
}

type HandshakeResponse struct {
	Name       string    `json:"name"`
	ApiVersion int32     `json:"apiVersion"`
	Features   []Feature `json:"features"`
}

type _List_Feature_ValueList []Feature

func (v _List_Feature_ValueList) ForEach(f func(wire.Value) error) error {
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

func (v _List_Feature_ValueList) Size() int {
	return len(v)
}

func (_List_Feature_ValueList) ValueType() wire.Type {
	return wire.TI32
}

func (_List_Feature_ValueList) Close() {
}

func (v *HandshakeResponse) ToWire() (wire.Value, error) {
	var (
		fields [3]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	w, err = wire.NewValueString(v.Name), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	w, err = wire.NewValueI32(v.ApiVersion), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 2, Value: w}
	i++
	if v.Features == nil {
		return w, errors.New("field Features of HandshakeResponse is required")
	}
	w, err = wire.NewValueList(_List_Feature_ValueList(v.Features)), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 3, Value: w}
	i++
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _Feature_Read(w wire.Value) (Feature, error) {
	var v Feature
	err := v.FromWire(w)
	return v, err
}

func _List_Feature_Read(l wire.ValueList) ([]Feature, error) {
	if l.ValueType() != wire.TI32 {
		return nil, nil
	}
	o := make([]Feature, 0, l.Size())
	err := l.ForEach(func(x wire.Value) error {
		i, err := _Feature_Read(x)
		if err != nil {
			return err
		}
		o = append(o, i)
		return nil
	})
	l.Close()
	return o, err
}

func (v *HandshakeResponse) FromWire(w wire.Value) error {
	var err error
	nameIsSet := false
	apiVersionIsSet := false
	featuresIsSet := false
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.Name, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				nameIsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TI32 {
				v.ApiVersion, err = field.Value.GetI32(), error(nil)
				if err != nil {
					return err
				}
				apiVersionIsSet = true
			}
		case 3:
			if field.Value.Type() == wire.TList {
				v.Features, err = _List_Feature_Read(field.Value.GetList())
				if err != nil {
					return err
				}
				featuresIsSet = true
			}
		}
	}
	if !nameIsSet {
		return errors.New("field Name of HandshakeResponse is required")
	}
	if !apiVersionIsSet {
		return errors.New("field ApiVersion of HandshakeResponse is required")
	}
	if !featuresIsSet {
		return errors.New("field Features of HandshakeResponse is required")
	}
	return nil
}

func (v *HandshakeResponse) String() string {
	var fields [3]string
	i := 0
	fields[i] = fmt.Sprintf("Name: %v", v.Name)
	i++
	fields[i] = fmt.Sprintf("ApiVersion: %v", v.ApiVersion)
	i++
	fields[i] = fmt.Sprintf("Features: %v", v.Features)
	i++
	return fmt.Sprintf("HandshakeResponse{%v}", strings.Join(fields[:i], ", "))
}

type Module struct {
	ImportPath string `json:"importPath"`
	Directory  string `json:"directory"`
}

func (v *Module) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	w, err = wire.NewValueString(v.ImportPath), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	w, err = wire.NewValueString(v.Directory), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 2, Value: w}
	i++
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *Module) FromWire(w wire.Value) error {
	var err error
	importPathIsSet := false
	directoryIsSet := false
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.ImportPath, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				importPathIsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TBinary {
				v.Directory, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				directoryIsSet = true
			}
		}
	}
	if !importPathIsSet {
		return errors.New("field ImportPath of Module is required")
	}
	if !directoryIsSet {
		return errors.New("field Directory of Module is required")
	}
	return nil
}

func (v *Module) String() string {
	var fields [2]string
	i := 0
	fields[i] = fmt.Sprintf("ImportPath: %v", v.ImportPath)
	i++
	fields[i] = fmt.Sprintf("Directory: %v", v.Directory)
	i++
	return fmt.Sprintf("Module{%v}", strings.Join(fields[:i], ", "))
}

type ModuleID int32

func (v ModuleID) ToWire() (wire.Value, error) {
	x := (int32)(v)
	return wire.NewValueI32(x), error(nil)
}

func (v ModuleID) String() string {
	x := (int32)(v)
	return fmt.Sprint(x)
}

func (v *ModuleID) FromWire(w wire.Value) error {
	x, err := w.GetI32(), error(nil)
	*v = (ModuleID)(x)
	return err
}

type Service struct {
	Name       string      `json:"name"`
	ImportPath string      `json:"importPath"`
	Directory  string      `json:"directory"`
	ParentID   *ServiceID  `json:"parentID,omitempty"`
	Functions  []*Function `json:"functions"`
	ModuleID   ModuleID    `json:"moduleID"`
}

type _List_Function_ValueList []*Function

func (v _List_Function_ValueList) ForEach(f func(wire.Value) error) error {
	for i, x := range v {
		if x == nil {
			return fmt.Errorf("invalid [%v]: value is nil", i)
		}
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

func (v _List_Function_ValueList) Size() int {
	return len(v)
}

func (_List_Function_ValueList) ValueType() wire.Type {
	return wire.TStruct
}

func (_List_Function_ValueList) Close() {
}

func (v *Service) ToWire() (wire.Value, error) {
	var (
		fields [6]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	w, err = wire.NewValueString(v.Name), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	w, err = wire.NewValueString(v.ImportPath), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 2, Value: w}
	i++
	w, err = wire.NewValueString(v.Directory), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 3, Value: w}
	i++
	if v.ParentID != nil {
		w, err = v.ParentID.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 4, Value: w}
		i++
	}
	if v.Functions == nil {
		return w, errors.New("field Functions of Service is required")
	}
	w, err = wire.NewValueList(_List_Function_ValueList(v.Functions)), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 5, Value: w}
	i++
	w, err = v.ModuleID.ToWire()
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 6, Value: w}
	i++
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _Function_Read(w wire.Value) (*Function, error) {
	var v Function
	err := v.FromWire(w)
	return &v, err
}

func _List_Function_Read(l wire.ValueList) ([]*Function, error) {
	if l.ValueType() != wire.TStruct {
		return nil, nil
	}
	o := make([]*Function, 0, l.Size())
	err := l.ForEach(func(x wire.Value) error {
		i, err := _Function_Read(x)
		if err != nil {
			return err
		}
		o = append(o, i)
		return nil
	})
	l.Close()
	return o, err
}

func (v *Service) FromWire(w wire.Value) error {
	var err error
	nameIsSet := false
	importPathIsSet := false
	directoryIsSet := false
	functionsIsSet := false
	moduleIDIsSet := false
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.Name, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				nameIsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TBinary {
				v.ImportPath, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				importPathIsSet = true
			}
		case 3:
			if field.Value.Type() == wire.TBinary {
				v.Directory, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				directoryIsSet = true
			}
		case 4:
			if field.Value.Type() == wire.TI32 {
				var x ServiceID
				x, err = _ServiceID_Read(field.Value)
				v.ParentID = &x
				if err != nil {
					return err
				}
			}
		case 5:
			if field.Value.Type() == wire.TList {
				v.Functions, err = _List_Function_Read(field.Value.GetList())
				if err != nil {
					return err
				}
				functionsIsSet = true
			}
		case 6:
			if field.Value.Type() == wire.TI32 {
				v.ModuleID, err = _ModuleID_Read(field.Value)
				if err != nil {
					return err
				}
				moduleIDIsSet = true
			}
		}
	}
	if !nameIsSet {
		return errors.New("field Name of Service is required")
	}
	if !importPathIsSet {
		return errors.New("field ImportPath of Service is required")
	}
	if !directoryIsSet {
		return errors.New("field Directory of Service is required")
	}
	if !functionsIsSet {
		return errors.New("field Functions of Service is required")
	}
	if !moduleIDIsSet {
		return errors.New("field ModuleID of Service is required")
	}
	return nil
}

func (v *Service) String() string {
	var fields [6]string
	i := 0
	fields[i] = fmt.Sprintf("Name: %v", v.Name)
	i++
	fields[i] = fmt.Sprintf("ImportPath: %v", v.ImportPath)
	i++
	fields[i] = fmt.Sprintf("Directory: %v", v.Directory)
	i++
	if v.ParentID != nil {
		fields[i] = fmt.Sprintf("ParentID: %v", *(v.ParentID))
		i++
	}
	fields[i] = fmt.Sprintf("Functions: %v", v.Functions)
	i++
	fields[i] = fmt.Sprintf("ModuleID: %v", v.ModuleID)
	i++
	return fmt.Sprintf("Service{%v}", strings.Join(fields[:i], ", "))
}

type ServiceID int32

func (v ServiceID) ToWire() (wire.Value, error) {
	x := (int32)(v)
	return wire.NewValueI32(x), error(nil)
}

func (v ServiceID) String() string {
	x := (int32)(v)
	return fmt.Sprint(x)
}

func (v *ServiceID) FromWire(w wire.Value) error {
	x, err := w.GetI32(), error(nil)
	*v = (ServiceID)(x)
	return err
}

type SimpleType int32

const (
	SimpleTypeBool        SimpleType = 1
	SimpleTypeByte        SimpleType = 2
	SimpleTypeInt8        SimpleType = 3
	SimpleTypeInt16       SimpleType = 4
	SimpleTypeInt32       SimpleType = 5
	SimpleTypeInt64       SimpleType = 6
	SimpleTypeFloat64     SimpleType = 7
	SimpleTypeString      SimpleType = 8
	SimpleTypeStructEmpty SimpleType = 9
)

func (v SimpleType) ToWire() (wire.Value, error) {
	return wire.NewValueI32(int32(v)), nil
}

func (v *SimpleType) FromWire(w wire.Value) error {
	*v = (SimpleType)(w.GetI32())
	return nil
}

func (v SimpleType) String() string {
	w := int32(v)
	switch w {
	case 1:
		return "BOOL"
	case 2:
		return "BYTE"
	case 3:
		return "INT8"
	case 4:
		return "INT16"
	case 5:
		return "INT32"
	case 6:
		return "INT64"
	case 7:
		return "FLOAT64"
	case 8:
		return "STRING"
	case 9:
		return "STRUCT_EMPTY"
	}
	return fmt.Sprintf("SimpleType(%d)", w)
}

type Type struct {
	SimpleType        *SimpleType    `json:"simpleType,omitempty"`
	SliceType         *Type          `json:"sliceType,omitempty"`
	KeyValueSliceType *TypePair      `json:"keyValueSliceType,omitempty"`
	MapType           *TypePair      `json:"mapType,omitempty"`
	ReferenceType     *TypeReference `json:"referenceType,omitempty"`
	PointerType       *Type          `json:"pointerType,omitempty"`
}

func (v *Type) ToWire() (wire.Value, error) {
	var (
		fields [6]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.SimpleType != nil {
		w, err = v.SimpleType.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	if v.SliceType != nil {
		w, err = v.SliceType.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}
	if v.KeyValueSliceType != nil {
		w, err = v.KeyValueSliceType.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 3, Value: w}
		i++
	}
	if v.MapType != nil {
		w, err = v.MapType.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 4, Value: w}
		i++
	}
	if v.ReferenceType != nil {
		w, err = v.ReferenceType.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 5, Value: w}
		i++
	}
	if v.PointerType != nil {
		w, err = v.PointerType.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 6, Value: w}
		i++
	}
	if i != 1 {
		return wire.Value{}, fmt.Errorf("Type should have exactly one field: got %v fields", i)
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _SimpleType_Read(w wire.Value) (SimpleType, error) {
	var v SimpleType
	err := v.FromWire(w)
	return v, err
}

func _TypePair_Read(w wire.Value) (*TypePair, error) {
	var v TypePair
	err := v.FromWire(w)
	return &v, err
}

func _TypeReference_Read(w wire.Value) (*TypeReference, error) {
	var v TypeReference
	err := v.FromWire(w)
	return &v, err
}

func (v *Type) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TI32 {
				var x SimpleType
				x, err = _SimpleType_Read(field.Value)
				v.SimpleType = &x
				if err != nil {
					return err
				}
			}
		case 2:
			if field.Value.Type() == wire.TStruct {
				v.SliceType, err = _Type_Read(field.Value)
				if err != nil {
					return err
				}
			}
		case 3:
			if field.Value.Type() == wire.TStruct {
				v.KeyValueSliceType, err = _TypePair_Read(field.Value)
				if err != nil {
					return err
				}
			}
		case 4:
			if field.Value.Type() == wire.TStruct {
				v.MapType, err = _TypePair_Read(field.Value)
				if err != nil {
					return err
				}
			}
		case 5:
			if field.Value.Type() == wire.TStruct {
				v.ReferenceType, err = _TypeReference_Read(field.Value)
				if err != nil {
					return err
				}
			}
		case 6:
			if field.Value.Type() == wire.TStruct {
				v.PointerType, err = _Type_Read(field.Value)
				if err != nil {
					return err
				}
			}
		}
	}
	count := 0
	if v.SimpleType != nil {
		count++
	}
	if v.SliceType != nil {
		count++
	}
	if v.KeyValueSliceType != nil {
		count++
	}
	if v.MapType != nil {
		count++
	}
	if v.ReferenceType != nil {
		count++
	}
	if v.PointerType != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("Type should have exactly one field: got %v fields", count)
	}
	return nil
}

func (v *Type) String() string {
	var fields [6]string
	i := 0
	if v.SimpleType != nil {
		fields[i] = fmt.Sprintf("SimpleType: %v", *(v.SimpleType))
		i++
	}
	if v.SliceType != nil {
		fields[i] = fmt.Sprintf("SliceType: %v", v.SliceType)
		i++
	}
	if v.KeyValueSliceType != nil {
		fields[i] = fmt.Sprintf("KeyValueSliceType: %v", v.KeyValueSliceType)
		i++
	}
	if v.MapType != nil {
		fields[i] = fmt.Sprintf("MapType: %v", v.MapType)
		i++
	}
	if v.ReferenceType != nil {
		fields[i] = fmt.Sprintf("ReferenceType: %v", v.ReferenceType)
		i++
	}
	if v.PointerType != nil {
		fields[i] = fmt.Sprintf("PointerType: %v", v.PointerType)
		i++
	}
	return fmt.Sprintf("Type{%v}", strings.Join(fields[:i], ", "))
}

type TypePair struct {
	Left  *Type `json:"left"`
	Right *Type `json:"right"`
}

func (v *TypePair) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Left == nil {
		return w, errors.New("field Left of TypePair is required")
	}
	w, err = v.Left.ToWire()
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	if v.Right == nil {
		return w, errors.New("field Right of TypePair is required")
	}
	w, err = v.Right.ToWire()
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 2, Value: w}
	i++
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *TypePair) FromWire(w wire.Value) error {
	var err error
	leftIsSet := false
	rightIsSet := false
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.Left, err = _Type_Read(field.Value)
				if err != nil {
					return err
				}
				leftIsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TStruct {
				v.Right, err = _Type_Read(field.Value)
				if err != nil {
					return err
				}
				rightIsSet = true
			}
		}
	}
	if !leftIsSet {
		return errors.New("field Left of TypePair is required")
	}
	if !rightIsSet {
		return errors.New("field Right of TypePair is required")
	}
	return nil
}

func (v *TypePair) String() string {
	var fields [2]string
	i := 0
	fields[i] = fmt.Sprintf("Left: %v", v.Left)
	i++
	fields[i] = fmt.Sprintf("Right: %v", v.Right)
	i++
	return fmt.Sprintf("TypePair{%v}", strings.Join(fields[:i], ", "))
}

type TypeReference struct {
	Name       string `json:"name"`
	ImportPath string `json:"importPath"`
}

func (v *TypeReference) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	w, err = wire.NewValueString(v.Name), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	w, err = wire.NewValueString(v.ImportPath), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 2, Value: w}
	i++
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *TypeReference) FromWire(w wire.Value) error {
	var err error
	nameIsSet := false
	importPathIsSet := false
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.Name, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				nameIsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TBinary {
				v.ImportPath, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				importPathIsSet = true
			}
		}
	}
	if !nameIsSet {
		return errors.New("field Name of TypeReference is required")
	}
	if !importPathIsSet {
		return errors.New("field ImportPath of TypeReference is required")
	}
	return nil
}

func (v *TypeReference) String() string {
	var fields [2]string
	i := 0
	fields[i] = fmt.Sprintf("Name: %v", v.Name)
	i++
	fields[i] = fmt.Sprintf("ImportPath: %v", v.ImportPath)
	i++
	return fmt.Sprintf("TypeReference{%v}", strings.Join(fields[:i], ", "))
}
