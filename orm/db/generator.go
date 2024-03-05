package db

import (
	"bytes"
	_ "embed"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"html/template"
	"os"
	"strings"

	"github.com/authink/inkstone/util"
)

//go:embed template/*.tpl
var tmpl string

type dbModel struct {
	AtDB      bool
	Tname     string
	AtEmbed   bool
	EmbedName string

	Model  string
	Name   string
	Fields []string
}

func writeFile(dbm *dbModel, dbPath string) {
	if _, err := os.Stat(dbPath); err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(dbPath, os.ModePerm)
		}
	}

	f, err := os.Create(fmt.Sprintf("%s/%s.go", dbPath, dbm.Name))
	if err != nil {
		panic(err)
	}

	t := template.Must(template.New("").Parse(tmpl))

	var buf bytes.Buffer
	if err = t.Execute(&buf, dbm); err != nil {
		panic(err)
	}

	p, err := format.Source(buf.Bytes())
	if err != nil {
		panic(err)
	}

	if _, err = f.Write(p); err != nil {
		panic(err)
	}
}

func outputDBModel(genDecl *ast.GenDecl, dbm *dbModel, dbPath string) {
	for _, spec := range genDecl.Specs {
		if t, ok := spec.(*ast.TypeSpec); ok {
			dbm.Model = t.Name.String()
			dbm.Name = util.ToLowerFirstLetter(dbm.Model)
			if st, ok := t.Type.(*ast.StructType); ok {
				for _, field := range st.Fields.List {
					if len(field.Names) == 0 {
						continue
					}

					dbm.Fields = append(dbm.Fields, field.Names[0].Name)
				}
			}
		}
	}
	writeFile(dbm, dbPath)
}

func GenByModels(mPath, dbPath string) {
	fset := token.NewFileSet()

	pkgs, err := parser.ParseDir(fset, mPath, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	if pkg, ok := pkgs["models"]; ok {
		for _, file := range pkg.Files {
			for _, decl := range file.Decls {
				if genDecl, ok := decl.(*ast.GenDecl); ok && genDecl.Tok == token.TYPE && genDecl.Doc != nil {
					if annotations := genDecl.Doc.List; len(annotations) == 2 {
						if !strings.Contains(annotations[0].Text, "@model") {
							continue
						}

						if strings.Contains(annotations[1].Text, "@db") {
							var dbm dbModel

							dbm.AtDB = true
							dbm.Tname = strings.TrimPrefix(annotations[1].Text, "// @db ")

							outputDBModel(genDecl, &dbm, dbPath)
							continue
						}

						if strings.Contains(annotations[1].Text, "@embed") {
							var dbm dbModel

							dbm.AtEmbed = true
							dbm.EmbedName = util.ToLowerFirstLetter(strings.TrimPrefix(annotations[1].Text, "// @embed "))

							outputDBModel(genDecl, &dbm, dbPath)
							continue
						}
					}
				}
			}
		}
	}
}
