// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"grpcserver/data/ent/predicate"
	"grpcserver/data/ent/userentity"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserEntityUpdate is the builder for updating UserEntity entities.
type UserEntityUpdate struct {
	config
	hooks    []Hook
	mutation *UserEntityMutation
}

// Where appends a list predicates to the UserEntityUpdate builder.
func (ueu *UserEntityUpdate) Where(ps ...predicate.UserEntity) *UserEntityUpdate {
	ueu.mutation.Where(ps...)
	return ueu
}

// SetUsername sets the "username" field.
func (ueu *UserEntityUpdate) SetUsername(s string) *UserEntityUpdate {
	ueu.mutation.SetUsername(s)
	return ueu
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (ueu *UserEntityUpdate) SetNillableUsername(s *string) *UserEntityUpdate {
	if s != nil {
		ueu.SetUsername(*s)
	}
	return ueu
}

// SetPassword sets the "password" field.
func (ueu *UserEntityUpdate) SetPassword(s string) *UserEntityUpdate {
	ueu.mutation.SetPassword(s)
	return ueu
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (ueu *UserEntityUpdate) SetNillablePassword(s *string) *UserEntityUpdate {
	if s != nil {
		ueu.SetPassword(*s)
	}
	return ueu
}

// Mutation returns the UserEntityMutation object of the builder.
func (ueu *UserEntityUpdate) Mutation() *UserEntityMutation {
	return ueu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ueu *UserEntityUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ueu.sqlSave, ueu.mutation, ueu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ueu *UserEntityUpdate) SaveX(ctx context.Context) int {
	affected, err := ueu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ueu *UserEntityUpdate) Exec(ctx context.Context) error {
	_, err := ueu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ueu *UserEntityUpdate) ExecX(ctx context.Context) {
	if err := ueu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ueu *UserEntityUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(userentity.Table, userentity.Columns, sqlgraph.NewFieldSpec(userentity.FieldID, field.TypeInt))
	if ps := ueu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ueu.mutation.Username(); ok {
		_spec.SetField(userentity.FieldUsername, field.TypeString, value)
	}
	if value, ok := ueu.mutation.Password(); ok {
		_spec.SetField(userentity.FieldPassword, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ueu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userentity.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ueu.mutation.done = true
	return n, nil
}

// UserEntityUpdateOne is the builder for updating a single UserEntity entity.
type UserEntityUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserEntityMutation
}

// SetUsername sets the "username" field.
func (ueuo *UserEntityUpdateOne) SetUsername(s string) *UserEntityUpdateOne {
	ueuo.mutation.SetUsername(s)
	return ueuo
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (ueuo *UserEntityUpdateOne) SetNillableUsername(s *string) *UserEntityUpdateOne {
	if s != nil {
		ueuo.SetUsername(*s)
	}
	return ueuo
}

// SetPassword sets the "password" field.
func (ueuo *UserEntityUpdateOne) SetPassword(s string) *UserEntityUpdateOne {
	ueuo.mutation.SetPassword(s)
	return ueuo
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (ueuo *UserEntityUpdateOne) SetNillablePassword(s *string) *UserEntityUpdateOne {
	if s != nil {
		ueuo.SetPassword(*s)
	}
	return ueuo
}

// Mutation returns the UserEntityMutation object of the builder.
func (ueuo *UserEntityUpdateOne) Mutation() *UserEntityMutation {
	return ueuo.mutation
}

// Where appends a list predicates to the UserEntityUpdate builder.
func (ueuo *UserEntityUpdateOne) Where(ps ...predicate.UserEntity) *UserEntityUpdateOne {
	ueuo.mutation.Where(ps...)
	return ueuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ueuo *UserEntityUpdateOne) Select(field string, fields ...string) *UserEntityUpdateOne {
	ueuo.fields = append([]string{field}, fields...)
	return ueuo
}

// Save executes the query and returns the updated UserEntity entity.
func (ueuo *UserEntityUpdateOne) Save(ctx context.Context) (*UserEntity, error) {
	return withHooks(ctx, ueuo.sqlSave, ueuo.mutation, ueuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ueuo *UserEntityUpdateOne) SaveX(ctx context.Context) *UserEntity {
	node, err := ueuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ueuo *UserEntityUpdateOne) Exec(ctx context.Context) error {
	_, err := ueuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ueuo *UserEntityUpdateOne) ExecX(ctx context.Context) {
	if err := ueuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ueuo *UserEntityUpdateOne) sqlSave(ctx context.Context) (_node *UserEntity, err error) {
	_spec := sqlgraph.NewUpdateSpec(userentity.Table, userentity.Columns, sqlgraph.NewFieldSpec(userentity.FieldID, field.TypeInt))
	id, ok := ueuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UserEntity.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ueuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userentity.FieldID)
		for _, f := range fields {
			if !userentity.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != userentity.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ueuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ueuo.mutation.Username(); ok {
		_spec.SetField(userentity.FieldUsername, field.TypeString, value)
	}
	if value, ok := ueuo.mutation.Password(); ok {
		_spec.SetField(userentity.FieldPassword, field.TypeString, value)
	}
	_node = &UserEntity{config: ueuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ueuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userentity.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ueuo.mutation.done = true
	return _node, nil
}
