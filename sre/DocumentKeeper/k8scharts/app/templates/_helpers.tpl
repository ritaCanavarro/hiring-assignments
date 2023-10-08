{{/*
 documentKeeper common labels
*/}}
{{- define "document-keeper.labels" -}}
app.kubernetes.io/component: documentKeeper
{{- end }}

{{/*
documentKeeper selector labels
*/}}
{{- define "document-keeper.selectorLabels" -}}
app.kubernetes.io/component: documentKeeper
{{- end }}

{{/*
documentKeeper fullname
*/}}
{{- define "document-keeper.fullname" -}}
document-keeper
{{- end }}
