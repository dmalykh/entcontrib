// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/dmalykh/entcontrib/entproto/internal/entprototest/ent/messagewithfieldone"
)

// MessageWithFieldOne is the model entity for the MessageWithFieldOne schema.
type MessageWithFieldOne struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// FieldOne holds the value of the "field_one" field.
	FieldOne     int32 `json:"field_one,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MessageWithFieldOne) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case messagewithfieldone.FieldID, messagewithfieldone.FieldFieldOne:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MessageWithFieldOne fields.
func (mwfo *MessageWithFieldOne) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case messagewithfieldone.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			mwfo.ID = int(value.Int64)
		case messagewithfieldone.FieldFieldOne:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field field_one", values[i])
			} else if value.Valid {
				mwfo.FieldOne = int32(value.Int64)
			}
		default:
			mwfo.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the MessageWithFieldOne.
// This includes values selected through modifiers, order, etc.
func (mwfo *MessageWithFieldOne) Value(name string) (ent.Value, error) {
	return mwfo.selectValues.Get(name)
}

// Update returns a builder for updating this MessageWithFieldOne.
// Note that you need to call MessageWithFieldOne.Unwrap() before calling this method if this MessageWithFieldOne
// was returned from a transaction, and the transaction was committed or rolled back.
func (mwfo *MessageWithFieldOne) Update() *MessageWithFieldOneUpdateOne {
	return NewMessageWithFieldOneClient(mwfo.config).UpdateOne(mwfo)
}

// Unwrap unwraps the MessageWithFieldOne entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (mwfo *MessageWithFieldOne) Unwrap() *MessageWithFieldOne {
	_tx, ok := mwfo.config.driver.(*txDriver)
	if !ok {
		panic("ent: MessageWithFieldOne is not a transactional entity")
	}
	mwfo.config.driver = _tx.drv
	return mwfo
}

// String implements the fmt.Stringer.
func (mwfo *MessageWithFieldOne) String() string {
	var builder strings.Builder
	builder.WriteString("MessageWithFieldOne(")
	builder.WriteString(fmt.Sprintf("id=%v, ", mwfo.ID))
	builder.WriteString("field_one=")
	builder.WriteString(fmt.Sprintf("%v", mwfo.FieldOne))
	builder.WriteByte(')')
	return builder.String()
}

// MessageWithFieldOnes is a parsable slice of MessageWithFieldOne.
type MessageWithFieldOnes []*MessageWithFieldOne
