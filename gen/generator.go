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

package gen

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
	"reflect"
	"text/template"

	"go.uber.org/thriftrw/compile"
	"go.uber.org/thriftrw/internal/curry"
)

const generatedByHeader = "// Code generated by thriftrw\n// @generated\n\n"

// Generator allows generating declarations and other templated text for a
// single Go package which may be spread across multiple files.
//
// Multiple files may be written by calling Write multiple times between calls
// to DeclareFromTemplate.
type Generator interface {
	// TextTemplate renders the given template with the given template
	// context.
	TextTemplate(s string, data interface{}, opts ...TemplateOption) (string, error)

	// DeclareFromTemplate renders the given template, and includes the
	// declarations from the generated Go code in the package being generated.
	// The next call to Write will write these declarations and any other
	// declarations that have not been written so far.
	DeclareFromTemplate(s string, data interface{}, opts ...TemplateOption) error

	// EnsureDeclared is similar to DeclareFromTemplate except that it simply
	// ignores conflicting definitions.
	EnsureDeclared(s string, data interface{}, opts ...TemplateOption) error

	// LookupTypeName returns the fully qualified name that should be used for
	// the given Thrift type. It imports the corresponding Go package if
	// necessary.
	//
	// This must be called with user-defined types only.
	LookupTypeName(compile.TypeSpec) (string, error)

	// LookupConstantName returns the fully qualified name that should be used
	// for the given Thrift constant. It imports the corresponding Go package if
	// necessary.
	LookupConstantName(*compile.Constant) (string, error)

	// Import ensures that the given package has been imported in the generated
	// code. Returns the name that should be used to reference the imported
	// module.
	Import(path string) string

	// Write the generated code to the given Writer and start a new file for
	// this package. All consecutive calls to DeclareFromTemplate will be
	// accumulated until the next Write call.
	//
	// A single Write must corresponds to a single file in the generated
	// package.
	Write(w io.Writer, fs *token.FileSet) error
}

var _typeOfGenerator = reflect.TypeOf((*Generator)(nil)).Elem()

// TemplateOption customizes templates.
type TemplateOption func(Generator, *template.Template) *template.Template

// TemplateFunc adds a function to the template.
//
// The function may be anything accepted by text/template. If the function
// accepts a Generator as its first argument, it will be provided automatically.
func TemplateFunc(name string, f interface{}) TemplateOption {
	return func(g Generator, t *template.Template) *template.Template {
		return t.Funcs(template.FuncMap{name: curryGenerator(f, g)})
	}
}

// curryGenerator adds g as the first argument to f if f accepts it. Otherwise f
// is returned unmodified.
func curryGenerator(f interface{}, g Generator) interface{} {
	typ := reflect.TypeOf(f)
	if typ.NumIn() > 0 && typ.In(0) == _typeOfGenerator {
		return curry.One(f, g)
	}
	return f
}

// generator tracks code generation state as we generate the output.
type generator struct {
	importer
	Namespace

	PackageName string
	ImportPath  string

	w              WireGenerator
	decls          []ast.Decl
	thriftImporter thriftPackageImporter

	// TODO use something to group related decls together
}

// NewGenerator sets up a new generator for Go code.
func NewGenerator(timport thriftPackageImporter, importPath string, packageName string) Generator {
	// TODO(abg): Determine package name from `namespace go` directive.
	namespace := NewNamespace()
	return &generator{
		PackageName:    packageName,
		ImportPath:     importPath,
		Namespace:      namespace,
		importer:       newImporter(namespace.Child()),
		thriftImporter: timport,
	}
}

func (g *generator) LookupTypeName(t compile.TypeSpec) (string, error) {
	if t.ThriftFile() == "" {
		return "", fmt.Errorf(
			"LookupTypeName called with native type (%T) %v", t, t)
	}

	importPath, err := g.thriftImporter.Package(t.ThriftFile())
	if err != nil {
		return "", err
	}

	name := goCase(t.ThriftName())
	if importPath != g.ImportPath {
		pkg := g.Import(importPath)
		name = pkg + "." + name
	}
	return name, nil
}

func (g *generator) LookupConstantName(c *compile.Constant) (string, error) {
	importPath, err := g.thriftImporter.Package(c.File)
	if err != nil {
		return "", err
	}

	name := goCase(c.Name)
	if importPath != g.ImportPath {
		pkg := g.Import(importPath)
		name = pkg + "." + name
	}
	return name, nil
}

// servicePackage imports the service package for the given service and returns
// the name that should be used to refer to that package.
func (g *generator) servicePackage(s *compile.ServiceSpec) (string, error) {
	pkg, err := g.thriftImporter.ServicePackage(s.ThriftFile(), s.Name)
	if err != nil {
		return "", err
	}

	return g.Import(pkg), nil
}

// TextTemplate renders the given template with the given template context.
func (g *generator) TextTemplate(s string, data interface{}, opts ...TemplateOption) (string, error) {
	templateFuncs := template.FuncMap{
		"goCase":           goCase,
		"import":           g.Import,
		"isHashable":       isHashable,
		"isPrimitiveType":  isPrimitiveType,
		"isStructType":     isStructType,
		"newNamespace":     g.Namespace.Child,
		"newVar":           g.Namespace.Child().NewName,
		"servicePackage":   g.servicePackage,
		"typeName":         curryGenerator(typeName, g),
		"typeReference":    curryGenerator(typeReference, g),
		"typeReferencePtr": curryGenerator(typeReferencePtr, g),
		"fromWire":         curryGenerator(g.w.FromWire, g),
		"fromWirePtr":      curryGenerator(g.w.FromWirePtr, g),
		"toWire":           curryGenerator(g.w.ToWire, g),
		"toWirePtr":        curryGenerator(g.w.ToWirePtr, g),
		"typeCode":         curryGenerator(TypeCode, g),
	}

	tmpl := template.New("thriftrw").Delims("<", ">").Funcs(templateFuncs)
	for _, opt := range opts {
		tmpl = opt(g, tmpl)
	}

	tmpl, err := tmpl.Parse(s)
	if err != nil {
		return "", err
	}

	buff := bytes.Buffer{}
	if err := tmpl.Execute(&buff, data); err != nil {
		return "", err
	}

	return buff.String(), nil

}

func (g *generator) renderTemplate(s string, data interface{}, opts ...TemplateOption) ([]byte, error) {
	buff := bytes.NewBufferString("package thriftrw\n\n")
	out, err := g.TextTemplate(s, data, opts...)
	if err != nil {
		return nil, err
	}
	if _, err := buff.WriteString(out); err != nil {
		return nil, err
	}

	return buff.Bytes(), nil
}

func (g *generator) recordGenDeclNames(ignoreConflicts bool, d *ast.GenDecl) error {
	switch d.Tok {
	case token.IMPORT:
		for _, spec := range d.Specs {
			if err := g.AddImportSpec(spec.(*ast.ImportSpec)); err != nil {
				return fmt.Errorf(
					"could not add explicit import %s: %v", spec, err,
				)
			}
		}
	case token.CONST:
		for _, spec := range d.Specs {
			for _, name := range spec.(*ast.ValueSpec).Names {
				if err := g.Reserve(name.Name); err != nil && !ignoreConflicts {
					return fmt.Errorf(
						"could not declare constant %q: %v", name.Name, err,
					)
				}
			}
		}
	case token.TYPE:
		for _, spec := range d.Specs {
			name := spec.(*ast.TypeSpec).Name.Name
			if err := g.Reserve(name); err != nil && !ignoreConflicts {
				return fmt.Errorf("could not declare type %q: %v", name, err)
			}
		}
	case token.VAR:
		for _, spec := range d.Specs {
			for _, name := range spec.(*ast.ValueSpec).Names {
				if err := g.Reserve(name.Name); err != nil && !ignoreConflicts {
					return fmt.Errorf(
						"could not declare var %q: %v", name.Name, err,
					)
				}
			}
		}
	default:
		return fmt.Errorf("unknown declaration: %v", d)
	}
	return nil
}

// DeclareFromTemplate renders a template (in the text/template format) that
// generates Go code and includes all declarations from the template in the code
// generated by the generator.
//
// An error is returned if anything went wrong while generating the template.
//
// For example,
//
// 	g.DeclareFromTemplate(
// 		'type <.Name> int32',
// 		struct{Name string}{Name: "myType"}
// 	)
//
// Will generate,
//
// 	type myType int32
//
// The following functions are available to templates:
//
// fromWire(TypeSpec, v): Returns an expression of type (T, error) where T is
// the type represented by TypeSpec, read from the given Value v.
//
// goCase(str): Accepts a string and returns it in CamelCase form and the
// first character upper-cased. The string may be ALLCAPS, snake_case, or
// already camelCase.
//
// import(str): Accepts a string and returns the name that should be used in
// the template to refer to that imported module. This helps avoid naming
// conflicts with imports.
//
// 	<$fmt := import "fmt">
// 	<$fmt>.Println("hello world")
//
// isHashable(TypeSpec): Returns true if the given TypeSpec is for a type that
// is hashable.
//
// isPrimitiveType(TypeSpec): Returns true if the given TypeSpec is for a
// primitive type.
//
// isStructType(TypeSpec): Returns true if the given TypeSpec is a StructSpec.
//
// newVar(s): Gets a new name that the template can use for a variable without
// worrying about shadowing any globals. Prefers the given string.
//
// 	<$x := newVar "x">
//
// toWire(TypeSpec, v): Returns an expression of type (Value, error) that
// contains the wire representation of the item "v" of type TypeSpec.

// toWirePtr(TypeSpec, v): Returns an expression of type (Value, error) that
// contains the wire representation of the item "v" which is a reference to a
// value of type TypeSpec.
//
// typeCode(TypeSpec): Gets the wire.Type for the given TypeSpec, importing
// the wire module if necessary.
//
// typeName(TypeSpec): Returns the Go name of the given type, regardless of
// whether it's native or custom.
//
// typeReference(TypeSpec): Takes any TypeSpec. Returns a string representing
// a reference to that type.
//
// 	<typeReference $someType>
//
// typeReferencePtr(TypeSpec): Takes any TypeSpec. Returns a string
// representing a reference to a pointer of that type, provided that the type
// itself is not a reference type.
//
// 	<typeReferencePtr $someType>
func (g *generator) declare(ignoreConflicts bool, s string, data interface{}, opts ...TemplateOption) error {
	bs, err := g.renderTemplate(s, data, opts...)
	if err != nil {
		return err
	}

	f, err := parser.ParseFile(token.NewFileSet(), "thriftrw.go", bs, 0)
	if err != nil {
		return fmt.Errorf("could not parse generated code: %v:\n%s", err, bs)
	}

	for _, decl := range f.Decls {
		switch d := decl.(type) {
		case *ast.FuncDecl:
			if d.Recv == nil {
				// top-level function
				if err := g.Reserve(d.Name.Name); err != nil {
					if ignoreConflicts {
						continue
					}
					return err
				}
			}
		case *ast.GenDecl:
			if err := g.recordGenDeclNames(ignoreConflicts, d); err != nil {
				return err
			}
		default:
			// No special behavior. Move along.
		}
		g.appendDecl(decl)
	}

	return nil
}

func (g *generator) DeclareFromTemplate(s string, data interface{}, opts ...TemplateOption) error {
	return g.declare(false, s, data, opts...)
}

func (g *generator) EnsureDeclared(s string, data interface{}, opts ...TemplateOption) error {
	return g.declare(true, s, data, opts...)
}

func (g *generator) Write(w io.Writer, fs *token.FileSet) error {
	// TODO constants first, types next, and functions after that

	if _, err := w.Write([]byte(generatedByHeader)); err != nil {
		return err
	}

	if _, err := fmt.Fprintf(w, "package %s\n\n", g.PackageName); err != nil {
		return err
	}

	cfg := printer.Config{
		Mode:     printer.UseSpaces | printer.TabIndent,
		Tabwidth: 8,
	}

	if importDecl := g.importDecl(); importDecl != nil {
		if err := cfg.Fprint(w, fs, importDecl); err != nil {
			return err
		}
	}

	if _, err := io.WriteString(w, "\n"); err != nil {
		return err
	}

	for _, decl := range g.decls {
		if _, err := io.WriteString(w, "\n"); err != nil {
			return err
		}

		if err := cfg.Fprint(w, fs, decl); err != nil {
			return err
		}

		if _, err := io.WriteString(w, "\n"); err != nil {
			return err
		}
	}

	g.decls = nil
	g.importer = newImporter(g.Namespace.Child())

	// init can appear multiple times in the same package across different
	// files
	g.Namespace.Forget("init")
	return nil
}

// appendDecl appends a new declaration to the generator.
func (g *generator) appendDecl(decl ast.Decl) {
	g.decls = append(g.decls, decl)
}
