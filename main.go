package main

import (
	"fmt"
	"net/http"
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
}

func main() {
	http.HandleFunc("/", healthCheckHandler)
	fmt.Println("Server berjalan di port 8090...")
	if err := http.ListenAndServe(":8090", nil); err != nil {
		panic(err)
	}
}