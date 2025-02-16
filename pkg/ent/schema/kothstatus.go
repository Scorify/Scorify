package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// KothStatus holds the schema definition for the KothStatus entity.
type KothStatus struct {
	ent.Schema
}

// Fields of the Koth.
func (KothStatus) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			StructTag(`json:"id"`).
			Comment("The uuid of a koth status").
			Unique().
			Immutable().
			Default(uuid.New),
		field.UUID("user_id", uuid.UUID{}).
			StructTag(`json:"user_id"`).
			Comment("The uuid of a user").
			Nillable().
			Optional(),
		field.UUID("round_id", uuid.UUID{}).
			StructTag(`json:"round_id"`).
			Comment("The uuid of a round").
			Immutable(),
		field.UUID("minion_id", uuid.UUID{}).
			StructTag(`json:"minion_id"`).
			Comment("The uuid of a minion").
			Optional(),
		field.UUID("check_id", uuid.UUID{}).
			StructTag(`json:"check_id"`).
			Comment("The uuid of a check").
			Immutable(),
		field.Int("points").
			StructTag(`json:"points"`).
			Comment("The points of a koth status").
			NonNegative(),
		field.String("error").
			StructTag(`json:"error"`).
			Comment("The error of a koth status").
			Optional(),
	}
}

// Indexes of the KothStatus.
func (KothStatus) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("round_id"),
	}
}

// Mixins of the KothStatus.
func (KothStatus) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Edges of the KothStatus.
func (KothStatus) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("kothStatuses").
			Field("user_id").
			Unique(),
		edge.From("round", Round.Type).
			Ref("kothStatuses").
			Field("round_id").
			Required().
			Immutable().
			Unique(),
		edge.To("minion", Minion.Type).
			Field("minion_id").
			Unique(),
		edge.From("check", KothCheck.Type).
			Ref("statuses").
			Field("check_id").
			Immutable().
			Required().
			Unique(),
	}
}
