package rabbitmq

const (
	KothTaskRequestQueue = "koth_task_request_queue_.*" // Match all queues with the prefix "koth_task_request_queue_"
	KothTaskRequestVhost = "koth_task_request_vhost"
)

/* ┌──────────────────────────────────────────────────────────────────┐
 * │                             topic_key:                           │
 * │                         koth.[check_name]                        │
 * │                                                                  │
 * │                 koth.check_1                                     │
 * │                 koth.check_2                                     │
 * │                      ┌─────────────►  Queue 1 ─────► Koth Minion │
 * │                      │                                           │
 * │                      │                                           │
 * │                      │                                           │
 * │                        koth.check_3                              │
 * │ Server ─────► Exchange ────────────►  Queue 2 ─────► Koth Minion │
 * │                                                                  │
 * │                      │                                           │
 * │                      │                                           │
 * │                      │                                           │
 * │                      └─────────────►  Queue 3 ─────► Koth Minion │
 * │                 koth.check_4                                     │
 * └──────────────────────────────────────────────────────────────────┘
 */

var (
	// Permissions for minions in koth_task_request vhosts
	KothTaskRequestConfigurePermissions   = regex(KothTaskRequestQueue)
	KothTaskRequestMinionWritePermissions = regex("")
	KothTaskRequestMinionReadPermissions  = regex(KothTaskRequestQueue)
)
