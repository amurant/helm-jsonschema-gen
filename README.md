# helm-jsonschema-gen

helm-jsonschema-gen is a cli tool and go libary for generating Helm `values.schema.json`, `values.yaml` and `values.docs.yaml`.
It uses a source of truth that is written in go.

# cert-manager example (`example/values`)

The go structs are defined in their [corresponding go source files](https://github.com/amurant/helm-jsonschema-gen/tree/main/example/values).

Below is a snippet of the `main.go` file that references the "root" object with the default values.
```go
package main

import (
	"github.com/amurant/helm-jsonschema-gen/pkg/schema"
)

//go:generate go run github.com/amurant/helm-jsonschema-gen/cmd/jsonschema-gen -p .

func main() {	
	if err := schema.GenSchema(&defaultValues, "../values.schema.json"); err != nil {
		panic(err)
	}
	if err := schema.GenDefaults(&defaultValues, "../values.yaml"); err != nil {
		panic(err)
	}
	if err := schema.GenDocs(&defaultValues, "../values.docs.yaml"); err != nil {
		panic(err)
	}
}

var defaultValues = Values{
    ...
}
```

Using the following commands, we can now generate the `values.schema.json`, `values.yaml` and `values.docs.yaml` files.

```bash
$ cd example/values
$ go generate . && go run .

$ cat ../values.schema.json
{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "required": [
    "installCRDs",
    "replicaCount",
    "podLabels",
    "strategy",
    "nodeSelector",
    "affinity",
    "tolerations",
    "image",
    "resources",
...

$ cat ../values.yaml
affinity: {}
cainjector:
  affinity: {}
  containerSecurityContext: {}
  enabled: true
  extraArgs: []
  extraEnv: []
...

$ cat ../values.docs.yaml
- default: false
  path: installCRDs
  required: true
  type: boolean
- default: 1
  path: replicaCount
  required: true
  type: integer
- path: deploymentAnnotations
  required: false
  type: map[string]string
...
```

# Generating markdown documentation

Use the [`tem`](https://github.com/amurant/tem) tool.
```bash
$ cd example
$ tem -f valuesdocs=./values.docs.yaml -t docs.template.md
| command | type | description | default |
| ------- | ---- | ----------- | ------- |
| --set installCRDs | boolean | If true, CRD resources will be installed as part of the Helm chart. If enabled, when uninstalling CRD resources will be deleted causing all installed custom resources to be DELETED. | `false` |
| --set replicaCount | integer | Number of replicas | `1` |
| --set deploymentAnnotations | map[string]string | Annotations to add to the deployment | `` |
| --set podAnnotations | map[string]string | Annotations to add to the pods | `` |
| --set podLabels | map[string]string | Annotations to add to the pods | `` |
...
```
