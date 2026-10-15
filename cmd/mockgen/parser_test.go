package main

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"testing"

	"golang.org/x/tools/go/analysis"
)

const anonymousStructFixture = "package fixture\n" + `
import clock "time"

type Embedded struct { Value string }
type ExampleAo struct {
 XML struct {
  Header struct { ID string ` + "`xml:\"EIid\"`" + ` } ` + "`xml:\"Header\"`" + `
  Items []struct {
   Name string ` + "`xml:\"ItemName\"`" + `
   Detail *struct { At clock.Time }
  } ` + "`xml:\"Items\"`" + `
  Embedded
  *clock.Location
  A, B string ` + "`json:\"value\"`" + `
  Empty struct{}
 }
}

func (*ExampleAo) Convert(input struct {
 Nested []struct { At clock.Time ` + "`json:\"at\"`" + ` }
 Embedded
}) struct { Result *struct { At clock.Time }; Empty struct{} } {
 return struct { Result *struct { At clock.Time }; Empty struct{} }{}
}
`

func parseFixture(t *testing.T) (*analysis.Pass, *ast.File) {
	t.Helper()
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "fixtureao.go", anonymousStructFixture, 0)
	if err != nil {
		t.Fatal(err)
	}
	info := &types.Info{
		Types: make(map[ast.Expr]types.TypeAndValue),
		Defs:  make(map[*ast.Ident]types.Object),
		Uses:  make(map[*ast.Ident]types.Object),
	}
	cfg := types.Config{Importer: importer.Default()}
	pkg, err := cfg.Check("fixture", fset, []*ast.File{file}, info)
	if err != nil {
		t.Fatal(err)
	}
	return &analysis.Pass{Fset: fset, Files: []*ast.File{file}, TypesInfo: info, Pkg: pkg}, file
}

func TestAnonymousStructType(t *testing.T) {
	pass, file := parseFixture(t)
	p := NewParser()
	pkgInfo := PackageInfo{}
	p.extractImportsFromFile([]*ast.File{file}, &pkgInfo)
	decl := file.Decls[2].(*ast.GenDecl).Specs[0].(*ast.TypeSpec)
	expr := decl.Type.(*ast.StructType).Fields.List[0].Type
	deps := map[string]ImportInfo{}
	result := p.extractTypeFromExpr(pass, expr, pass.Pkg.Path(), pkgInfo.Imports, deps)
	// Type identity checks nested fields, embedding, order, and tags together.
	got, err := types.Eval(pass.Fset, pass.Pkg, expr.Pos(), result.Code)
	if err != nil {
		t.Fatalf("invalid generated type: %v\n%s", err, result.Code)
	}
	if !types.Identical(got.Type, pass.TypesInfo.TypeOf(expr)) {
		t.Fatalf("type changed: %s", result.Code)
	}
	if len(result.DepImports) != 1 || deps["time"].PkgName != "clock" {
		t.Fatalf("missing aliased nested import: %#v", result.DepImports)
	}
}

func TestGenerateAnonymousStructMethods(t *testing.T) {
	pass, file := parseFixture(t)
	p := NewParser()
	pkgInfo := PackageInfo{Name: "fixture", Path: "fixture"}
	p.extractImportsFromFile([]*ast.File{file}, &pkgInfo)
	decl := file.Decls[2].(*ast.GenDecl)
	spec := decl.Specs[0].(*ast.TypeSpec)
	structInfo := p.extractStructInfo(pass, decl, spec, spec.Type.(*ast.StructType), pkgInfo.Imports)
	structInfo.Methods = []MethodInfo{p.extractMethodInfo(pass, file.Decls[3].(*ast.FuncDecl), pkgInfo.Imports)}
	pkgInfo.Structs = []StructInfo{structInfo}
	g := NewGenerator(".*", ".*")
	code, err := g.generateContent(pkgInfo)
	if err != nil {
		t.Fatal(err)
	}
	code, err = g.formatSource(code)
	if err != nil {
		t.Fatal(err)
	}
	fset := token.NewFileSet()
	source, err := parser.ParseFile(fset, "fixtureao.go", anonymousStructFixture, 0)
	if err != nil {
		t.Fatal(err)
	}
	generated, err := parser.ParseFile(fset, "fixture_mock.go", code+"\nvar _ IExampleAo = (*ExampleAo)(nil)\nvar _ IExampleAo = (*ExampleAoMock)(nil)\n", 0)
	if err != nil {
		t.Fatal(err)
	}
	cfg := types.Config{Importer: mockgenTestImporter{Importer: importer.Default()}}
	if _, err := cfg.Check("fixture", fset, []*ast.File{source, generated}, nil); err != nil {
		t.Fatalf("generated mock does not type-check: %v\n%s", err, code)
	}
}

// The generated registration call only needs this signature for type checking;
// all other imports are checked against their real exported types.
type mockgenTestImporter struct{ types.Importer }

func (i mockgenTestImporter) Import(path string) (*types.Package, error) {
	if path != "github.com/fishedee/fishgo-boost/app/proxy" {
		return i.Importer.Import(path)
	}
	pkg := types.NewPackage(path, "proxy")
	param := types.NewVar(token.NoPos, pkg, "mocks", types.NewInterfaceType(nil, nil).Complete())
	sig := types.NewSignatureType(nil, nil, nil, types.NewTuple(param), nil, false)
	pkg.Scope().Insert(types.NewFunc(token.NoPos, pkg, "RegisterProxyMock", sig))
	pkg.MarkComplete()
	return pkg, nil
}
