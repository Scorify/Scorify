package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// CheckConfig holds the schema definition for the CheckConfig entity.
type CheckConfig struct {
	ent.Schema
}

// Fields of the CheckConfig.
func (CheckConfig) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			StructTag(`json:"id"`).
			Comment("The uuid of a check configuration").
			Unique().
			Immutable().
			Default(uuid.New),
		field.JSON("config", map[string]interface{}{}).
			StructTag(`json:"config"`).
			Comment("The configuration of a check"),
		field.UUID("check_id", uuid.UUID{}).
			StructTag(`json:"check_id"`).
			Comment("The check this configuration belongs to").
			Immutable(),
		field.UUID("user_id", uuid.UUID{}).
			StructTag(`json:"user_id"`).
			Comment("The user this configuration belongs to").
			Immutable(),
	}
}

// Indexes of the CheckConfig.
func (CheckConfig) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("check_id", "user_id"),
	}
}

// Mixins of the CheckConfig.
func (CheckConfig) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Edges of the CheckConfig.
func (CheckConfig) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("check", Check.Type).
			StructTag(`json:"check"`).
			Comment("The check this configuration belongs to").
			Field("check_id").
			Immutable().
			Required().
			Unique().
			Ref("configs"),
		edge.From("user", User.Type).
			StructTag(`json:"user"`).
			Comment("The user this configuration belongs to").
			Field("user_id").
			Immutable().
			Required().
			Unique().
			Ref("configs"),
	}
}
