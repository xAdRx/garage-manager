package server

import (
	"net/http"
	"os"
	"path"
	"path/filepath"
)

// spaHandler serves the built React app. Paths that match a file (JS, CSS,
// images) get that file. Other paths, such as /owner, get index.html so the
// React router can show the right page.
type spaHandler struct {
	dir   string
	files http.Handler
}

func newSPAHandler(dir string) http.Handler {
	return &spaHandler{dir: dir, files: http.FileServer(http.Dir(dir))}
}

func (h *spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodHead {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// path.Clean on a rooted path removes any "..", so the lookup stays inside dir.
	clean := path.Clean("/" + r.URL.Path)
	info, err := os.Stat(filepath.Join(h.dir, filepath.FromSlash(clean)))
	if err == nil && !info.IsDir() {
		h.files.ServeHTTP(w, r)
		return
	}

	// A missing file with an extension (e.g. /assets/old.js) is a real 404,
	// not a page of the app.
	if path.Ext(clean) != "" {
		http.NotFound(w, r)
		return
	}

	// index.html must not be cached, so a new deploy is picked up at once.
	w.Header().Set("Cache-Control", "no-cache")
	http.ServeFile(w, r, filepath.Join(h.dir, "index.html"))
}
