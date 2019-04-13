package main

import (
	"fmt"
	"net/http"
)

func redirTo(to string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		redirURL := fmt.Sprintf("%s/%s", to, path)
		http.Redirect(w, r, redirURL, http.StatusFound)
	})
}
