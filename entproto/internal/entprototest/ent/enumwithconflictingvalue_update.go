// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dmalykh/entcontrib/entproto/internal/entprototest/ent/enumwithconflictingvalue"
	"github.com/dmalykh/entcontrib/entproto/internal/entprototest/ent/predicate"
)

// EnumWithConflictingValueUpdate is the builder for updating EnumWithConflictingValue entities.
type EnumWithConflictingValueUpdate struct {
	config
	hooks    []Hook
	mutation *EnumWithConflictingValueMutation
}

// Where appends a list predicates to the EnumWithConflictingValueUpdate builder.
func (ewcvu *EnumWithConflictingValueUpdate) Where(ps ...predicate.EnumWithConflictingValue) *EnumWithConflictingValueUpdate {
	ewcvu.mutation.Where(ps...)
	return ewcvu
}

// SetEnum sets the "enum" field.
func (ewcvu *EnumWithConflictingValueUpdate) SetEnum(e enumwithconflictingvalue.Enum) *EnumWithConflictingValueUpdate {
	ewcvu.mutation.SetEnum(e)
	return ewcvu
}

// SetNillableEnum sets the "enum" field if the given value is not nil.
func (ewcvu *EnumWithConflictingValueUpdate) SetNillableEnum(e *enumwithconflictingvalue.Enum) *EnumWithConflictingValueUpdate {
	if e != nil {
		ewcvu.SetEnum(*e)
	}
	return ewcvu
}

// Mutation returns the EnumWithConflictingValueMutation object of the builder.
func (ewcvu *EnumWithConflictingValueUpdate) Mutation() *EnumWithConflictingValueMutation {
	return ewcvu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ewcvu *EnumWithConflictingValueUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ewcvu.sqlSave, ewcvu.mutation, ewcvu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ewcvu *EnumWithConflictingValueUpdate) SaveX(ctx context.Context) int {
	affected, err := ewcvu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ewcvu *EnumWithConflictingValueUpdate) Exec(ctx context.Context) error {
	_, err := ewcvu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ewcvu *EnumWithConflictingValueUpdate) ExecX(ctx context.Context) {
	if err := ewcvu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ewcvu *EnumWithConflictingValueUpdate) check() error {
	if v, ok := ewcvu.mutation.Enum(); ok {
		if err := enumwithconflictingvalue.EnumValidator(v); err != nil {
			return &ValidationError{Name: "enum", err: fmt.Errorf(`ent: validator failed for field "EnumWithConflictingValue.enum": %w`, err)}
		}
	}
	return nil
}

func (ewcvu *EnumWithConflictingValueUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ewcvu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(enumwithconflictingvalue.Table, enumwithconflictingvalue.Columns, sqlgraph.NewFieldSpec(enumwithconflictingvalue.FieldID, field.TypeInt))
	if ps := ewcvu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ewcvu.mutation.Enum(); ok {
		_spec.SetField(enumwithconflictingvalue.FieldEnum, field.TypeEnum, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ewcvu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{enumwithconflictingvalue.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ewcvu.mutation.done = true
	return n, nil
}

// EnumWithConflictingValueUpdateOne is the builder for updating a single EnumWithConflictingValue entity.
type EnumWithConflictingValueUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EnumWithConflictingValueMutation
}

// SetEnum sets the "enum" field.
func (ewcvuo *EnumWithConflictingValueUpdateOne) SetEnum(e enumwithconflictingvalue.Enum) *EnumWithConflictingValueUpdateOne {
	ewcvuo.mutation.SetEnum(e)
	return ewcvuo
}

// SetNillableEnum sets the "enum" field if the given value is not nil.
func (ewcvuo *EnumWithConflictingValueUpdateOne) SetNillableEnum(e *enumwithconflictingvalue.Enum) *EnumWithConflictingValueUpdateOne {
	if e != nil {
		ewcvuo.SetEnum(*e)
	}
	return ewcvuo
}

// Mutation returns the EnumWithConflictingValueMutation object of the builder.
func (ewcvuo *EnumWithConflictingValueUpdateOne) Mutation() *EnumWithConflictingValueMutation {
	return ewcvuo.mutation
}

// Where appends a list predicates to the EnumWithConflictingValueUpdate builder.
func (ewcvuo *EnumWithConflictingValueUpdateOne) Where(ps ...predicate.EnumWithConflictingValue) *EnumWithConflictingValueUpdateOne {
	ewcvuo.mutation.Where(ps...)
	return ewcvuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ewcvuo *EnumWithConflictingValueUpdateOne) Select(field string, fields ...string) *EnumWithConflictingValueUpdateOne {
	ewcvuo.fields = append([]string{field}, fields...)
	return ewcvuo
}

// Save executes the query and returns the updated EnumWithConflictingValue entity.
func (ewcvuo *EnumWithConflictingValueUpdateOne) Save(ctx context.Context) (*EnumWithConflictingValue, error) {
	return withHooks(ctx, ewcvuo.sqlSave, ewcvuo.mutation, ewcvuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ewcvuo *EnumWithConflictingValueUpdateOne) SaveX(ctx context.Context) *EnumWithConflictingValue {
	node, err := ewcvuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ewcvuo *EnumWithConflictingValueUpdateOne) Exec(ctx context.Context) error {
	_, err := ewcvuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ewcvuo *EnumWithConflictingValueUpdateOne) ExecX(ctx context.Context) {
	if err := ewcvuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ewcvuo *EnumWithConflictingValueUpdateOne) check() error {
	if v, ok := ewcvuo.mutation.Enum(); ok {
		if err := enumwithconflictingvalue.EnumValidator(v); err != nil {
			return &ValidationError{Name: "enum", err: fmt.Errorf(`ent: validator failed for field "EnumWithConflictingValue.enum": %w`, err)}
		}
	}
	return nil
}

func (ewcvuo *EnumWithConflictingValueUpdateOne) sqlSave(ctx context.Context) (_node *EnumWithConflictingValue, err error) {
	if err := ewcvuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(enumwithconflictingvalue.Table, enumwithconflictingvalue.Columns, sqlgraph.NewFieldSpec(enumwithconflictingvalue.FieldID, field.TypeInt))
	id, ok := ewcvuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "EnumWithConflictingValue.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ewcvuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, enumwithconflictingvalue.FieldID)
		for _, f := range fields {
			if !enumwithconflictingvalue.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != enumwithconflictingvalue.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ewcvuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ewcvuo.mutation.Enum(); ok {
		_spec.SetField(enumwithconflictingvalue.FieldEnum, field.TypeEnum, value)
	}
	_node = &EnumWithConflictingValue{config: ewcvuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ewcvuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{enumwithconflictingvalue.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ewcvuo.mutation.done = true
	return _node, nil
}
