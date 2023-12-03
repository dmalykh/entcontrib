// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dmalykh/entcontrib/entproto/internal/entprototest/ent/blogpost"
	"github.com/dmalykh/entcontrib/entproto/internal/entprototest/ent/category"
	"github.com/dmalykh/entcontrib/entproto/internal/entprototest/ent/predicate"
	"github.com/dmalykh/entcontrib/entproto/internal/entprototest/ent/user"
)

// BlogPostUpdate is the builder for updating BlogPost entities.
type BlogPostUpdate struct {
	config
	hooks    []Hook
	mutation *BlogPostMutation
}

// Where appends a list predicates to the BlogPostUpdate builder.
func (bpu *BlogPostUpdate) Where(ps ...predicate.BlogPost) *BlogPostUpdate {
	bpu.mutation.Where(ps...)
	return bpu
}

// SetTitle sets the "title" field.
func (bpu *BlogPostUpdate) SetTitle(s string) *BlogPostUpdate {
	bpu.mutation.SetTitle(s)
	return bpu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (bpu *BlogPostUpdate) SetNillableTitle(s *string) *BlogPostUpdate {
	if s != nil {
		bpu.SetTitle(*s)
	}
	return bpu
}

// SetBody sets the "body" field.
func (bpu *BlogPostUpdate) SetBody(s string) *BlogPostUpdate {
	bpu.mutation.SetBody(s)
	return bpu
}

// SetNillableBody sets the "body" field if the given value is not nil.
func (bpu *BlogPostUpdate) SetNillableBody(s *string) *BlogPostUpdate {
	if s != nil {
		bpu.SetBody(*s)
	}
	return bpu
}

// SetExternalID sets the "external_id" field.
func (bpu *BlogPostUpdate) SetExternalID(i int) *BlogPostUpdate {
	bpu.mutation.ResetExternalID()
	bpu.mutation.SetExternalID(i)
	return bpu
}

// SetNillableExternalID sets the "external_id" field if the given value is not nil.
func (bpu *BlogPostUpdate) SetNillableExternalID(i *int) *BlogPostUpdate {
	if i != nil {
		bpu.SetExternalID(*i)
	}
	return bpu
}

// AddExternalID adds i to the "external_id" field.
func (bpu *BlogPostUpdate) AddExternalID(i int) *BlogPostUpdate {
	bpu.mutation.AddExternalID(i)
	return bpu
}

// SetAuthorID sets the "author" edge to the User entity by ID.
func (bpu *BlogPostUpdate) SetAuthorID(id int) *BlogPostUpdate {
	bpu.mutation.SetAuthorID(id)
	return bpu
}

// SetNillableAuthorID sets the "author" edge to the User entity by ID if the given value is not nil.
func (bpu *BlogPostUpdate) SetNillableAuthorID(id *int) *BlogPostUpdate {
	if id != nil {
		bpu = bpu.SetAuthorID(*id)
	}
	return bpu
}

// SetAuthor sets the "author" edge to the User entity.
func (bpu *BlogPostUpdate) SetAuthor(u *User) *BlogPostUpdate {
	return bpu.SetAuthorID(u.ID)
}

// AddCategoryIDs adds the "categories" edge to the Category entity by IDs.
func (bpu *BlogPostUpdate) AddCategoryIDs(ids ...int) *BlogPostUpdate {
	bpu.mutation.AddCategoryIDs(ids...)
	return bpu
}

// AddCategories adds the "categories" edges to the Category entity.
func (bpu *BlogPostUpdate) AddCategories(c ...*Category) *BlogPostUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return bpu.AddCategoryIDs(ids...)
}

// Mutation returns the BlogPostMutation object of the builder.
func (bpu *BlogPostUpdate) Mutation() *BlogPostMutation {
	return bpu.mutation
}

// ClearAuthor clears the "author" edge to the User entity.
func (bpu *BlogPostUpdate) ClearAuthor() *BlogPostUpdate {
	bpu.mutation.ClearAuthor()
	return bpu
}

// ClearCategories clears all "categories" edges to the Category entity.
func (bpu *BlogPostUpdate) ClearCategories() *BlogPostUpdate {
	bpu.mutation.ClearCategories()
	return bpu
}

// RemoveCategoryIDs removes the "categories" edge to Category entities by IDs.
func (bpu *BlogPostUpdate) RemoveCategoryIDs(ids ...int) *BlogPostUpdate {
	bpu.mutation.RemoveCategoryIDs(ids...)
	return bpu
}

// RemoveCategories removes "categories" edges to Category entities.
func (bpu *BlogPostUpdate) RemoveCategories(c ...*Category) *BlogPostUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return bpu.RemoveCategoryIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bpu *BlogPostUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, bpu.sqlSave, bpu.mutation, bpu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bpu *BlogPostUpdate) SaveX(ctx context.Context) int {
	affected, err := bpu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bpu *BlogPostUpdate) Exec(ctx context.Context) error {
	_, err := bpu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bpu *BlogPostUpdate) ExecX(ctx context.Context) {
	if err := bpu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bpu *BlogPostUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(blogpost.Table, blogpost.Columns, sqlgraph.NewFieldSpec(blogpost.FieldID, field.TypeInt))
	if ps := bpu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bpu.mutation.Title(); ok {
		_spec.SetField(blogpost.FieldTitle, field.TypeString, value)
	}
	if value, ok := bpu.mutation.Body(); ok {
		_spec.SetField(blogpost.FieldBody, field.TypeString, value)
	}
	if value, ok := bpu.mutation.ExternalID(); ok {
		_spec.SetField(blogpost.FieldExternalID, field.TypeInt, value)
	}
	if value, ok := bpu.mutation.AddedExternalID(); ok {
		_spec.AddField(blogpost.FieldExternalID, field.TypeInt, value)
	}
	if bpu.mutation.AuthorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   blogpost.AuthorTable,
			Columns: []string{blogpost.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bpu.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   blogpost.AuthorTable,
			Columns: []string{blogpost.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bpu.mutation.CategoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   blogpost.CategoriesTable,
			Columns: blogpost.CategoriesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bpu.mutation.RemovedCategoriesIDs(); len(nodes) > 0 && !bpu.mutation.CategoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   blogpost.CategoriesTable,
			Columns: blogpost.CategoriesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bpu.mutation.CategoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   blogpost.CategoriesTable,
			Columns: blogpost.CategoriesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bpu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{blogpost.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bpu.mutation.done = true
	return n, nil
}

// BlogPostUpdateOne is the builder for updating a single BlogPost entity.
type BlogPostUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BlogPostMutation
}

// SetTitle sets the "title" field.
func (bpuo *BlogPostUpdateOne) SetTitle(s string) *BlogPostUpdateOne {
	bpuo.mutation.SetTitle(s)
	return bpuo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (bpuo *BlogPostUpdateOne) SetNillableTitle(s *string) *BlogPostUpdateOne {
	if s != nil {
		bpuo.SetTitle(*s)
	}
	return bpuo
}

// SetBody sets the "body" field.
func (bpuo *BlogPostUpdateOne) SetBody(s string) *BlogPostUpdateOne {
	bpuo.mutation.SetBody(s)
	return bpuo
}

// SetNillableBody sets the "body" field if the given value is not nil.
func (bpuo *BlogPostUpdateOne) SetNillableBody(s *string) *BlogPostUpdateOne {
	if s != nil {
		bpuo.SetBody(*s)
	}
	return bpuo
}

// SetExternalID sets the "external_id" field.
func (bpuo *BlogPostUpdateOne) SetExternalID(i int) *BlogPostUpdateOne {
	bpuo.mutation.ResetExternalID()
	bpuo.mutation.SetExternalID(i)
	return bpuo
}

// SetNillableExternalID sets the "external_id" field if the given value is not nil.
func (bpuo *BlogPostUpdateOne) SetNillableExternalID(i *int) *BlogPostUpdateOne {
	if i != nil {
		bpuo.SetExternalID(*i)
	}
	return bpuo
}

// AddExternalID adds i to the "external_id" field.
func (bpuo *BlogPostUpdateOne) AddExternalID(i int) *BlogPostUpdateOne {
	bpuo.mutation.AddExternalID(i)
	return bpuo
}

// SetAuthorID sets the "author" edge to the User entity by ID.
func (bpuo *BlogPostUpdateOne) SetAuthorID(id int) *BlogPostUpdateOne {
	bpuo.mutation.SetAuthorID(id)
	return bpuo
}

// SetNillableAuthorID sets the "author" edge to the User entity by ID if the given value is not nil.
func (bpuo *BlogPostUpdateOne) SetNillableAuthorID(id *int) *BlogPostUpdateOne {
	if id != nil {
		bpuo = bpuo.SetAuthorID(*id)
	}
	return bpuo
}

// SetAuthor sets the "author" edge to the User entity.
func (bpuo *BlogPostUpdateOne) SetAuthor(u *User) *BlogPostUpdateOne {
	return bpuo.SetAuthorID(u.ID)
}

// AddCategoryIDs adds the "categories" edge to the Category entity by IDs.
func (bpuo *BlogPostUpdateOne) AddCategoryIDs(ids ...int) *BlogPostUpdateOne {
	bpuo.mutation.AddCategoryIDs(ids...)
	return bpuo
}

// AddCategories adds the "categories" edges to the Category entity.
func (bpuo *BlogPostUpdateOne) AddCategories(c ...*Category) *BlogPostUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return bpuo.AddCategoryIDs(ids...)
}

// Mutation returns the BlogPostMutation object of the builder.
func (bpuo *BlogPostUpdateOne) Mutation() *BlogPostMutation {
	return bpuo.mutation
}

// ClearAuthor clears the "author" edge to the User entity.
func (bpuo *BlogPostUpdateOne) ClearAuthor() *BlogPostUpdateOne {
	bpuo.mutation.ClearAuthor()
	return bpuo
}

// ClearCategories clears all "categories" edges to the Category entity.
func (bpuo *BlogPostUpdateOne) ClearCategories() *BlogPostUpdateOne {
	bpuo.mutation.ClearCategories()
	return bpuo
}

// RemoveCategoryIDs removes the "categories" edge to Category entities by IDs.
func (bpuo *BlogPostUpdateOne) RemoveCategoryIDs(ids ...int) *BlogPostUpdateOne {
	bpuo.mutation.RemoveCategoryIDs(ids...)
	return bpuo
}

// RemoveCategories removes "categories" edges to Category entities.
func (bpuo *BlogPostUpdateOne) RemoveCategories(c ...*Category) *BlogPostUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return bpuo.RemoveCategoryIDs(ids...)
}

// Where appends a list predicates to the BlogPostUpdate builder.
func (bpuo *BlogPostUpdateOne) Where(ps ...predicate.BlogPost) *BlogPostUpdateOne {
	bpuo.mutation.Where(ps...)
	return bpuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (bpuo *BlogPostUpdateOne) Select(field string, fields ...string) *BlogPostUpdateOne {
	bpuo.fields = append([]string{field}, fields...)
	return bpuo
}

// Save executes the query and returns the updated BlogPost entity.
func (bpuo *BlogPostUpdateOne) Save(ctx context.Context) (*BlogPost, error) {
	return withHooks(ctx, bpuo.sqlSave, bpuo.mutation, bpuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bpuo *BlogPostUpdateOne) SaveX(ctx context.Context) *BlogPost {
	node, err := bpuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (bpuo *BlogPostUpdateOne) Exec(ctx context.Context) error {
	_, err := bpuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bpuo *BlogPostUpdateOne) ExecX(ctx context.Context) {
	if err := bpuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bpuo *BlogPostUpdateOne) sqlSave(ctx context.Context) (_node *BlogPost, err error) {
	_spec := sqlgraph.NewUpdateSpec(blogpost.Table, blogpost.Columns, sqlgraph.NewFieldSpec(blogpost.FieldID, field.TypeInt))
	id, ok := bpuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "BlogPost.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := bpuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, blogpost.FieldID)
		for _, f := range fields {
			if !blogpost.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != blogpost.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := bpuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bpuo.mutation.Title(); ok {
		_spec.SetField(blogpost.FieldTitle, field.TypeString, value)
	}
	if value, ok := bpuo.mutation.Body(); ok {
		_spec.SetField(blogpost.FieldBody, field.TypeString, value)
	}
	if value, ok := bpuo.mutation.ExternalID(); ok {
		_spec.SetField(blogpost.FieldExternalID, field.TypeInt, value)
	}
	if value, ok := bpuo.mutation.AddedExternalID(); ok {
		_spec.AddField(blogpost.FieldExternalID, field.TypeInt, value)
	}
	if bpuo.mutation.AuthorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   blogpost.AuthorTable,
			Columns: []string{blogpost.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bpuo.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   blogpost.AuthorTable,
			Columns: []string{blogpost.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bpuo.mutation.CategoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   blogpost.CategoriesTable,
			Columns: blogpost.CategoriesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bpuo.mutation.RemovedCategoriesIDs(); len(nodes) > 0 && !bpuo.mutation.CategoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   blogpost.CategoriesTable,
			Columns: blogpost.CategoriesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bpuo.mutation.CategoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   blogpost.CategoriesTable,
			Columns: blogpost.CategoriesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &BlogPost{config: bpuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, bpuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{blogpost.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	bpuo.mutation.done = true
	return _node, nil
}
