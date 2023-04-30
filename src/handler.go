package main

import "net/http"

func addHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cross-Origin-Opener-Policy", "same-origin")
		w.Header().Set("Cross-Origin-Embedder-Policy", "require-corp")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}

func (app *application) handler() *http.ServeMux {
	mux := http.NewServeMux()
	fsHandler := addHeaders(http.FileServer(http.Dir(app.config.directory)))

	mux.Handle("/", fsHandler)
	return mux
}
