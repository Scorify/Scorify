package schema

import (
	"fmt"
	"net"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/scorify/scorify/pkg/structs"
)

// Audit holds the schema definition for the Audit entity.
type Audit struct {
	ent.Schema
}

// Fields of the Audit.
func (Audit) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			StructTag(`json:"id"`).
			Comment("The uuid of the Audit Log").
			Unique().
			Immutable().
			Default(uuid.New),
		field.Enum("action").
			StructTag(`json:"action"`).
			Comment("The action of the audit log").
			Values(
				"auth_login",
				"auth_logout",
				"auth_failed_login",
				"admin_login",
				"admin_become",
				"user_change_password",
				"user_create",
				"user_update",
				"user_delete",
				"check_create",
				"check_update",
				"check_delete",
				"check_validate",
				"check_config",
				"koth_check_create",
				"koth_check_update",
				"koth_check_delete",
				"notification_create",
				"engine_start",
				"engine_stop",
				"inject_create",
				"inject_update",
				"inject_delete",
				"inject_submit",
				"inject_grade",
				"minion_register",
				"minion_deactivate",
				"minion_activate",
				"wipe_all",
				"wipe_check_configs",
				"wipe_inject_submissions",
				"wipe_statuses",
				"wipe_scores",
				"wipe_round",
				"wipe_cache",
			),
		field.String("ip").
			GoType(&structs.Inet{}).
			SchemaType(map[string]string{
				dialect.Postgres: "inet",
			}).
			Validate(func(s string) error {
				if net.ParseIP(s) == nil {
					return fmt.Errorf("invalid value for ip %q", s)
				}
				return nil
			}),
		field.Time("timestamp").
			StructTag(`json:"timestamp"`).
			Comment("The timestamp of the audit log").
			Immutable().
			Default(time.Now),
		field.String("message").
			StructTag(`json:"message"`).
			Comment("The message of the audit log").
			NotEmpty(),
		field.UUID("user_id", uuid.UUID{}).
			StructTag(`json:"user_id"`).
			Comment("The user responsible for the audit log").
			Optional().
			Immutable(),
	}
}

// Edges of the Audit.
func (Audit) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			StructTag(`json:"user"`).
			Comment("The user responsible for the audit log").
			Field("user_id").
			Immutable().
			Unique(),
	}
}
