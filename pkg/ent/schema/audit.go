package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
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
		field.Enum("resource").
			StructTag(`json:"resource"`).
			Comment("The resource of the audit log").
			Values(
				"authentication", // User authentication (logins)
				"checks",         // Edits to checks
				"database",       // Database changes (wipes)
				"engine_state",   // Engine state changes (start, stop)
				"injects",        // Edits to injects
				"notifications",  // Notifications sent
				"other",          // Other changes
				"users",          // Edits to users (add, remove, edit, change password, ...)
			),
		field.String("log").
			StructTag(`json:"log"`).
			Comment("The log message of the audit log").
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
