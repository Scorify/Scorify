// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/scorify/scorify/pkg/ent/kothcheck"
	"github.com/scorify/scorify/pkg/ent/kothstatus"
)

// KothCheckCreate is the builder for creating a KothCheck entity.
type KothCheckCreate struct {
	config
	mutation *KothCheckMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (kcc *KothCheckCreate) SetCreateTime(t time.Time) *KothCheckCreate {
	kcc.mutation.SetCreateTime(t)
	return kcc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (kcc *KothCheckCreate) SetNillableCreateTime(t *time.Time) *KothCheckCreate {
	if t != nil {
		kcc.SetCreateTime(*t)
	}
	return kcc
}

// SetUpdateTime sets the "update_time" field.
func (kcc *KothCheckCreate) SetUpdateTime(t time.Time) *KothCheckCreate {
	kcc.mutation.SetUpdateTime(t)
	return kcc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (kcc *KothCheckCreate) SetNillableUpdateTime(t *time.Time) *KothCheckCreate {
	if t != nil {
		kcc.SetUpdateTime(*t)
	}
	return kcc
}

// SetName sets the "name" field.
func (kcc *KothCheckCreate) SetName(s string) *KothCheckCreate {
	kcc.mutation.SetName(s)
	return kcc
}

// SetFile sets the "file" field.
func (kcc *KothCheckCreate) SetFile(s string) *KothCheckCreate {
	kcc.mutation.SetFile(s)
	return kcc
}

// SetHost sets the "host" field.
func (kcc *KothCheckCreate) SetHost(s string) *KothCheckCreate {
	kcc.mutation.SetHost(s)
	return kcc
}

// SetTopic sets the "topic" field.
func (kcc *KothCheckCreate) SetTopic(s string) *KothCheckCreate {
	kcc.mutation.SetTopic(s)
	return kcc
}

// SetWeight sets the "weight" field.
func (kcc *KothCheckCreate) SetWeight(i int) *KothCheckCreate {
	kcc.mutation.SetWeight(i)
	return kcc
}

// SetID sets the "id" field.
func (kcc *KothCheckCreate) SetID(u uuid.UUID) *KothCheckCreate {
	kcc.mutation.SetID(u)
	return kcc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (kcc *KothCheckCreate) SetNillableID(u *uuid.UUID) *KothCheckCreate {
	if u != nil {
		kcc.SetID(*u)
	}
	return kcc
}

// AddStatusIDs adds the "statuses" edge to the KothStatus entity by IDs.
func (kcc *KothCheckCreate) AddStatusIDs(ids ...uuid.UUID) *KothCheckCreate {
	kcc.mutation.AddStatusIDs(ids...)
	return kcc
}

// AddStatuses adds the "statuses" edges to the KothStatus entity.
func (kcc *KothCheckCreate) AddStatuses(k ...*KothStatus) *KothCheckCreate {
	ids := make([]uuid.UUID, len(k))
	for i := range k {
		ids[i] = k[i].ID
	}
	return kcc.AddStatusIDs(ids...)
}

// Mutation returns the KothCheckMutation object of the builder.
func (kcc *KothCheckCreate) Mutation() *KothCheckMutation {
	return kcc.mutation
}

// Save creates the KothCheck in the database.
func (kcc *KothCheckCreate) Save(ctx context.Context) (*KothCheck, error) {
	kcc.defaults()
	return withHooks(ctx, kcc.sqlSave, kcc.mutation, kcc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (kcc *KothCheckCreate) SaveX(ctx context.Context) *KothCheck {
	v, err := kcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (kcc *KothCheckCreate) Exec(ctx context.Context) error {
	_, err := kcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kcc *KothCheckCreate) ExecX(ctx context.Context) {
	if err := kcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (kcc *KothCheckCreate) defaults() {
	if _, ok := kcc.mutation.CreateTime(); !ok {
		v := kothcheck.DefaultCreateTime()
		kcc.mutation.SetCreateTime(v)
	}
	if _, ok := kcc.mutation.UpdateTime(); !ok {
		v := kothcheck.DefaultUpdateTime()
		kcc.mutation.SetUpdateTime(v)
	}
	if _, ok := kcc.mutation.ID(); !ok {
		v := kothcheck.DefaultID()
		kcc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (kcc *KothCheckCreate) check() error {
	if _, ok := kcc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "KothCheck.create_time"`)}
	}
	if _, ok := kcc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "KothCheck.update_time"`)}
	}
	if _, ok := kcc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "KothCheck.name"`)}
	}
	if v, ok := kcc.mutation.Name(); ok {
		if err := kothcheck.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "KothCheck.name": %w`, err)}
		}
	}
	if _, ok := kcc.mutation.File(); !ok {
		return &ValidationError{Name: "file", err: errors.New(`ent: missing required field "KothCheck.file"`)}
	}
	if v, ok := kcc.mutation.File(); ok {
		if err := kothcheck.FileValidator(v); err != nil {
			return &ValidationError{Name: "file", err: fmt.Errorf(`ent: validator failed for field "KothCheck.file": %w`, err)}
		}
	}
	if _, ok := kcc.mutation.Host(); !ok {
		return &ValidationError{Name: "host", err: errors.New(`ent: missing required field "KothCheck.host"`)}
	}
	if v, ok := kcc.mutation.Host(); ok {
		if err := kothcheck.HostValidator(v); err != nil {
			return &ValidationError{Name: "host", err: fmt.Errorf(`ent: validator failed for field "KothCheck.host": %w`, err)}
		}
	}
	if _, ok := kcc.mutation.Topic(); !ok {
		return &ValidationError{Name: "topic", err: errors.New(`ent: missing required field "KothCheck.topic"`)}
	}
	if v, ok := kcc.mutation.Topic(); ok {
		if err := kothcheck.TopicValidator(v); err != nil {
			return &ValidationError{Name: "topic", err: fmt.Errorf(`ent: validator failed for field "KothCheck.topic": %w`, err)}
		}
	}
	if _, ok := kcc.mutation.Weight(); !ok {
		return &ValidationError{Name: "weight", err: errors.New(`ent: missing required field "KothCheck.weight"`)}
	}
	if v, ok := kcc.mutation.Weight(); ok {
		if err := kothcheck.WeightValidator(v); err != nil {
			return &ValidationError{Name: "weight", err: fmt.Errorf(`ent: validator failed for field "KothCheck.weight": %w`, err)}
		}
	}
	return nil
}

func (kcc *KothCheckCreate) sqlSave(ctx context.Context) (*KothCheck, error) {
	if err := kcc.check(); err != nil {
		return nil, err
	}
	_node, _spec := kcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, kcc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	kcc.mutation.id = &_node.ID
	kcc.mutation.done = true
	return _node, nil
}

func (kcc *KothCheckCreate) createSpec() (*KothCheck, *sqlgraph.CreateSpec) {
	var (
		_node = &KothCheck{config: kcc.config}
		_spec = sqlgraph.NewCreateSpec(kothcheck.Table, sqlgraph.NewFieldSpec(kothcheck.FieldID, field.TypeUUID))
	)
	if id, ok := kcc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := kcc.mutation.CreateTime(); ok {
		_spec.SetField(kothcheck.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := kcc.mutation.UpdateTime(); ok {
		_spec.SetField(kothcheck.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := kcc.mutation.Name(); ok {
		_spec.SetField(kothcheck.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := kcc.mutation.File(); ok {
		_spec.SetField(kothcheck.FieldFile, field.TypeString, value)
		_node.File = value
	}
	if value, ok := kcc.mutation.Host(); ok {
		_spec.SetField(kothcheck.FieldHost, field.TypeString, value)
		_node.Host = value
	}
	if value, ok := kcc.mutation.Topic(); ok {
		_spec.SetField(kothcheck.FieldTopic, field.TypeString, value)
		_node.Topic = value
	}
	if value, ok := kcc.mutation.Weight(); ok {
		_spec.SetField(kothcheck.FieldWeight, field.TypeInt, value)
		_node.Weight = value
	}
	if nodes := kcc.mutation.StatusesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   kothcheck.StatusesTable,
			Columns: []string{kothcheck.StatusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(kothstatus.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// KothCheckCreateBulk is the builder for creating many KothCheck entities in bulk.
type KothCheckCreateBulk struct {
	config
	err      error
	builders []*KothCheckCreate
}

// Save creates the KothCheck entities in the database.
func (kccb *KothCheckCreateBulk) Save(ctx context.Context) ([]*KothCheck, error) {
	if kccb.err != nil {
		return nil, kccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(kccb.builders))
	nodes := make([]*KothCheck, len(kccb.builders))
	mutators := make([]Mutator, len(kccb.builders))
	for i := range kccb.builders {
		func(i int, root context.Context) {
			builder := kccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*KothCheckMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, kccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, kccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, kccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (kccb *KothCheckCreateBulk) SaveX(ctx context.Context) []*KothCheck {
	v, err := kccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (kccb *KothCheckCreateBulk) Exec(ctx context.Context) error {
	_, err := kccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kccb *KothCheckCreateBulk) ExecX(ctx context.Context) {
	if err := kccb.Exec(ctx); err != nil {
		panic(err)
	}
}
