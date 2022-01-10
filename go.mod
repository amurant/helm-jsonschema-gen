module github.com/amurant/helm-jsonschema-gen

go 1.17

require (
	github.com/alecthomas/jsonschema v0.0.0-20211022214203-8b29eab41725
	github.com/spf13/cobra v1.3.0
	sigs.k8s.io/yaml v1.3.0
)

require (
	github.com/iancoleman/orderedmap v0.2.0 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
	sigs.k8s.io/json v0.0.0-20211020170558-c049b76a60c6 // indirect
)

replace sigs.k8s.io/yaml => github.com/amurant/yaml v1.2.1-0.20211023101056-a4ea27ae8bb4
