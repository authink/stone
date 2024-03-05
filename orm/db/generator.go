package db

// import (
// 	"go/ast"
// 	"go/build"
// 	"go/parser"
// 	"go/token"
// 	"os"
// 	"path/filepath"
// )

// func GenByModels(packageName string) {
// 	pkg, err := build.Import(packageName, "", build.FindOnly)
// 	if err != nil {
// 		panic(err)
// 	}

// 	if err = filepath.Walk(pkg.Dir, func(path string, info os.FileInfo, err error) error {
// 		if err != nil {
// 			return err
// 		}

// 		if !info.IsDir() && filepath.Ext(path) == ".go" {
// 			parseFile(path)
// 		}

// 		return nil
// 	}); err != nil {
// 		panic(err)
// 	}
// }

// func parseFile(filename string) {
// 	fset := token.NewFileSet()
// 	node, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
// 	if err != nil {
// 		panic(err)
// 	}

// 	var comments []string

// 	for _, decl := range node.Decls {
// 		genDecl, ok := decl.(*ast.GenDecl)
// 		if !ok || genDecl.Tok != token.TYPE {
// 			continue
// 		}

// 		for _, spec := range genDecl.Specs {
// 			typeSpec, ok := spec.(*ast.TypeSpec)
// 			if !ok {
// 				continue
// 			}

// 			if typeSpec.Name.Name == structName {
// 				// 获取结构体定义之前的注释
// 				for _, c := range genDecl.Doc.List {
// 					comments = append(comments, c.Text)
// 				}
// 			}
// 		}
// 	}
// }
