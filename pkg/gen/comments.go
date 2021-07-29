package gen

import (
	"go/ast"
	"go/doc"
	"path"
	"strconv"
	"strings"
)

type TemplateCommentValues struct {
	PackageName string
	Values      map[string]map[string]string
}

func GenComments(rootPath string, headerPath string, pkgName string, pkg *ast.Package) error {
	p := doc.New(pkg, "", doc.PreserveAST)

	docs := map[string]map[string]string{}
	for _, t := range p.Types {
		doc := map[string]string{}

		ast.Inspect(t.Decl, func(n ast.Node) bool {
			if x, ok := n.(*ast.Field); ok {
				s := x.Doc.Text()
				s = strings.TrimSuffix(s, "\n")
				if (len(x.Names) == 1) && (s != "") {
					if strconv.CanBackquote(s) {
						s = "`" + s + "`"
					} else {
						s = strconv.Quote(s)
					}
					doc[x.Names[0].String()] = s
				}
			}

			return true
		})

		if len(doc) > 0 {
			docs[t.Name] = doc
		}
	}

	v := TemplateCommentValues{
		pkgName,
		docs,
	}

	return teplateAndWriteToFile(goTemplate, v, path.Join(rootPath, "comments_generated.go"), headerPath)
}

const goTemplate = `
/* DO NOT EDIT, this file was auto-generated */

package {{ .PackageName }}

{{- range $typName, $typ := .Values }}
func ({{ $typName }}) GetFieldDocString(fieldName string) string {
	switch fieldName {
	{{- range $fieldName, $comment := $typ }}
	case "{{ $fieldName }}":
		return {{ $comment }}
	{{- end }}
	default:
		return ""
	}
}
{{- end }}
`
