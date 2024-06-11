package domain

import "fsm/pkg/ent"

type SyncTaskRepository interface {
	Create(sync ent.SyncTask) error
	Delete(userID, syncID string) error
	Get(userID, syncID string) ent.SyncTask
	GetAll(userID string) ([]ent.SyncTask, error)
}
