package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// ScoreCache holds the schema definition for the ScoreCache entity.
type ScoreCache struct {
	ent.Schema
}

// Fields of the ScoreCache.
func (ScoreCache) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			StructTag(`json:"id"`).
			Comment("The uuid of a score cache").
			Unique().
			Immutable().
			Default(uuid.New),
		field.Int("points").
			StructTag(`json:"points"`).
			Comment("The points of the round").
			NonNegative(),
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

// Indexes of the ScoreCache.
func (ScoreCache) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("round_id", "user_id"),
	}
}

// Mixins of the ScoreCache.
func (ScoreCache) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Edges of the ScoreCache.
func (ScoreCache) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("round", Round.Type).
			StructTag(`json:"round"`).
			Comment("The round of a score cache").
			Field("round_id").
			Immutable().
			Required().
			Unique().
			Ref("scoreCaches"),
		edge.From("user", User.Type).
			StructTag(`json:"user"`).
			Comment("The user of a score cache").
			Field("user_id").
			Immutable().
			Required().
			Unique().
			Ref("scoreCaches"),
	}
}
