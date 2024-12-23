// Code generated by ent, DO NOT EDIT.

package audit

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/scorify/scorify/pkg/ent/predicate"
	"github.com/scorify/scorify/pkg/structs"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Audit {
	return predicate.Audit(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Audit {
	return predicate.Audit(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Audit {
	return predicate.Audit(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Audit {
	return predicate.Audit(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Audit {
	return predicate.Audit(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Audit {
	return predicate.Audit(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Audit {
	return predicate.Audit(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Audit {
	return predicate.Audit(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Audit {
	return predicate.Audit(sql.FieldLTE(FieldID, id))
}

// IP applies equality check predicate on the "ip" field. It's identical to IPEQ.
func IP(v *structs.Inet) predicate.Audit {
	return predicate.Audit(sql.FieldEQ(FieldIP, v))
}

// Timestamp applies equality check predicate on the "timestamp" field. It's identical to TimestampEQ.
func Timestamp(v time.Time) predicate.Audit {
	return predicate.Audit(sql.FieldEQ(FieldTimestamp, v))
}

// Message applies equality check predicate on the "message" field. It's identical to MessageEQ.
func Message(v string) predicate.Audit {
	return predicate.Audit(sql.FieldEQ(FieldMessage, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uuid.UUID) predicate.Audit {
	return predicate.Audit(sql.FieldEQ(FieldUserID, v))
}

// ActionEQ applies the EQ predicate on the "action" field.
func ActionEQ(v Action) predicate.Audit {
	return predicate.Audit(sql.FieldEQ(FieldAction, v))
}

// ActionNEQ applies the NEQ predicate on the "action" field.
func ActionNEQ(v Action) predicate.Audit {
	return predicate.Audit(sql.FieldNEQ(FieldAction, v))
}

// ActionIn applies the In predicate on the "action" field.
func ActionIn(vs ...Action) predicate.Audit {
	return predicate.Audit(sql.FieldIn(FieldAction, vs...))
}

// ActionNotIn applies the NotIn predicate on the "action" field.
func ActionNotIn(vs ...Action) predicate.Audit {
	return predicate.Audit(sql.FieldNotIn(FieldAction, vs...))
}

// IPEQ applies the EQ predicate on the "ip" field.
func IPEQ(v *structs.Inet) predicate.Audit {
	return predicate.Audit(sql.FieldEQ(FieldIP, v))
}

// IPNEQ applies the NEQ predicate on the "ip" field.
func IPNEQ(v *structs.Inet) predicate.Audit {
	return predicate.Audit(sql.FieldNEQ(FieldIP, v))
}

// IPIn applies the In predicate on the "ip" field.
func IPIn(vs ...*structs.Inet) predicate.Audit {
	return predicate.Audit(sql.FieldIn(FieldIP, vs...))
}

// IPNotIn applies the NotIn predicate on the "ip" field.
func IPNotIn(vs ...*structs.Inet) predicate.Audit {
	return predicate.Audit(sql.FieldNotIn(FieldIP, vs...))
}

// IPGT applies the GT predicate on the "ip" field.
func IPGT(v *structs.Inet) predicate.Audit {
	return predicate.Audit(sql.FieldGT(FieldIP, v))
}

// IPGTE applies the GTE predicate on the "ip" field.
func IPGTE(v *structs.Inet) predicate.Audit {
	return predicate.Audit(sql.FieldGTE(FieldIP, v))
}

// IPLT applies the LT predicate on the "ip" field.
func IPLT(v *structs.Inet) predicate.Audit {
	return predicate.Audit(sql.FieldLT(FieldIP, v))
}

// IPLTE applies the LTE predicate on the "ip" field.
func IPLTE(v *structs.Inet) predicate.Audit {
	return predicate.Audit(sql.FieldLTE(FieldIP, v))
}

// IPContains applies the Contains predicate on the "ip" field.
func IPContains(v *structs.Inet) predicate.Audit {
	vc := v.String()
	return predicate.Audit(sql.FieldContains(FieldIP, vc))
}

// IPHasPrefix applies the HasPrefix predicate on the "ip" field.
func IPHasPrefix(v *structs.Inet) predicate.Audit {
	vc := v.String()
	return predicate.Audit(sql.FieldHasPrefix(FieldIP, vc))
}

// IPHasSuffix applies the HasSuffix predicate on the "ip" field.
func IPHasSuffix(v *structs.Inet) predicate.Audit {
	vc := v.String()
	return predicate.Audit(sql.FieldHasSuffix(FieldIP, vc))
}

// IPEqualFold applies the EqualFold predicate on the "ip" field.
func IPEqualFold(v *structs.Inet) predicate.Audit {
	vc := v.String()
	return predicate.Audit(sql.FieldEqualFold(FieldIP, vc))
}

// IPContainsFold applies the ContainsFold predicate on the "ip" field.
func IPContainsFold(v *structs.Inet) predicate.Audit {
	vc := v.String()
	return predicate.Audit(sql.FieldContainsFold(FieldIP, vc))
}

// TimestampEQ applies the EQ predicate on the "timestamp" field.
func TimestampEQ(v time.Time) predicate.Audit {
	return predicate.Audit(sql.FieldEQ(FieldTimestamp, v))
}

// TimestampNEQ applies the NEQ predicate on the "timestamp" field.
func TimestampNEQ(v time.Time) predicate.Audit {
	return predicate.Audit(sql.FieldNEQ(FieldTimestamp, v))
}

// TimestampIn applies the In predicate on the "timestamp" field.
func TimestampIn(vs ...time.Time) predicate.Audit {
	return predicate.Audit(sql.FieldIn(FieldTimestamp, vs...))
}

// TimestampNotIn applies the NotIn predicate on the "timestamp" field.
func TimestampNotIn(vs ...time.Time) predicate.Audit {
	return predicate.Audit(sql.FieldNotIn(FieldTimestamp, vs...))
}

// TimestampGT applies the GT predicate on the "timestamp" field.
func TimestampGT(v time.Time) predicate.Audit {
	return predicate.Audit(sql.FieldGT(FieldTimestamp, v))
}

// TimestampGTE applies the GTE predicate on the "timestamp" field.
func TimestampGTE(v time.Time) predicate.Audit {
	return predicate.Audit(sql.FieldGTE(FieldTimestamp, v))
}

// TimestampLT applies the LT predicate on the "timestamp" field.
func TimestampLT(v time.Time) predicate.Audit {
	return predicate.Audit(sql.FieldLT(FieldTimestamp, v))
}

// TimestampLTE applies the LTE predicate on the "timestamp" field.
func TimestampLTE(v time.Time) predicate.Audit {
	return predicate.Audit(sql.FieldLTE(FieldTimestamp, v))
}

// MessageEQ applies the EQ predicate on the "message" field.
func MessageEQ(v string) predicate.Audit {
	return predicate.Audit(sql.FieldEQ(FieldMessage, v))
}

// MessageNEQ applies the NEQ predicate on the "message" field.
func MessageNEQ(v string) predicate.Audit {
	return predicate.Audit(sql.FieldNEQ(FieldMessage, v))
}

// MessageIn applies the In predicate on the "message" field.
func MessageIn(vs ...string) predicate.Audit {
	return predicate.Audit(sql.FieldIn(FieldMessage, vs...))
}

// MessageNotIn applies the NotIn predicate on the "message" field.
func MessageNotIn(vs ...string) predicate.Audit {
	return predicate.Audit(sql.FieldNotIn(FieldMessage, vs...))
}

// MessageGT applies the GT predicate on the "message" field.
func MessageGT(v string) predicate.Audit {
	return predicate.Audit(sql.FieldGT(FieldMessage, v))
}

// MessageGTE applies the GTE predicate on the "message" field.
func MessageGTE(v string) predicate.Audit {
	return predicate.Audit(sql.FieldGTE(FieldMessage, v))
}

// MessageLT applies the LT predicate on the "message" field.
func MessageLT(v string) predicate.Audit {
	return predicate.Audit(sql.FieldLT(FieldMessage, v))
}

// MessageLTE applies the LTE predicate on the "message" field.
func MessageLTE(v string) predicate.Audit {
	return predicate.Audit(sql.FieldLTE(FieldMessage, v))
}

// MessageContains applies the Contains predicate on the "message" field.
func MessageContains(v string) predicate.Audit {
	return predicate.Audit(sql.FieldContains(FieldMessage, v))
}

// MessageHasPrefix applies the HasPrefix predicate on the "message" field.
func MessageHasPrefix(v string) predicate.Audit {
	return predicate.Audit(sql.FieldHasPrefix(FieldMessage, v))
}

// MessageHasSuffix applies the HasSuffix predicate on the "message" field.
func MessageHasSuffix(v string) predicate.Audit {
	return predicate.Audit(sql.FieldHasSuffix(FieldMessage, v))
}

// MessageEqualFold applies the EqualFold predicate on the "message" field.
func MessageEqualFold(v string) predicate.Audit {
	return predicate.Audit(sql.FieldEqualFold(FieldMessage, v))
}

// MessageContainsFold applies the ContainsFold predicate on the "message" field.
func MessageContainsFold(v string) predicate.Audit {
	return predicate.Audit(sql.FieldContainsFold(FieldMessage, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uuid.UUID) predicate.Audit {
	return predicate.Audit(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uuid.UUID) predicate.Audit {
	return predicate.Audit(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uuid.UUID) predicate.Audit {
	return predicate.Audit(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uuid.UUID) predicate.Audit {
	return predicate.Audit(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDIsNil applies the IsNil predicate on the "user_id" field.
func UserIDIsNil() predicate.Audit {
	return predicate.Audit(sql.FieldIsNull(FieldUserID))
}

// UserIDNotNil applies the NotNil predicate on the "user_id" field.
func UserIDNotNil() predicate.Audit {
	return predicate.Audit(sql.FieldNotNull(FieldUserID))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Audit {
	return predicate.Audit(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Audit {
	return predicate.Audit(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Audit) predicate.Audit {
	return predicate.Audit(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Audit) predicate.Audit {
	return predicate.Audit(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Audit) predicate.Audit {
	return predicate.Audit(sql.NotPredicates(p))
}
