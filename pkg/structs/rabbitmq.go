package structs

import (
	"fmt"
	"slices"
	"time"

	"github.com/google/uuid"
	"github.com/scorify/scorify/pkg/ent/minion"
	"github.com/scorify/scorify/pkg/ent/status"
)

type RabbitMQErrorResponse struct {
	Error  string `json:"error"`
	Reason string `json:"reason"`
}

type TaskRequest struct {
	StatusID   uuid.UUID `json:"status_id"`
	SourceName string    `json:"source_name"`
	Config     string    `json:"config"`
}
type TaskResponse struct {
	StatusID uuid.UUID     `json:"status_id"`
	MinionID uuid.UUID     `json:"minion_id"`
	Status   status.Status `json:"status"`
	Error    string        `json:"error"`
}

type WorkerEnroll struct {
	MinionID uuid.UUID   `json:"minion_id"`
	Hostname string      `json:"hostname"`
	Role     minion.Role `json:"role"`
}

type WorkerStatus []uuid.UUID

func (w WorkerStatus) Disabled(minionID uuid.UUID) bool {
	return slices.Contains(w, minionID)
}

type Heartbeat struct {
	MinionID    uuid.UUID `json:"minion_id"`
	Timestamp   time.Time `json:"timestamp"`
	MemoryUsage int64     `json:"memory_usage"`
	MemoryTotal int64     `json:"memory_total"`
	CPUUsage    float64   `json:"cpu_usage"`
	Goroutines  int64     `json:"goroutines"`
}

func (h Heartbeat) String() string {
	return fmt.Sprintf("minionID=%s, timestamp=%s, memoryUsage=%s, memoryTotal=%s, cpuUsage=%.2f%%, goroutines=%d",
		h.MinionID,
		h.Timestamp,
		readableBytes(h.MemoryUsage),
		readableBytes(h.MemoryTotal),
		h.CPUUsage,
		h.Goroutines,
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
