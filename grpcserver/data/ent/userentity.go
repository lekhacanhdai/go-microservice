// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"grpcserver/data/ent/userentity"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// UserEntity is the model entity for the UserEntity schema.
type UserEntity struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// Password holds the value of the "password" field.
	Password     string `json:"password,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserEntity) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case userentity.FieldID:
			values[i] = new(sql.NullInt64)
		case userentity.FieldUsername, userentity.FieldPassword:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserEntity fields.
func (ue *UserEntity) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case userentity.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ue.ID = int(value.Int64)
		case userentity.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				ue.Username = value.String
			}
		case userentity.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				ue.Password = value.String
			}
		default:
			ue.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the UserEntity.
// This includes values selected through modifiers, order, etc.
func (ue *UserEntity) Value(name string) (ent.Value, error) {
	return ue.selectValues.Get(name)
}

// Update returns a builder for updating this UserEntity.
// Note that you need to call UserEntity.Unwrap() before calling this method if this UserEntity
// was returned from a transaction, and the transaction was committed or rolled back.
func (ue *UserEntity) Update() *UserEntityUpdateOne {
	return NewUserEntityClient(ue.config).UpdateOne(ue)
}

// Unwrap unwraps the UserEntity entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ue *UserEntity) Unwrap() *UserEntity {
	_tx, ok := ue.config.driver.(*txDriver)
	if !ok {
		panic("ent: UserEntity is not a transactional entity")
	}
	ue.config.driver = _tx.drv
	return ue
}

// String implements the fmt.Stringer.
func (ue *UserEntity) String() string {
	var builder strings.Builder
	builder.WriteString("UserEntity(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ue.ID))
	builder.WriteString("username=")
	builder.WriteString(ue.Username)
	builder.WriteString(", ")
	builder.WriteString("password=")
	builder.WriteString(ue.Password)
	builder.WriteByte(')')
	return builder.String()
}

// UserEntities is a parsable slice of UserEntity.
type UserEntities []*UserEntity