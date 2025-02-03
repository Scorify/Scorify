package rabbitmq

const (
	KothTaskResponseQueue = "koth_task_response_queue"
	KothTaskResponseVhost = "koth_task_response_vhost"
)

var (
	// Permissions for minions in koth_task_response vhosts
	KothTaskResponseConfigurePermissions   = regex(KothTaskResponseQueue)
	KothTaskResponseMinionWritePermissions = regex_amq_default(KothTaskResponseQueue)
	KothTaskResponseMinionReadPermissions  = regex("")
)
