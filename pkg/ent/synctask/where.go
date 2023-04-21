// Code generated by ent, DO NOT EDIT.

package synctask

import (
	"fsm/pkg/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// RootDir applies equality check predicate on the "root_dir" field. It's identical to RootDirEQ.
func RootDir(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRootDir), v))
	})
}

// Ignore applies equality check predicate on the "ignore" field. It's identical to IgnoreEQ.
func Ignore(v bool) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIgnore), v))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v int64) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserID), v))
	})
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...string) predicate.SyncTask {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUserID), v...))
	})
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...string) predicate.SyncTask {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUserID), v...))
	})
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserID), v))
	})
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserID), v))
	})
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserID), v))
	})
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserID), v))
	})
}

// UserIDContains applies the Contains predicate on the "user_id" field.
func UserIDContains(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUserID), v))
	})
}

// UserIDHasPrefix applies the HasPrefix predicate on the "user_id" field.
func UserIDHasPrefix(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUserID), v))
	})
}

// UserIDHasSuffix applies the HasSuffix predicate on the "user_id" field.
func UserIDHasSuffix(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUserID), v))
	})
}

// UserIDEqualFold applies the EqualFold predicate on the "user_id" field.
func UserIDEqualFold(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUserID), v))
	})
}

// UserIDContainsFold applies the ContainsFold predicate on the "user_id" field.
func UserIDContainsFold(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUserID), v))
	})
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldType), v))
	})
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.SyncTask {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldType), v...))
	})
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.SyncTask {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldType), v...))
	})
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldType), v))
	})
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldType), v))
	})
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldType), v))
	})
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldType), v))
	})
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldType), v))
	})
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldType), v))
	})
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldType), v))
	})
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldType), v))
	})
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldType), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.SyncTask {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.SyncTask {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// RootDirEQ applies the EQ predicate on the "root_dir" field.
func RootDirEQ(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRootDir), v))
	})
}

// RootDirNEQ applies the NEQ predicate on the "root_dir" field.
func RootDirNEQ(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRootDir), v))
	})
}

// RootDirIn applies the In predicate on the "root_dir" field.
func RootDirIn(vs ...string) predicate.SyncTask {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldRootDir), v...))
	})
}

// RootDirNotIn applies the NotIn predicate on the "root_dir" field.
func RootDirNotIn(vs ...string) predicate.SyncTask {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldRootDir), v...))
	})
}

// RootDirGT applies the GT predicate on the "root_dir" field.
func RootDirGT(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRootDir), v))
	})
}

// RootDirGTE applies the GTE predicate on the "root_dir" field.
func RootDirGTE(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRootDir), v))
	})
}

// RootDirLT applies the LT predicate on the "root_dir" field.
func RootDirLT(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRootDir), v))
	})
}

// RootDirLTE applies the LTE predicate on the "root_dir" field.
func RootDirLTE(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRootDir), v))
	})
}

// RootDirContains applies the Contains predicate on the "root_dir" field.
func RootDirContains(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRootDir), v))
	})
}

// RootDirHasPrefix applies the HasPrefix predicate on the "root_dir" field.
func RootDirHasPrefix(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRootDir), v))
	})
}

// RootDirHasSuffix applies the HasSuffix predicate on the "root_dir" field.
func RootDirHasSuffix(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRootDir), v))
	})
}

// RootDirEqualFold applies the EqualFold predicate on the "root_dir" field.
func RootDirEqualFold(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRootDir), v))
	})
}

// RootDirContainsFold applies the ContainsFold predicate on the "root_dir" field.
func RootDirContainsFold(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRootDir), v))
	})
}

// IgnoreEQ applies the EQ predicate on the "ignore" field.
func IgnoreEQ(v bool) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIgnore), v))
	})
}

// IgnoreNEQ applies the NEQ predicate on the "ignore" field.
func IgnoreNEQ(v bool) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIgnore), v))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v int64) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v int64) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...int64) predicate.SyncTask {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...int64) predicate.SyncTask {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v int64) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v int64) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v int64) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v int64) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SyncTask) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SyncTask) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.SyncTask) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		p(s.Not())
	})
}
