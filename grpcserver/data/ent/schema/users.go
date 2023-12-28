package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type UserEntity struct {
	ent.Schema
}

func (UserEntity) Fields() []ent.Field {
	return []ent.Field{
		field.String("username"),
		field.String("password"),
	}
}

func (UserEntity) Edges() []ent.Edge {
	return nil
}

func (UserEntity) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("username").Unique(),
	}
}
