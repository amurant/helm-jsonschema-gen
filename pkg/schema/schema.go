package schema

import (
	"encoding/json"
	"os"

	"github.com/ghodss/yaml"

	"github.com/alecthomas/jsonschema"
)

func GenSchema(val interface{}, path string) error {
	r := &jsonschema.Reflector{
		ExpandedStruct:            true,
		AllowAdditionalProperties: false,
	}
	schema := r.Reflect(val)
	s, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		return err
	}

	return saveToFile(s, path)
}

func GenDefaults(defaultValues interface{}, path string) error {
	s, err := json.MarshalIndent(defaultValues, "", "  ")
	if err != nil {
		return err
	}
	s, err = yaml.JSONToYAML(s)
	if err != nil {
		return err
	}

	return saveToFile(s, path)
}

func saveToFile(content []byte, path string) error {
	f, err := os.Create(path)

	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.Write(content)

	return err
}
