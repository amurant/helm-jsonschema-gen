package schema

import (
	"encoding/json"
	"strings"

	"github.com/ghodss/yaml"

	"github.com/alecthomas/jsonschema"
)

type generatedDocItem struct {
	Path        string      `json:"path"`
	Type        string      `json:"type"`
	Description string      `json:"description,omitempty"`
	Required    bool        `json:"required"`
	Default     interface{} `json:"default,omitempty"`
}

func GenDocs(val interface{}, path string) error {
	r := &jsonschema.Reflector{
		ExpandedStruct:            true,
		AllowAdditionalProperties: false,
		DoNotReference:            true,
	}
	schema := r.Reflect(val)

	s, err := json.Marshal(val)
	if err != nil {
		return err
	}

	var jsonMap interface{}
	err = json.Unmarshal(s, &jsonMap)
	if err != nil {
		return err
	}

	docs := extractDocs("", schema.Type, jsonMap, true)

	s, err = json.MarshalIndent(docs, "", "  ")
	if err != nil {
		return err
	}
	s, err = yaml.JSONToYAML(s)
	if err != nil {
		return err
	}

	return saveToFile(s, path)
}

func extractDocs(path string, t *jsonschema.Type, val interface{}, required bool) []generatedDocItem {
	genDocs := []generatedDocItem{}

	if t.Properties != nil {
		keys := t.Properties.Keys()
		for _, k := range keys {
			v, _ := t.Properties.Get(k)

			p := strings.TrimPrefix(path+".", ".")
			fieldDocs := extractDocs(p+k, v.(*jsonschema.Type), val.(map[string]interface{})[k], stringInSlice(k, t.Required))
			genDocs = append(genDocs, fieldDocs...)
		}
	} else {
		var typ string
		if _, ok := t.PatternProperties[".*"]; ok {
			typ = "map[string]string"
		} else if (t.Type == "array") && (t.Items != nil) {
			typ = "[]" + typeName(t.Items.Type)
		} else {
			typ = typeName(t.Type)
		}

		genDocs = append(genDocs, generatedDocItem{
			Path:        path,
			Type:        typ,
			Description: t.Description,
			Required:    required,
			Default:     val,
		})
	}

	return genDocs
}

func typeName(rawType string) string {
	if rawType == "" {
		return "object"
	}
	return rawType
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
