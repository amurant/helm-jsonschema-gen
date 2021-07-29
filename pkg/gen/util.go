package gen

import (
	"bytes"
	"go/format"
	"io"
	"os"
	"text/template"
)

func teplateAndWriteToFile(templateString string, values interface{}, path string, headerPath string) error {
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

	fh, err := os.Open(headerPath)

	if err != nil {
		return err
	}

	defer fh.Close()

	f, err := os.Create(path)

	if err != nil {
		return err
	}

	defer f.Close()

	if _, err = io.Copy(f, fh); err != nil {
		return err
	}

	_, err = f.Write(p)

	return err
}
