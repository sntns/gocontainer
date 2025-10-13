package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Call the worker binary
		out, err := exec.Command("/usr/local/bin/worker", "Hello from server").Output()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "Server response: %s", string(out))
	})

	log.Printf("Multi-binary server starting on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}