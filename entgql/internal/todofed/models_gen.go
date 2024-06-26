// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package todofed

import (
	"github.com/dmalykh/entcontrib/entgql/internal/todofed/ent/todo"
)

type TodoInput struct {
	Status     todo.Status `json:"status"`
	Priority   *int        `json:"priority,omitempty"`
	Text       string      `json:"text"`
	Parent     *int        `json:"parent,omitempty"`
	CategoryID *int        `json:"category_id,omitempty"`
}
