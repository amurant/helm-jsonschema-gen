package gen

import (
	"bytes"
	"go/format"
	"os"
	"text/template"
)

func teplateAndWriteToFile(templateString string, values interface{}, path string) error {
	tmpl, err := template.New("render").Parse(templateString)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	if err = tmpl.Execute(&buf, values); err != nil {
		return err
	}
	p, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}

	f, err := os.Create(path)

	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.Write(p)

	return err
}
