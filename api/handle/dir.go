package handle

import (
	"net/http"
	"time"

	"fsm/pkg/domain"
	"fsm/pkg/ent"
	"fsm/pkg/sync"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

	clientID := c.GetHeader("clientID")
	dir.UserID = c.GetHeader("userID")

	if err := d.S.DirCreate(c, &dir, clientID); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, NewErrorApiResult(501, "创建目录失败"))
		return
	}

	c.JSON(http.StatusOK, NewApiJsonResult(201, "创建文件夹成功", dir))
}

func (d *Dir) Delete(c *gin.Context) {

	var dir ent.Dir
	if err := c.ShouldBindJSON(&dir); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, NewErrorApiResult(501, "解析请求数据失败"))
		return
	}

	clientID := c.GetHeader("clientID")
	dir.UserID = c.GetHeader("userID")

	if err := d.S.DirDelete(c, dir, clientID); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, NewErrorApiResult(501, "删除目录失败"))
		return
	}
	c.JSON(http.StatusOK, NewApiJsonResult(201, "删除文件夹成功", nil))
}

// todo
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

// GetAllDirByPath todo
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
	userID := c.GetHeader("userID")
	syncID := c.Param("syncID")
	if syncID == "" {
		c.AbortWithStatusJSON(http.StatusOK, NewErrorApiResult(501, "解析请求数据失败"))
		return
	}

	if dirs, err := d.D.WalkDirBySyncID(c, userID, syncID); err == nil {
		c.JSON(http.StatusOK, NewApiJsonResult(201, "获取所有文件夹信息成功", dirs))
		return
	}
	c.JSON(http.StatusOK, NewErrorApiResult(501, "获取syncID下所有文件夹信息失败"))
}
