package sync

import (
	"encoding/json"

	"fsm/pkg/domain"
	"fsm/pkg/ent"
	"fsm/pkg/types"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

type Syncer struct {
	DR                   domain.DirRepository
	Redis                *redis.Client
	WebsocketConnChannel chan types.SyncClient
	FR                   domain.FileRepository
	Min                  *minio.Client
	ST                   domain.SyncTaskRepository
}

func NewSyncer(dr domain.DirRepository, rc *redis.Client, fr domain.FileRepository, min *minio.Client, st domain.SyncTaskRepository) *Syncer {
	return &Syncer{
		ST:                   st,
		DR:                   dr,
		Redis:                rc,
		FR:                   fr,
		Min:                  min,
		WebsocketConnChannel: make(chan types.SyncClient),
	}
}

func (s *Syncer) FileCreate(c *gin.Context, file *ent.File, ClientID string) error {

	file.ID = uuid.New().String()
	object, err := s.Min.PutObject(c, file.UserID, file.ID, c.Request.Body, c.Request.ContentLength,
		minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		return err
	}

	file.Size = object.Size
	file.Hash = object.ETag

	if err := s.FR.Create(c, file); err != nil {
		return err
	}

	fileMas, _ := json.Marshal(file)
	marshal, err := json.Marshal(types.PubSubMessage{
		Type:     "file",
		Action:   "create",
		SyncID:   file.SyncID,
		ClientID: ClientID,
		Data:     fileMas,
	})
	s.Redis.Publish(c, file.UserID, marshal)
	return err
}

func (s *Syncer) FileDelete(c *gin.Context, file ent.File, ClientID string) error {

	if err := s.FR.Delete(c, file); err != nil {
		return err
	}

	if err := s.Min.RemoveObject(c, file.UserID, file.ID, minio.RemoveObjectOptions{}); err != nil {
		return err
	}

	fileMas, _ := json.Marshal(file)
	marshal, err := json.Marshal(types.PubSubMessage{
		Type:     "file",
		Action:   "delete",
		SyncID:   file.SyncID,
		ClientID: ClientID,
		Data:     fileMas,
	})
	s.Redis.Publish(c, file.UserID, marshal)
	return err
}

func (s *Syncer) FileUpdate(c *gin.Context, file ent.File, ClientID string) error {
	object, err := s.Min.PutObject(c, file.UserID, file.ID, c.Request.Body, c.Request.ContentLength,
		minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		return err
	}

	file.Size = object.Size
	file.Hash = object.ETag

	if err := s.FR.Update(c, file); err != nil {
		return err
	}

	fileMas, _ := json.Marshal(file)
	marshal, err := json.Marshal(types.PubSubMessage{
		Type:     "file",
		Action:   "update",
		ClientID: ClientID,
		SyncID:   file.SyncID,
		Data:     fileMas,
	})

	s.Redis.Publish(c, file.UserID, marshal)
	return err
}

func (s *Syncer) DirCreate(c *gin.Context, dir *ent.Dir, ClientID string) error {

	dir.ID = uuid.New().String()
	if err := s.DR.Create(c, *dir); err != nil {
		return err
	}

	dirMas, _ := json.Marshal(dir)
	marshal, err := json.Marshal(types.PubSubMessage{
		Type:     "dir",
		Action:   "create",
		SyncID:   dir.SyncID,
		ClientID: ClientID,
		Data:     dirMas,
	})

	s.Redis.Publish(c, dir.UserID, marshal)
	return err

}

func (s *Syncer) DirDelete(c *gin.Context, dir ent.Dir, ClientID string) error {
	if err := s.DR.Delete(c, dir); err != nil {
		return err
	}

	dirMas, _ := json.Marshal(dir)
	marshal, err := json.Marshal(types.PubSubMessage{
		Type:     "dir",
		Action:   "delete",
		SyncID:   dir.SyncID,
		ClientID: ClientID,
		Data:     dirMas,
	})
	s.Redis.Publish(c, dir.UserID, marshal)
	return err
}

func (s *Syncer) SyncTaskCreate(c *gin.Context, syncTask *ent.SyncTask, ClientID string) error {

	syncTask.ID = uuid.New().String()
	if err := s.ST.Create(*syncTask); err != nil {
		return err
	}

	syncTaskMas, _ := json.Marshal(syncTask)
	marshal, err := json.Marshal(types.PubSubMessage{
		Type:     "syncTask",
		Action:   "create",
		ClientID: ClientID,
		Data:     syncTaskMas,
	})

	s.Redis.Publish(c, syncTask.UserID, marshal)
	return err

}

func (s *Syncer) SyncTaskDelete(c *gin.Context, userID string, syncID string, ClientID string) error {
	if err := s.ST.Delete(userID, syncID); err != nil {
		return err
	}

	var st ent.SyncTask
	st.ID = syncID

	syncTaskMas, _ := json.Marshal(st)
	marshal, err := json.Marshal(types.PubSubMessage{
		Type:     "syncTask",
		Action:   "delete",
		ClientID: ClientID,
		Data:     syncTaskMas,
	})
	s.Redis.Publish(c, userID, marshal)
	return err
}
