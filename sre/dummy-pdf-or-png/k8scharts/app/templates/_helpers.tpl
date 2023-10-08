{{/*
 dummy-pdf-or-png common labels
*/}}
{{- define "dummy-pdf-or-png .Labels" -}}
{{ include "dummy-pdf-or-png .labels" . }}
app.kubernetes.io/component: dummypdfOrpng 
{{- end }}

{{/*
dummy-pdf-or-png selector labels
*/}}
{{- define "dummy-pdf-or-png.SelectorLabels" -}}
{{ include "dummy-pdf-or-png.selectorLabels" . }}
app.kubernetes.io/component: dummypdfOrpng
{{- end }}

{{/*
dummy-pdf-or-png fullname
*/}}
{{- define "dummy-pdf-or-png.fullname" -}}
dummy-pdf-or-png
{{- end }}