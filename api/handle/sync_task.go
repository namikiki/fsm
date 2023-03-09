package handle

import (
	"log"
	"net/http"

	"fsm/pkg/domain"
	"fsm/pkg/ent"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SyncTask struct {
	ST domain.SyncTask
}

func NewSyncTask(st domain.SyncTask) SyncTask {
	return SyncTask{ST: st}
}

func (i *SyncTask) Create(c *gin.Context) {

	var st ent.SyncTask
	st.ID = uuid.New().String()

	if err := c.ShouldBindJSON(&st); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, NewErrorApiResult(501, "解析请求数据失败"))
		return
	}

	log.Println(st)

	if err := i.ST.Create(st); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, NewErrorApiResult(501, "创建同步任务失败"))
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, st)
}

func (i *SyncTask) Delete(c *gin.Context) {

	if err := i.ST.Delete("xyn233", "sync1"); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, NewErrorApiResult(501, "删除同步人物失败"))
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, nil)
}

func (i *SyncTask) Get(c *gin.Context) {
	var syncTask ent.SyncTask
	if syncTask = i.ST.Get("xyn233", "sync1"); syncTask.ID == "" {
		c.AbortWithStatusJSON(http.StatusOK, NewErrorApiResult(501, "获取同步任务失败"))
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, syncTask)
}

func (i *SyncTask) GetAll(c *gin.Context) {
	if syncTasks, err := i.ST.GetAll("xyn233"); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, syncTasks)
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, NewErrorApiResult(501, "获取失败"))
}
