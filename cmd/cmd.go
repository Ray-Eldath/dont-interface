package cmd

import (
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
)

type Visitor struct {
	TotalParams int
	EvilParams  int

	TotalResults int
	EvilResults  int

	TotalValueDecl int
	EvilValueDecl  int

	TotalStructField int
	EvilStructField  int
}

func (v *Visitor) Visit(node ast.Node) ast.Visitor {
	switch n := node.(type) {
	case *ast.FuncType:
		if n.Params != nil {
			v.TotalParams += n.Params.NumFields()
			for _, param := range n.Params.List {
				if checkIfEvil(param) {
					v.EvilParams += 1
				}
				if ellipsis, ok := param.Type.(*ast.Ellipsis); ok {
					if _, ok = ellipsis.Elt.(*ast.InterfaceType); ok {
						v.EvilParams += 1
					}
				}
			}
		}
		if n.Results != nil {
			v.TotalResults += n.Results.NumFields()
			for _, result := range n.Results.List {
				if checkIfEvil(result) {
					v.EvilResults += 1
				}
			}
		}
	case *ast.ValueSpec:
		v.TotalValueDecl += 1
		if _, ok := n.Type.(*ast.InterfaceType); ok {
			v.EvilValueDecl += 1
		}
	case *ast.StructType:
		fields := n.Fields
		if fields != nil {
			v.TotalStructField += fields.NumFields()
			for _, field := range fields.List {
				if checkIfEvil(field) {
					v.EvilStructField += 1
				}
			}
		}
	}
	return v
}

func checkIfEvil(field *ast.Field) bool {
	if field == nil {
		return false
	}
	switch f := field.Type.(type) {
	case *ast.InterfaceType:
		return true
	case *ast.ArrayType:
		_, ok := f.Elt.(*ast.InterfaceType)
		return ok
	}
	return false
}

func Calculate(paths []string) (*Visitor, error) {
	fset := token.NewFileSet()
	visitor := &Visitor{}
	for _, path := range paths {
		p, _ := filepath.Abs(path)
		f, err := parser.ParseFile(fset, p, nil, parser.AllErrors)
		if err != nil {
			return nil, err
		}
		ast.Walk(visitor, f)
	}
	return visitor, nil
}
