package model

import (
	"time"
)

type RateLimitInfo struct {
	Count     int
	LastReset time.Time
}
