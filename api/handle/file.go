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
	"github.com/google/uuid"
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
	//var fs types.File
	//
	//if err := c.ShouldBindQuery(&fs); err != nil {
	//	log.Println("bind err") //数据类型错误
	//	return
	//}
	//
	//if err := f.V.Struct(fs); err != nil {
	//	log.Printf("val %v", err)
	//	return
	//}

	//userID := c.GetHeader("userid")

	var file ent.File
	if err := c.ShouldBindQuery(&file); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, NewErrorApiResult(501, "解析请求数据失败"))
		return
	}

	clientID := c.Request.Header.Get("client")
	if clientID == "" {
		c.AbortWithStatusJSON(http.StatusOK, NewErrorApiResult(501, "解析请求数据失败"))
		return
	}
	//clientclientID := "client1"
	file.UserID = "xyn233"
	file.ID = uuid.New().String()

	fi, err := f.S.FileCreate(c, file, clientID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, NewErrorApiResult(501, "创建文件失败"))
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, fi)

	//map[string]map[int]map[string]map[string]string
	// 上传文件回调 文件上传客户端id  文件下载地址 文件夹同步ID， 文件id(防止文件重复) 文件名，复文件夹名，文件层级， 哈希值，修改时间  创建时间
	// 文件上传 文件上传客户端id 同步id  文件名  层级  文件修改时间
}

func (f *File) Delete(c *gin.Context) {
	var file ent.File
	clientID := c.GetHeader("clientid")
	err := f.S.FileDelete(c, file, clientID)
	if err != nil {
		return
	}

	//id := c.Param("id")
	//if id == "" {
	//	log.Println("error id not nil")
	//	return
	//}
	//user := c.MustGet("user").(*ent.User)
	//err := f.FS.Remove(c, id, user.ID)
	//if err != nil {
	//	log.Printf("remove %v\n", err)
	//	return
	//}
	//c.JSON(http.StatusOK, gin.H{
	//	"code": 200,
	//	"data": "success",
	//})
}

func (f *File) Update(c *gin.Context) {

	var file ent.File
	clientID := c.GetHeader("clientid")

	if err := f.S.FileUpdate(c, file, clientID); err != nil {
		return
	}
}

func (f *File) GetMetadata(c *gin.Context) {

	fileMeta, err := f.F.GetMetadataByID(c, "user1", "sync1", "123")
	if err != nil {
		c.JSON(400, "fail")
		return
	}
	c.JSON(200, fileMeta)
}

func (f *File) Open(c *gin.Context) {
	fileID := c.Param("fileID")
	if fileID == "" {
		log.Println()
		c.AbortWithStatusJSON(http.StatusOK, NewErrorApiResult(501, "error"))
		return
	}

	obj, err := f.FS.Open(c, "xyn233", fileID)
	if err != nil {
		c.JSON(400, "fail")
	}
	defer obj.Close()

	if _, err = io.Copy(c.Writer, obj); err != nil {
		c.JSON(400, "fail")
	}

	//c.Writer.Header().Set("Content-Disposition", "attachment; filename=protoc")
	//file, err := os.Open("123.pdf")
	//if err != nil {
	//	c.JSON(400, "fail")
	//	return
	//}

	//if _, err = io.Copy(c.Writer, file); err != nil {
	//	c.JSON(400, "fail")
	//}

	//id := c.Param("id")
	//if id == "" {
	//	log.Println("error id not nil")
	//	return
	//}
	//
	//user := c.MustGet("user").(*ent.User)
	//
	//file, err := f.FS.Get(c, id, user.ID)
	//if err != nil {
	//	log.Printf("get %v\n", err)
	//	return
	//}
	//
	//c.JSON(http.StatusOK, gin.H{
	//	"code": 200,
	//	"data": file.ID,
	//})
}

func (f *File) GetAllFileBySyncID(c *gin.Context) {
	syncID := c.Param("syncID")
	if syncID == "" {
		c.AbortWithStatusJSON(http.StatusOK, NewErrorApiResult(501, "解析请求数据失败"))
		return
	}

	if files, err := f.F.GetAllBySyncID(c, "xyn233", syncID); err == nil {
		c.AbortWithStatusJSON(http.StatusOK, files)
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, NewErrorApiResult(501, "获取所有文件信息失败"))
}
