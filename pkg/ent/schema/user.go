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

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			StructTag(`json:"id"`).
			Comment("The uuid of the user").
			Unique().
			Immutable().
			Default(uuid.New),
		field.String("username").
			StructTag(`json:"username"`).
			Comment("The username of the user").
			Unique().
			NotEmpty(),
		field.String("password").
			Comment("The password hash of user password").
			Sensitive().
			NotEmpty(),
		field.Enum("role").
			StructTag(`json:"role"`).
			Comment("The role of the user").
			Values("admin", "user").
			Default("user").
			Immutable(),
		field.Int("number").
			StructTag(`json:"number"`).
			Comment("The number of the user").
			Optional().
			Unique().
			Positive(),
	}
}

// Indexes of the User.
func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("username"),
	}
}

// Mixins of the User.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("configs", CheckConfig.Type).
			StructTag(`json:"config"`).
			Comment("The configuration of a check").
			Annotations(
				entsql.OnDelete(
					entsql.Cascade,
				),
			),
		edge.To("statuses", Status.Type).
			StructTag(`json:"status"`).
			Comment("The status of a user").
			Annotations(
				entsql.OnDelete(
					entsql.Cascade,
				),
			),
		edge.To("scoreCaches", ScoreCache.Type).
			StructTag(`json:"score_caches"`).
			Comment("The score caches of a user").
			Annotations(
				entsql.OnDelete(
					entsql.Cascade,
				),
			),
		edge.To("submissions", InjectSubmission.Type).
			StructTag(`json:"submissions"`).
			Comment("The submissions of a user").
			Annotations(
				entsql.OnDelete(
					entsql.Cascade,
				),
			),
	}
}
