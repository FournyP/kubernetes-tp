package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Text holds the schema definition for the Text entity.
type Text struct {
	ent.Schema
}

// Fields of the Text.
func (Text) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Default("unknown"),
		field.String("text"),
	}
}

// Edges of the Text.
func (Text) Edges() []ent.Edge {
	return nil
}
