package ent

import (
	"context"

	"fsm/pkg/ent"
	"fsm/pkg/ent/dir"
)

type mysqlFolderRepository struct {
	Conn *ent.Client
}

func (m *mysqlFolderRepository) Create(ctx context.Context, f ent.Dir) error {
	//TODO implement me
	panic("implement me")
}

func (m *mysqlFolderRepository) Delete(ctx context.Context, f ent.Dir) error {
	//TODO implement me
	panic("implement me")
}

func (m *mysqlFolderRepository) Get() error {
	//TODO implement me
	panic("implement me")
}

func (m *mysqlFolderRepository) Rename() error {
	//TODO implement me
	panic("implement me")
}

func (m *mysqlFolderRepository) GetPathAll(ctx context.Context, syncPath string) {
	//TODO implement me
	panic("implement me")
}

//func NewMysqlFolderRepository(Conn *ent.Client) domain.DirRepository {
//	return &mysqlFolderRepository{Conn: Conn}
//}

func (m *mysqlFolderRepository) Store(ctx context.Context, f ent.Dir) error {
	//_, err := m.Conn.Folder.Create().
	//	SetID(f.ID).
	//	SetName(f.Name).
	//	SetParentID(f.ParentID).
	//	SetUserID(f.UserID).
	//	SetDeleted(false).
	//	SetCreateTime(f.CreateTime).
	//	SetModTime(f.ModTime).Save(ctx)
	//return err
	return nil
}

func (m *mysqlFolderRepository) Remove(ctx context.Context, id, uid string) error {
	_, err := m.Conn.Dir.Delete().Where(dir.ID(id), dir.UserID(uid)).Exec(ctx)
	return err
}
