package controllers

import (
	"fsm/services"
	"github.com/gin-gonic/gin"
	"io"
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
		ErrorResponse(c, 1, "下载文件失败", err)
	}
}

// GetFileMeta 获取文件元数据
func (f *FileController) GetFileMeta(c *gin.Context) {
	UserID := c.GetHeader("userID")
	fileID := c.Param("file-id")

	fileMeta, err := f.fileService.GetFileMeta(c, fileID, UserID)
	if err != nil {
		ErrorResponse(c, 1, "获取文件元数据失败", err)
	}
	SuccessResponse(c, 1, "获取文件元数据成功", fileMeta)
}

// CreateFile 创建文件
func (f *FileController) CreateFile(c *gin.Context) {
	var file services.File
	if err := c.ShouldBindQuery(&file); err != nil {
		ErrorResponse(c, 1, "解析请求数据失败", err)
	}

	ClientID := c.GetHeader("clientID")
	UserID := c.GetHeader("userID")

	fileMeta, err := f.fileService.CreateFile(c, file, UserID, ClientID)
	if err != nil {
		ErrorResponse(c, 1, "文件创建失败", err)
	}
	SuccessResponse(c, 1, "创建文件成功", fileMeta)
}

// DeleteFile 删除文件
func (f *FileController) DeleteFile(c *gin.Context) {

	fileID := c.Param("file-id")
	ClientID := c.GetHeader("clientID")
	UserID := c.GetHeader("userID")

	if err := f.fileService.DeleteFile(c, fileID, UserID, ClientID); err != nil {
		ErrorResponse(c, 1, "删除文件失败", err)
	}
	SuccessResponse(c, 1, "删除文件成功", nil)
}

// UpdateFile 更新文件
func (f *FileController) UpdateFile(c *gin.Context) {
	var file services.File
	if err := c.ShouldBindQuery(&file); err != nil {
		ErrorResponse(c, 1, "解析请求数据失败", err)
	}

	ClientID := c.GetHeader("clientID")
	UserID := c.GetHeader("userID")

	if err := f.fileService.UpdateFile(c, file, UserID, ClientID); err != nil {
		ErrorResponse(c, 1, "更新文件失败", err)
	}
	SuccessResponse(c, 1, "更新文件成功", nil)
}

// RenameFile 重命名文件
func (f *FileController) RenameFile(c *gin.Context) {
	var file services.File
	if err := c.ShouldBindQuery(&file); err != nil {
		ErrorResponse(c, 1, "解析请求数据失败", err)
	}

	ClientID := c.GetHeader("clientID")
	UserID := c.GetHeader("userID")

	if err := f.fileService.RenameFile(c, file, UserID, ClientID); err != nil {
		ErrorResponse(c, 1, "重命名文件失败", err)
	}
	SuccessResponse(c, 1, "重命名文件成功", nil)
}
