package handle

import (
	"net/http"

	"fsm/pkg/domain"
	"fsm/pkg/ent"
	"fsm/pkg/sync"

	"github.com/gin-gonic/gin"
)

type SyncTask struct {
	ST domain.SyncTaskRepository
	S  *sync.Syncer
}

func NewSyncTask(st domain.SyncTaskRepository, s *sync.Syncer) SyncTask {
	return SyncTask{
		ST: st,
		S:  s,
	}
}

func (i *SyncTask) Create(c *gin.Context) {

	var syncTask ent.SyncTask
	if err := c.ShouldBindJSON(&syncTask); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, NewErrorApiResult(501, "解析请求数据失败"))
		return
	}

	clientID := c.GetHeader("clientID")
	syncTask.UserID = c.GetHeader("userID")

	if err := i.S.SyncTaskCreate(c, &syncTask, clientID); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, NewErrorApiResult(501, "创建同步任务失败"))
		return
	}
	c.JSON(http.StatusOK, NewApiJsonResult(201, "创建同步任务成功", syncTask))
}

// Delete todo 文件和文件夹删除
func (i *SyncTask) Delete(c *gin.Context) {

	clientID := c.GetHeader("clientID")
	userID := c.GetHeader("userID")
	syncID := c.Param("syncID")
	if syncID == "" {
		c.AbortWithStatusJSON(http.StatusOK, NewErrorApiResult(501, "解析请求数据失败"))
		return
	}

	if err := i.S.SyncTaskDelete(c, userID, syncID, clientID); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, NewErrorApiResult(501, "删除同步人物失败"))
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, NewApiResult(201, "删除同步任务成功", nil))
}

func (i *SyncTask) Get(c *gin.Context) {

	userID := c.GetHeader("userID")
	syncID := c.Param("syncID")
	if syncID == "" {
		c.AbortWithStatusJSON(http.StatusOK, NewErrorApiResult(501, "解析请求数据失败"))
		return
	}

	var syncTask ent.SyncTask
	if syncTask = i.ST.Get(userID, syncID); syncTask.ID == "" {
		c.AbortWithStatusJSON(http.StatusOK, NewErrorApiResult(501, "获取同步任务失败"))
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, syncTask)
}

func (i *SyncTask) GetAllSyncTask(c *gin.Context) {

	if syncTasks, err := i.ST.GetAll(c.GetHeader("userID")); err == nil {
		c.JSON(http.StatusOK, NewApiJsonResult(201, "获取所有的同步任务成功", syncTasks))
		return
	}
	c.JSON(http.StatusOK, NewErrorApiResult(501, "获取失败"))
}
