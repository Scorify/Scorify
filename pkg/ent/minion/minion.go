// Code generated by ent, DO NOT EDIT.

package minion

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the minion type in the database.
	Label = "minion"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldIP holds the string denoting the ip field in the database.
	FieldIP = "ip"
	// FieldDeactivated holds the string denoting the deactivated field in the database.
	FieldDeactivated = "deactivated"
	// FieldRole holds the string denoting the role field in the database.
	FieldRole = "role"
	// EdgeStatuses holds the string denoting the statuses edge name in mutations.
	EdgeStatuses = "statuses"
	// EdgeKothStatuses holds the string denoting the kothstatuses edge name in mutations.
	EdgeKothStatuses = "kothStatuses"
	// Table holds the table name of the minion in the database.
	Table = "minions"
	// StatusesTable is the table that holds the statuses relation/edge.
	StatusesTable = "status"
	// StatusesInverseTable is the table name for the Status entity.
	// It exists in this package in order to avoid circular dependency with the "status" package.
	StatusesInverseTable = "status"
	// StatusesColumn is the table column denoting the statuses relation/edge.
	StatusesColumn = "minion_id"
	// KothStatusesTable is the table that holds the kothStatuses relation/edge.
	KothStatusesTable = "koth_status"
	// KothStatusesInverseTable is the table name for the KothStatus entity.
	// It exists in this package in order to avoid circular dependency with the "kothstatus" package.
	KothStatusesInverseTable = "koth_status"
	// KothStatusesColumn is the table column denoting the kothStatuses relation/edge.
	KothStatusesColumn = "minion_id"
)

// Columns holds all SQL columns for minion fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldName,
	FieldIP,
	FieldDeactivated,
	FieldRole,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// DefaultDeactivated holds the default value on creation for the "deactivated" field.
	DefaultDeactivated bool
)

// Role defines the type for the "role" enum field.
type Role string

// Role values.
const (
	RoleKoth    Role = "koth"
	RoleService Role = "service"
)

func (r Role) String() string {
	return string(r)
}

// RoleValidator is a validator for the "role" field enum values. It is called by the builders before save.
func RoleValidator(r Role) error {
	switch r {
	case RoleKoth, RoleService:
		return nil
	default:
		return fmt.Errorf("minion: invalid enum value for role field: %q", r)
	}
}

// OrderOption defines the ordering options for the Minion queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByUpdateTime orders the results by the update_time field.
func ByUpdateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateTime, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByIP orders the results by the ip field.
func ByIP(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIP, opts...).ToFunc()
}

// ByDeactivated orders the results by the deactivated field.
func ByDeactivated(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeactivated, opts...).ToFunc()
}

// ByRole orders the results by the role field.
func ByRole(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRole, opts...).ToFunc()
}

// ByStatusesCount orders the results by statuses count.
func ByStatusesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newStatusesStep(), opts...)
	}
}

// ByStatuses orders the results by statuses terms.
func ByStatuses(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newStatusesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByKothStatusesCount orders the results by kothStatuses count.
func ByKothStatusesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newKothStatusesStep(), opts...)
	}
}

// ByKothStatuses orders the results by kothStatuses terms.
func ByKothStatuses(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newKothStatusesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newStatusesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(StatusesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, StatusesTable, StatusesColumn),
	)
}
func newKothStatusesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(KothStatusesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, KothStatusesTable, KothStatusesColumn),
	)
}
