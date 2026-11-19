package handlers

import (
	"net/http"

	"restapi/docs"
)

// scalarPage loads the Scalar API reference UI and points it at /docs/openapi.json.
// The integrity hash makes the browser refuse the script if the CDN serves a changed
// file. When you change the Scalar version, compute the new hash:
//
//	curl -sL <script URL> | openssl dgst -sha384 -binary | base64
const scalarPage = `<!doctype html>
<html>
<head>
	<title>API Reference</title>
	<meta charset="utf-8" />
	<meta name="viewport" content="width=device-width, initial-scale=1" />
</head>
<body>
	<script id="api-reference" data-url="/docs/openapi.json"></script>
	<script src="https://cdn.jsdelivr.net/npm/@scalar/api-reference@1.76.4/dist/browser/standalone.min.js"
		integrity="sha384-28+p9NyLFAl2b71X2XembwoB3f1K1iY0PyATQvkWJJ9k8E3sFW8FeGMbguBzasxv"
		crossorigin="anonymous"></script>
</body>
</html>`

// DocsUI handles GET /docs.
func DocsUI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(scalarPage))
}

// DocsSpec handles GET /docs/openapi.json.
func DocsSpec(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(docs.Spec)
}
