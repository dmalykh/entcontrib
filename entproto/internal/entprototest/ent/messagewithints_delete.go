// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dmalykh/entcontrib/entproto/internal/entprototest/ent/messagewithints"
	"github.com/dmalykh/entcontrib/entproto/internal/entprototest/ent/predicate"
)

// MessageWithIntsDelete is the builder for deleting a MessageWithInts entity.
type MessageWithIntsDelete struct {
	config
	hooks    []Hook
	mutation *MessageWithIntsMutation
}

// Where appends a list predicates to the MessageWithIntsDelete builder.
func (mwid *MessageWithIntsDelete) Where(ps ...predicate.MessageWithInts) *MessageWithIntsDelete {
	mwid.mutation.Where(ps...)
	return mwid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (mwid *MessageWithIntsDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, mwid.sqlExec, mwid.mutation, mwid.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (mwid *MessageWithIntsDelete) ExecX(ctx context.Context) int {
	n, err := mwid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (mwid *MessageWithIntsDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(messagewithints.Table, sqlgraph.NewFieldSpec(messagewithints.FieldID, field.TypeInt))
	if ps := mwid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, mwid.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	mwid.mutation.done = true
	return affected, err
}

// MessageWithIntsDeleteOne is the builder for deleting a single MessageWithInts entity.
type MessageWithIntsDeleteOne struct {
	mwid *MessageWithIntsDelete
}

// Where appends a list predicates to the MessageWithIntsDelete builder.
func (mwido *MessageWithIntsDeleteOne) Where(ps ...predicate.MessageWithInts) *MessageWithIntsDeleteOne {
	mwido.mwid.mutation.Where(ps...)
	return mwido
}

// Exec executes the deletion query.
func (mwido *MessageWithIntsDeleteOne) Exec(ctx context.Context) error {
	n, err := mwido.mwid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{messagewithints.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (mwido *MessageWithIntsDeleteOne) ExecX(ctx context.Context) {
	if err := mwido.Exec(ctx); err != nil {
		panic(err)
	}
}
