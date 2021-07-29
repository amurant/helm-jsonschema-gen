package gen

import (
	"errors"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
)

func LoadPackage(path string) (error, string, *ast.Package) {
	// Create the AST by parsing src and test.
	fset := token.NewFileSet()

	pkgs, err := parser.ParseDir(fset, path, func(file fs.FileInfo) bool {
		return file.Name() != "tools.go"
	}, parser.ParseComments)
	if err != nil {
		return err, "", nil
	}

	if len(pkgs) == 0 {
		return errors.New("no package found"), "", nil
	}

	if len(pkgs) > 1 {
		return errors.New("more than 1 package found"), "", nil
	}

	var pkgName string
	var pkg *ast.Package
	for pkgName, pkg = range pkgs {
		break
	}

	return nil, pkgName, pkg
}
