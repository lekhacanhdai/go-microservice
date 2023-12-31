// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"grpcserver/data/ent/predicate"
	"grpcserver/data/ent/userentity"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserEntityDelete is the builder for deleting a UserEntity entity.
type UserEntityDelete struct {
	config
	hooks    []Hook
	mutation *UserEntityMutation
}

// Where appends a list predicates to the UserEntityDelete builder.
func (ued *UserEntityDelete) Where(ps ...predicate.UserEntity) *UserEntityDelete {
	ued.mutation.Where(ps...)
	return ued
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ued *UserEntityDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ued.sqlExec, ued.mutation, ued.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ued *UserEntityDelete) ExecX(ctx context.Context) int {
	n, err := ued.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ued *UserEntityDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(userentity.Table, sqlgraph.NewFieldSpec(userentity.FieldID, field.TypeInt))
	if ps := ued.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ued.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ued.mutation.done = true
	return affected, err
}

// UserEntityDeleteOne is the builder for deleting a single UserEntity entity.
type UserEntityDeleteOne struct {
	ued *UserEntityDelete
}

// Where appends a list predicates to the UserEntityDelete builder.
func (uedo *UserEntityDeleteOne) Where(ps ...predicate.UserEntity) *UserEntityDeleteOne {
	uedo.ued.mutation.Where(ps...)
	return uedo
}

// Exec executes the deletion query.
func (uedo *UserEntityDeleteOne) Exec(ctx context.Context) error {
	n, err := uedo.ued.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{userentity.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (uedo *UserEntityDeleteOne) ExecX(ctx context.Context) {
	if err := uedo.Exec(ctx); err != nil {
		panic(err)
	}
}
