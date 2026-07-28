package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
		if GreetingEnabled() {
			fmt.Fprintln(w, "The greeting variable is true")
			fmt.Fprintf(w, "Max items: %d\n", MaxItems())
			fmt.Fprintf(w, "Retry count: %d\n", RetryCount())
		}
	})

	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}
