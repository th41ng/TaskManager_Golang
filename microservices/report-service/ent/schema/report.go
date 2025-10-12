package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Report lưu thông tin thống kê toàn hệ thống
type Report struct {
	ent.Schema
}

func (Report) Fields() []ent.Field {
	return []ent.Field{
		field.Time("date"),
		field.Int("total_tasks"),
		field.Int("completed_tasks"),
		field.Int("pending_tasks"),
		field.Int("total_users"),
		field.Int("total_projects"),
	}
}
