package ent

//type mysqlFileRepository struct {
//	Conn *ent.Client
//}
//
//func NewMysqlFileRepository(conn *ent.Client) domain.FileRepository {
//	return &mysqlFileRepository{Conn: conn}
//}
//
//func (m *mysqlFileRepository) Store(ctx context.Context, f ent.File) error {
//
//	_, err := m.Conn.File.Create().
//		SetID(f.ID).SetUserID(f.UserID).SetSyncID(f.SyncID).
//		SetName(f.Name).SetParentName(f.ParentName).SetHash(f.Hash).SetLevel(f.Level).
//		SetSize(f.Size).
//		SetDeleted(false).
//		SetCreateTime(f.CreateTime).
//		SetModTime(f.ModTime).Save(ctx)
//	return err
//}
//
//func (m *mysqlFileRepository) Remove(ctx context.Context, id, uid string) error {
//	_, err := m.Conn.File.Update().Where(file.ID(id), file.UserID(uid)).SetDeleted(true).Save(ctx)
//	return err
//}
//
//func (m *mysqlFileRepository) Get(ctx context.Context, id, uid string) (*ent.File, error) {
//	return m.Conn.File.Query().Where(file.ID(id), file.UserID(uid)).Only(ctx)
//}
//
//func (m *mysqlFileRepository) Update(ctx context.Context, id, uid string, size int64, mod time.Time) error {
//	_, err := m.Conn.File.Update().Where(file.ID(id), file.UserID(uid)).SetSize(size).SetModTime(mod).Save(ctx)
//	return err
//}
