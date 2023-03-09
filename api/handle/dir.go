package handle

import (
	"net/http"
	"time"

	"fsm/pkg/domain"
	"fsm/pkg/ent"
	"fsm/pkg/sync"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Dir struct {
	V *validator.Validate
	D domain.DirRepository
	S *sync.Syncer
}

func NewDir(s *sync.Syncer, d domain.DirRepository, v *validator.Validate) Dir {
	return Dir{
		V: v,
		S: s,
		D: d,
	}
}

func (d *Dir) Create(c *gin.Context) {

	var dir ent.Dir
	if err := c.ShouldBindJSON(&dir); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, NewErrorApiResult(501, "解析请求数据失败"))
		return
	}

	dir.ID = uuid.New().String()
	dir.UserID = "xyn233"
	clientID := c.Request.Header.Get("client")
	if clientID == "" {
		c.AbortWithStatusJSON(http.StatusOK, NewErrorApiResult(501, "解析请求数据失败"))
		return
	}

	if err := d.S.DirCreate(c, dir, clientID); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, NewErrorApiResult(501, "创建目录失败"))
		return
	}

	//NewApiResult(201, "成功", dir)
	c.AbortWithStatusJSON(http.StatusOK, dir)
}

func (d *Dir) Delete(c *gin.Context) {
	//var dirs ent.Dir
	//if err := c.Bind(&dirs); err != nil {
	//	c.JSON(400, "fail")
	//}
	clientID := c.Query("client_id")
	dir := ent.Dir{ID: c.Query("id"),
		UserID:     "user1",
		SyncID:     "sync1",
		Dir:        uuid.New().String(),
		Level:      0,
		Deleted:    false,
		CreateTime: time.Now(),
		ModTime:    time.Now(),
	}

	if err := d.S.DirDelete(c, dir, clientID); err != nil {
		c.JSON(400, "fail")
		return
	}
	c.JSON(200, "success")
}

func (d *Dir) ReadDir(c *gin.Context) {
	dir := ent.Dir{ID: "123",
		UserID:     "user1",
		SyncID:     "sync1",
		Dir:        "root",
		Level:      0,
		Deleted:    false,
		CreateTime: time.Now(),
		ModTime:    time.Now(),
	}
	readDir, err := d.D.ReadDir(c, dir)
	if err != nil {
		c.JSON(400, "fail")
		return
	}
	c.JSON(200, readDir)
}

func (d *Dir) GetAllDirByPath(c *gin.Context) {

	dir := ent.Dir{ID: "123",
		UserID:     "user1",
		SyncID:     "sync1",
		Dir:        "root",
		Level:      0,
		Deleted:    false,
		CreateTime: time.Now(),
		ModTime:    time.Now(),
	}
	readDir, err := d.D.WalkDirByPath(c, dir)
	if err != nil {
		c.JSON(400, "fail")
		return
	}
	c.JSON(200, readDir)
}

func (d *Dir) GetAllDirBySyncID(c *gin.Context) {
	if dirs, err := d.D.WalkDirBySyncID(c, "xyn233", "sync1"); err == nil {
		c.AbortWithStatusJSON(http.StatusOK, dirs)
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, NewErrorApiResult(501, "获取syncID下所有文件夹失败"))
}
