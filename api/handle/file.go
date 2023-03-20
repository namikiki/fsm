package handle

import (
	"io"
	"log"
	"net/http"

	"fsm/pkg/domain"
	"fsm/pkg/ent"
	"fsm/pkg/sync"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type File struct {
	F  domain.FileRepository
	V  *validator.Validate
	FS domain.FileStorageService
	S  *sync.Syncer
}

func NewFile(f domain.FileRepository, s *sync.Syncer, fs domain.FileStorageService, v *validator.Validate) File {
	return File{
		V:  v,
		F:  f,
		FS: fs,
		S:  s,
	}
}

func (f *File) Create(c *gin.Context) {

	var file ent.File
	if err := c.ShouldBindQuery(&file); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, NewErrorApiResult(501, "解析请求数据失败"))
		return
	}

	clientID := c.GetHeader("clientID")
	file.UserID = c.GetHeader("userID")

	if err := f.S.FileCreate(c, &file, clientID); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, NewErrorApiResult(501, "创建文件失败"+err.Error()))
		return
	}

	c.JSON(http.StatusOK, NewApiJsonResult(201, "创建文件成功", file))

	//map[string]map[int]map[string]map[string]string
	// 上传文件回调 文件上传客户端id  文件下载地址 文件夹同步ID， 文件id(防止文件重复) 文件名，复文件夹名，文件层级， 哈希值，修改时间  创建时间
	// 文件上传 文件上传客户端id 同步id  文件名  层级  文件修改时间
}

func (f *File) Delete(c *gin.Context) {

	var file ent.File
	if err := c.ShouldBindJSON(&file); err != nil || file.ID == "" {
		c.AbortWithStatusJSON(http.StatusOK, NewErrorApiResult(501, "解析请求数据失败"))
		return
	}

	clientID := c.GetHeader("clientID")
	file.UserID = c.GetHeader("userID")

	if err := f.S.FileDelete(c, file, clientID); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, NewErrorApiResult(501, "删除文件失败"))
		return
	}

	c.JSON(http.StatusOK, NewApiResult(201, "删除文件成功", nil))
}

func (f *File) Update(c *gin.Context) {

	var file ent.File
	if err := c.ShouldBindQuery(&file); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, NewErrorApiResult(501, "解析请求数据失败"))
		return
	}

	clientID := c.GetHeader("clientID")
	file.UserID = c.GetHeader("userID")
	log.Println("client=", clientID)

	if err := f.S.FileUpdate(c, file, clientID); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, NewErrorApiResult(501, "更新文件失败"))
		return
	}

	c.JSON(http.StatusOK, NewApiJsonResult(201, "更新文件成功", file))
}

// todo
func (f *File) GetMetadata(c *gin.Context) {

	fileMeta, err := f.F.GetMetadataByID(c, "user1", "123")
	if err != nil {
		c.JSON(400, "fail")
		return
	}
	c.JSON(200, fileMeta)
}

// todo
func (f *File) Open(c *gin.Context) {

	userID := c.GetHeader("userID")
	var fileID string
	if fileID = c.Param("fileID"); fileID == "" {
		c.AbortWithStatusJSON(http.StatusOK, NewErrorApiResult(501, "error"))
		return
	}

	file, err := f.F.GetMetadataByID(c, userID, fileID)
	if file.ID == "" {
		c.AbortWithStatusJSON(http.StatusOK, NewErrorApiResult(501, "此文件不存在"))
		return
	}

	object, err := f.FS.Open(c, userID, fileID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, NewErrorApiResult(501, "error"))
		return
	}
	defer object.Close()

	if _, err = io.Copy(c.Writer, object); err != nil {
		c.JSON(http.StatusOK, NewErrorApiResult(501, "error"))
	}
}

func (f *File) GetAllFileBySyncID(c *gin.Context) {
	syncID := c.Param("syncID")
	userID := c.GetHeader("userID")
	if syncID == "" {
		c.AbortWithStatusJSON(http.StatusOK, NewErrorApiResult(501, "解析请求数据失败"))
		return
	}

	log.Println(userID, syncID)

	files, err := f.F.GetAllBySyncID(c, userID, syncID)
	if err != nil {
		c.JSON(http.StatusOK, NewErrorApiResult(501, "获取所有文件信息失败"))
		return
	}
	c.JSON(http.StatusOK, NewApiJsonResult(201, "获取所有文件信息成功", files))
}

func (f *File) Rename(c *gin.Context) {

	var file ent.File
	if err := c.ShouldBindJSON(&file); err != nil || file.ID == "" {
		c.AbortWithStatusJSON(http.StatusOK, NewErrorApiResult(501, "解析请求数据失败"))
		return
	}

	clientID := c.GetHeader("clientID")
	file.UserID = c.GetHeader("userID")

	log.Println(file)

	if err := f.S.FileRename(c, file, clientID); err != nil {
		c.JSON(http.StatusOK, NewErrorApiResult(501, "更新文件名失败"))
		return
	}

	c.JSON(http.StatusOK, NewApiJsonResult(201, "更新文件名成功", nil))
}
