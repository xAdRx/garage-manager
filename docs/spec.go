// Package docs holds the OpenAPI spec that swag generates from the handler
// comments. Regenerate swagger.json and swagger.yaml with `swag init --outputTypes json,yaml`.
package docs

import _ "embed"

// Spec is the generated OpenAPI (Swagger 2.0) spec, built into the binary.
//
//go:embed swagger.json
var Spec []byte
