// Code generated by ent, DO NOT EDIT.

package audit

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/scorify/scorify/pkg/ent/predicate"
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

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Audit {
	return predicate.Audit(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Audit {
	return predicate.Audit(sql.FieldEQ(FieldUpdateTime, v))
}

// Log applies equality check predicate on the "log" field. It's identical to LogEQ.
func Log(v string) predicate.Audit {
	return predicate.Audit(sql.FieldEQ(FieldLog, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Audit {
	return predicate.Audit(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Audit {
	return predicate.Audit(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Audit {
	return predicate.Audit(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Audit {
	return predicate.Audit(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Audit {
	return predicate.Audit(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Audit {
	return predicate.Audit(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Audit {
	return predicate.Audit(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Audit {
	return predicate.Audit(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Audit {
	return predicate.Audit(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Audit {
	return predicate.Audit(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Audit {
	return predicate.Audit(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Audit {
	return predicate.Audit(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.Audit {
	return predicate.Audit(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Audit {
	return predicate.Audit(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Audit {
	return predicate.Audit(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Audit {
	return predicate.Audit(sql.FieldLTE(FieldUpdateTime, v))
}

// ResourceEQ applies the EQ predicate on the "resource" field.
func ResourceEQ(v Resource) predicate.Audit {
	return predicate.Audit(sql.FieldEQ(FieldResource, v))
}

// ResourceNEQ applies the NEQ predicate on the "resource" field.
func ResourceNEQ(v Resource) predicate.Audit {
	return predicate.Audit(sql.FieldNEQ(FieldResource, v))
}

// ResourceIn applies the In predicate on the "resource" field.
func ResourceIn(vs ...Resource) predicate.Audit {
	return predicate.Audit(sql.FieldIn(FieldResource, vs...))
}

// ResourceNotIn applies the NotIn predicate on the "resource" field.
func ResourceNotIn(vs ...Resource) predicate.Audit {
	return predicate.Audit(sql.FieldNotIn(FieldResource, vs...))
}

// LogEQ applies the EQ predicate on the "log" field.
func LogEQ(v string) predicate.Audit {
	return predicate.Audit(sql.FieldEQ(FieldLog, v))
}

// LogNEQ applies the NEQ predicate on the "log" field.
func LogNEQ(v string) predicate.Audit {
	return predicate.Audit(sql.FieldNEQ(FieldLog, v))
}

// LogIn applies the In predicate on the "log" field.
func LogIn(vs ...string) predicate.Audit {
	return predicate.Audit(sql.FieldIn(FieldLog, vs...))
}

// LogNotIn applies the NotIn predicate on the "log" field.
func LogNotIn(vs ...string) predicate.Audit {
	return predicate.Audit(sql.FieldNotIn(FieldLog, vs...))
}

// LogGT applies the GT predicate on the "log" field.
func LogGT(v string) predicate.Audit {
	return predicate.Audit(sql.FieldGT(FieldLog, v))
}

// LogGTE applies the GTE predicate on the "log" field.
func LogGTE(v string) predicate.Audit {
	return predicate.Audit(sql.FieldGTE(FieldLog, v))
}

// LogLT applies the LT predicate on the "log" field.
func LogLT(v string) predicate.Audit {
	return predicate.Audit(sql.FieldLT(FieldLog, v))
}

// LogLTE applies the LTE predicate on the "log" field.
func LogLTE(v string) predicate.Audit {
	return predicate.Audit(sql.FieldLTE(FieldLog, v))
}

// LogContains applies the Contains predicate on the "log" field.
func LogContains(v string) predicate.Audit {
	return predicate.Audit(sql.FieldContains(FieldLog, v))
}

// LogHasPrefix applies the HasPrefix predicate on the "log" field.
func LogHasPrefix(v string) predicate.Audit {
	return predicate.Audit(sql.FieldHasPrefix(FieldLog, v))
}

// LogHasSuffix applies the HasSuffix predicate on the "log" field.
func LogHasSuffix(v string) predicate.Audit {
	return predicate.Audit(sql.FieldHasSuffix(FieldLog, v))
}

// LogEqualFold applies the EqualFold predicate on the "log" field.
func LogEqualFold(v string) predicate.Audit {
	return predicate.Audit(sql.FieldEqualFold(FieldLog, v))
}

// LogContainsFold applies the ContainsFold predicate on the "log" field.
func LogContainsFold(v string) predicate.Audit {
	return predicate.Audit(sql.FieldContainsFold(FieldLog, v))
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
