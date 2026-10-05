package main

import (
	"fmt"
	"net/http"
	"os"
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
		writeLog(ip, info.Count, "BLOCKED")
		return
	}

	fmt.Fprintf(w, "hello request ke-%d dari ip %s", info.Count, ip)
	writeLog(ip, info.Count, "OK")
}

func writeLog(ip string, count int, status string) {
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error buka file:", err)
		return
	}
	defer file.Close()

	waktu := time.Now().Format("2006-01-02 15:04:05")

	logLine := fmt.Sprintf("[%s] IP: %s | Request ke-%d | Status: %s\n", waktu, ip, count, status)

	file.WriteString(logLine)

	fmt.Print(logLine)

}

func main() {
	//daftarkan route
	http.HandleFunc("/hello", rateLimiter)

	//jalankan di server
	fmt.Println("server berjalan di http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
