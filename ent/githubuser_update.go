// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"slackbot/ent/githubuser"
	"slackbot/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// GithubUserUpdate is the builder for updating GithubUser entities.
type GithubUserUpdate struct {
	config
	hooks    []Hook
	mutation *GithubUserMutation
}

// Where appends a list predicates to the GithubUserUpdate builder.
func (guu *GithubUserUpdate) Where(ps ...predicate.GithubUser) *GithubUserUpdate {
	guu.mutation.Where(ps...)
	return guu
}

// SetGhUsername sets the "gh_username" field.
func (guu *GithubUserUpdate) SetGhUsername(s string) *GithubUserUpdate {
	guu.mutation.SetGhUsername(s)
	return guu
}

// SetNillableGhUsername sets the "gh_username" field if the given value is not nil.
func (guu *GithubUserUpdate) SetNillableGhUsername(s *string) *GithubUserUpdate {
	if s != nil {
		guu.SetGhUsername(*s)
	}
	return guu
}

// SetSlackID sets the "slack_id" field.
func (guu *GithubUserUpdate) SetSlackID(s string) *GithubUserUpdate {
	guu.mutation.SetSlackID(s)
	return guu
}

// SetNillableSlackID sets the "slack_id" field if the given value is not nil.
func (guu *GithubUserUpdate) SetNillableSlackID(s *string) *GithubUserUpdate {
	if s != nil {
		guu.SetSlackID(*s)
	}
	return guu
}

// SetGhAccessToken sets the "gh_access_token" field.
func (guu *GithubUserUpdate) SetGhAccessToken(s string) *GithubUserUpdate {
	guu.mutation.SetGhAccessToken(s)
	return guu
}

// SetNillableGhAccessToken sets the "gh_access_token" field if the given value is not nil.
func (guu *GithubUserUpdate) SetNillableGhAccessToken(s *string) *GithubUserUpdate {
	if s != nil {
		guu.SetGhAccessToken(*s)
	}
	return guu
}

// SetCreatedAt sets the "created_at" field.
func (guu *GithubUserUpdate) SetCreatedAt(t time.Time) *GithubUserUpdate {
	guu.mutation.SetCreatedAt(t)
	return guu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (guu *GithubUserUpdate) SetNillableCreatedAt(t *time.Time) *GithubUserUpdate {
	if t != nil {
		guu.SetCreatedAt(*t)
	}
	return guu
}

// SetUpdatedAt sets the "updated_at" field.
func (guu *GithubUserUpdate) SetUpdatedAt(t time.Time) *GithubUserUpdate {
	guu.mutation.SetUpdatedAt(t)
	return guu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (guu *GithubUserUpdate) SetNillableUpdatedAt(t *time.Time) *GithubUserUpdate {
	if t != nil {
		guu.SetUpdatedAt(*t)
	}
	return guu
}

// Mutation returns the GithubUserMutation object of the builder.
func (guu *GithubUserUpdate) Mutation() *GithubUserMutation {
	return guu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (guu *GithubUserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, guu.sqlSave, guu.mutation, guu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (guu *GithubUserUpdate) SaveX(ctx context.Context) int {
	affected, err := guu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (guu *GithubUserUpdate) Exec(ctx context.Context) error {
	_, err := guu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (guu *GithubUserUpdate) ExecX(ctx context.Context) {
	if err := guu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (guu *GithubUserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(githubuser.Table, githubuser.Columns, sqlgraph.NewFieldSpec(githubuser.FieldID, field.TypeUUID))
	if ps := guu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := guu.mutation.GhUsername(); ok {
		_spec.SetField(githubuser.FieldGhUsername, field.TypeString, value)
	}
	if value, ok := guu.mutation.SlackID(); ok {
		_spec.SetField(githubuser.FieldSlackID, field.TypeString, value)
	}
	if value, ok := guu.mutation.GhAccessToken(); ok {
		_spec.SetField(githubuser.FieldGhAccessToken, field.TypeString, value)
	}
	if value, ok := guu.mutation.CreatedAt(); ok {
		_spec.SetField(githubuser.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := guu.mutation.UpdatedAt(); ok {
		_spec.SetField(githubuser.FieldUpdatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, guu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{githubuser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	guu.mutation.done = true
	return n, nil
}

// GithubUserUpdateOne is the builder for updating a single GithubUser entity.
type GithubUserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GithubUserMutation
}

// SetGhUsername sets the "gh_username" field.
func (guuo *GithubUserUpdateOne) SetGhUsername(s string) *GithubUserUpdateOne {
	guuo.mutation.SetGhUsername(s)
	return guuo
}

// SetNillableGhUsername sets the "gh_username" field if the given value is not nil.
func (guuo *GithubUserUpdateOne) SetNillableGhUsername(s *string) *GithubUserUpdateOne {
	if s != nil {
		guuo.SetGhUsername(*s)
	}
	return guuo
}

// SetSlackID sets the "slack_id" field.
func (guuo *GithubUserUpdateOne) SetSlackID(s string) *GithubUserUpdateOne {
	guuo.mutation.SetSlackID(s)
	return guuo
}

// SetNillableSlackID sets the "slack_id" field if the given value is not nil.
func (guuo *GithubUserUpdateOne) SetNillableSlackID(s *string) *GithubUserUpdateOne {
	if s != nil {
		guuo.SetSlackID(*s)
	}
	return guuo
}

// SetGhAccessToken sets the "gh_access_token" field.
func (guuo *GithubUserUpdateOne) SetGhAccessToken(s string) *GithubUserUpdateOne {
	guuo.mutation.SetGhAccessToken(s)
	return guuo
}

// SetNillableGhAccessToken sets the "gh_access_token" field if the given value is not nil.
func (guuo *GithubUserUpdateOne) SetNillableGhAccessToken(s *string) *GithubUserUpdateOne {
	if s != nil {
		guuo.SetGhAccessToken(*s)
	}
	return guuo
}

// SetCreatedAt sets the "created_at" field.
func (guuo *GithubUserUpdateOne) SetCreatedAt(t time.Time) *GithubUserUpdateOne {
	guuo.mutation.SetCreatedAt(t)
	return guuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (guuo *GithubUserUpdateOne) SetNillableCreatedAt(t *time.Time) *GithubUserUpdateOne {
	if t != nil {
		guuo.SetCreatedAt(*t)
	}
	return guuo
}

// SetUpdatedAt sets the "updated_at" field.
func (guuo *GithubUserUpdateOne) SetUpdatedAt(t time.Time) *GithubUserUpdateOne {
	guuo.mutation.SetUpdatedAt(t)
	return guuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (guuo *GithubUserUpdateOne) SetNillableUpdatedAt(t *time.Time) *GithubUserUpdateOne {
	if t != nil {
		guuo.SetUpdatedAt(*t)
	}
	return guuo
}

// Mutation returns the GithubUserMutation object of the builder.
func (guuo *GithubUserUpdateOne) Mutation() *GithubUserMutation {
	return guuo.mutation
}

// Where appends a list predicates to the GithubUserUpdate builder.
func (guuo *GithubUserUpdateOne) Where(ps ...predicate.GithubUser) *GithubUserUpdateOne {
	guuo.mutation.Where(ps...)
	return guuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (guuo *GithubUserUpdateOne) Select(field string, fields ...string) *GithubUserUpdateOne {
	guuo.fields = append([]string{field}, fields...)
	return guuo
}

// Save executes the query and returns the updated GithubUser entity.
func (guuo *GithubUserUpdateOne) Save(ctx context.Context) (*GithubUser, error) {
	return withHooks(ctx, guuo.sqlSave, guuo.mutation, guuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (guuo *GithubUserUpdateOne) SaveX(ctx context.Context) *GithubUser {
	node, err := guuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (guuo *GithubUserUpdateOne) Exec(ctx context.Context) error {
	_, err := guuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (guuo *GithubUserUpdateOne) ExecX(ctx context.Context) {
	if err := guuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (guuo *GithubUserUpdateOne) sqlSave(ctx context.Context) (_node *GithubUser, err error) {
	_spec := sqlgraph.NewUpdateSpec(githubuser.Table, githubuser.Columns, sqlgraph.NewFieldSpec(githubuser.FieldID, field.TypeUUID))
	id, ok := guuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "GithubUser.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := guuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, githubuser.FieldID)
		for _, f := range fields {
			if !githubuser.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != githubuser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := guuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := guuo.mutation.GhUsername(); ok {
		_spec.SetField(githubuser.FieldGhUsername, field.TypeString, value)
	}
	if value, ok := guuo.mutation.SlackID(); ok {
		_spec.SetField(githubuser.FieldSlackID, field.TypeString, value)
	}
	if value, ok := guuo.mutation.GhAccessToken(); ok {
		_spec.SetField(githubuser.FieldGhAccessToken, field.TypeString, value)
	}
	if value, ok := guuo.mutation.CreatedAt(); ok {
		_spec.SetField(githubuser.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := guuo.mutation.UpdatedAt(); ok {
		_spec.SetField(githubuser.FieldUpdatedAt, field.TypeTime, value)
	}
	_node = &GithubUser{config: guuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, guuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{githubuser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	guuo.mutation.done = true
	return _node, nil
}