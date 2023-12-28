// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"grpcserver/data/ent/predicate"
	"grpcserver/data/ent/userentity"
	"sync"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeUserEntity = "UserEntity"
)

// UserEntityMutation represents an operation that mutates the UserEntity nodes in the graph.
type UserEntityMutation struct {
	config
	op            Op
	typ           string
	id            *int
	username      *string
	password      *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*UserEntity, error)
	predicates    []predicate.UserEntity
}

var _ ent.Mutation = (*UserEntityMutation)(nil)

// userentityOption allows management of the mutation configuration using functional options.
type userentityOption func(*UserEntityMutation)

// newUserEntityMutation creates new mutation for the UserEntity entity.
func newUserEntityMutation(c config, op Op, opts ...userentityOption) *UserEntityMutation {
	m := &UserEntityMutation{
		config:        c,
		op:            op,
		typ:           TypeUserEntity,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserEntityID sets the ID field of the mutation.
func withUserEntityID(id int) userentityOption {
	return func(m *UserEntityMutation) {
		var (
			err   error
			once  sync.Once
			value *UserEntity
		)
		m.oldValue = func(ctx context.Context) (*UserEntity, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().UserEntity.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUserEntity sets the old UserEntity of the mutation.
func withUserEntity(node *UserEntity) userentityOption {
	return func(m *UserEntityMutation) {
		m.oldValue = func(context.Context) (*UserEntity, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserEntityMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserEntityMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserEntityMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *UserEntityMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().UserEntity.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetUsername sets the "username" field.
func (m *UserEntityMutation) SetUsername(s string) {
	m.username = &s
}

// Username returns the value of the "username" field in the mutation.
func (m *UserEntityMutation) Username() (r string, exists bool) {
	v := m.username
	if v == nil {
		return
	}
	return *v, true
}

// OldUsername returns the old "username" field's value of the UserEntity entity.
// If the UserEntity object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserEntityMutation) OldUsername(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUsername is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUsername requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUsername: %w", err)
	}
	return oldValue.Username, nil
}

// ResetUsername resets all changes to the "username" field.
func (m *UserEntityMutation) ResetUsername() {
	m.username = nil
}

// SetPassword sets the "password" field.
func (m *UserEntityMutation) SetPassword(s string) {
	m.password = &s
}

// Password returns the value of the "password" field in the mutation.
func (m *UserEntityMutation) Password() (r string, exists bool) {
	v := m.password
	if v == nil {
		return
	}
	return *v, true
}

// OldPassword returns the old "password" field's value of the UserEntity entity.
// If the UserEntity object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserEntityMutation) OldPassword(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPassword is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPassword requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPassword: %w", err)
	}
	return oldValue.Password, nil
}

// ResetPassword resets all changes to the "password" field.
func (m *UserEntityMutation) ResetPassword() {
	m.password = nil
}

// Where appends a list predicates to the UserEntityMutation builder.
func (m *UserEntityMutation) Where(ps ...predicate.UserEntity) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the UserEntityMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *UserEntityMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.UserEntity, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *UserEntityMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *UserEntityMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (UserEntity).
func (m *UserEntityMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserEntityMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.username != nil {
		fields = append(fields, userentity.FieldUsername)
	}
	if m.password != nil {
		fields = append(fields, userentity.FieldPassword)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserEntityMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case userentity.FieldUsername:
		return m.Username()
	case userentity.FieldPassword:
		return m.Password()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserEntityMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case userentity.FieldUsername:
		return m.OldUsername(ctx)
	case userentity.FieldPassword:
		return m.OldPassword(ctx)
	}
	return nil, fmt.Errorf("unknown UserEntity field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserEntityMutation) SetField(name string, value ent.Value) error {
	switch name {
	case userentity.FieldUsername:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUsername(v)
		return nil
	case userentity.FieldPassword:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPassword(v)
		return nil
	}
	return fmt.Errorf("unknown UserEntity field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserEntityMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserEntityMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserEntityMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown UserEntity numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserEntityMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserEntityMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserEntityMutation) ClearField(name string) error {
	return fmt.Errorf("unknown UserEntity nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserEntityMutation) ResetField(name string) error {
	switch name {
	case userentity.FieldUsername:
		m.ResetUsername()
		return nil
	case userentity.FieldPassword:
		m.ResetPassword()
		return nil
	}
	return fmt.Errorf("unknown UserEntity field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserEntityMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserEntityMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserEntityMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserEntityMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserEntityMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserEntityMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserEntityMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown UserEntity unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserEntityMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown UserEntity edge %s", name)
}