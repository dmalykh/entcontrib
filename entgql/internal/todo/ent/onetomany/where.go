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

package onetomany

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/dmalykh/entcontrib/entgql/internal/todo/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.OneToMany {
	return predicate.OneToMany(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.OneToMany {
	return predicate.OneToMany(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.OneToMany {
	return predicate.OneToMany(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.OneToMany {
	return predicate.OneToMany(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.OneToMany {
	return predicate.OneToMany(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.OneToMany {
	return predicate.OneToMany(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.OneToMany {
	return predicate.OneToMany(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.OneToMany {
	return predicate.OneToMany(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.OneToMany {
	return predicate.OneToMany(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.OneToMany {
	return predicate.OneToMany(sql.FieldEQ(FieldName, v))
}

// Field2 applies equality check predicate on the "field2" field. It's identical to Field2EQ.
func Field2(v string) predicate.OneToMany {
	return predicate.OneToMany(sql.FieldEQ(FieldField2, v))
}

// ParentID applies equality check predicate on the "parent_id" field. It's identical to ParentIDEQ.
func ParentID(v int) predicate.OneToMany {
	return predicate.OneToMany(sql.FieldEQ(FieldParentID, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.OneToMany {
	return predicate.OneToMany(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.OneToMany {
	return predicate.OneToMany(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.OneToMany {
	return predicate.OneToMany(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.OneToMany {
	return predicate.OneToMany(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.OneToMany {
	return predicate.OneToMany(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.OneToMany {
	return predicate.OneToMany(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.OneToMany {
	return predicate.OneToMany(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.OneToMany {
	return predicate.OneToMany(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.OneToMany {
	return predicate.OneToMany(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.OneToMany {
	return predicate.OneToMany(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.OneToMany {
	return predicate.OneToMany(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.OneToMany {
	return predicate.OneToMany(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.OneToMany {
	return predicate.OneToMany(sql.FieldContainsFold(FieldName, v))
}

// Field2EQ applies the EQ predicate on the "field2" field.
func Field2EQ(v string) predicate.OneToMany {
	return predicate.OneToMany(sql.FieldEQ(FieldField2, v))
}

// Field2NEQ applies the NEQ predicate on the "field2" field.
func Field2NEQ(v string) predicate.OneToMany {
	return predicate.OneToMany(sql.FieldNEQ(FieldField2, v))
}

// Field2In applies the In predicate on the "field2" field.
func Field2In(vs ...string) predicate.OneToMany {
	return predicate.OneToMany(sql.FieldIn(FieldField2, vs...))
}

// Field2NotIn applies the NotIn predicate on the "field2" field.
func Field2NotIn(vs ...string) predicate.OneToMany {
	return predicate.OneToMany(sql.FieldNotIn(FieldField2, vs...))
}

// Field2GT applies the GT predicate on the "field2" field.
func Field2GT(v string) predicate.OneToMany {
	return predicate.OneToMany(sql.FieldGT(FieldField2, v))
}

// Field2GTE applies the GTE predicate on the "field2" field.
func Field2GTE(v string) predicate.OneToMany {
	return predicate.OneToMany(sql.FieldGTE(FieldField2, v))
}

// Field2LT applies the LT predicate on the "field2" field.
func Field2LT(v string) predicate.OneToMany {
	return predicate.OneToMany(sql.FieldLT(FieldField2, v))
}

// Field2LTE applies the LTE predicate on the "field2" field.
func Field2LTE(v string) predicate.OneToMany {
	return predicate.OneToMany(sql.FieldLTE(FieldField2, v))
}

// Field2Contains applies the Contains predicate on the "field2" field.
func Field2Contains(v string) predicate.OneToMany {
	return predicate.OneToMany(sql.FieldContains(FieldField2, v))
}

// Field2HasPrefix applies the HasPrefix predicate on the "field2" field.
func Field2HasPrefix(v string) predicate.OneToMany {
	return predicate.OneToMany(sql.FieldHasPrefix(FieldField2, v))
}

// Field2HasSuffix applies the HasSuffix predicate on the "field2" field.
func Field2HasSuffix(v string) predicate.OneToMany {
	return predicate.OneToMany(sql.FieldHasSuffix(FieldField2, v))
}

// Field2IsNil applies the IsNil predicate on the "field2" field.
func Field2IsNil() predicate.OneToMany {
	return predicate.OneToMany(sql.FieldIsNull(FieldField2))
}

// Field2NotNil applies the NotNil predicate on the "field2" field.
func Field2NotNil() predicate.OneToMany {
	return predicate.OneToMany(sql.FieldNotNull(FieldField2))
}

// Field2EqualFold applies the EqualFold predicate on the "field2" field.
func Field2EqualFold(v string) predicate.OneToMany {
	return predicate.OneToMany(sql.FieldEqualFold(FieldField2, v))
}

// Field2ContainsFold applies the ContainsFold predicate on the "field2" field.
func Field2ContainsFold(v string) predicate.OneToMany {
	return predicate.OneToMany(sql.FieldContainsFold(FieldField2, v))
}

// ParentIDEQ applies the EQ predicate on the "parent_id" field.
func ParentIDEQ(v int) predicate.OneToMany {
	return predicate.OneToMany(sql.FieldEQ(FieldParentID, v))
}

// ParentIDNEQ applies the NEQ predicate on the "parent_id" field.
func ParentIDNEQ(v int) predicate.OneToMany {
	return predicate.OneToMany(sql.FieldNEQ(FieldParentID, v))
}

// ParentIDIn applies the In predicate on the "parent_id" field.
func ParentIDIn(vs ...int) predicate.OneToMany {
	return predicate.OneToMany(sql.FieldIn(FieldParentID, vs...))
}

// ParentIDNotIn applies the NotIn predicate on the "parent_id" field.
func ParentIDNotIn(vs ...int) predicate.OneToMany {
	return predicate.OneToMany(sql.FieldNotIn(FieldParentID, vs...))
}

// ParentIDIsNil applies the IsNil predicate on the "parent_id" field.
func ParentIDIsNil() predicate.OneToMany {
	return predicate.OneToMany(sql.FieldIsNull(FieldParentID))
}

// ParentIDNotNil applies the NotNil predicate on the "parent_id" field.
func ParentIDNotNil() predicate.OneToMany {
	return predicate.OneToMany(sql.FieldNotNull(FieldParentID))
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.OneToMany {
	return predicate.OneToMany(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.OneToMany) predicate.OneToMany {
	return predicate.OneToMany(func(s *sql.Selector) {
		step := newParentStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasChildren applies the HasEdge predicate on the "children" edge.
func HasChildren() predicate.OneToMany {
	return predicate.OneToMany(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ChildrenTable, ChildrenColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChildrenWith applies the HasEdge predicate on the "children" edge with a given conditions (other predicates).
func HasChildrenWith(preds ...predicate.OneToMany) predicate.OneToMany {
	return predicate.OneToMany(func(s *sql.Selector) {
		step := newChildrenStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.OneToMany) predicate.OneToMany {
	return predicate.OneToMany(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.OneToMany) predicate.OneToMany {
	return predicate.OneToMany(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.OneToMany) predicate.OneToMany {
	return predicate.OneToMany(sql.NotPredicates(p))
}
