package main

import (
	"fmt"
	"net/http"
	"strings"
)

var requestCount = make(map[string]int)

func rateLimiter(w http.ResponseWriter, r *http.Request) {

	ip := strings.Split(r.RemoteAddr, ":")[0]

	requestCount[ip]++

	if requestCount[ip] > 5 {
		w.WriteHeader(http.StatusTooManyRequests)
		fmt.Fprintf(w, "terlalu banyak request %s", ip)
		return
	}

	fmt.Fprintf(w, "hello request ke-%d dari ip %s", requestCount[ip], ip)
}

/*
// func handler
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}
*/

func main() {
	//daftarkan route
	http.HandleFunc("/hello", rateLimiter)

	//jalankan di server
	fmt.Println("server berjalan di http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
