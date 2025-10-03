package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Project struct {
	ent.Schema
}

func (Project) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

func (Project) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("projects").Unique(), // N project → 1 user
		edge.To("tasks", Task.Type),                            // 1 project → N task
	}
}
