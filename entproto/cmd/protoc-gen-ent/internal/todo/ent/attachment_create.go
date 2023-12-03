// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dmalykh/entcontrib/entproto/cmd/protoc-gen-ent/internal/todo/ent/attachment"
)

// AttachmentCreate is the builder for creating a Attachment entity.
type AttachmentCreate struct {
	config
	mutation *AttachmentMutation
	hooks    []Hook
}

// SetContents sets the "contents" field.
func (ac *AttachmentCreate) SetContents(s string) *AttachmentCreate {
	ac.mutation.SetContents(s)
	return ac
}

// SetID sets the "id" field.
func (ac *AttachmentCreate) SetID(s string) *AttachmentCreate {
	ac.mutation.SetID(s)
	return ac
}

// Mutation returns the AttachmentMutation object of the builder.
func (ac *AttachmentCreate) Mutation() *AttachmentMutation {
	return ac.mutation
}

// Save creates the Attachment in the database.
func (ac *AttachmentCreate) Save(ctx context.Context) (*Attachment, error) {
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AttachmentCreate) SaveX(ctx context.Context) *Attachment {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AttachmentCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AttachmentCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AttachmentCreate) check() error {
	if _, ok := ac.mutation.Contents(); !ok {
		return &ValidationError{Name: "contents", err: errors.New(`ent: missing required field "Attachment.contents"`)}
	}
	return nil
}

func (ac *AttachmentCreate) sqlSave(ctx context.Context) (*Attachment, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Attachment.ID type: %T", _spec.ID.Value)
		}
	}
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *AttachmentCreate) createSpec() (*Attachment, *sqlgraph.CreateSpec) {
	var (
		_node = &Attachment{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(attachment.Table, sqlgraph.NewFieldSpec(attachment.FieldID, field.TypeString))
	)
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ac.mutation.Contents(); ok {
		_spec.SetField(attachment.FieldContents, field.TypeString, value)
		_node.Contents = value
	}
	return _node, _spec
}

// AttachmentCreateBulk is the builder for creating many Attachment entities in bulk.
type AttachmentCreateBulk struct {
	config
	err      error
	builders []*AttachmentCreate
}

// Save creates the Attachment entities in the database.
func (acb *AttachmentCreateBulk) Save(ctx context.Context) ([]*Attachment, error) {
	if acb.err != nil {
		return nil, acb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Attachment, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AttachmentMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AttachmentCreateBulk) SaveX(ctx context.Context) []*Attachment {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AttachmentCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AttachmentCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
