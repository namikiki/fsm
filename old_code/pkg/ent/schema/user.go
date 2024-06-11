package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		//field.UUID("id", uuid.UUID{}),
		field.String("id"),
		field.String("email").MinLen(10).Unique(),
		field.String("pass_word").MinLen(10),
		field.Bytes("salt"),
		field.String("user_name").MinLen(6),
		field.String("bucket_name"),
		field.Int64("current_store_cap"),
		field.Int64("max_store_cap"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
