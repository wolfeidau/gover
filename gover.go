package gover

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

func ReadVersion(filename string, src []byte) (string, error) {

	fset := token.NewFileSet()

	file, err := parser.ParseFile(fset, "", src, parser.SpuriousErrors)

	if err != nil {
		return "", err
	}

	for _, d := range file.Decls {
		switch n := d.(type) {
		case *ast.GenDecl:
			log.Printf("n = %+v", n)
			for _, s := range n.Specs {
				log.Printf("s = %+v", s)
				switch ss := s.(type) {
				case *ast.ValueSpec:
					log.Printf("ss = %+v", ss.Names)
					for _, v := range ss.Values {
						log.Printf("v = %+v", v)
					}
				}

			}

		}
		//		log.Printf("d = %+v", d)
	}
	return "0.2.0", nil
}
