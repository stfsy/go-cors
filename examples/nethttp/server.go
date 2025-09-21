package main

import (
	"net/http"

	"github.com/stfsy/go-cors"
)

func main() {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://foo.com"},
	})

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"hello\": \"world\"}"))
	})

	http.ListenAndServe(":8080", c.Handler(handler))
}
