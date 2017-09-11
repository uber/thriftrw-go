// Code generated by thriftrw v1.7.0. DO NOT EDIT.
// @generated

package enums

import "go.uber.org/thriftrw/thriftreflect"

// ThriftModule represents the IDL file used to generate this package.
var ThriftModule = &thriftreflect.ThriftModule{
	Name:     "enums",
	Package:  "go.uber.org/thriftrw/gen/testdata/enums",
	FilePath: "enums.thrift",
	SHA1:     "29d968038855257994d5accefed40594eeccfc08",
	Raw:      rawIDL,
}

const rawIDL = "enum EmptyEnum {}\n\nenum EnumDefault {\n    Foo, Bar, Baz\n}\n\nenum EnumWithValues {\n    X = 123,\n    Y = 456,\n    Z = 789,\n}\n\nenum EnumWithDuplicateValues {\n    P, // 0\n    Q = -1,\n    R, // 0\n}\n\n// enum with item names conflicting with those of another enum\nenum EnumWithDuplicateName {\n    A, B, C, P, Q, R, X, Y, Z\n}\n\n// Enum treated as optional inside a struct\nstruct StructWithOptionalEnum {\n    1: optional EnumDefault e\n}\n\nenum RecordType {\n  NAME,\n  HOME_ADDRESS,\n  WORK_ADDRESS\n}\n\nenum lowerCaseEnum {\n    containing, lower_case, items\n}\n\n// collision with RecordType_Values() function.\nenum RecordType_Values { FOO, BAR }\n"
