package main

import (
	"fmt"
	"net/http"

	"rate-limiter-mini/handler"
)

func main() {
	//daftarkan route
	http.HandleFunc("/hello", handler.RateLimiter)

	//jalankan di server
	fmt.Println("server berjalan di http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
