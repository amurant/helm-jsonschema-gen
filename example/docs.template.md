| command | type | description | default |
| ------- | ---- | ----------- | ------- |
{{- range $index, $doc := .valuesdocs }}
{{- $desc := $doc.description | toString | replace "\n" " " }}
{{- $flag := (eq $doc.type "string") | ternary "--set-string" "--set" }}
| {{ $flag }} {{ $doc.path }} | {{ $doc.type }} | {{ $desc }} | `{{ kindIs "invalid" $doc.default | ternary "" ($doc.default | toJson) }}` |
{{- end }}
