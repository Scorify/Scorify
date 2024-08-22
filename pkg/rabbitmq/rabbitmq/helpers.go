package rabbitmq

import "fmt"

func regex(resource string) string {
	return fmt.Sprintf("^%s$", resource)
}

func regex_amq_default(resource string) string {
	return fmt.Sprintf("^(amq\\.default|%s)$", resource)
}

func regex_amq_gen(resource string) string {
	return fmt.Sprintf("^(amq\\.gen.*|%s)$", resource)
}
