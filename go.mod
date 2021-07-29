module github.com/amurant/helm-jsonschema-gen

go 1.16

require (
	github.com/alecthomas/jsonschema v0.0.0-20210526225647-edb03dcab7bc
	github.com/ghodss/yaml v1.0.0
	github.com/iancoleman/orderedmap v0.2.0 // indirect
	github.com/spf13/cobra v1.2.1
)

replace github.com/alecthomas/jsonschema => github.com/amurant/jsonschema v0.0.0-20210729111349-9c856f912aa1
