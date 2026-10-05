package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

type RateLimitInfo struct {
	Count     int
	LastReset time.Time
}

var requestCount = make(map[string]RateLimitInfo)

func rateLimiter(w http.ResponseWriter, r *http.Request) {

	ip := strings.Split(r.RemoteAddr, ":")[0]

	info := requestCount[ip]

	if info.LastReset.IsZero() || time.Since(info.LastReset) > time.Second*10 {
		info.Count = 0
		info.LastReset = time.Now()
	}

	info.Count++
	requestCount[ip] = info

	if info.Count > 6 {
		w.WriteHeader(http.StatusTooManyRequests)
		fmt.Fprintf(w, "terlalu banyak request dari IP %s", ip)
		return
	}

	fmt.Fprintf(w, "hello request ke-%d dari ip %s", info.Count, ip)
}

func main() {
	//daftarkan route
	http.HandleFunc("/hello", rateLimiter)

	//jalankan di server
	fmt.Println("server berjalan di http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
