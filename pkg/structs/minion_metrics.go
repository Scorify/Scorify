package structs

import (
	"time"

	"github.com/google/uuid"
)

type MinionMetrics struct {
	MinionID    uuid.UUID `json:"minion_id"`
	Timestamp   time.Time `json:"timestamp"`
	IP          string    `json:"ip"`
	MemoryUsage int64     `json:"memory_usage"`
	MemoryTotal int64     `json:"memory_total"`
	CPUUsage    int64     `json:"cpu_usage"`
	CPUTotal    int64     `json:"cpu_total"`
	DiskUsage   int64     `json:"disk_usage"`
	DiskTotal   int64     `json:"disk_total"`
}
