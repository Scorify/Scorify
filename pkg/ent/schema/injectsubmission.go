package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
	"github.com/scorify/scorify/pkg/structs"
)

// InjectSubmission holds the schema definition for the InjectSubmission entity.
type InjectSubmission struct {
	ent.Schema
}

// Fields of the InjectSubmission.
func (InjectSubmission) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			StructTag(`json:"id"`).
			Comment("The uuid of an inject submission").
			Unique().
			Immutable().
			Default(uuid.New),
		field.JSON("files", []structs.File{}).
			StructTag(`json:"files"`).
			Comment("The files of the inject submission"),
		field.UUID("inject_id", uuid.UUID{}).
			StructTag(`json:"inject_id"`).
			Comment("The inject this submission belongs to").
			Immutable(),
		field.UUID("user_id", uuid.UUID{}).
			StructTag(`json:"user_id"`).
			Comment("The user this submission belongs to").
			Immutable(),
		field.String("notes").
			StructTag(`json:"notes"`).
			Comment("The notes of the inject submission"),
		field.JSON("rubric", &structs.Rubric{}).
			StructTag(`json:"rubric"`).
			Comment("The rubric of the inject submission").
			Optional(),
		field.Bool("graded").
			StructTag(`json:"graded"`).
			Comment("The graded status of the inject submission").
			Default(false),
	}
}

// Mixins of the InjectSubmission.
func (InjectSubmission) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Edges of the InjectSubmission.
func (InjectSubmission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("inject", Inject.Type).
			StructTag(`json:"inject"`).
			Comment("The inject this submission belongs to").
			Field("inject_id").
			Immutable().
			Required().
			Unique().
			Ref("submissions"),
		edge.From("user", User.Type).
			StructTag(`json:"user"`).
			Comment("The user this submission belongs to").
			Field("user_id").
			Immutable().
			Required().
			Unique().
			Ref("submissions"),
	}
}
