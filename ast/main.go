package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "../sync/pool/main.go", nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	ast.Inspect(file, func(n ast.Node) bool {
		funcCall, ok := n.(*ast.CallExpr)
		if ok {
			mthd, ok := funcCall.Fun.(*ast.SelectorExpr)
			if ok {
				fmt.Printf("mthd %+v\n", mthd)
				id, ok := mthd.X.(*ast.Ident)
				if ok {
					if id.Name == "bufPool" {
						line := fset.Position(mthd.Pos()).Line
						fmt.Printf("line %v\n", line)
						for _, cg := range file.Comments {
							fmt.Printf("cg %v\n", cg.Text())
						}
					}
				}
			}
		}
		return true
	})
}
