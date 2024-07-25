package structs

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type MinionMetrics struct {
	MinionID    uuid.UUID `json:"minion_id"`
	Timestamp   time.Time `json:"timestamp"`
	IP          string    `json:"ip"`
	MemoryUsage int64     `json:"memory_usage"`
	MemoryTotal int64     `json:"memory_total"`
	CPUUsage    float64   `json:"cpu_usage"`
	Goroutines  int64     `json:"goroutines"`
}

func (m MinionMetrics) String() string {
	return fmt.Sprintf("minionID=%s, timestamp=%s, ip=%s, memoryUsage=%s, memoryTotal=%s, cpuUsage=%.2f%%, goroutines=%d",
		m.MinionID,
		m.Timestamp,
		m.IP,
		readableBytes(m.MemoryUsage),
		readableBytes(m.MemoryTotal),
		m.CPUUsage,
		m.Goroutines,
	)
}

const (
	KB = 1024
	MB = KB * 1024
	GB = MB * 1024
	TB = GB * 1024
)

func readableBytes(bytes int64) string {
	bytes = bytes * 1024

	if bytes < KB {
		return fmt.Sprintf("%dB", bytes)
	}
	if bytes < MB {
		return fmt.Sprintf("%.2fKB", float64(bytes)/KB)
	}
	if bytes < GB {
		return fmt.Sprintf("%.2fMB", float64(bytes)/MB)
	}
	return fmt.Sprintf("%.2fGB", float64(bytes)/TB)
}
