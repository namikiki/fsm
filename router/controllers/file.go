package controllers

import (
	"fsm/services"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

// FileController 控制器结构体，包含一个文件服务实例
type FileController struct {
	fileService *services.FileService // 存储服务
}

// NewFileController 创建一个新的 FileController 实例
func NewFileController(fileService *services.FileService) *FileController {
	return &FileController{fileService: fileService}
}

// GetFileContent 获取文件内容
func (f *FileController) GetFileContent(c *gin.Context) {
	userID := c.GetHeader("userID")
	fileID := c.Param("file-id")

	file, err := f.fileService.GetFileContent(c, fileID, userID)
	if err != nil {
		return
	}

	if _, err = io.Copy(c.Writer, file); err != nil {
		c.JSON(http.StatusOK, NewErrorApiResult(501, "error"))
	}
}

// GetFileMeta 获取文件元数据
func (f *FileController) GetFileMeta(c *gin.Context) {
	UserID := c.GetHeader("userID")
	fileID := c.Param("file-id")

	fileMeta, err := f.fileService.GetFileMeta(c, fileID, UserID)
	if err != nil {
		c.JSON(http.StatusOK, NewErrorApiResult(501, "error"))
		return
	}
	c.JSON(http.StatusOK, NewApiResult(501, "", fileMeta))
}

// CreateFile 创建文件
func (f *FileController) CreateFile(c *gin.Context) {
	var file services.File
	if err := c.ShouldBindQuery(&file); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, NewErrorApiResult(501, "解析请求数据失败"))
		return
	}

	ClientID := c.GetHeader("clientID")
	UserID := c.GetHeader("userID")

	fileMeta, err := f.fileService.CreateFile(c, file, UserID, ClientID)
	if err != nil {
		c.JSON(http.StatusOK, NewErrorApiResult(210, "文件创建失败"))
		return
	}

	c.JSON(http.StatusOK, NewApiResult(201, "创建文件成功", fileMeta))
}

// DeleteFile 删除文件
func (f *FileController) DeleteFile(c *gin.Context) {

	fileID := c.Param("file-id")
	ClientID := c.GetHeader("clientID")
	UserID := c.GetHeader("userID")

	if err := f.fileService.DeleteFile(c, fileID, UserID, ClientID); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, NewErrorApiResult(501, "删除文件失败"))
		return
	}

	c.JSON(http.StatusOK, NewApiResult(201, "删除文件成功", nil))
}

// UpdateFile 更新文件
func (f *FileController) UpdateFile(c *gin.Context) {
	var file services.File
	if err := c.ShouldBindQuery(&file); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, NewErrorApiResult(501, "解析请求数据失败"))
		return
	}

	ClientID := c.GetHeader("clientID")
	UserID := c.GetHeader("userID")

	if err := f.fileService.UpdateFile(c, file, UserID, ClientID); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, NewErrorApiResult(501, "删除文件失败"))
		return
	}

}

// RenameFile 重命名文件
func (f *FileController) RenameFile(c *gin.Context) {
	var file services.File
	if err := c.ShouldBindQuery(&file); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, NewErrorApiResult(501, "解析请求数据失败"))
		return
	}

	ClientID := c.GetHeader("clientID")
	UserID := c.GetHeader("userID")

	if err := f.fileService.RenameFile(c, file, UserID, ClientID); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, NewErrorApiResult(501, "删除文件失败"))
		return
	}

}
