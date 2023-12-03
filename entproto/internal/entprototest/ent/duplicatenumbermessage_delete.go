// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dmalykh/entcontrib/entproto/internal/entprototest/ent/duplicatenumbermessage"
	"github.com/dmalykh/entcontrib/entproto/internal/entprototest/ent/predicate"
)

// DuplicateNumberMessageDelete is the builder for deleting a DuplicateNumberMessage entity.
type DuplicateNumberMessageDelete struct {
	config
	hooks    []Hook
	mutation *DuplicateNumberMessageMutation
}

// Where appends a list predicates to the DuplicateNumberMessageDelete builder.
func (dnmd *DuplicateNumberMessageDelete) Where(ps ...predicate.DuplicateNumberMessage) *DuplicateNumberMessageDelete {
	dnmd.mutation.Where(ps...)
	return dnmd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (dnmd *DuplicateNumberMessageDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, dnmd.sqlExec, dnmd.mutation, dnmd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (dnmd *DuplicateNumberMessageDelete) ExecX(ctx context.Context) int {
	n, err := dnmd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (dnmd *DuplicateNumberMessageDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(duplicatenumbermessage.Table, sqlgraph.NewFieldSpec(duplicatenumbermessage.FieldID, field.TypeInt))
	if ps := dnmd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, dnmd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	dnmd.mutation.done = true
	return affected, err
}

// DuplicateNumberMessageDeleteOne is the builder for deleting a single DuplicateNumberMessage entity.
type DuplicateNumberMessageDeleteOne struct {
	dnmd *DuplicateNumberMessageDelete
}

// Where appends a list predicates to the DuplicateNumberMessageDelete builder.
func (dnmdo *DuplicateNumberMessageDeleteOne) Where(ps ...predicate.DuplicateNumberMessage) *DuplicateNumberMessageDeleteOne {
	dnmdo.dnmd.mutation.Where(ps...)
	return dnmdo
}

// Exec executes the deletion query.
func (dnmdo *DuplicateNumberMessageDeleteOne) Exec(ctx context.Context) error {
	n, err := dnmdo.dnmd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{duplicatenumbermessage.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (dnmdo *DuplicateNumberMessageDeleteOne) ExecX(ctx context.Context) {
	if err := dnmdo.Exec(ctx); err != nil {
		panic(err)
	}
}
