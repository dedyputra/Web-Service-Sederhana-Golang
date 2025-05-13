package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Requested from %s: %s %s\n",
			r.UserAgent(),
			r.Method,
			r.RequestURI)
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(time.Now().Format(time.RFC1123Z)))
	})

	fmt.Println("Server started at :8888")
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Printf("Server error: %s\n", err)
	}
}
