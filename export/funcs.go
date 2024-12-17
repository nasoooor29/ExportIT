package export

import (
	"go/ast"
	"go/token"
	"strings"

	"golang.org/x/tools/go/packages"
)

type ExportedFunc struct {
	Func *ast.FuncDecl
	Pkg  *packages.Package
}

func GetExportedFuncs() ([]ExportedFunc, error) {
	cfg := &packages.Config{Mode: packages.NeedSyntax | packages.NeedTypes | packages.NeedName, Fset: token.NewFileSet()}
	pkgs, err := packages.Load(cfg, "./...") // Current package
	if err != nil {
		return nil, err
	}
	funcs := []ExportedFunc{}

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
				doc := funcDecl.Doc
				if doc == nil {
					continue
				}
				comment := doc.Text()
				comment = strings.TrimSpace(comment)

				if !strings.HasPrefix(strings.ToLower(comment), "export:") {
					continue
				}

				funcs = append(funcs, ExportedFunc{
					Func: funcDecl,
					Pkg:  pkg,
				})

			}
		}
	}

	return funcs, nil
}

func GetFuncParams(f *ast.FuncDecl) []string {
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
