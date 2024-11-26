package main

import (
	"WebSpire/core"
	"fmt"
	"net/http"
)

func main() {
	router := core.NewRouter()

	router.AddRoute("GET", "/", func(w http.ResponseWriter, r *http.Request, params map[string]string) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("<h1>Welcome to the Framework!</h1>"))
	}, "home")

	router.AddRoute("GET", "/example", func(w http.ResponseWriter, r *http.Request, params map[string]string) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("<h1>example route feel free to build ur route, wait for more</h1>"))
	}, "home")

	// Start the server
	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", router)
}
