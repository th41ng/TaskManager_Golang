package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Task struct {
	ent.Schema
}

func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.Bool("done").Default(false),
		field.Int("priority"),
		field.Int("project_id"),
	}
}
