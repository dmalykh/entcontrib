// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dmalykh/entcontrib/entproto/internal/entprototest/ent/enumwithconflictingvalue"
	"github.com/dmalykh/entcontrib/entproto/internal/entprototest/ent/predicate"
)

// EnumWithConflictingValueDelete is the builder for deleting a EnumWithConflictingValue entity.
type EnumWithConflictingValueDelete struct {
	config
	hooks    []Hook
	mutation *EnumWithConflictingValueMutation
}

// Where appends a list predicates to the EnumWithConflictingValueDelete builder.
func (ewcvd *EnumWithConflictingValueDelete) Where(ps ...predicate.EnumWithConflictingValue) *EnumWithConflictingValueDelete {
	ewcvd.mutation.Where(ps...)
	return ewcvd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ewcvd *EnumWithConflictingValueDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ewcvd.sqlExec, ewcvd.mutation, ewcvd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ewcvd *EnumWithConflictingValueDelete) ExecX(ctx context.Context) int {
	n, err := ewcvd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ewcvd *EnumWithConflictingValueDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(enumwithconflictingvalue.Table, sqlgraph.NewFieldSpec(enumwithconflictingvalue.FieldID, field.TypeInt))
	if ps := ewcvd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ewcvd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ewcvd.mutation.done = true
	return affected, err
}

// EnumWithConflictingValueDeleteOne is the builder for deleting a single EnumWithConflictingValue entity.
type EnumWithConflictingValueDeleteOne struct {
	ewcvd *EnumWithConflictingValueDelete
}

// Where appends a list predicates to the EnumWithConflictingValueDelete builder.
func (ewcvdo *EnumWithConflictingValueDeleteOne) Where(ps ...predicate.EnumWithConflictingValue) *EnumWithConflictingValueDeleteOne {
	ewcvdo.ewcvd.mutation.Where(ps...)
	return ewcvdo
}

// Exec executes the deletion query.
func (ewcvdo *EnumWithConflictingValueDeleteOne) Exec(ctx context.Context) error {
	n, err := ewcvdo.ewcvd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{enumwithconflictingvalue.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ewcvdo *EnumWithConflictingValueDeleteOne) ExecX(ctx context.Context) {
	if err := ewcvdo.Exec(ctx); err != nil {
		panic(err)
	}
}
