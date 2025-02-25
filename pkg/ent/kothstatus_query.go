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
	"github.com/scorify/scorify/pkg/ent/kothcheck"
	"github.com/scorify/scorify/pkg/ent/kothstatus"
	"github.com/scorify/scorify/pkg/ent/minion"
	"github.com/scorify/scorify/pkg/ent/predicate"
	"github.com/scorify/scorify/pkg/ent/round"
	"github.com/scorify/scorify/pkg/ent/user"
)

// KothStatusQuery is the builder for querying KothStatus entities.
type KothStatusQuery struct {
	config
	ctx        *QueryContext
	order      []kothstatus.OrderOption
	inters     []Interceptor
	predicates []predicate.KothStatus
	withUser   *UserQuery
	withRound  *RoundQuery
	withMinion *MinionQuery
	withCheck  *KothCheckQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the KothStatusQuery builder.
func (ksq *KothStatusQuery) Where(ps ...predicate.KothStatus) *KothStatusQuery {
	ksq.predicates = append(ksq.predicates, ps...)
	return ksq
}

// Limit the number of records to be returned by this query.
func (ksq *KothStatusQuery) Limit(limit int) *KothStatusQuery {
	ksq.ctx.Limit = &limit
	return ksq
}

// Offset to start from.
func (ksq *KothStatusQuery) Offset(offset int) *KothStatusQuery {
	ksq.ctx.Offset = &offset
	return ksq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ksq *KothStatusQuery) Unique(unique bool) *KothStatusQuery {
	ksq.ctx.Unique = &unique
	return ksq
}

// Order specifies how the records should be ordered.
func (ksq *KothStatusQuery) Order(o ...kothstatus.OrderOption) *KothStatusQuery {
	ksq.order = append(ksq.order, o...)
	return ksq
}

// QueryUser chains the current query on the "user" edge.
func (ksq *KothStatusQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: ksq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ksq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ksq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(kothstatus.Table, kothstatus.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, kothstatus.UserTable, kothstatus.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(ksq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRound chains the current query on the "round" edge.
func (ksq *KothStatusQuery) QueryRound() *RoundQuery {
	query := (&RoundClient{config: ksq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ksq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ksq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(kothstatus.Table, kothstatus.FieldID, selector),
			sqlgraph.To(round.Table, round.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, kothstatus.RoundTable, kothstatus.RoundColumn),
		)
		fromU = sqlgraph.SetNeighbors(ksq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMinion chains the current query on the "minion" edge.
func (ksq *KothStatusQuery) QueryMinion() *MinionQuery {
	query := (&MinionClient{config: ksq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ksq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ksq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(kothstatus.Table, kothstatus.FieldID, selector),
			sqlgraph.To(minion.Table, minion.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, kothstatus.MinionTable, kothstatus.MinionColumn),
		)
		fromU = sqlgraph.SetNeighbors(ksq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCheck chains the current query on the "check" edge.
func (ksq *KothStatusQuery) QueryCheck() *KothCheckQuery {
	query := (&KothCheckClient{config: ksq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ksq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ksq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(kothstatus.Table, kothstatus.FieldID, selector),
			sqlgraph.To(kothcheck.Table, kothcheck.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, kothstatus.CheckTable, kothstatus.CheckColumn),
		)
		fromU = sqlgraph.SetNeighbors(ksq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first KothStatus entity from the query.
// Returns a *NotFoundError when no KothStatus was found.
func (ksq *KothStatusQuery) First(ctx context.Context) (*KothStatus, error) {
	nodes, err := ksq.Limit(1).All(setContextOp(ctx, ksq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{kothstatus.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ksq *KothStatusQuery) FirstX(ctx context.Context) *KothStatus {
	node, err := ksq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first KothStatus ID from the query.
// Returns a *NotFoundError when no KothStatus ID was found.
func (ksq *KothStatusQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ksq.Limit(1).IDs(setContextOp(ctx, ksq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{kothstatus.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ksq *KothStatusQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := ksq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single KothStatus entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one KothStatus entity is found.
// Returns a *NotFoundError when no KothStatus entities are found.
func (ksq *KothStatusQuery) Only(ctx context.Context) (*KothStatus, error) {
	nodes, err := ksq.Limit(2).All(setContextOp(ctx, ksq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{kothstatus.Label}
	default:
		return nil, &NotSingularError{kothstatus.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ksq *KothStatusQuery) OnlyX(ctx context.Context) *KothStatus {
	node, err := ksq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only KothStatus ID in the query.
// Returns a *NotSingularError when more than one KothStatus ID is found.
// Returns a *NotFoundError when no entities are found.
func (ksq *KothStatusQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ksq.Limit(2).IDs(setContextOp(ctx, ksq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{kothstatus.Label}
	default:
		err = &NotSingularError{kothstatus.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ksq *KothStatusQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := ksq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of KothStatusSlice.
func (ksq *KothStatusQuery) All(ctx context.Context) ([]*KothStatus, error) {
	ctx = setContextOp(ctx, ksq.ctx, "All")
	if err := ksq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*KothStatus, *KothStatusQuery]()
	return withInterceptors[[]*KothStatus](ctx, ksq, qr, ksq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ksq *KothStatusQuery) AllX(ctx context.Context) []*KothStatus {
	nodes, err := ksq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of KothStatus IDs.
func (ksq *KothStatusQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if ksq.ctx.Unique == nil && ksq.path != nil {
		ksq.Unique(true)
	}
	ctx = setContextOp(ctx, ksq.ctx, "IDs")
	if err = ksq.Select(kothstatus.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ksq *KothStatusQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := ksq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ksq *KothStatusQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ksq.ctx, "Count")
	if err := ksq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ksq, querierCount[*KothStatusQuery](), ksq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ksq *KothStatusQuery) CountX(ctx context.Context) int {
	count, err := ksq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ksq *KothStatusQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ksq.ctx, "Exist")
	switch _, err := ksq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ksq *KothStatusQuery) ExistX(ctx context.Context) bool {
	exist, err := ksq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the KothStatusQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ksq *KothStatusQuery) Clone() *KothStatusQuery {
	if ksq == nil {
		return nil
	}
	return &KothStatusQuery{
		config:     ksq.config,
		ctx:        ksq.ctx.Clone(),
		order:      append([]kothstatus.OrderOption{}, ksq.order...),
		inters:     append([]Interceptor{}, ksq.inters...),
		predicates: append([]predicate.KothStatus{}, ksq.predicates...),
		withUser:   ksq.withUser.Clone(),
		withRound:  ksq.withRound.Clone(),
		withMinion: ksq.withMinion.Clone(),
		withCheck:  ksq.withCheck.Clone(),
		// clone intermediate query.
		sql:  ksq.sql.Clone(),
		path: ksq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (ksq *KothStatusQuery) WithUser(opts ...func(*UserQuery)) *KothStatusQuery {
	query := (&UserClient{config: ksq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ksq.withUser = query
	return ksq
}

// WithRound tells the query-builder to eager-load the nodes that are connected to
// the "round" edge. The optional arguments are used to configure the query builder of the edge.
func (ksq *KothStatusQuery) WithRound(opts ...func(*RoundQuery)) *KothStatusQuery {
	query := (&RoundClient{config: ksq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ksq.withRound = query
	return ksq
}

// WithMinion tells the query-builder to eager-load the nodes that are connected to
// the "minion" edge. The optional arguments are used to configure the query builder of the edge.
func (ksq *KothStatusQuery) WithMinion(opts ...func(*MinionQuery)) *KothStatusQuery {
	query := (&MinionClient{config: ksq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ksq.withMinion = query
	return ksq
}

// WithCheck tells the query-builder to eager-load the nodes that are connected to
// the "check" edge. The optional arguments are used to configure the query builder of the edge.
func (ksq *KothStatusQuery) WithCheck(opts ...func(*KothCheckQuery)) *KothStatusQuery {
	query := (&KothCheckClient{config: ksq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ksq.withCheck = query
	return ksq
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
//	client.KothStatus.Query().
//		GroupBy(kothstatus.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ksq *KothStatusQuery) GroupBy(field string, fields ...string) *KothStatusGroupBy {
	ksq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &KothStatusGroupBy{build: ksq}
	grbuild.flds = &ksq.ctx.Fields
	grbuild.label = kothstatus.Label
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
//	client.KothStatus.Query().
//		Select(kothstatus.FieldCreateTime).
//		Scan(ctx, &v)
func (ksq *KothStatusQuery) Select(fields ...string) *KothStatusSelect {
	ksq.ctx.Fields = append(ksq.ctx.Fields, fields...)
	sbuild := &KothStatusSelect{KothStatusQuery: ksq}
	sbuild.label = kothstatus.Label
	sbuild.flds, sbuild.scan = &ksq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a KothStatusSelect configured with the given aggregations.
func (ksq *KothStatusQuery) Aggregate(fns ...AggregateFunc) *KothStatusSelect {
	return ksq.Select().Aggregate(fns...)
}

func (ksq *KothStatusQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ksq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ksq); err != nil {
				return err
			}
		}
	}
	for _, f := range ksq.ctx.Fields {
		if !kothstatus.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ksq.path != nil {
		prev, err := ksq.path(ctx)
		if err != nil {
			return err
		}
		ksq.sql = prev
	}
	return nil
}

func (ksq *KothStatusQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*KothStatus, error) {
	var (
		nodes       = []*KothStatus{}
		_spec       = ksq.querySpec()
		loadedTypes = [4]bool{
			ksq.withUser != nil,
			ksq.withRound != nil,
			ksq.withMinion != nil,
			ksq.withCheck != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*KothStatus).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &KothStatus{config: ksq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ksq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ksq.withUser; query != nil {
		if err := ksq.loadUser(ctx, query, nodes, nil,
			func(n *KothStatus, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := ksq.withRound; query != nil {
		if err := ksq.loadRound(ctx, query, nodes, nil,
			func(n *KothStatus, e *Round) { n.Edges.Round = e }); err != nil {
			return nil, err
		}
	}
	if query := ksq.withMinion; query != nil {
		if err := ksq.loadMinion(ctx, query, nodes, nil,
			func(n *KothStatus, e *Minion) { n.Edges.Minion = e }); err != nil {
			return nil, err
		}
	}
	if query := ksq.withCheck; query != nil {
		if err := ksq.loadCheck(ctx, query, nodes, nil,
			func(n *KothStatus, e *KothCheck) { n.Edges.Check = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ksq *KothStatusQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*KothStatus, init func(*KothStatus), assign func(*KothStatus, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*KothStatus)
	for i := range nodes {
		if nodes[i].UserID == nil {
			continue
		}
		fk := *nodes[i].UserID
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
func (ksq *KothStatusQuery) loadRound(ctx context.Context, query *RoundQuery, nodes []*KothStatus, init func(*KothStatus), assign func(*KothStatus, *Round)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*KothStatus)
	for i := range nodes {
		fk := nodes[i].RoundID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(round.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "round_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (ksq *KothStatusQuery) loadMinion(ctx context.Context, query *MinionQuery, nodes []*KothStatus, init func(*KothStatus), assign func(*KothStatus, *Minion)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*KothStatus)
	for i := range nodes {
		fk := nodes[i].MinionID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(minion.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "minion_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (ksq *KothStatusQuery) loadCheck(ctx context.Context, query *KothCheckQuery, nodes []*KothStatus, init func(*KothStatus), assign func(*KothStatus, *KothCheck)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*KothStatus)
	for i := range nodes {
		fk := nodes[i].CheckID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(kothcheck.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "check_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (ksq *KothStatusQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ksq.querySpec()
	_spec.Node.Columns = ksq.ctx.Fields
	if len(ksq.ctx.Fields) > 0 {
		_spec.Unique = ksq.ctx.Unique != nil && *ksq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ksq.driver, _spec)
}

func (ksq *KothStatusQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(kothstatus.Table, kothstatus.Columns, sqlgraph.NewFieldSpec(kothstatus.FieldID, field.TypeUUID))
	_spec.From = ksq.sql
	if unique := ksq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ksq.path != nil {
		_spec.Unique = true
	}
	if fields := ksq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, kothstatus.FieldID)
		for i := range fields {
			if fields[i] != kothstatus.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if ksq.withUser != nil {
			_spec.Node.AddColumnOnce(kothstatus.FieldUserID)
		}
		if ksq.withRound != nil {
			_spec.Node.AddColumnOnce(kothstatus.FieldRoundID)
		}
		if ksq.withMinion != nil {
			_spec.Node.AddColumnOnce(kothstatus.FieldMinionID)
		}
		if ksq.withCheck != nil {
			_spec.Node.AddColumnOnce(kothstatus.FieldCheckID)
		}
	}
	if ps := ksq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ksq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ksq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ksq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ksq *KothStatusQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ksq.driver.Dialect())
	t1 := builder.Table(kothstatus.Table)
	columns := ksq.ctx.Fields
	if len(columns) == 0 {
		columns = kothstatus.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ksq.sql != nil {
		selector = ksq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ksq.ctx.Unique != nil && *ksq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ksq.predicates {
		p(selector)
	}
	for _, p := range ksq.order {
		p(selector)
	}
	if offset := ksq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ksq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// KothStatusGroupBy is the group-by builder for KothStatus entities.
type KothStatusGroupBy struct {
	selector
	build *KothStatusQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ksgb *KothStatusGroupBy) Aggregate(fns ...AggregateFunc) *KothStatusGroupBy {
	ksgb.fns = append(ksgb.fns, fns...)
	return ksgb
}

// Scan applies the selector query and scans the result into the given value.
func (ksgb *KothStatusGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ksgb.build.ctx, "GroupBy")
	if err := ksgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*KothStatusQuery, *KothStatusGroupBy](ctx, ksgb.build, ksgb, ksgb.build.inters, v)
}

func (ksgb *KothStatusGroupBy) sqlScan(ctx context.Context, root *KothStatusQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ksgb.fns))
	for _, fn := range ksgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ksgb.flds)+len(ksgb.fns))
		for _, f := range *ksgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ksgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ksgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// KothStatusSelect is the builder for selecting fields of KothStatus entities.
type KothStatusSelect struct {
	*KothStatusQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (kss *KothStatusSelect) Aggregate(fns ...AggregateFunc) *KothStatusSelect {
	kss.fns = append(kss.fns, fns...)
	return kss
}

// Scan applies the selector query and scans the result into the given value.
func (kss *KothStatusSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, kss.ctx, "Select")
	if err := kss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*KothStatusQuery, *KothStatusSelect](ctx, kss.KothStatusQuery, kss, kss.inters, v)
}

func (kss *KothStatusSelect) sqlScan(ctx context.Context, root *KothStatusQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(kss.fns))
	for _, fn := range kss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*kss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := kss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
