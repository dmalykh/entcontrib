// Copyright 2019-present Facebook
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/dmalykh/entcontrib/entgql/internal/todo/ent/billproduct"
	"github.com/dmalykh/entcontrib/entgql/internal/todo/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BillProductDelete is the builder for deleting a BillProduct entity.
type BillProductDelete struct {
	config
	hooks    []Hook
	mutation *BillProductMutation
}

// Where appends a list predicates to the BillProductDelete builder.
func (bpd *BillProductDelete) Where(ps ...predicate.BillProduct) *BillProductDelete {
	bpd.mutation.Where(ps...)
	return bpd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bpd *BillProductDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, bpd.sqlExec, bpd.mutation, bpd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (bpd *BillProductDelete) ExecX(ctx context.Context) int {
	n, err := bpd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (bpd *BillProductDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(billproduct.Table, sqlgraph.NewFieldSpec(billproduct.FieldID, field.TypeInt))
	if ps := bpd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, bpd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	bpd.mutation.done = true
	return affected, err
}

// BillProductDeleteOne is the builder for deleting a single BillProduct entity.
type BillProductDeleteOne struct {
	bpd *BillProductDelete
}

// Where appends a list predicates to the BillProductDelete builder.
func (bpdo *BillProductDeleteOne) Where(ps ...predicate.BillProduct) *BillProductDeleteOne {
	bpdo.bpd.mutation.Where(ps...)
	return bpdo
}

// Exec executes the deletion query.
func (bpdo *BillProductDeleteOne) Exec(ctx context.Context) error {
	n, err := bpdo.bpd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{billproduct.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bpdo *BillProductDeleteOne) ExecX(ctx context.Context) {
	if err := bpdo.Exec(ctx); err != nil {
		panic(err)
	}
}
