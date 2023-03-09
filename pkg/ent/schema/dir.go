package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Dir holds the schema definition for the Folder entity.
type Dir struct {
	ent.Schema
}

// Fields of the Dir.
func (Dir) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("user_id"),
		field.String("sync_id"),
		field.String("dir"),
		//field.String("parent_dir"),
		field.Uint64("level"),
		field.Bool("deleted"),
		field.Time("create_time"),
		field.Time("mod_time"),
	}
}