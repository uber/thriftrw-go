// Copyright (c) 2015 Uber Technologies, Inc.
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

package gen

import (
	"encoding"
	"encoding/json"
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"
	"time"

	tl "go.uber.org/thriftrw/gen/internal/tests/collision"
	tc "go.uber.org/thriftrw/gen/internal/tests/containers"
	tle "go.uber.org/thriftrw/gen/internal/tests/enum_conflict"
	te "go.uber.org/thriftrw/gen/internal/tests/enums"
	tx "go.uber.org/thriftrw/gen/internal/tests/exceptions"
	tz "go.uber.org/thriftrw/gen/internal/tests/nozap"
	tf "go.uber.org/thriftrw/gen/internal/tests/services"
	ts "go.uber.org/thriftrw/gen/internal/tests/structs"
	td "go.uber.org/thriftrw/gen/internal/tests/typedefs"
	tu "go.uber.org/thriftrw/gen/internal/tests/unions"
	tul "go.uber.org/thriftrw/gen/internal/tests/uuid_conflict"
	"go.uber.org/thriftrw/wire"
	"go.uber.org/zap/zapcore"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func thriftTypeIsValid(v thriftType) bool {
	// TODO(abg): ToWire + EvaluateValue to validate here means we end
	// up serializing this value twice. We may want to include a
	// Validate method on generated types.

	// We validate while serializing.
	w, err := v.ToWire()
	if err != nil {
		return false
	}

	// Because we evaluate collections lazily, validation issues with items in
	// them won't be known until we try to serialize it or explicitly evaluate
	// the lazy lists with wire.EvaluateValue.
	if err := wire.EvaluateValue(w); err != nil {
		return false
	}

	return true
}

func defaultValueGenerator(typ reflect.Type) func(*testing.T, *rand.Rand) thriftType {
	return func(t *testing.T, rand *rand.Rand) thriftType {
		for {
			// We will keep trying to generate a value until a valid one
			// is found.

			v, ok := quick.Value(typ, rand)
			require.True(t, ok, "failed to generate a value")

			tval := v.Addr().Interface().(thriftType)
			if !thriftTypeIsValid(tval) {
				// Value fails validity check. Try again.
				continue
			}

			return tval
		}
	}
}

// Version fo defaultValueGenerator that sets only one of the fields of a
// union.
func unionValueGenerator(sample interface{}) func(*testing.T, *rand.Rand) thriftType {
	typ := reflect.TypeOf(sample)
	return func(t *testing.T, rand *rand.Rand) thriftType {
		for {
			// We will keep trying to generate a value until a valid one
			// is found.

			v := reflect.New(typ)
			if typ.NumField() == 0 {
				return v.Interface().(thriftType)
			}

			field := typ.Field(rand.Intn(typ.NumField()))
			fieldValue, ok := quick.Value(field.Type, rand)
			require.True(t, ok, "failed to generate a value for field %q", field.Name)

			v.Elem().FieldByIndex(field.Index).Set(fieldValue)

			tval := v.Interface().(thriftType)
			if !thriftTypeIsValid(tval) {
				// Value fails validity check. Try again.
				continue
			}

			return tval
		}
	}
}

// enumValueGenerator builds a generator for random enum values given the
// `*_Values` function for that enum.
func enumValueGenerator(valuesFunc interface{}) func(*testing.T, *rand.Rand) thriftType {
	vfunc := reflect.ValueOf(valuesFunc)
	typ := vfunc.Type().Out(0).Elem() // Foo_Values() []Foo -> Foo
	return func(t *testing.T, rand *rand.Rand) thriftType {
		knownValues := vfunc.Call(nil)[0]

		var giveV reflect.Value
		// Flip a coin to decide whether we're evaluating a known or
		// unknown value.
		if rand.Int()%2 == 0 && knownValues.Len() > 0 {
			// Pick a known value at random
			giveV = knownValues.Index(rand.Intn(knownValues.Len()))
		} else {
			// give = MyEnum($randomValue)
			giveV = reflect.New(typ).Elem()
			giveV.Set(reflect.ValueOf(rand.Int31()).Convert(typ))
		}

		return giveV.Addr().Interface().(thriftType)
	}
}

// Returns true for Go types that are nillable from Thrift's point-of-view.
func isThriftNillable(typ reflect.Type) bool {
	// Only struct types and typedefs of nillable Go native types are
	// nillable.
	switch typ.Kind() {
	case reflect.Map, reflect.Ptr, reflect.Slice, reflect.Struct:
		return true
	}
	return false
}

func TestQuickSuite(t *testing.T) {
	type testCase struct {
		// Sample value of the type to be tested.
		Sample interface{}

		// Specifies how we generate valid values of this type. Defaults to
		// defaultValueGenerator(Type) if unspecified.
		Generator func(*testing.T, *rand.Rand) thriftType

		// If set, logging for this type will not be tested. This is needed
		// for typedefs of primitives which can't implement ArrayMarshaler or
		// ObjectMarshaler.
		NoLog bool

		// If set, the Equals check will not be performed. The check should be
		// disabled for types for which we cannot reliably generate random
		// values at this time: maps with unhashable keys. The randomly
		// generated values will have duplicate keys.
		NoEquals bool
		// TODO(abg): Use a custom generator for these types^

		// If this is an enum, we run a series of additional tests that aren't
		// valid for other types.
		IsEnum bool
	}

	// The following types from our tests have been skipped because they have
	// recursive references (ArbitraryValue) and we don't have a way of
	// addressing that at this time.
	//
	// - services.KeyValue_SetValue_Args
	// - services.KeyValue_SetValueV2_Args
	// - services.KeyValue_GetManyValues_Result
	// - services.KeyValue_GetValue_Result

	// TODO(abg): ^Use custom generators to make this not-a-problem.

	tests := []testCase{
		// structs, unions, and exceptions
		{Sample: tc.ContainersOfContainers{}, NoEquals: true},
		{Sample: tc.EnumContainers{}},
		{Sample: tc.ListOfConflictingEnums{}},
		{Sample: tc.ListOfConflictingUUIDs{}},
		{Sample: tc.MapOfBinaryAndString{}, NoEquals: true},
		{Sample: tc.PrimitiveContainersRequired{}},
		{Sample: tc.PrimitiveContainers{}},
		{Sample: td.DefaultPrimitiveTypedef{}},
		{Sample: td.Event{}},
		{Sample: td.I128{}},
		{Sample: td.Transition{}},
		{Sample: te.StructWithOptionalEnum{}},
		{Sample: tf.Cache_Clear_Args{}},
		{Sample: tf.Cache_ClearAfter_Args{}},
		{Sample: tf.ConflictingNames_SetValue_Args{}},
		{Sample: tf.ConflictingNames_SetValue_Result{}},
		{Sample: tf.ConflictingNamesSetValueArgs{}},
		{Sample: tf.InternalError{}},
		{Sample: tf.KeyValue_DeleteValue_Args{}},
		{Sample: tf.KeyValue_DeleteValue_Result{}},
		{Sample: tf.KeyValue_GetManyValues_Args{}},
		{Sample: tf.KeyValue_GetValue_Args{}},
		{Sample: tf.KeyValue_SetValue_Result{}},
		{Sample: tf.KeyValue_SetValueV2_Result{}},
		{Sample: tf.KeyValue_Size_Args{}},
		{Sample: tf.KeyValue_Size_Result{}},
		{Sample: tf.NonStandardServiceName_NonStandardFunctionName_Args{}},
		{Sample: tf.NonStandardServiceName_NonStandardFunctionName_Result{}},
		{Sample: tl.AccessorConflict{}},
		{Sample: tl.AccessorNoConflict{}},
		{Sample: tl.PrimitiveContainers{}},
		{Sample: tl.StructCollision2{}},
		{Sample: tl.StructCollision{}},
		{Sample: tl.UnionCollision2{}, Generator: unionValueGenerator(tl.UnionCollision2{})},
		{Sample: tl.UnionCollision{}, Generator: unionValueGenerator(tl.UnionCollision{})},
		{Sample: tl.WithDefault{}},
		{Sample: tle.Records{}},
		{Sample: ts.ContactInfo{}},
		{Sample: ts.DefaultsStruct{}},
		{Sample: ts.Edge{}},
		{Sample: ts.EmptyStruct{}},
		{Sample: ts.Frame{}},
		{Sample: ts.GoTags{}},
		{Sample: ts.Graph{}},
		{Sample: ts.Node{}},
		{Sample: ts.Omit{}},
		{Sample: ts.Point{}},
		{Sample: ts.PrimitiveOptionalStruct{}},
		{Sample: ts.PrimitiveRequiredStruct{}},
		{Sample: ts.Rename{}},
		{Sample: ts.Size{}},
		{Sample: ts.StructLabels{}},
		{Sample: ts.User{}},
		{Sample: ts.ZapOptOutStruct{}},
		{Sample: tu.ArbitraryValue{}, Generator: unionValueGenerator(tu.ArbitraryValue{})},
		{Sample: tu.Document{}, Generator: unionValueGenerator(tu.Document{})},
		{Sample: tu.EmptyUnion{}, Generator: unionValueGenerator(tu.EmptyUnion{})},
		{Sample: tul.UUIDConflict{}},
		{Sample: tx.DoesNotExistException{}},
		{Sample: tx.EmptyException{}},
		{Sample: tz.PrimitiveRequiredStruct{}, NoLog: true},

		// typedefs
		{Sample: td.BinarySet{}},
		{Sample: td.EdgeMap{}},
		{Sample: td.FrameGroup{}},
		{Sample: td.MyEnum(0)},
		{Sample: td.PDF{}, NoLog: true},
		{Sample: td.PointMap{}},
		{Sample: td.State(""), NoLog: true},
		{Sample: td.StateMap{}},
		{Sample: td.Timestamp(0), NoLog: true},
		{Sample: td.UUID{}},
		{Sample: tl.LittlePotatoe(0), NoLog: true},
		{Sample: tl.LittlePotatoe2(0.0), NoLog: true},
		{Sample: tul.UUID(""), NoLog: true},
		{Sample: tz.StringMap{}, NoLog: true},
		{Sample: tz.Primitives{}, NoLog: true},
		{Sample: tz.StringList{}, NoLog: true},

		// enums
		{
			Sample:    te.EmptyEnum(0),
			Generator: enumValueGenerator(te.EmptyEnum_Values),
			IsEnum:    true,
		},
		{
			Sample:    te.EnumDefault(0),
			Generator: enumValueGenerator(te.EnumDefault_Values),
			IsEnum:    true,
		},
		{
			Sample:    te.EnumWithDuplicateName(0),
			Generator: enumValueGenerator(te.EnumWithDuplicateName_Values),
			IsEnum:    true,
		},
		{
			Sample:    te.EnumWithDuplicateValues(0),
			Generator: enumValueGenerator(te.EnumWithDuplicateValues_Values),
			IsEnum:    true,
		},
		{
			Sample:    te.EnumWithLabel(0),
			Generator: enumValueGenerator(te.EnumWithLabel_Values),
			IsEnum:    true,
		},
		{
			Sample:    te.EnumWithValues(0),
			Generator: enumValueGenerator(te.EnumWithValues_Values),
			IsEnum:    true,
		},
		{
			Sample:    te.LowerCaseEnum(0),
			Generator: enumValueGenerator(te.LowerCaseEnum_Values),
			IsEnum:    true,
		},
		{
			Sample:    te.RecordType(0),
			Generator: enumValueGenerator(te.RecordType_Values),
			IsEnum:    true,
		},
		{
			Sample:    te.RecordTypeValues(0),
			Generator: enumValueGenerator(te.RecordTypeValues_Values),
			IsEnum:    true,
		},
		{
			Sample:    tl.MyEnum(0),
			Generator: enumValueGenerator(tl.MyEnum_Values),
			IsEnum:    true,
		},
		{
			Sample:    tl.MyEnum2(0),
			Generator: enumValueGenerator(tl.MyEnum2_Values),
			IsEnum:    true,
		},
		{
			Sample:    tle.RecordType(0),
			Generator: enumValueGenerator(tle.RecordType_Values),
			IsEnum:    true,
		},
		{
			Sample:    tz.EnumDefault(0),
			Generator: enumValueGenerator(tz.EnumDefault_Values),
			IsEnum:    true,
			NoLog:     true,
		},
	}

	// Log the seed so that we can reproduce this if it ever fails.
	seed := time.Now().UnixNano()
	rand := rand.New(rand.NewSource(seed))
	t.Logf("Using seed %v for testing/quick", seed)

	const numValues = 1000 // number of values to test against
	for _, tt := range tests {
		typ := reflect.TypeOf(tt.Sample)
		suite := quickSuite{Type: typ}

		t.Run(typ.Name(), func(t *testing.T) {
			generator := tt.Generator
			if generator == nil {
				generator = defaultValueGenerator(typ)
			}

			values := make([]thriftType, numValues)
			for i := range values {
				values[i] = generator(t, rand)
			}

			t.Run("Thrift", func(t *testing.T) {
				for _, give := range values {
					suite.testThriftRoundTrip(t, give)
				}
			})

			t.Run("String", func(t *testing.T) {
				for _, give := range values {
					suite.testString(t, give)
				}
			})

			if isThriftNillable(typ) {
				t.Run("StringNil", suite.testStringNil)
			}

			if tt.IsEnum {
				t.Run("JSON", func(t *testing.T) {
					for _, give := range values {
						suite.testJSONRoundTrip(t, give)
					}
				})

				t.Run("Text", func(t *testing.T) {
					for _, give := range values {
						suite.testTextRoundtrip(t, give)
					}
				})

				t.Run("Ptr", func(t *testing.T) {
					for _, give := range values {
						suite.testEnumPtr(t, give)
					}
				})
			}

			if !tt.NoLog {
				t.Run("Zap", func(t *testing.T) {
					for _, give := range values {
						suite.testLogging(t, give)
					}
				})
			}

			if !tt.NoEquals {
				t.Run("Equals", func(t *testing.T) {
					for _, give := range values {
						suite.testEquals(t, give)
					}
				})

				if typ.Kind() == reflect.Struct {
					t.Run("EqualsNil", suite.testEqualsNil)
				}
			}

		})
	}
}

type quickSuite struct {
	Type reflect.Type
}

// Builds a new empty value of the underlying type.
//
// For structs, returns a pointer to an empty struct. For containers, an empty
// container (not nil) is returned. For everything else, zero value is
// returned.
func (q *quickSuite) newEmpty() reflect.Value {
	v := reflect.New(q.Type).Elem()
	switch q.Type.Kind() {
	case reflect.Map:
		v.Set(reflect.MakeMap(q.Type))
		return v
	case reflect.Slice:
		v.Set(reflect.MakeSlice(q.Type, 0, 0))
		return v
	case reflect.Struct:
		v.Set(reflect.Zero(q.Type))
		return v.Addr()
	default:
		return v
	}
}

// Same as newEmpty but a pointer to the value is returned.
func (q *quickSuite) newEmptyPtr() reflect.Value {
	if q.Type.Kind() == reflect.Struct {
		return q.newEmpty()
	}
	return q.newEmpty().Addr()
}

// Builds a new nil value of the underlying type.
//
// For structs, a nil pointer to the struct is returned. For containers, a nil
// pointer to the container is returned. Using this with non-nillable types is
// invalid.
func (q *quickSuite) newNil(t *testing.T) (v reflect.Value) {
	defer func() {
		require.True(t, v.IsNil(), "bug: newNil generated non-nil value") // sanity check
	}()

	switch q.Type.Kind() {
	case reflect.Struct:
		return reflect.Zero(reflect.PtrTo(q.Type))
	default:
		return reflect.Zero(q.Type)
	}
}

// Tests that the provided value round-trips successfully with wire.Value.
func (q *quickSuite) testThriftRoundTrip(t *testing.T, give thriftType) {
	w, err := give.ToWire()
	require.NoError(t, err, "failed to Thrift encode %v", give)

	got := q.newEmptyPtr().Interface().(thriftType)
	require.NoError(t, got.FromWire(w), "failed to Thrift decode from %v", w)

	assert.Equal(t, got, give)
}

// Tests that String() works on any valid value of this type.
func (q *quickSuite) testString(t *testing.T, give thriftType) {
	assert.NotPanics(t, func() {
		_ = give.String()
	}, "failed to String %#v", give)
}

// Tests that String does not panic with a nil value of this type.
func (q *quickSuite) testStringNil(t *testing.T) {
	v := q.newNil(t).Interface().(fmt.Stringer)
	assert.NotPanics(t, func() {
		_ = v.String()
	})
}

// For types that support it (enums only at this time), tests that JSON
// representations round-trip successfully.
func (q *quickSuite) testJSONRoundTrip(t *testing.T, giveVal thriftType) {
	give, ok := giveVal.(json.Marshaler)
	require.True(t, ok, "does not implement json.Marshaler")

	bs, err := give.MarshalJSON()
	require.NoError(t, err, "failed to encode %v", give)

	got, ok := q.newEmptyPtr().Interface().(json.Unmarshaler)
	require.True(t, ok, "does not implement json.Unmarshaler")

	require.NoError(t, got.UnmarshalJSON(bs), "failed to decode from %q", bs)
	assert.Equal(t, got, give, "could not round-trip")
}

// For types that support it (enums only at this time), tests that
// encoding.TextMarshaler representations round-trip successfully.
func (q *quickSuite) testTextRoundtrip(t *testing.T, giveVal thriftType) {
	give, ok := giveVal.(encoding.TextMarshaler)
	require.True(t, ok, "does not implement encoding.TextMarshaler")

	bs, err := give.MarshalText()
	require.NoError(t, err, "failed to encode %v", give)

	got, ok := q.newEmptyPtr().Interface().(encoding.TextUnmarshaler)
	require.True(t, ok, "does not implement encoding.TextUnmarshaler")

	require.NoError(t, got.UnmarshalText(bs), "failed to decode from %q", bs)
	assert.Equal(t, got, give, "could not round-trip")
}

// Tests that the object can be logged by Zap.
func (q *quickSuite) testLogging(t *testing.T, give thriftType) {
	enc := zapcore.NewMapObjectEncoder()

	if obj, ok := give.(zapcore.ObjectMarshaler); ok {
		assert.NoErrorf(t, obj.MarshalLogObject(enc), "failed to log %v", give)
		return
	}

	if arr, ok := give.(zapcore.ArrayMarshaler); ok {
		assert.NoErrorf(t, enc.AddArray("values", arr), "failed to log %v", give)
		return
	}

	t.Fatal(
		"Type does not implement zapcore.ObjectMarshaler or zapcore.ArrayMarshaler. "+
			"Did you mean to add NoLog?", q.Type)
}

// Tests that the v.Equals(v) always returns true.
func (q *quickSuite) testEquals(t *testing.T, giveVal thriftType) {
	give := reflect.ValueOf(giveVal)
	rhs := give

	equals := give.MethodByName("Equals")
	require.True(t, equals.IsValid(), "Type does not implement Equals()")

	if equals.Type().In(0) != rhs.Type() {
		// We were passing the objects around by pointer but
		// we need the value-form here.
		rhs = rhs.Elem()
	}

	assert.True(t,
		equals.Call([]reflect.Value{rhs})[0].Bool(),
		"%v should be equal to itself", giveVal)
}

// Tests that Equals methods work with nil values and receivers.
func (q *quickSuite) testEqualsNil(t *testing.T) {
	t.Run("both nil", func(t *testing.T) {
		// var x, y *Type
		x := q.newNil(t)
		y := q.newNil(t)

		// x.Equals(y)
		result := x.MethodByName("Equals").
			Call([]reflect.Value{y})[0].
			Interface().(bool)

		assert.True(t, result)
	})

	t.Run("lhs not nil", func(t *testing.T) {
		// var x Type
		// var y *Type
		x := q.newEmpty()
		y := q.newNil(t)

		// x.Equals(y)
		result := x.MethodByName("Equals").
			Call([]reflect.Value{y})[0].
			Interface().(bool)

		assert.False(t, result)
	})

	t.Run("rhs not nil", func(t *testing.T) {
		// var x *Type
		// var y Type
		x := q.newNil(t)
		y := q.newEmpty()

		// x.Equals(y)
		result := x.MethodByName("Equals").
			Call([]reflect.Value{y})[0].
			Interface().(bool)

		assert.False(t, result)
	})
}

// Tests that Ptr methods on enums return the same value back.
func (q *quickSuite) testEnumPtr(t *testing.T, give thriftType) {
	// TODO(abg): should we generate Ptr and _Values for typedefs of enums?
	v := reflect.ValueOf(give)
	ptr := v.MethodByName("Ptr").Call(nil)[0]
	require.Equal(t, reflect.Ptr, ptr.Kind(), "must be a pointer")
	assert.Equal(t, give, ptr.Interface(),
		"pointer must point back to original value")
}
