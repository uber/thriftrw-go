// Code generated by thriftrw v1.2.0
// @generated

// Copyright (c) 2017 Uber Technologies, Inc.
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

import "go.uber.org/thriftrw/thriftreflect"

var ThriftModule = &thriftreflect.ThriftModule{Name: "api", Package: "go.uber.org/thriftrw/plugin/api", FilePath: "api.thrift", SHA1: "d31dcccbc1fd417d282feaed8f9d07e1d8eb2b8a", Raw: rawIDL}

const rawIDL = "/**\n * API_VERSION is the version of the plugin API.\n *\n * This MUST be provided in the HandshakeResponse.\n */\nconst i32 API_VERSION = 3\n\n/**\n * ServiceID is an arbitrary unique identifier to reference the different\n * services in this request.\n */\ntypedef i32 ServiceID\n\n/**\n * ModuleID is an arbitrary unique identifier to reference the different\n * modules in this request.\n */\ntypedef i32 ModuleID\n\n/**\n * TypeReference is a reference to a user-defined type.\n */\nstruct TypeReference {\n    1: required string name\n    /**\n     * Import path for the package defining this type.\n     */\n    2: required string importPath\n\n    // TODO(abg): Should this just be using ModuleID instead of a package?\n}\n\n/**\n * SimpleType is a standalone native Go type.\n */\nenum SimpleType {\n    BOOL = 1,     // bool\n    BYTE,         // byte\n    INT8,         // int8\n    INT16,        // int16\n    INT32,        // int32\n    INT64,        // int64\n    FLOAT64,      // float64\n    STRING,       // string\n    STRUCT_EMPTY, // struct{}\n}\n\n/**\n * TypePair is a pair of two types.\n */\nstruct TypePair {\n    1: required Type left\n    2: required Type right\n}\n\n/**\n * Type is a reference to a Go type which may be native or user defined.\n */\nunion Type {\n    1: SimpleType simpleType\n    /**\n     * Slice of a type\n     *\n     * []$sliceType\n     */\n    2: Type sliceType\n    /**\n     * Slice of key-value pairs of a pair of types.\n     *\n     * []struct{Key $left, Value $right}\n     */\n    3: TypePair keyValueSliceType\n    /**\n     * Map of a pair of types.\n     *\n     * map[$left]$right\n     */\n    4: TypePair mapType\n    /**\n     * Reference to a user-defined type.\n     */\n    5: TypeReference referenceType\n    /**\n     * Pointer to a type.\n     */\n    6: Type pointerType\n}\n\n/**\n * Argument is a single Argument inside a Function.\n * For,\n *\n *      void setValue(1: string key, 2: string value)\n *\n * You get the arguments,\n *\n *      Argument{Name: \"Key\", Type: Type{SimpleType: SimpleTypeString}}\n *\n *      Argument{Name: \"Value\", Type: Type{SimpleType: SimpleTypeString}}\n */\nstruct Argument {\n    /**\n     * Name of the argument. This is also the name of the argument field\n     * inside the args/result struct for that function.\n     */\n    1: required string name\n    /**\n     * Argument type.\n     */\n    2: required Type type\n}\n\n/**\n * Function is a single function on a Thrift service.\n */\nstruct Function {\n    /**\n     * Name of the Go function.\n     */\n    1: required string name\n    /**\n     * Name of the function as defined in the Thrift file.\n     */\n    2: required string thriftName\n    /**\n     * List of arguments accepted by the function.\n     *\n     * This list is in the order specified by the user in the Thrift file.\n     */\n    3: required list<Argument> arguments\n    /**\n     * Return type of the function, if any. If this is not set, the function\n     * is a void function.\n     */\n    4: optional Type returnType\n    /**\n     * List of exceptions raised by the function.\n     *\n     * This list is in the order specified by the user in the Thrift file.\n     */\n    5: optional list<Argument> exceptions\n    /**\n     * Whether this function is oneway or not. This should be assumed to be\n     * false unless explicitly stated otherwise. If this is true, the\n     * returnType and exceptions will be null or empty.\n     */\n    6: optional bool oneWay\n}\n\n/**\n * Service is a service defined by the user in the Thrift file.\n */\nstruct Service {\n    /**\n     * Name of the Thrift service in Go code.\n     */\n    7: required string name\n    /**\n     * Name of the service as defined in the Thrift file.\n     */\n    1: required string thriftName\n    /**\n     * ID of the parent service.\n     */\n    4: optional ServiceID parentID\n    /**\n     * List of functions defined for this service.\n     */\n    5: required list<Function> functions\n    /**\n     * ID of the module where this service was declared.\n     */\n    6: required ModuleID moduleID\n}\n\n/**\n * Module is a module generated from a single Thrift file. Each module\n * corresponds to exactly one Thrift file and contains all the types and\n * constants defined in that Thrift file.\n */\nstruct Module {\n    /**\n     * Import path for the package defining the types for this module.\n     */\n    1: required string importPath\n    /**\n     * Path to the directory containing the code for this module.\n     *\n     * The path is relative to the output directory into which ThriftRW is\n     * generating code. Plugins SHOULD NOT make any assumptions about the\n     * absolute location of the directory.\n     */\n    2: required string directory\n}\n\n//////////////////////////////////////////////////////////////////////////////\n\n/**\n * Feature is a functionality offered by a ThriftRW plugin.\n */\nenum Feature {\n    /**\n     * SERVICE_GENERATOR specifies that the plugin may generate arbitrary code\n     * for services defined in the Thrift file.\n     *\n     * If a plugin provides this, it MUST implement the ServiceGenerator\n     * service.\n     */\n    SERVICE_GENERATOR = 1,\n\n    // TODO: TAGGER for struct-tagging plugins\n}\n\n/**\n * HandshakeRequest is the initial request sent to the plugin as part of\n * establishing communication and feature negotiation.\n */\nstruct HandshakeRequest {\n}\n\n/**\n * HandshakeResponse is the response from the plugin for a HandshakeRequest.\n */\nstruct HandshakeResponse {\n    /**\n     * Name of the plugin. This MUST match the name of the plugin specified\n     * over the command line or the program will fail.\n     */\n    1: required string name\n    /**\n     * Version of the plugin API.\n     *\n     * This MUST be set to API_VERSION by the plugin.\n     */\n    2: required i32 apiVersion (go.name = \"APIVersion\")\n    /**\n     * List of features the plugin provides.\n     */\n    3: required list<Feature> features\n    /**\n     * Version of ThriftRW with which the plugin was built.\n     *\n     * This MUST be set to go.uber.org/thriftrw/version.Version by the plugin\n     * explicitly.\n     */\n    4: optional string libraryVersion\n}\n\nservice Plugin {\n    /**\n     * handshake performs a handshake with the plugin to negotiate the\n     * features provided by it and the version of the plugin API it expects.\n     */\n    HandshakeResponse handshake(1: HandshakeRequest request)\n\n    /**\n     * Informs the plugin process that it will not receive any more requests\n     * and it is safe for it to exit.\n     */\n    void goodbye()\n}\n\n//////////////////////////////////////////////////////////////////////////////\n\n/**\n * GenerateServiceRequest is a request to generate code for zero or more\n * Thrift services.\n */\nstruct GenerateServiceRequest {\n    /**\n     * IDs of services for which code should be generated.\n     *\n     * Note that the services map contains information about both, the\n     * services being generated and their transitive dependencies. Code should\n     * only be generated for service IDs listed here.\n     */\n    1: required list<ServiceID> rootServices\n    /**\n     * Map of service ID to service.\n     *\n     * Any service IDs present in this request will have a corresponding\n     * service definition in this map, including services for which code does\n     * not need to be generated.\n     */\n    2: required map<ServiceID, Service> services\n    /**\n     * Map of module ID to module.\n     *\n     * Any module IDs present in the request will have a corresponding module\n     * definition in this map.\n     */\n    3: required map<ModuleID, Module> modules\n}\n\n/**\n * GenerateServiceResponse is response to a GenerateServiceRequest.\n */\nstruct GenerateServiceResponse {\n    /**\n     * Map of file path to file contents.\n     *\n     * All paths MUST be relative to the output directory into which ThriftRW\n     * is generating code. Plugins SHOULD NOT make any assumptions about the\n     * absolute location of the directory.\n     *\n     * The paths MUST NOT contain the string \"..\" or the request will fail.\n     */\n    1: optional map<string, binary> files\n}\n\n/**\n * ServiceGenerator generates arbitrary code for services.\n *\n * This MUST be implemented if the SERVICE_GENERATOR feature is enabled.\n */\nservice ServiceGenerator {\n    /**\n     * Generates code for requested services.\n     */\n    GenerateServiceResponse generate(1: GenerateServiceRequest request)\n}\n"
