package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

// Handler function for "/"
// Accepts an http ResponseWriter which is used to write info back to the client
// and a *http.Request which is used to et info about the request
func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello world!")
	io.WriteString(w, "Request received!\n")
}

func main() {
	// Sets up the handler function for a SPECIFIC request path
	http.HandleFunc("/", getRoot)

	// Tells global HTTP server to listen for incoming requests on port
	err := http.ListenAndServe(":8080", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server %s\n", err)
		os.Exit(1)
	}
}
