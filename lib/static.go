package lib

import (
	"net/http"
	"os"
	"path"
	"path/filepath"
)

// NewStaticMiddleware creates a router that:
//  - serves static content from file if available
//  - passes request to next handler if not
// Note that this middleware terminates if the file was found!
func NewStaticMiddleware(root string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			p := path.Join(root, r.URL.Path)

			// Normalize directory routes to index.html
			info, err := os.Stat(filepath.FromSlash(p))
			if err == nil && info.IsDir() {
				p = path.Join(p, "index.html")
			}

			_, err = os.Stat(filepath.FromSlash(p))
			if err != nil {
				if next != nil {
					next.ServeHTTP(w, r)
				} else {
					http.Error(w, "File not found", http.StatusNotFound)
				}
			} else {
				fs := http.FileServer(http.Dir(root))
				fs.ServeHTTP(w, r)
			}
		})
	}
}
