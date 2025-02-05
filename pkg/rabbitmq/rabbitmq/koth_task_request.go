package rabbitmq

const (
	KothTaskRequestExchange = "koth_task_request_exchange"
	KothTaskRequestVhost    = "koth_task_request_vhost"
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
	KothTaskRequestConfigurePermissions   = regex_amq_gen(KothTaskRequestExchange)
	KothTaskRequestMinionWritePermissions = regex("amq\\.gen-.*")
	KothTaskRequestMinionReadPermissions  = regex_amq_gen(KothTaskRequestExchange)
)
