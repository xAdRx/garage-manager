package handlers

import (
	"net/http"

	"github.com/swaggo/swag"
)

// scalarPage loads the Scalar API reference UI and points it at /docs/openapi.json.
const scalarPage = `<!doctype html>
<html>
<head>
	<title>API Reference</title>
	<meta charset="utf-8" />
	<meta name="viewport" content="width=device-width, initial-scale=1" />
</head>
<body>
	<script id="api-reference" data-url="/docs/openapi.json"></script>
	<script src="https://cdn.jsdelivr.net/npm/@scalar/api-reference@1.25.0"></script>
</body>
</html>`

// DocsHandler serves the generated OpenAPI spec and the Scalar UI.
type DocsHandler struct{}

// NewDocsHandler creates a DocsHandler.
func NewDocsHandler() *DocsHandler {
	return &DocsHandler{}
}

// UI handles GET /docs.
func (h *DocsHandler) UI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(scalarPage))
}

// Spec handles GET /docs/openapi.json.
func (h *DocsHandler) Spec(w http.ResponseWriter, r *http.Request) {
	doc, err := swag.ReadDoc()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(doc))
}
