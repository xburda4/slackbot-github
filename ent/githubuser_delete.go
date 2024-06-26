// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"slackbot/ent/githubuser"
	"slackbot/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// GithubUserDelete is the builder for deleting a GithubUser entity.
type GithubUserDelete struct {
	config
	hooks    []Hook
	mutation *GithubUserMutation
}

// Where appends a list predicates to the GithubUserDelete builder.
func (gud *GithubUserDelete) Where(ps ...predicate.GithubUser) *GithubUserDelete {
	gud.mutation.Where(ps...)
	return gud
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (gud *GithubUserDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, gud.sqlExec, gud.mutation, gud.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (gud *GithubUserDelete) ExecX(ctx context.Context) int {
	n, err := gud.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (gud *GithubUserDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(githubuser.Table, sqlgraph.NewFieldSpec(githubuser.FieldID, field.TypeUUID))
	if ps := gud.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, gud.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	gud.mutation.done = true
	return affected, err
}

// GithubUserDeleteOne is the builder for deleting a single GithubUser entity.
type GithubUserDeleteOne struct {
	gud *GithubUserDelete
}

// Where appends a list predicates to the GithubUserDelete builder.
func (gudo *GithubUserDeleteOne) Where(ps ...predicate.GithubUser) *GithubUserDeleteOne {
	gudo.gud.mutation.Where(ps...)
	return gudo
}

// Exec executes the deletion query.
func (gudo *GithubUserDeleteOne) Exec(ctx context.Context) error {
	n, err := gudo.gud.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{githubuser.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (gudo *GithubUserDeleteOne) ExecX(ctx context.Context) {
	if err := gudo.Exec(ctx); err != nil {
		panic(err)
	}
}
