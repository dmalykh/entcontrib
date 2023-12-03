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
	"fmt"
	"strings"

	"github.com/dmalykh/entcontrib/entgql/internal/todopulid/ent/billproduct"
	"github.com/dmalykh/entcontrib/entgql/internal/todopulid/ent/schema/pulid"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// BillProduct is the model entity for the BillProduct schema.
type BillProduct struct {
	config `json:"-"`
	// ID of the ent.
	ID pulid.ID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Sku holds the value of the "sku" field.
	Sku string `json:"sku,omitempty"`
	// Quantity holds the value of the "quantity" field.
	Quantity     uint64 `json:"quantity,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*BillProduct) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case billproduct.FieldID:
			values[i] = new(pulid.ID)
		case billproduct.FieldQuantity:
			values[i] = new(sql.NullInt64)
		case billproduct.FieldName, billproduct.FieldSku:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the BillProduct fields.
func (bp *BillProduct) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case billproduct.FieldID:
			if value, ok := values[i].(*pulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				bp.ID = *value
			}
		case billproduct.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				bp.Name = value.String
			}
		case billproduct.FieldSku:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sku", values[i])
			} else if value.Valid {
				bp.Sku = value.String
			}
		case billproduct.FieldQuantity:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field quantity", values[i])
			} else if value.Valid {
				bp.Quantity = uint64(value.Int64)
			}
		default:
			bp.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the BillProduct.
// This includes values selected through modifiers, order, etc.
func (bp *BillProduct) Value(name string) (ent.Value, error) {
	return bp.selectValues.Get(name)
}

// Update returns a builder for updating this BillProduct.
// Note that you need to call BillProduct.Unwrap() before calling this method if this BillProduct
// was returned from a transaction, and the transaction was committed or rolled back.
func (bp *BillProduct) Update() *BillProductUpdateOne {
	return NewBillProductClient(bp.config).UpdateOne(bp)
}

// Unwrap unwraps the BillProduct entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (bp *BillProduct) Unwrap() *BillProduct {
	_tx, ok := bp.config.driver.(*txDriver)
	if !ok {
		panic("ent: BillProduct is not a transactional entity")
	}
	bp.config.driver = _tx.drv
	return bp
}

// String implements the fmt.Stringer.
func (bp *BillProduct) String() string {
	var builder strings.Builder
	builder.WriteString("BillProduct(")
	builder.WriteString(fmt.Sprintf("id=%v, ", bp.ID))
	builder.WriteString("name=")
	builder.WriteString(bp.Name)
	builder.WriteString(", ")
	builder.WriteString("sku=")
	builder.WriteString(bp.Sku)
	builder.WriteString(", ")
	builder.WriteString("quantity=")
	builder.WriteString(fmt.Sprintf("%v", bp.Quantity))
	builder.WriteByte(')')
	return builder.String()
}

// BillProducts is a parsable slice of BillProduct.
type BillProducts []*BillProduct
