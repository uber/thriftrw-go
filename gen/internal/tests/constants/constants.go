// Code generated by thriftrw v1.33.0. DO NOT EDIT.
// @generated

package constants

import (
	containers "go.uber.org/thriftrw/gen/internal/tests/containers"
	enums "go.uber.org/thriftrw/gen/internal/tests/enums"
	exceptions "go.uber.org/thriftrw/gen/internal/tests/exceptions"
	other_constants "go.uber.org/thriftrw/gen/internal/tests/other_constants"
	structs "go.uber.org/thriftrw/gen/internal/tests/structs"
	typedefs "go.uber.org/thriftrw/gen/internal/tests/typedefs"
	unions "go.uber.org/thriftrw/gen/internal/tests/unions"
	ptr "go.uber.org/thriftrw/ptr"
	thriftreflect "go.uber.org/thriftrw/thriftreflect"
)

const Home enums.RecordType = enums.RecordTypeHomeAddress

const Name enums.RecordType = enums.RecordTypeName

const WorkAddress enums.RecordType = enums.RecordTypeWorkAddress

var ArbitraryValue *unions.ArbitraryValue = &unions.ArbitraryValue{
	ListValue: []*unions.ArbitraryValue{
		&unions.ArbitraryValue{
			BoolValue: ptr.Bool(true),
		},
		&unions.ArbitraryValue{
			Int64Value: ptr.Int64(2),
		},
		&unions.ArbitraryValue{
			StringValue: ptr.String("hello"),
		},
		&unions.ArbitraryValue{
			MapValue: map[string]*unions.ArbitraryValue{
				"foo": &unions.ArbitraryValue{
					StringValue: ptr.String("bar"),
				},
			},
		},
	},
}

// Timestamp at which time began.
const BeginningOfTime typedefs.Timestamp = typedefs.Timestamp(0)

var ContainersOfContainers *containers.ContainersOfContainers = &containers.ContainersOfContainers{
	ListOfLists: [][]int32{
		[]int32{
			1,
			2,
			3,
		},
		[]int32{
			4,
			5,
			6,
		},
	},
	ListOfMaps: []map[int32]int32{
		map[int32]int32{
			1: 2,
			3: 4,
			5: 6,
		},
		map[int32]int32{
			7:  8,
			9:  10,
			11: 12,
		},
	},
	ListOfSets: []map[int32]struct{}{
		map[int32]struct{}{
			1: struct{}{},
			2: struct{}{},
			3: struct{}{},
		},
		map[int32]struct{}{
			4: struct{}{},
			5: struct{}{},
			6: struct{}{},
		},
	},
	MapOfListToSet: []struct {
		Key   []int32
		Value map[int64]struct{}
	}{
		{
			Key: []int32{
				1,
				2,
				3,
			},
			Value: map[int64]struct{}{
				1: struct{}{},
				2: struct{}{},
				3: struct{}{},
			},
		},
		{
			Key: []int32{
				4,
				5,
				6,
			},
			Value: map[int64]struct{}{
				4: struct{}{},
				5: struct{}{},
				6: struct{}{},
			},
		},
	},
	MapOfMapToInt: []struct {
		Key   map[string]int32
		Value int64
	}{
		{
			Key: map[string]int32{
				"1": 1,
				"2": 2,
				"3": 3,
			},
			Value: 100,
		},
		{
			Key: map[string]int32{
				"4": 4,
				"5": 5,
				"6": 6,
			},
			Value: 200,
		},
	},
	MapOfSetToListOfDouble: []struct {
		Key   map[int32]struct{}
		Value []float64
	}{
		{
			Key: map[int32]struct{}{
				1: struct{}{},
				2: struct{}{},
				3: struct{}{},
			},
			Value: []float64{
				1.2,
				3.4,
			},
		},
		{
			Key: map[int32]struct{}{
				4: struct{}{},
				5: struct{}{},
				6: struct{}{},
			},
			Value: []float64{
				5.6,
				7.8,
			},
		},
	},
	SetOfLists: [][]string{
		[]string{
			"1",
			"2",
			"3",
		},
		[]string{
			"4",
			"5",
			"6",
		},
	},
	SetOfMaps: []map[string]string{
		map[string]string{
			"1": "2",
			"3": "4",
			"5": "6",
		},
		map[string]string{
			"7":  "8",
			"9":  "10",
			"11": "12",
		},
	},
	SetOfSets: []map[string]struct{}{
		map[string]struct{}{
			"1": struct{}{},
			"2": struct{}{},
			"3": struct{}{},
		},
		map[string]struct{}{
			"4": struct{}{},
			"5": struct{}{},
			"6": struct{}{},
		},
	},
}

var EmptyException *exceptions.EmptyException = &exceptions.EmptyException{}

var EnumContainers *containers.EnumContainers = &containers.EnumContainers{
	ListOfEnums: []enums.EnumDefault{
		enums.EnumDefaultBar,
		enums.EnumDefaultFoo,
	},
	MapOfEnums: map[enums.EnumWithDuplicateValues]int32{
		enums.EnumWithDuplicateValuesP: 1,
		enums.EnumWithDuplicateValuesQ: 2,
	},
	SetOfEnums: map[enums.EnumWithValues]struct{}{
		enums.EnumWithValuesX: struct{}{},
		enums.EnumWithValuesY: struct{}{},
	},
}

// An example frame group.
//
// Contains two frames.
var FrameGroup typedefs.FrameGroup = typedefs.FrameGroup{
	&structs.Frame{
		Size: &structs.Size{
			Height: 200,
			Width:  100,
		},
		TopLeft: &structs.Point{
			X: 1,
			Y: 2,
		},
	},
	&structs.Frame{
		Size: &structs.Size{
			Height: 400,
			Width:  300,
		},
		TopLeft: &structs.Point{
			X: 3,
			Y: 4,
		},
	},
}

var Graph *structs.Graph = &structs.Graph{
	Edges: []*structs.Edge{
		&structs.Edge{
			EndPoint: &structs.Point{
				X: 3,
				Y: 4,
			},
			StartPoint: &structs.Point{
				X: 1,
				Y: 2,
			},
		},
		&structs.Edge{
			EndPoint: &structs.Point{
				X: 7,
				Y: 8,
			},
			StartPoint: &structs.Point{
				X: 5,
				Y: 6,
			},
		},
	},
}

const Hex16 int16 = 4660

const Hex32 int32 = 305419896

const Hex64 int64 = 1311768467294899695

var I128 *typedefs.I128 = &typedefs.I128{
	High: 1234,
	Low:  5678,
}

const Int16 int16 = 12345

const Int32 int32 = 1234567890

const Int64 int64 = 1234567890123456789

var LastNode *structs.Node = &structs.Node{
	Value: 3,
}

const Lower enums.LowerCaseEnum = enums.LowerCaseEnumItems

const MyEnum typedefs.MyEnum = typedefs.MyEnum(enums.EnumWithValuesY)

var Node *structs.Node = &structs.Node{
	Tail: &structs.List{
		Tail: &structs.List{
			Value: 3,
		},
		Value: 2,
	},
	Value: 1,
}

var PrimitiveContainers *containers.PrimitiveContainers = &containers.PrimitiveContainers{
	ListOfInts: []int64{
		1,
		2,
		3,
	},
	MapOfIntToString: map[int32]string{
		1: "1",
		2: "2",
		3: "3",
	},
	MapOfStringToBool: map[string]bool{
		"1": false,
		"2": true,
		"3": true,
	},
	SetOfBytes: map[int8]struct{}{
		1: struct{}{},
		2: struct{}{},
		3: struct{}{},
	},
	SetOfStrings: map[string]struct{}{
		"foo": struct{}{},
		"bar": struct{}{},
	},
}

func _EnumDefault_ptr(v enums.EnumDefault) *enums.EnumDefault {
	return &v
}

var StructWithOptionalEnum *enums.StructWithOptionalEnum = &enums.StructWithOptionalEnum{
	E: _EnumDefault_ptr(enums.EnumDefaultBaz),
}

var UUID *typedefs.UUID = &typedefs.UUID{
	High: 1234,
	Low:  5678,
}

// ThriftModule represents the IDL file used to generate this package.
var ThriftModule = &thriftreflect.ThriftModule{
	Name:     "constants",
	Package:  "go.uber.org/thriftrw/gen/internal/tests/constants",
	FilePath: "constants.thrift",
	SHA1:     "9dcab618ffb2d35b6baacfaa01c23461e63dbf06",
	Includes: []*thriftreflect.ThriftModule{
		containers.ThriftModule,
		enums.ThriftModule,
		exceptions.ThriftModule,
		other_constants.ThriftModule,
		structs.ThriftModule,
		typedefs.ThriftModule,
		unions.ThriftModule,
	},
	Raw: rawIDL,
}

const rawIDL = "include \"./other_constants.thrift\"\ninclude \"./containers.thrift\"\ninclude \"./enums.thrift\"\ninclude \"./exceptions.thrift\"\ninclude \"./structs.thrift\"\ninclude \"./unions.thrift\"\ninclude \"./typedefs.thrift\"\n\nconst i16 int16 = 12345\nconst i32 int32 = 1234567890\nconst i64 int64 = 1234567890123456789\n\nconst i16 hex16 = 0x1234\nconst i32 hex32 = 0x12345678\nconst i64 hex64 = 0x1234567890abcdef\n\nconst containers.PrimitiveContainers primitiveContainers = {\n    \"listOfInts\": other_constants.listOfInts, // imported constant\n    \"setOfStrings\": [\"foo\", \"bar\"],\n    \"setOfBytes\": other_constants.listOfInts, // imported constant with type casting\n    \"mapOfIntToString\": {\n        1: \"1\",\n        2: \"2\",\n        3: \"3\",\n    },\n    \"mapOfStringToBool\": {\n        \"1\": 0,\n        \"2\": 1,\n        \"3\": 1,\n    }\n}\n\nconst containers.EnumContainers enumContainers = {\n    \"listOfEnums\": [1, enums.EnumDefault.Foo],\n    \"setOfEnums\": [123, enums.EnumWithValues.Y],\n    \"mapOfEnums\": {\n        0: 1,\n        enums.EnumWithDuplicateValues.Q: 2,\n    },\n}\n\nconst containers.ContainersOfContainers containersOfContainers = {\n    \"listOfLists\": [[1, 2, 3], [4, 5, 6]],\n    \"listOfSets\": [[1, 2, 3], [4, 5, 6]],\n    \"listOfMaps\": [{1: 2, 3: 4, 5: 6}, {7: 8, 9: 10, 11: 12}],\n    \"setOfSets\": [[\"1\", \"2\", \"3\"], [\"4\", \"5\", \"6\"]],\n    \"setOfLists\": [[\"1\", \"2\", \"3\"], [\"4\", \"5\", \"6\"]],\n    \"setOfMaps\": [\n        {\"1\": \"2\", \"3\": \"4\", \"5\": \"6\"},\n        {\"7\": \"8\", \"9\": \"10\", \"11\": \"12\"},\n    ],\n    \"mapOfMapToInt\": {\n        {\"1\": 1, \"2\": 2, \"3\": 3}: 100,\n        {\"4\": 4, \"5\": 5, \"6\": 6}: 200,\n    },\n    \"mapOfListToSet\": {\n        // more type casting\n        other_constants.listOfInts: other_constants.listOfInts,\n        [4, 5, 6]: [4, 5, 6],\n    },\n    \"mapOfSetToListOfDouble\": {\n        [1, 2, 3]: [1.2, 3.4],\n        [4, 5, 6]: [5.6, 7.8],\n    },\n}\n\nconst enums.StructWithOptionalEnum structWithOptionalEnum = {\n    \"e\": enums.EnumDefault.Baz\n}\n\nconst exceptions.EmptyException emptyException = {}\n\nconst structs.Graph graph = {\n    \"edges\": [\n        {\"startPoint\": other_constants.some_point, \"endPoint\": {\"x\": 3, \"y\": 4}},\n        {\"startPoint\": {\"x\": 5, \"y\": 6}, \"endPoint\": {\"x\": 7, \"y\": 8}},\n    ]\n}\n\nconst structs.Node lastNode = {\"value\": 3}\nconst structs.Node node = {\n    \"value\": 1,\n    \"tail\": {\"value\": 2, \"tail\": lastNode},\n}\n\nconst unions.ArbitraryValue arbitraryValue = {\n    \"listValue\": [\n        {\"boolValue\": 1},\n        {\"int64Value\": 2},\n        {\"stringValue\": \"hello\"},\n        {\"mapValue\": {\"foo\": {\"stringValue\": \"bar\"}}},\n    ],\n}\n// TODO: union validation for constants?\n\nconst typedefs.i128 i128 = uuid\nconst typedefs.UUID uuid = {\"high\": 1234, \"low\": 5678}\n\n/** Timestamp at which time began. */\nconst typedefs.Timestamp beginningOfTime = 0\n\n/**\n * An example frame group.\n *\n * Contains two frames.\n */\nconst typedefs.FrameGroup frameGroup = [\n    {\n        \"topLeft\": {\"x\": 1, \"y\": 2},\n        \"size\": {\"width\": 100, \"height\": 200},\n    }\n    {\n        \"topLeft\": {\"x\": 3, \"y\": 4},\n        \"size\": {\"width\": 300, \"height\": 400},\n    },\n]\n\nconst typedefs.MyEnum myEnum = enums.EnumWithValues.Y\n\nconst enums.RecordType NAME = enums.RecordType.NAME\nconst enums.RecordType HOME = enums.RecordType.HOME_ADDRESS\nconst enums.RecordType WORK_ADDRESS = enums.RecordType.WORK_ADDRESS\n\nconst enums.lowerCaseEnum lower = enums.lowerCaseEnum.items\n"
