// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// UserEntitiesColumns holds the columns for the "user_entities" table.
	UserEntitiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "username", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
	}
	// UserEntitiesTable holds the schema information for the "user_entities" table.
	UserEntitiesTable = &schema.Table{
		Name:       "user_entities",
		Columns:    UserEntitiesColumns,
		PrimaryKey: []*schema.Column{UserEntitiesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "userentity_username",
				Unique:  true,
				Columns: []*schema.Column{UserEntitiesColumns[1]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		UserEntitiesTable,
	}
)

func init() {
}