package rabbitmq

const (
	KothTaskRequestQueue = "koth_task_request_queue_.*" // Match all queues with the prefix "koth_task_request_queue_"
	KothTaskRequestVhost = "koth_task_request_vhost"
)

var (
	// Permissions for minions in koth_task_request vhosts
	KothTaskRequestConfigurePermissions   = regex(KothTaskRequestQueue)
	KothTaskRequestMinionWritePermissions = regex("")
	KothTaskRequestMinionReadPermissions  = regex(KothTaskRequestQueue)
)
