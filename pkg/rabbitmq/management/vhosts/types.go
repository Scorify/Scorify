package vhosts

type QueueType string

const (
	Classic QueueType = "classic"
	Quorum  QueueType = "quorum"
	Steam   QueueType = "stream"
)

type vhostsRequest struct {
	DefaultQueueType QueueType `json:"default_queue_type"`
	Description      string    `json:"description"`
	Name             string    `json:"name"`
	Tags             string    `json:"tags"`
}

type vhostsResponse struct {
	ClusterState     map[string]string      `json:"cluster_state"`
	DefaultQueueType string                 `json:"default_queue_type"`
	Description      string                 `json:"description"`
	Metadata         map[string]interface{} `json:"metadata"`
	Name             string                 `json:"name"`
	Tags             []string               `json:"tags"`
	Tracing          bool                   `json:"tracing"`
}
