{{/*
 query common labels
*/}}
{{- define "document-keeper.Labels" -}}
{{ include "document-keeper.labels" . }}
app.kubernetes.io/component: documentKeeper
{{- end }}

{{/*
query selector labels
*/}}
{{- define "document-keeper.SelectorLabels" -}}
{{ include "document-keeper.selectorLabels" . }}
app.kubernetes.io/component: documentKeeper
{{- end }}