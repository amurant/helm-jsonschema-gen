module main

go 1.16

require github.com/amurant/helm-jsonschema-gen v0.0.0-20210729124805-36ac7569c1ed

replace github.com/amurant/helm-jsonschema-gen => ../../

replace github.com/alecthomas/jsonschema => github.com/amurant/jsonschema v0.0.0-20210729111349-9c856f912aa1
