package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Project struct {
	ent.Schema
}

func (Project) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int("owner_id"),
	}
}
