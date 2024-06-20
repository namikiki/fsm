package controllers

import (
	"fsm/services"
	"github.com/gin-gonic/gin"
)

type FolderController struct {
	folderService *services.FolderService // 存储服务
}

// NewFolderController 创建一个新的 FolderController 实例
func NewFolderController(folderService *services.FolderService) *FolderController {
	return &FolderController{folderService: folderService}
}

// ListFolder 获取文件夹列表
func (f *FolderController) ListFolder(c *gin.Context) {
	// 从请求头中获取 userID 调用文件夹服务获取文件夹列表
	listFolder, err := f.folderService.ListFolder(c, c.GetHeader("userID"))
	if err != nil {
		ErrorResponse(c, 1, "获取文件夹列表失败", err)
	}
	SuccessResponse(c, 1, "获取成功", listFolder)
}

// CreatFolder 创建文件夹
func (f *FolderController) CreatFolder(c *gin.Context) {
	// 绑定请求参数到 folder 结构体
	var folder services.Folder
	if err := c.ShouldBindQuery(&folder); err != nil {
		ErrorResponse(c, 1, "解析请求数据失败", err)
	}

	// 调用文件夹服务创建文件夹
	folderMeta, err := f.folderService.CreatFolder(c, folder, c.GetHeader("userID"), c.GetHeader("clientID"))
	if err != nil {
		ErrorResponse(c, 1, "创建文件夹失败", err)
	}
	SuccessResponse(c, 1, "创建文件夹成功", folderMeta)
}

func (f *FolderController) DeleteFolder(c *gin.Context) {
	// 从路径参数中获取 folderID
	folderID := c.Param("folder-id")

	//从请求头中获取 userID 和 clientID, 调用文件夹服务删除文件夹
	err := f.folderService.DeleteFolder(c, folderID, c.GetHeader("userID"), c.GetHeader("clientID"))
	if err != nil {
		ErrorResponse(c, 1, "删除文件夹失败", err)
	}
	SuccessResponse(c, 1, "删除文件夹成功", nil)
}

// RenameFolder 重命名文件夹
func (f *FolderController) RenameFolder(c *gin.Context) {
	// 绑定请求参数到 folder 结构体
	var folder services.Folder
	if err := c.ShouldBindQuery(&folder); err != nil {
		ErrorResponse(c, 1, "解析请求数据失败", err)
	}

	ClientID := c.GetHeader("clientID")
	UserID := c.GetHeader("userID")

	if err := f.folderService.RenameFolder(c, folder, UserID, ClientID); err != nil {
		ErrorResponse(c, 1, "重命名文件夹失败", err)
	}
	SuccessResponse(c, 1, "重命名文件夹成功", nil)
}
