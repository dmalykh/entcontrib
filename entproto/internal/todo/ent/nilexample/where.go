// Code generated by ent, DO NOT EDIT.

package nilexample

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/dmalykh/entcontrib/entproto/internal/todo/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.NilExample {
	return predicate.NilExample(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.NilExample {
	return predicate.NilExample(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.NilExample {
	return predicate.NilExample(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.NilExample {
	return predicate.NilExample(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.NilExample {
	return predicate.NilExample(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.NilExample {
	return predicate.NilExample(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.NilExample {
	return predicate.NilExample(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.NilExample {
	return predicate.NilExample(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.NilExample {
	return predicate.NilExample(sql.FieldLTE(FieldID, id))
}

// StrNil applies equality check predicate on the "str_nil" field. It's identical to StrNilEQ.
func StrNil(v string) predicate.NilExample {
	return predicate.NilExample(sql.FieldEQ(FieldStrNil, v))
}

// TimeNil applies equality check predicate on the "time_nil" field. It's identical to TimeNilEQ.
func TimeNil(v time.Time) predicate.NilExample {
	return predicate.NilExample(sql.FieldEQ(FieldTimeNil, v))
}

// StrNilEQ applies the EQ predicate on the "str_nil" field.
func StrNilEQ(v string) predicate.NilExample {
	return predicate.NilExample(sql.FieldEQ(FieldStrNil, v))
}

// StrNilNEQ applies the NEQ predicate on the "str_nil" field.
func StrNilNEQ(v string) predicate.NilExample {
	return predicate.NilExample(sql.FieldNEQ(FieldStrNil, v))
}

// StrNilIn applies the In predicate on the "str_nil" field.
func StrNilIn(vs ...string) predicate.NilExample {
	return predicate.NilExample(sql.FieldIn(FieldStrNil, vs...))
}

// StrNilNotIn applies the NotIn predicate on the "str_nil" field.
func StrNilNotIn(vs ...string) predicate.NilExample {
	return predicate.NilExample(sql.FieldNotIn(FieldStrNil, vs...))
}

// StrNilGT applies the GT predicate on the "str_nil" field.
func StrNilGT(v string) predicate.NilExample {
	return predicate.NilExample(sql.FieldGT(FieldStrNil, v))
}

// StrNilGTE applies the GTE predicate on the "str_nil" field.
func StrNilGTE(v string) predicate.NilExample {
	return predicate.NilExample(sql.FieldGTE(FieldStrNil, v))
}

// StrNilLT applies the LT predicate on the "str_nil" field.
func StrNilLT(v string) predicate.NilExample {
	return predicate.NilExample(sql.FieldLT(FieldStrNil, v))
}

// StrNilLTE applies the LTE predicate on the "str_nil" field.
func StrNilLTE(v string) predicate.NilExample {
	return predicate.NilExample(sql.FieldLTE(FieldStrNil, v))
}

// StrNilContains applies the Contains predicate on the "str_nil" field.
func StrNilContains(v string) predicate.NilExample {
	return predicate.NilExample(sql.FieldContains(FieldStrNil, v))
}

// StrNilHasPrefix applies the HasPrefix predicate on the "str_nil" field.
func StrNilHasPrefix(v string) predicate.NilExample {
	return predicate.NilExample(sql.FieldHasPrefix(FieldStrNil, v))
}

// StrNilHasSuffix applies the HasSuffix predicate on the "str_nil" field.
func StrNilHasSuffix(v string) predicate.NilExample {
	return predicate.NilExample(sql.FieldHasSuffix(FieldStrNil, v))
}

// StrNilIsNil applies the IsNil predicate on the "str_nil" field.
func StrNilIsNil() predicate.NilExample {
	return predicate.NilExample(sql.FieldIsNull(FieldStrNil))
}

// StrNilNotNil applies the NotNil predicate on the "str_nil" field.
func StrNilNotNil() predicate.NilExample {
	return predicate.NilExample(sql.FieldNotNull(FieldStrNil))
}

// StrNilEqualFold applies the EqualFold predicate on the "str_nil" field.
func StrNilEqualFold(v string) predicate.NilExample {
	return predicate.NilExample(sql.FieldEqualFold(FieldStrNil, v))
}

// StrNilContainsFold applies the ContainsFold predicate on the "str_nil" field.
func StrNilContainsFold(v string) predicate.NilExample {
	return predicate.NilExample(sql.FieldContainsFold(FieldStrNil, v))
}

// TimeNilEQ applies the EQ predicate on the "time_nil" field.
func TimeNilEQ(v time.Time) predicate.NilExample {
	return predicate.NilExample(sql.FieldEQ(FieldTimeNil, v))
}

// TimeNilNEQ applies the NEQ predicate on the "time_nil" field.
func TimeNilNEQ(v time.Time) predicate.NilExample {
	return predicate.NilExample(sql.FieldNEQ(FieldTimeNil, v))
}

// TimeNilIn applies the In predicate on the "time_nil" field.
func TimeNilIn(vs ...time.Time) predicate.NilExample {
	return predicate.NilExample(sql.FieldIn(FieldTimeNil, vs...))
}

// TimeNilNotIn applies the NotIn predicate on the "time_nil" field.
func TimeNilNotIn(vs ...time.Time) predicate.NilExample {
	return predicate.NilExample(sql.FieldNotIn(FieldTimeNil, vs...))
}

// TimeNilGT applies the GT predicate on the "time_nil" field.
func TimeNilGT(v time.Time) predicate.NilExample {
	return predicate.NilExample(sql.FieldGT(FieldTimeNil, v))
}

// TimeNilGTE applies the GTE predicate on the "time_nil" field.
func TimeNilGTE(v time.Time) predicate.NilExample {
	return predicate.NilExample(sql.FieldGTE(FieldTimeNil, v))
}

// TimeNilLT applies the LT predicate on the "time_nil" field.
func TimeNilLT(v time.Time) predicate.NilExample {
	return predicate.NilExample(sql.FieldLT(FieldTimeNil, v))
}

// TimeNilLTE applies the LTE predicate on the "time_nil" field.
func TimeNilLTE(v time.Time) predicate.NilExample {
	return predicate.NilExample(sql.FieldLTE(FieldTimeNil, v))
}

// TimeNilIsNil applies the IsNil predicate on the "time_nil" field.
func TimeNilIsNil() predicate.NilExample {
	return predicate.NilExample(sql.FieldIsNull(FieldTimeNil))
}

// TimeNilNotNil applies the NotNil predicate on the "time_nil" field.
func TimeNilNotNil() predicate.NilExample {
	return predicate.NilExample(sql.FieldNotNull(FieldTimeNil))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.NilExample) predicate.NilExample {
	return predicate.NilExample(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.NilExample) predicate.NilExample {
	return predicate.NilExample(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.NilExample) predicate.NilExample {
	return predicate.NilExample(sql.NotPredicates(p))
}
