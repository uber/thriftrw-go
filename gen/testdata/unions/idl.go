// Code generated by thriftrw v1.11.0. DO NOT EDIT.
// @generated

package unions

import (
	"go.uber.org/thriftrw/gen/testdata/typedefs"
	"go.uber.org/thriftrw/thriftreflect"
)

// ThriftModule represents the IDL file used to generate this package.
var ThriftModule = &thriftreflect.ThriftModule{
	Name:     "unions",
	Package:  "go.uber.org/thriftrw/gen/testdata/unions",
	FilePath: "unions.thrift",
	SHA1:     "ae222d8ff3a8efe55b3d2ebb0c70b4565255623c",
	Includes: []*thriftreflect.ThriftModule{
		typedefs.ThriftModule,
	},
	Raw: rawIDL,
}

const rawIDL = "include \"./typedefs.thrift\"\n\nunion EmptyUnion {}\n\nunion Document {\n    1: typedefs.PDF pdf\n    2: string plainText\n}\n\n/**\n * ArbitraryValue allows constructing complex values without a schema.\n *\n * A value is one of,\n *\n * * Boolean\n * * Integer\n * * String\n * * A list of other values\n * * A dictionary of other values\n */\nunion ArbitraryValue {\n    1: bool boolValue\n    2: i64 int64Value\n    3: string stringValue\n    4: list<ArbitraryValue> listValue\n    5: map<string, ArbitraryValue> mapValue\n}\n"
