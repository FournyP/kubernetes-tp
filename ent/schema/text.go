package schema

import "entgo.io/ent"

// Text holds the schema definition for the Text entity.
type Text struct {
	ent.Schema
}

// Fields of the Text.
func (Text) Fields() []ent.Field {
	return nil
}

// Edges of the Text.
func (Text) Edges() []ent.Edge {
	return nil
}
