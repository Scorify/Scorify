package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
	"github.com/scorify/scorify/pkg/structs"
)

// Inject holds the schema definition for the Inject entity.
type Inject struct {
	ent.Schema
}

// Fields of the Inject.
func (Inject) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			StructTag(`json:"id"`).
			Comment("The uuid of an inject").
			Unique().
			Immutable().
			Default(uuid.New),
		field.String("title").
			StructTag(`json:"title"`).
			Comment("The title of the inject").
			Unique().
			NotEmpty(),
		field.Time("start_time").
			StructTag(`json:"start_time"`).
			Comment("The start time of the inject"),
		field.Time("end_time").
			StructTag(`json:"end_time"`).
			Comment("The end time of the inject"),
		field.JSON("files", []structs.File{}).
			StructTag(`json:"files"`).
			Comment("The files of the inject"),
		field.JSON("rubric", structs.RubricTemplate{}).
			StructTag(`json:"rubric"`).
			Comment("The rubric of the inject"),
	}
}

// Indexes of the Inject.
func (Inject) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("title"),
	}
}

// Mixins of the Inject.
func (Inject) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Edges of the Inject.
func (Inject) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("submissions", InjectSubmission.Type).
			StructTag(`json:"submissions"`).
			Comment("The submissions of the inject").
			Annotations(
				entsql.OnDelete(
					entsql.Cascade,
				),
			),
	}
}
