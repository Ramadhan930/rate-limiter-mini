package handler

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"rate-limiter-mini/detector"
	"rate-limiter-mini/logger"
	"rate-limiter-mini/model"
)

var requestCount = make(map[string]model.RateLimitInfo)

var blacklist = make(map[string]bool)

func RateLimiter(w http.ResponseWriter, r *http.Request) {

	ip := strings.Split(r.RemoteAddr, ":")[0]

	if blacklist[ip] {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, "IP %s diblokir karena terdeteksi menyerang.", ip)
		logger.WriteLog(ip, 0, "BLACKLISTED")
		return
	}

	name := r.URL.Query().Get("name")
	if detector.IsSQLInjection(name) {
		blacklist[ip] = true
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, "serangan terdeteksi! IP %s diblokir.", ip)
		logger.WriteLog(ip, 0, "ATTACK_DETECTED")
		return
	}

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
		logger.WriteLog(ip, info.Count, "BLOCKED")
		return
	}

	fmt.Fprintf(w, "hello request ke-%d dari ip %s", info.Count, ip)
	logger.WriteLog(ip, info.Count, "OK")
}
