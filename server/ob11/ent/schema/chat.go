package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Chat holds the schema definition for the Chat entity.
type Chat struct {
	ent.Schema
}

// Fields of the Registry.
func (Chat) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Int64("group_id"),
		field.Int64("user_id"),
		field.Int64("time"),
		field.Text("message_type"),
		field.Text("message_content"),
	}
}

// Edges of the Registry.
func (Chat) Edges() []ent.Edge {
	return nil
}

func (Chat) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Mixin{},
	}
}

func (Chat) Indexes() []ent.Index {
	return []ent.Index{}
}
