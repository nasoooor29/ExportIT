package ExportIT

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
	"runtime"
	"strings"

	"golang.org/x/tools/go/packages"
)

func getFuncByName(name string) (*ast.FuncDecl, error) {
	if strings.Contains(name, "/") {
		name = strings.Split(name, "/")[1]

	}
	cfg := &packages.Config{Mode: packages.NeedSyntax | packages.NeedTypes | packages.NeedName, Fset: token.NewFileSet()}
	pkgs, err := packages.Load(cfg, "./...") // Current package
	if err != nil {
		return nil, err
	}

	for _, pkg := range pkgs {
		for _, syntax := range pkg.Syntax {
			for _, decl := range syntax.Decls {
				funcDecl, ok := decl.(*ast.FuncDecl)
				if !ok {
					continue
				}

				ast.Inspect(funcDecl, func(n ast.Node) bool {
					return true
				})

				fname := pkg.Name + "." + funcDecl.Name.Name
				if name == fname {
					return funcDecl, nil
				}
			}
		}
	}

	return nil, fmt.Errorf("func not found")
}
func getFunctionName(fn interface{}) string {
	if reflect.TypeOf(fn).Kind() != reflect.Func {
		return ""
	}
	n := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	return n
}

func getFuncParams(f *ast.FuncDecl) []string {
	var params []string
	if f.Type.Params != nil {
		for _, param := range f.Type.Params.List {
			for _, name := range param.Names {
				params = append(params, name.Name)
			}
		}
	}
	return params
}
