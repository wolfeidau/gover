package gover

import (
	"errors"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

func ReadValue(filename string, src []byte, name string) (string, error) {

	fset := token.NewFileSet()

	f, err := parser.ParseFile(fset, "", src, parser.SpuriousErrors)

	if err != nil {
		return "", err
	}

	//	ast.Print(fset, f)

	for _, d := range f.Decls {
		switch d := d.(type) {
		case *ast.GenDecl:
			for _, s := range d.Specs {
				switch s := s.(type) {
				case *ast.ValueSpec:
					for _, n := range s.Names {
						if n.Name == name {
							for _, v := range s.Values {
								switch vv := v.(type) {
								case *ast.BasicLit:
									log.Printf("Value %v", vv.Value)
									return vv.Value, nil
								}

							}
						}

					}
				}
			}
		}
	}

	return "", errors.New("value not found")
}
