package ent

//func initCase() (domain.FolderRepository, ent.Folder) {
//	repository := NewMysqlFolderRepository(connect.NewEntSQLiteConnect())
//	folder := ent.Folder{
//		ID:         "di1231",
//		ParentName: "123123",
//		UserID:     "1231231",
//		Name:       "1231231",
//		Deleted:    false,
//		CreateTime: time.Now(),
//		ModTime:    time.Now(),
//	}
//	return repository, folder
//}
//
//func TestNewMysqlFolderRepositoryStore(t *testing.T) {
//	repository, folder := initCase()
//	err := repository.Store(context.Background(), folder)
//	if err != nil {
//		log.Printf("err %v", err)
//	}
//}
//
//func TestNewMysqlFolderRepositoryRemove(t *testing.T) {
//	repository, folder := initCase()
//	err := repository.Remove(context.Background(), folder.ID, folder.UserID)
//	if err != nil {
//		log.Printf("err %v", err)
//	}
//
//}
