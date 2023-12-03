// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/dmalykh/entcontrib/schemast/internal/mutatetest/ent/withoutfields"
)

// WithoutFields is the model entity for the WithoutFields schema.
type WithoutFields struct {
	config
	// ID of the ent.
	ID           int `json:"id,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*WithoutFields) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case withoutfields.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the WithoutFields fields.
func (wf *WithoutFields) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case withoutfields.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			wf.ID = int(value.Int64)
		default:
			wf.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the WithoutFields.
// This includes values selected through modifiers, order, etc.
func (wf *WithoutFields) Value(name string) (ent.Value, error) {
	return wf.selectValues.Get(name)
}

// Update returns a builder for updating this WithoutFields.
// Note that you need to call WithoutFields.Unwrap() before calling this method if this WithoutFields
// was returned from a transaction, and the transaction was committed or rolled back.
func (wf *WithoutFields) Update() *WithoutFieldsUpdateOne {
	return NewWithoutFieldsClient(wf.config).UpdateOne(wf)
}

// Unwrap unwraps the WithoutFields entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (wf *WithoutFields) Unwrap() *WithoutFields {
	_tx, ok := wf.config.driver.(*txDriver)
	if !ok {
		panic("ent: WithoutFields is not a transactional entity")
	}
	wf.config.driver = _tx.drv
	return wf
}

// String implements the fmt.Stringer.
func (wf *WithoutFields) String() string {
	var builder strings.Builder
	builder.WriteString("WithoutFields(")
	builder.WriteString(fmt.Sprintf("id=%v", wf.ID))
	builder.WriteByte(')')
	return builder.String()
}

// WithoutFieldsSlice is a parsable slice of WithoutFields.
type WithoutFieldsSlice []*WithoutFields
