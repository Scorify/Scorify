package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// Minion holds the schema definition for the Minion entity.
type Minion struct {
	ent.Schema
}

// Fields of the Minion.
func (Minion) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			StructTag(`json:"id"`).
			Comment("The uuid of the minion").
			Unique().
			Immutable(),
		field.String("name").
			StructTag(`json:"name"`).
			Comment("The name of the minion"),
		field.String("ip").
			StructTag(`json:"ip"`).
			Comment("The ip of the minion"),
		field.Bool("deactivated").
			StructTag(`json:"deactivated"`).
			Comment("The deactivation status of the minion").
			Default(false),
		field.Enum("role").
			StructTag(`json:"role"`).
			Comment("The role of the minion").
			Values("koth", "service").
			Immutable(),
	}
}

// Indexes of the Minion.
func (Minion) Indexes() []ent.Index {
	return []ent.Index{}
}

// Mixins of the Minion.
func (Minion) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Edges of the Minion.
func (Minion) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("statuses", Status.Type).
			StructTag(`json:"status"`).
			Ref("minion"),
		edge.From("kothStatuses", KothStatus.Type).
			StructTag(`json:"kothStatuses"`).
			Ref("minion"),
	}
}
