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
	"time"

	"github.com/dmalykh/entcontrib/entgql/internal/todo/ent/category"
	"github.com/dmalykh/entcontrib/entgql/internal/todo/ent/friendship"
	"github.com/dmalykh/entcontrib/entgql/internal/todo/ent/group"
	"github.com/dmalykh/entcontrib/entgql/internal/todo/ent/onetomany"
	"github.com/dmalykh/entcontrib/entgql/internal/todo/ent/schema"
	"github.com/dmalykh/entcontrib/entgql/internal/todo/ent/todo"
	"github.com/dmalykh/entcontrib/entgql/internal/todo/ent/user"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	categoryFields := schema.Category{}.Fields()
	_ = categoryFields
	// categoryDescText is the schema descriptor for text field.
	categoryDescText := categoryFields[1].Descriptor()
	// category.TextValidator is a validator for the "text" field. It is called by the builders before save.
	category.TextValidator = categoryDescText.Validators[0].(func(string) error)
	friendshipFields := schema.Friendship{}.Fields()
	_ = friendshipFields
	// friendshipDescCreatedAt is the schema descriptor for created_at field.
	friendshipDescCreatedAt := friendshipFields[0].Descriptor()
	// friendship.DefaultCreatedAt holds the default value on creation for the created_at field.
	friendship.DefaultCreatedAt = friendshipDescCreatedAt.Default.(func() time.Time)
	groupFields := schema.Group{}.Fields()
	_ = groupFields
	// groupDescName is the schema descriptor for name field.
	groupDescName := groupFields[0].Descriptor()
	// group.DefaultName holds the default value on creation for the name field.
	group.DefaultName = groupDescName.Default.(string)
	onetomanyFields := schema.OneToMany{}.Fields()
	_ = onetomanyFields
	// onetomanyDescName is the schema descriptor for name field.
	onetomanyDescName := onetomanyFields[0].Descriptor()
	// onetomany.NameValidator is a validator for the "name" field. It is called by the builders before save.
	onetomany.NameValidator = onetomanyDescName.Validators[0].(func(string) error)
	todoFields := schema.Todo{}.Fields()
	_ = todoFields
	// todoDescCreatedAt is the schema descriptor for created_at field.
	todoDescCreatedAt := todoFields[0].Descriptor()
	// todo.DefaultCreatedAt holds the default value on creation for the created_at field.
	todo.DefaultCreatedAt = todoDescCreatedAt.Default.(func() time.Time)
	// todoDescPriority is the schema descriptor for priority field.
	todoDescPriority := todoFields[2].Descriptor()
	// todo.DefaultPriority holds the default value on creation for the priority field.
	todo.DefaultPriority = todoDescPriority.Default.(int)
	// todoDescText is the schema descriptor for text field.
	todoDescText := todoFields[3].Descriptor()
	// todo.TextValidator is a validator for the "text" field. It is called by the builders before save.
	todo.TextValidator = todoDescText.Validators[0].(func(string) error)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[0].Descriptor()
	// user.DefaultName holds the default value on creation for the name field.
	user.DefaultName = userDescName.Default.(string)
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[1].Descriptor()
	// user.DefaultUsername holds the default value on creation for the username field.
	user.DefaultUsername = userDescUsername.Default.(func() uuid.UUID)
}
