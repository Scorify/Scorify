package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// Status holds the schema definition for the Status entity.
type Status struct {
	ent.Schema
}

// Fields of the Status.
func (Status) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			StructTag(`json:"id"`).
			Comment("The uuid of a status").
			Unique().
			Immutable().
			Default(uuid.New),
		field.String("error").
			StructTag(`json:"error"`).
			Comment("The error message of the status").
			Optional(),
		field.Enum("status").
			StructTag(`json:"status"`).
			Comment("The status of the status").
			Values("up", "down", "unknown").
			Default("unknown"),
		field.Int("points").
			StructTag(`json:"points"`).
			Comment("The points of the status").
			NonNegative(),
		field.UUID("check_id", uuid.UUID{}).
			StructTag(`json:"check_id"`).
			Comment("The uuid of a check").
			Immutable(),
		field.UUID("round_id", uuid.UUID{}).
			StructTag(`json:"round_id"`).
			Comment("The uuid of a round").
			Immutable(),
		field.UUID("user_id", uuid.UUID{}).
			StructTag(`json:"user_id"`).
			Comment("The uuid of a user").
			Immutable(),
	}
}

// Indexes of the Status.
func (Status) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("check_id", "round_id", "user_id"),
	}
}

// Mixins of the Status.
func (Status) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Edges of the Status.
func (Status) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("check", Check.Type).
			StructTag(`json:"check"`).
			Comment("The check of a status").
			Field("check_id").
			Immutable().
			Required().
			Unique().
			Ref("statuses"),
		edge.From("round", Round.Type).
			StructTag(`json:"round"`).
			Comment("The round of a status").
			Field("round_id").
			Immutable().
			Required().
			Unique().
			Ref("statuses"),
		edge.From("user", User.Type).
			StructTag(`json:"user"`).
			Comment("The user of a status").
			Field("user_id").
			Immutable().
			Required().
			Unique().
			Ref("statuses"),
	}
}
