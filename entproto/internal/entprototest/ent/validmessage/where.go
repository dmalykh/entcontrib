// Code generated by ent, DO NOT EDIT.

package validmessage

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/dmalykh/entcontrib/entproto/internal/entprototest/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ValidMessage {
	return predicate.ValidMessage(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ValidMessage {
	return predicate.ValidMessage(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ValidMessage {
	return predicate.ValidMessage(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ValidMessage {
	return predicate.ValidMessage(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.ValidMessage {
	return predicate.ValidMessage(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.ValidMessage {
	return predicate.ValidMessage(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ValidMessage {
	return predicate.ValidMessage(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ValidMessage {
	return predicate.ValidMessage(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ValidMessage {
	return predicate.ValidMessage(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.ValidMessage {
	return predicate.ValidMessage(sql.FieldEQ(FieldName, v))
}

// Ts applies equality check predicate on the "ts" field. It's identical to TsEQ.
func Ts(v time.Time) predicate.ValidMessage {
	return predicate.ValidMessage(sql.FieldEQ(FieldTs, v))
}

// UUID applies equality check predicate on the "uuid" field. It's identical to UUIDEQ.
func UUID(v uuid.UUID) predicate.ValidMessage {
	return predicate.ValidMessage(sql.FieldEQ(FieldUUID, v))
}

// U8 applies equality check predicate on the "u8" field. It's identical to U8EQ.
func U8(v uint8) predicate.ValidMessage {
	return predicate.ValidMessage(sql.FieldEQ(FieldU8, v))
}

// Opti8 applies equality check predicate on the "opti8" field. It's identical to Opti8EQ.
func Opti8(v int8) predicate.ValidMessage {
	return predicate.ValidMessage(sql.FieldEQ(FieldOpti8, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.ValidMessage {
	return predicate.ValidMessage(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.ValidMessage {
	return predicate.ValidMessage(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.ValidMessage {
	return predicate.ValidMessage(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.ValidMessage {
	return predicate.ValidMessage(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.ValidMessage {
	return predicate.ValidMessage(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.ValidMessage {
	return predicate.ValidMessage(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.ValidMessage {
	return predicate.ValidMessage(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.ValidMessage {
	return predicate.ValidMessage(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.ValidMessage {
	return predicate.ValidMessage(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.ValidMessage {
	return predicate.ValidMessage(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.ValidMessage {
	return predicate.ValidMessage(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.ValidMessage {
	return predicate.ValidMessage(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.ValidMessage {
	return predicate.ValidMessage(sql.FieldContainsFold(FieldName, v))
}

// TsEQ applies the EQ predicate on the "ts" field.
func TsEQ(v time.Time) predicate.ValidMessage {
	return predicate.ValidMessage(sql.FieldEQ(FieldTs, v))
}

// TsNEQ applies the NEQ predicate on the "ts" field.
func TsNEQ(v time.Time) predicate.ValidMessage {
	return predicate.ValidMessage(sql.FieldNEQ(FieldTs, v))
}

// TsIn applies the In predicate on the "ts" field.
func TsIn(vs ...time.Time) predicate.ValidMessage {
	return predicate.ValidMessage(sql.FieldIn(FieldTs, vs...))
}

// TsNotIn applies the NotIn predicate on the "ts" field.
func TsNotIn(vs ...time.Time) predicate.ValidMessage {
	return predicate.ValidMessage(sql.FieldNotIn(FieldTs, vs...))
}

// TsGT applies the GT predicate on the "ts" field.
func TsGT(v time.Time) predicate.ValidMessage {
	return predicate.ValidMessage(sql.FieldGT(FieldTs, v))
}

// TsGTE applies the GTE predicate on the "ts" field.
func TsGTE(v time.Time) predicate.ValidMessage {
	return predicate.ValidMessage(sql.FieldGTE(FieldTs, v))
}

// TsLT applies the LT predicate on the "ts" field.
func TsLT(v time.Time) predicate.ValidMessage {
	return predicate.ValidMessage(sql.FieldLT(FieldTs, v))
}

// TsLTE applies the LTE predicate on the "ts" field.
func TsLTE(v time.Time) predicate.ValidMessage {
	return predicate.ValidMessage(sql.FieldLTE(FieldTs, v))
}

// UUIDEQ applies the EQ predicate on the "uuid" field.
func UUIDEQ(v uuid.UUID) predicate.ValidMessage {
	return predicate.ValidMessage(sql.FieldEQ(FieldUUID, v))
}

// UUIDNEQ applies the NEQ predicate on the "uuid" field.
func UUIDNEQ(v uuid.UUID) predicate.ValidMessage {
	return predicate.ValidMessage(sql.FieldNEQ(FieldUUID, v))
}

// UUIDIn applies the In predicate on the "uuid" field.
func UUIDIn(vs ...uuid.UUID) predicate.ValidMessage {
	return predicate.ValidMessage(sql.FieldIn(FieldUUID, vs...))
}

// UUIDNotIn applies the NotIn predicate on the "uuid" field.
func UUIDNotIn(vs ...uuid.UUID) predicate.ValidMessage {
	return predicate.ValidMessage(sql.FieldNotIn(FieldUUID, vs...))
}

// UUIDGT applies the GT predicate on the "uuid" field.
func UUIDGT(v uuid.UUID) predicate.ValidMessage {
	return predicate.ValidMessage(sql.FieldGT(FieldUUID, v))
}

// UUIDGTE applies the GTE predicate on the "uuid" field.
func UUIDGTE(v uuid.UUID) predicate.ValidMessage {
	return predicate.ValidMessage(sql.FieldGTE(FieldUUID, v))
}

// UUIDLT applies the LT predicate on the "uuid" field.
func UUIDLT(v uuid.UUID) predicate.ValidMessage {
	return predicate.ValidMessage(sql.FieldLT(FieldUUID, v))
}

// UUIDLTE applies the LTE predicate on the "uuid" field.
func UUIDLTE(v uuid.UUID) predicate.ValidMessage {
	return predicate.ValidMessage(sql.FieldLTE(FieldUUID, v))
}

// U8EQ applies the EQ predicate on the "u8" field.
func U8EQ(v uint8) predicate.ValidMessage {
	return predicate.ValidMessage(sql.FieldEQ(FieldU8, v))
}

// U8NEQ applies the NEQ predicate on the "u8" field.
func U8NEQ(v uint8) predicate.ValidMessage {
	return predicate.ValidMessage(sql.FieldNEQ(FieldU8, v))
}

// U8In applies the In predicate on the "u8" field.
func U8In(vs ...uint8) predicate.ValidMessage {
	return predicate.ValidMessage(sql.FieldIn(FieldU8, vs...))
}

// U8NotIn applies the NotIn predicate on the "u8" field.
func U8NotIn(vs ...uint8) predicate.ValidMessage {
	return predicate.ValidMessage(sql.FieldNotIn(FieldU8, vs...))
}

// U8GT applies the GT predicate on the "u8" field.
func U8GT(v uint8) predicate.ValidMessage {
	return predicate.ValidMessage(sql.FieldGT(FieldU8, v))
}

// U8GTE applies the GTE predicate on the "u8" field.
func U8GTE(v uint8) predicate.ValidMessage {
	return predicate.ValidMessage(sql.FieldGTE(FieldU8, v))
}

// U8LT applies the LT predicate on the "u8" field.
func U8LT(v uint8) predicate.ValidMessage {
	return predicate.ValidMessage(sql.FieldLT(FieldU8, v))
}

// U8LTE applies the LTE predicate on the "u8" field.
func U8LTE(v uint8) predicate.ValidMessage {
	return predicate.ValidMessage(sql.FieldLTE(FieldU8, v))
}

// Opti8EQ applies the EQ predicate on the "opti8" field.
func Opti8EQ(v int8) predicate.ValidMessage {
	return predicate.ValidMessage(sql.FieldEQ(FieldOpti8, v))
}

// Opti8NEQ applies the NEQ predicate on the "opti8" field.
func Opti8NEQ(v int8) predicate.ValidMessage {
	return predicate.ValidMessage(sql.FieldNEQ(FieldOpti8, v))
}

// Opti8In applies the In predicate on the "opti8" field.
func Opti8In(vs ...int8) predicate.ValidMessage {
	return predicate.ValidMessage(sql.FieldIn(FieldOpti8, vs...))
}

// Opti8NotIn applies the NotIn predicate on the "opti8" field.
func Opti8NotIn(vs ...int8) predicate.ValidMessage {
	return predicate.ValidMessage(sql.FieldNotIn(FieldOpti8, vs...))
}

// Opti8GT applies the GT predicate on the "opti8" field.
func Opti8GT(v int8) predicate.ValidMessage {
	return predicate.ValidMessage(sql.FieldGT(FieldOpti8, v))
}

// Opti8GTE applies the GTE predicate on the "opti8" field.
func Opti8GTE(v int8) predicate.ValidMessage {
	return predicate.ValidMessage(sql.FieldGTE(FieldOpti8, v))
}

// Opti8LT applies the LT predicate on the "opti8" field.
func Opti8LT(v int8) predicate.ValidMessage {
	return predicate.ValidMessage(sql.FieldLT(FieldOpti8, v))
}

// Opti8LTE applies the LTE predicate on the "opti8" field.
func Opti8LTE(v int8) predicate.ValidMessage {
	return predicate.ValidMessage(sql.FieldLTE(FieldOpti8, v))
}

// Opti8IsNil applies the IsNil predicate on the "opti8" field.
func Opti8IsNil() predicate.ValidMessage {
	return predicate.ValidMessage(sql.FieldIsNull(FieldOpti8))
}

// Opti8NotNil applies the NotNil predicate on the "opti8" field.
func Opti8NotNil() predicate.ValidMessage {
	return predicate.ValidMessage(sql.FieldNotNull(FieldOpti8))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ValidMessage) predicate.ValidMessage {
	return predicate.ValidMessage(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ValidMessage) predicate.ValidMessage {
	return predicate.ValidMessage(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ValidMessage) predicate.ValidMessage {
	return predicate.ValidMessage(sql.NotPredicates(p))
}
