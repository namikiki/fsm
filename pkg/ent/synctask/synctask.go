// Code generated by ent, DO NOT EDIT.

package synctask

const (
	// Label holds the string label denoting the synctask type in the database.
	Label = "sync_task"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldRootDir holds the string denoting the root_dir field in the database.
	FieldRootDir = "root_dir"
	// FieldIgnore holds the string denoting the ignore field in the database.
	FieldIgnore = "ignore"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// Table holds the table name of the synctask in the database.
	Table = "sync_tasks"
)

// Columns holds all SQL columns for synctask fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldType,
	FieldName,
	FieldRootDir,
	FieldIgnore,
	FieldCreateTime,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
