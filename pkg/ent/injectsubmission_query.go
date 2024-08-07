// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/scorify/scorify/pkg/ent/inject"
	"github.com/scorify/scorify/pkg/ent/injectsubmission"
	"github.com/scorify/scorify/pkg/ent/predicate"
	"github.com/scorify/scorify/pkg/ent/user"
)

// InjectSubmissionQuery is the builder for querying InjectSubmission entities.
type InjectSubmissionQuery struct {
	config
	ctx        *QueryContext
	order      []injectsubmission.OrderOption
	inters     []Interceptor
	predicates []predicate.InjectSubmission
	withInject *InjectQuery
	withUser   *UserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the InjectSubmissionQuery builder.
func (isq *InjectSubmissionQuery) Where(ps ...predicate.InjectSubmission) *InjectSubmissionQuery {
	isq.predicates = append(isq.predicates, ps...)
	return isq
}

// Limit the number of records to be returned by this query.
func (isq *InjectSubmissionQuery) Limit(limit int) *InjectSubmissionQuery {
	isq.ctx.Limit = &limit
	return isq
}

// Offset to start from.
func (isq *InjectSubmissionQuery) Offset(offset int) *InjectSubmissionQuery {
	isq.ctx.Offset = &offset
	return isq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (isq *InjectSubmissionQuery) Unique(unique bool) *InjectSubmissionQuery {
	isq.ctx.Unique = &unique
	return isq
}

// Order specifies how the records should be ordered.
func (isq *InjectSubmissionQuery) Order(o ...injectsubmission.OrderOption) *InjectSubmissionQuery {
	isq.order = append(isq.order, o...)
	return isq
}

// QueryInject chains the current query on the "inject" edge.
func (isq *InjectSubmissionQuery) QueryInject() *InjectQuery {
	query := (&InjectClient{config: isq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := isq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := isq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(injectsubmission.Table, injectsubmission.FieldID, selector),
			sqlgraph.To(inject.Table, inject.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, injectsubmission.InjectTable, injectsubmission.InjectColumn),
		)
		fromU = sqlgraph.SetNeighbors(isq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUser chains the current query on the "user" edge.
func (isq *InjectSubmissionQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: isq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := isq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := isq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(injectsubmission.Table, injectsubmission.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, injectsubmission.UserTable, injectsubmission.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(isq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first InjectSubmission entity from the query.
// Returns a *NotFoundError when no InjectSubmission was found.
func (isq *InjectSubmissionQuery) First(ctx context.Context) (*InjectSubmission, error) {
	nodes, err := isq.Limit(1).All(setContextOp(ctx, isq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{injectsubmission.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (isq *InjectSubmissionQuery) FirstX(ctx context.Context) *InjectSubmission {
	node, err := isq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first InjectSubmission ID from the query.
// Returns a *NotFoundError when no InjectSubmission ID was found.
func (isq *InjectSubmissionQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = isq.Limit(1).IDs(setContextOp(ctx, isq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{injectsubmission.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (isq *InjectSubmissionQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := isq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single InjectSubmission entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one InjectSubmission entity is found.
// Returns a *NotFoundError when no InjectSubmission entities are found.
func (isq *InjectSubmissionQuery) Only(ctx context.Context) (*InjectSubmission, error) {
	nodes, err := isq.Limit(2).All(setContextOp(ctx, isq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{injectsubmission.Label}
	default:
		return nil, &NotSingularError{injectsubmission.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (isq *InjectSubmissionQuery) OnlyX(ctx context.Context) *InjectSubmission {
	node, err := isq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only InjectSubmission ID in the query.
// Returns a *NotSingularError when more than one InjectSubmission ID is found.
// Returns a *NotFoundError when no entities are found.
func (isq *InjectSubmissionQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = isq.Limit(2).IDs(setContextOp(ctx, isq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{injectsubmission.Label}
	default:
		err = &NotSingularError{injectsubmission.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (isq *InjectSubmissionQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := isq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of InjectSubmissions.
func (isq *InjectSubmissionQuery) All(ctx context.Context) ([]*InjectSubmission, error) {
	ctx = setContextOp(ctx, isq.ctx, "All")
	if err := isq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*InjectSubmission, *InjectSubmissionQuery]()
	return withInterceptors[[]*InjectSubmission](ctx, isq, qr, isq.inters)
}

// AllX is like All, but panics if an error occurs.
func (isq *InjectSubmissionQuery) AllX(ctx context.Context) []*InjectSubmission {
	nodes, err := isq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of InjectSubmission IDs.
func (isq *InjectSubmissionQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if isq.ctx.Unique == nil && isq.path != nil {
		isq.Unique(true)
	}
	ctx = setContextOp(ctx, isq.ctx, "IDs")
	if err = isq.Select(injectsubmission.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (isq *InjectSubmissionQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := isq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (isq *InjectSubmissionQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, isq.ctx, "Count")
	if err := isq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, isq, querierCount[*InjectSubmissionQuery](), isq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (isq *InjectSubmissionQuery) CountX(ctx context.Context) int {
	count, err := isq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (isq *InjectSubmissionQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, isq.ctx, "Exist")
	switch _, err := isq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (isq *InjectSubmissionQuery) ExistX(ctx context.Context) bool {
	exist, err := isq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the InjectSubmissionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (isq *InjectSubmissionQuery) Clone() *InjectSubmissionQuery {
	if isq == nil {
		return nil
	}
	return &InjectSubmissionQuery{
		config:     isq.config,
		ctx:        isq.ctx.Clone(),
		order:      append([]injectsubmission.OrderOption{}, isq.order...),
		inters:     append([]Interceptor{}, isq.inters...),
		predicates: append([]predicate.InjectSubmission{}, isq.predicates...),
		withInject: isq.withInject.Clone(),
		withUser:   isq.withUser.Clone(),
		// clone intermediate query.
		sql:  isq.sql.Clone(),
		path: isq.path,
	}
}

// WithInject tells the query-builder to eager-load the nodes that are connected to
// the "inject" edge. The optional arguments are used to configure the query builder of the edge.
func (isq *InjectSubmissionQuery) WithInject(opts ...func(*InjectQuery)) *InjectSubmissionQuery {
	query := (&InjectClient{config: isq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	isq.withInject = query
	return isq
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (isq *InjectSubmissionQuery) WithUser(opts ...func(*UserQuery)) *InjectSubmissionQuery {
	query := (&UserClient{config: isq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	isq.withUser = query
	return isq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.InjectSubmission.Query().
//		GroupBy(injectsubmission.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (isq *InjectSubmissionQuery) GroupBy(field string, fields ...string) *InjectSubmissionGroupBy {
	isq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &InjectSubmissionGroupBy{build: isq}
	grbuild.flds = &isq.ctx.Fields
	grbuild.label = injectsubmission.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.InjectSubmission.Query().
//		Select(injectsubmission.FieldCreateTime).
//		Scan(ctx, &v)
func (isq *InjectSubmissionQuery) Select(fields ...string) *InjectSubmissionSelect {
	isq.ctx.Fields = append(isq.ctx.Fields, fields...)
	sbuild := &InjectSubmissionSelect{InjectSubmissionQuery: isq}
	sbuild.label = injectsubmission.Label
	sbuild.flds, sbuild.scan = &isq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a InjectSubmissionSelect configured with the given aggregations.
func (isq *InjectSubmissionQuery) Aggregate(fns ...AggregateFunc) *InjectSubmissionSelect {
	return isq.Select().Aggregate(fns...)
}

func (isq *InjectSubmissionQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range isq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, isq); err != nil {
				return err
			}
		}
	}
	for _, f := range isq.ctx.Fields {
		if !injectsubmission.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if isq.path != nil {
		prev, err := isq.path(ctx)
		if err != nil {
			return err
		}
		isq.sql = prev
	}
	return nil
}

func (isq *InjectSubmissionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*InjectSubmission, error) {
	var (
		nodes       = []*InjectSubmission{}
		_spec       = isq.querySpec()
		loadedTypes = [2]bool{
			isq.withInject != nil,
			isq.withUser != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*InjectSubmission).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &InjectSubmission{config: isq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, isq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := isq.withInject; query != nil {
		if err := isq.loadInject(ctx, query, nodes, nil,
			func(n *InjectSubmission, e *Inject) { n.Edges.Inject = e }); err != nil {
			return nil, err
		}
	}
	if query := isq.withUser; query != nil {
		if err := isq.loadUser(ctx, query, nodes, nil,
			func(n *InjectSubmission, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (isq *InjectSubmissionQuery) loadInject(ctx context.Context, query *InjectQuery, nodes []*InjectSubmission, init func(*InjectSubmission), assign func(*InjectSubmission, *Inject)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*InjectSubmission)
	for i := range nodes {
		fk := nodes[i].InjectID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(inject.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "inject_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (isq *InjectSubmissionQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*InjectSubmission, init func(*InjectSubmission), assign func(*InjectSubmission, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*InjectSubmission)
	for i := range nodes {
		fk := nodes[i].UserID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (isq *InjectSubmissionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := isq.querySpec()
	_spec.Node.Columns = isq.ctx.Fields
	if len(isq.ctx.Fields) > 0 {
		_spec.Unique = isq.ctx.Unique != nil && *isq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, isq.driver, _spec)
}

func (isq *InjectSubmissionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(injectsubmission.Table, injectsubmission.Columns, sqlgraph.NewFieldSpec(injectsubmission.FieldID, field.TypeUUID))
	_spec.From = isq.sql
	if unique := isq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if isq.path != nil {
		_spec.Unique = true
	}
	if fields := isq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, injectsubmission.FieldID)
		for i := range fields {
			if fields[i] != injectsubmission.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if isq.withInject != nil {
			_spec.Node.AddColumnOnce(injectsubmission.FieldInjectID)
		}
		if isq.withUser != nil {
			_spec.Node.AddColumnOnce(injectsubmission.FieldUserID)
		}
	}
	if ps := isq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := isq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := isq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := isq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (isq *InjectSubmissionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(isq.driver.Dialect())
	t1 := builder.Table(injectsubmission.Table)
	columns := isq.ctx.Fields
	if len(columns) == 0 {
		columns = injectsubmission.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if isq.sql != nil {
		selector = isq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if isq.ctx.Unique != nil && *isq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range isq.predicates {
		p(selector)
	}
	for _, p := range isq.order {
		p(selector)
	}
	if offset := isq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := isq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// InjectSubmissionGroupBy is the group-by builder for InjectSubmission entities.
type InjectSubmissionGroupBy struct {
	selector
	build *InjectSubmissionQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (isgb *InjectSubmissionGroupBy) Aggregate(fns ...AggregateFunc) *InjectSubmissionGroupBy {
	isgb.fns = append(isgb.fns, fns...)
	return isgb
}

// Scan applies the selector query and scans the result into the given value.
func (isgb *InjectSubmissionGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, isgb.build.ctx, "GroupBy")
	if err := isgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*InjectSubmissionQuery, *InjectSubmissionGroupBy](ctx, isgb.build, isgb, isgb.build.inters, v)
}

func (isgb *InjectSubmissionGroupBy) sqlScan(ctx context.Context, root *InjectSubmissionQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(isgb.fns))
	for _, fn := range isgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*isgb.flds)+len(isgb.fns))
		for _, f := range *isgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*isgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := isgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// InjectSubmissionSelect is the builder for selecting fields of InjectSubmission entities.
type InjectSubmissionSelect struct {
	*InjectSubmissionQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (iss *InjectSubmissionSelect) Aggregate(fns ...AggregateFunc) *InjectSubmissionSelect {
	iss.fns = append(iss.fns, fns...)
	return iss
}

// Scan applies the selector query and scans the result into the given value.
func (iss *InjectSubmissionSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, iss.ctx, "Select")
	if err := iss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*InjectSubmissionQuery, *InjectSubmissionSelect](ctx, iss.InjectSubmissionQuery, iss, iss.inters, v)
}

func (iss *InjectSubmissionSelect) sqlScan(ctx context.Context, root *InjectSubmissionQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(iss.fns))
	for _, fn := range iss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*iss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := iss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
