package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// Round holds the schema definition for the Round entity.
type Round struct {
	ent.Schema
}

// Fields of the Round.
func (Round) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			StructTag(`json:"id"`).
			Comment("The uuid of a round").
			Unique().
			Immutable().
			Default(uuid.New),
		field.Int("number").
			StructTag(`json:"number"`).
			Comment("The number of the round").
			Unique().
			NonNegative(),
		field.Bool("complete").
			StructTag(`json:"complete"`).
			Comment("The completion status of the round").
			Default(false),
	}
}

// Indexes of the Round.
func (Round) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("number"),
	}
}

// Mixins of the Round.
func (Round) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Edges of the Round.
func (Round) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("statuses", Status.Type).
			StructTag(`json:"statuses"`).
			Comment("The statuses of a round").
			Annotations(
				entsql.OnDelete(
					entsql.Cascade,
				),
			),
		edge.To("scoreCaches", ScoreCache.Type).
			StructTag(`json:"score_caches"`).
			Comment("The score caches of a round").
			Annotations(
				entsql.OnDelete(
					entsql.Cascade,
				),
			),
	}
}
