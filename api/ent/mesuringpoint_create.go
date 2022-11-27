// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/hmdyt/satori_codingtest-2/ent/mesuringpoint"
)

// MesuringPointCreate is the builder for creating a MesuringPoint entity.
type MesuringPointCreate struct {
	config
	mutation *MesuringPointMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (mpc *MesuringPointCreate) SetUserID(i int) *MesuringPointCreate {
	mpc.mutation.SetUserID(i)
	return mpc
}

// SetBodyMass sets the "body_mass" field.
func (mpc *MesuringPointCreate) SetBodyMass(f float64) *MesuringPointCreate {
	mpc.mutation.SetBodyMass(f)
	return mpc
}

// SetCreatedAt sets the "created_at" field.
func (mpc *MesuringPointCreate) SetCreatedAt(t time.Time) *MesuringPointCreate {
	mpc.mutation.SetCreatedAt(t)
	return mpc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mpc *MesuringPointCreate) SetNillableCreatedAt(t *time.Time) *MesuringPointCreate {
	if t != nil {
		mpc.SetCreatedAt(*t)
	}
	return mpc
}

// Mutation returns the MesuringPointMutation object of the builder.
func (mpc *MesuringPointCreate) Mutation() *MesuringPointMutation {
	return mpc.mutation
}

// Save creates the MesuringPoint in the database.
func (mpc *MesuringPointCreate) Save(ctx context.Context) (*MesuringPoint, error) {
	var (
		err  error
		node *MesuringPoint
	)
	mpc.defaults()
	if len(mpc.hooks) == 0 {
		if err = mpc.check(); err != nil {
			return nil, err
		}
		node, err = mpc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MesuringPointMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mpc.check(); err != nil {
				return nil, err
			}
			mpc.mutation = mutation
			if node, err = mpc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(mpc.hooks) - 1; i >= 0; i-- {
			if mpc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mpc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, mpc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*MesuringPoint)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from MesuringPointMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mpc *MesuringPointCreate) SaveX(ctx context.Context) *MesuringPoint {
	v, err := mpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mpc *MesuringPointCreate) Exec(ctx context.Context) error {
	_, err := mpc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mpc *MesuringPointCreate) ExecX(ctx context.Context) {
	if err := mpc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mpc *MesuringPointCreate) defaults() {
	if _, ok := mpc.mutation.CreatedAt(); !ok {
		v := mesuringpoint.DefaultCreatedAt()
		mpc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mpc *MesuringPointCreate) check() error {
	if _, ok := mpc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "MesuringPoint.user_id"`)}
	}
	if v, ok := mpc.mutation.UserID(); ok {
		if err := mesuringpoint.UserIDValidator(v); err != nil {
			return &ValidationError{Name: "user_id", err: fmt.Errorf(`ent: validator failed for field "MesuringPoint.user_id": %w`, err)}
		}
	}
	if _, ok := mpc.mutation.BodyMass(); !ok {
		return &ValidationError{Name: "body_mass", err: errors.New(`ent: missing required field "MesuringPoint.body_mass"`)}
	}
	if v, ok := mpc.mutation.BodyMass(); ok {
		if err := mesuringpoint.BodyMassValidator(v); err != nil {
			return &ValidationError{Name: "body_mass", err: fmt.Errorf(`ent: validator failed for field "MesuringPoint.body_mass": %w`, err)}
		}
	}
	if _, ok := mpc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "MesuringPoint.created_at"`)}
	}
	return nil
}

func (mpc *MesuringPointCreate) sqlSave(ctx context.Context) (*MesuringPoint, error) {
	_node, _spec := mpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mpc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (mpc *MesuringPointCreate) createSpec() (*MesuringPoint, *sqlgraph.CreateSpec) {
	var (
		_node = &MesuringPoint{config: mpc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: mesuringpoint.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: mesuringpoint.FieldID,
			},
		}
	)
	if value, ok := mpc.mutation.UserID(); ok {
		_spec.SetField(mesuringpoint.FieldUserID, field.TypeInt, value)
		_node.UserID = value
	}
	if value, ok := mpc.mutation.BodyMass(); ok {
		_spec.SetField(mesuringpoint.FieldBodyMass, field.TypeFloat64, value)
		_node.BodyMass = value
	}
	if value, ok := mpc.mutation.CreatedAt(); ok {
		_spec.SetField(mesuringpoint.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	return _node, _spec
}

// MesuringPointCreateBulk is the builder for creating many MesuringPoint entities in bulk.
type MesuringPointCreateBulk struct {
	config
	builders []*MesuringPointCreate
}

// Save creates the MesuringPoint entities in the database.
func (mpcb *MesuringPointCreateBulk) Save(ctx context.Context) ([]*MesuringPoint, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mpcb.builders))
	nodes := make([]*MesuringPoint, len(mpcb.builders))
	mutators := make([]Mutator, len(mpcb.builders))
	for i := range mpcb.builders {
		func(i int, root context.Context) {
			builder := mpcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MesuringPointMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, mpcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mpcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, mpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mpcb *MesuringPointCreateBulk) SaveX(ctx context.Context) []*MesuringPoint {
	v, err := mpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mpcb *MesuringPointCreateBulk) Exec(ctx context.Context) error {
	_, err := mpcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mpcb *MesuringPointCreateBulk) ExecX(ctx context.Context) {
	if err := mpcb.Exec(ctx); err != nil {
		panic(err)
	}
}
