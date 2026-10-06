package logger

import (
	"fmt"
	"os"
	"time"
)

func WriteLog(ip string, count int, status string) {
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
