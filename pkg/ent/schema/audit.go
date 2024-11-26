package schema

import (
	"database/sql/driver"
	"fmt"
	"net"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

type Inet struct {
	net.IP
}

func (i *Inet) Scan(value any) (err error) {
	switch v := value.(type) {
	case nil:
	case []byte:
		if i.IP = net.ParseIP(string(v)); i.IP == nil {
			err = fmt.Errorf("invalid value for ip %q", v)
		}
	case string:
		if i.IP = net.ParseIP(v); i.IP == nil {
			err = fmt.Errorf("invalid value for ip %q", v)
		}
	default:
		err = fmt.Errorf("unexpected type %T", v)
	}
	return
}

func (i Inet) Value() (driver.Value, error) {
	return i.IP.String(), nil
}

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
				"auth:login",
				"auth:logout",
				"auth:failed_login",
				"admin:login",
				"admin:logout",
				"admin:become",
				"user:change_password",
				"user:create",
				"user:edit",
				"user:delete",
				"check:create",
				"check:edit",
				"check:delete",
				"check:validate",
				"check:config",
				"notification:create",
				"engine:start",
				"engine:stop",
				"inject:create",
				"inject:edit",
				"inject:delete",
				"inject:submit",
				"inject:grade",
				"minion:register",
				"minion:deactivate",
				"minion:activate",
				"wipe:all",
				"wipe:check_configs",
				"wipe:inject_submissions",
				"wipe:statuses",
				"wipe:scores",
				"wipe:round",
				"wipe:cache",
			),
		field.String("ip").
			GoType(&Inet{}).
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
	}
}

// Mixins of the Audit.
func (Audit) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Edges of the Audit.
func (Audit) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			StructTag(`json:"user"`).
			Comment("The user responsible for the audit log").
			Unique(),
	}
}
