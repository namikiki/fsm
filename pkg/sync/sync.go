package sync

import (
	"encoding/json"
	"time"

	"fsm/pkg/domain"
	"fsm/pkg/ent"
	"fsm/pkg/types"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/minio/minio-go/v7"
)

type Syncer struct {
	DR                   domain.DirRepository
	Redis                *redis.Client
	WebsocketConnChannel chan types.SyncClient
	FR                   domain.FileRepository
	Min                  *minio.Client
	ST                   domain.SyncTask
}

func NewSyncer(dr domain.DirRepository, rc *redis.Client, fr domain.FileRepository, min *minio.Client, st domain.SyncTask) *Syncer {
	return &Syncer{
		ST:                   st,
		DR:                   dr,
		Redis:                rc,
		FR:                   fr,
		Min:                  min,
		WebsocketConnChannel: make(chan types.SyncClient),
	}
}

func (s *Syncer) FileCreate(c *gin.Context, file ent.File, ClientID string) (*ent.File, error) {

	object, err := s.Min.PutObject(c, file.UserID, file.ID, c.Request.Body, c.Request.ContentLength, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		return nil, err
	}
	file.Hash = object.ETag

	if err := s.FR.Create(c, file); err != nil {
		return nil, err
	}

	marshal, err := json.Marshal(types.PubSubMessage{
		Type:     "file",
		Action:   "create",
		ClientID: ClientID,
		Data:     file,
	})
	s.Redis.Publish(c, file.UserID, marshal)
	return &file, err
}

func (s *Syncer) FileDelete(c *gin.Context, f ent.File, ClientID string) error {

	if err := s.FR.Delete(c, f); err != nil {
		return err
	}

	if err := s.Min.RemoveObject(c, f.UserID, f.ID, minio.RemoveObjectOptions{}); err != nil {
		return err
	}

	marshal, err := json.Marshal(types.PubSubMessage{
		Type:     "file",
		Action:   "delete",
		ClientID: ClientID,
		Data:     f,
	})
	s.Redis.Publish(c, f.UserID, marshal)
	return err

}

func (s *Syncer) FileUpdate(c *gin.Context, file ent.File, ClientID string) error {

	object, err := s.Min.PutObject(c, file.UserID, file.ID, c.Request.Body, c.Request.ContentLength, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		return err
	}

	f := ent.File{
		ID:          file.ID,
		UserID:      file.UserID,
		SyncID:      file.SyncID,
		Name:        file.Name,
		ParentDirID: file.ParentDirID,
		Level:       file.Level,
		Hash:        object.ETag,
		Size:        object.Size,
		Deleted:     false,
		CreateTime:  time.Now(),
		ModTime:     time.Now(),
	}
	if err := s.FR.Update(c, f); err != nil {
		return err
	}

	marshal, err := json.Marshal(types.PubSubMessage{
		Type:     "file",
		Action:   "update",
		ClientID: ClientID,
		Data:     f,
	})
	s.Redis.Publish(c, file.UserID, marshal)
	return err
}

func (s *Syncer) DirCreate(c *gin.Context, dir ent.Dir, ClientID string) error {

	if err := s.DR.Create(c, dir); err != nil {
		return err
	}

	marshal, err := json.Marshal(types.PubSubMessage{
		Type:     "dir",
		Action:   "create",
		ClientID: ClientID,
		Data:     dir,
	})

	s.Redis.Publish(c, dir.UserID, marshal)
	return err

}

func (s *Syncer) DirDelete(c *gin.Context, dir ent.Dir, ClientID string) error {
	if err := s.DR.Delete(c, dir); err != nil {
		return err
	}

	marshal, err := json.Marshal(types.PubSubMessage{
		Type:     "dir",
		Action:   "delete",
		ClientID: ClientID,
		Data:     dir,
	})
	s.Redis.Publish(c, dir.UserID, marshal)
	return err
}

func (s *Syncer) SyncTaskCreate(c *gin.Context, st ent.SyncTask, ClientID string) error {
	if err := s.ST.Create(st); err != nil {
		return err
	}

	marshal, err := json.Marshal(types.PubSubMessage{
		Type:     "syncTask",
		Action:   "create",
		ClientID: ClientID,
		Data:     st,
	})
	s.Redis.Publish(c, st.UserID, marshal)
	return err

}

func (s *Syncer) SyncTaskDelete(c *gin.Context, userID string, syncID string, ClientID string) error {
	if err := s.ST.Delete(userID, syncID); err != nil {
		return err
	}
	marshal, err := json.Marshal(types.PubSubMessage{
		Type:     "syncTask",
		Action:   "delete",
		ClientID: ClientID,
		Data:     syncID,
	})
	s.Redis.Publish(c, userID, marshal)
	return err

}
