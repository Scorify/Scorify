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

// KothCheck holds the schema definition for the KothCheck entity.
type KothCheck struct {
	ent.Schema
}

// Fields of the Koth.
func (KothCheck) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			StructTag(`json:"id"`).
			Comment("The uuid of a koth status").
			Unique().
			Immutable().
			Default(uuid.New),
		field.String("name").
			StructTag(`json:"name"`).
			Comment("The name of the check").
			NotEmpty().
			Unique(),
		field.String("file").
			StructTag(`json:"file"`).
			Comment("The file of the check").
			NotEmpty(),
	}
}

// Indexes of the KothCheck.
func (KothCheck) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name"),
	}
}

// Mixins of the KothCheck.
func (KothCheck) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Edges of the KothCheck.
func (KothCheck) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("statuses", KothStatus.Type).
			StructTag(`json:"statuses"`).
			Comment("The statuses of a check").
			Annotations(
				entsql.OnDelete(
					entsql.Cascade,
				),
			),
	}
}
