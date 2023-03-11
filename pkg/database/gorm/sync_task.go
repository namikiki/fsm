package gorm

import (
	"fsm/pkg/domain"
	"fsm/pkg/ent"

	"gorm.io/gorm"
)

type SyncTaskRepository struct {
	Conn *gorm.DB
}

func NewSyncRepository(conn *gorm.DB) domain.SyncTaskRepository {
	return &SyncTaskRepository{Conn: conn}
}

func (s *SyncTaskRepository) Create(sync ent.SyncTask) error {
	s.Conn.Create(&sync)
	return nil
}

func (s *SyncTaskRepository) Delete(userID, syncID string) error {
	s.Conn.Where("id = ? and user_id =?", syncID, userID).Delete(&ent.SyncTask{})
	return nil
}

func (s *SyncTaskRepository) Get(userID, syncID string) ent.SyncTask {
	var sync ent.SyncTask
	s.Conn.Where("id = ? and user_id =?", syncID, userID).Find(&sync)
	return sync
}

func (s *SyncTaskRepository) GetAll(userID string) ([]ent.SyncTask, error) {
	var syncs []ent.SyncTask
	s.Conn.Where(" user_id = ?", userID).Find(&syncs)
	return syncs, nil
}
