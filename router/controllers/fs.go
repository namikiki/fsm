package controllers

import (
	"fsm/services"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type FSController struct {
	storageService *services.StorageService
}

func NewFSController(storageService *services.StorageService) *FSController {
	return &FSController{storageService: storageService}
}

func (fs *FSController) GetFileContent(c *gin.Context) {
	UserID := c.GetHeader("userID")
	fileID := c.Param("file-id")

	file, err := fs.storageService.GetFileContent(c, fileID, UserID)
	if err != nil {
		return
	}

	if _, err = io.Copy(c.Writer, file); err != nil {
		c.JSON(http.StatusOK, NewErrorApiResult(501, "error"))
	}
}

func (fs *FSController) GetFileMeta(c *gin.Context) {
	UserID := c.GetHeader("userID")
	fileID := c.Param("file-id")

	fileMeta, err := fs.storageService.GetFileMeta(c, fileID, UserID)
	if err != nil {
		c.JSON(http.StatusOK, NewErrorApiResult(501, "error"))
		return
	}
	c.JSON(http.StatusOK, NewApiResult(501, "", fileMeta))
}

func (fs *FSController) CreateFile(c *gin.Context) {
	var file services.File
	if err := c.ShouldBindQuery(&file); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, NewErrorApiResult(501, "解析请求数据失败"))
		return
	}

	ClientID := c.GetHeader("clientID")
	UserID := c.GetHeader("userID")

	fileMeta, err := fs.storageService.CreateFile(c, file, UserID, ClientID)
	if err != nil {
		c.JSON(http.StatusOK, NewErrorApiResult(210, "文件创建失败"))
		return
	}

	c.JSON(http.StatusOK, NewApiResult(201, "创建文件成功", fileMeta))
}

func (fs *FSController) DeleteFile(c *gin.Context) {

	fileID := c.Param("file-id")
	ClientID := c.GetHeader("clientID")
	UserID := c.GetHeader("userID")

	if err := fs.storageService.DeleteFile(c, fileID, UserID, ClientID); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, NewErrorApiResult(501, "删除文件失败"))
		return
	}

	c.JSON(http.StatusOK, NewApiResult(201, "删除文件成功", nil))
}

func (fs *FSController) UpdateFile(c *gin.Context) {
	var file services.File
	if err := c.ShouldBindQuery(&file); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, NewErrorApiResult(501, "解析请求数据失败"))
		return
	}

	ClientID := c.GetHeader("clientID")
	UserID := c.GetHeader("userID")

	if err := fs.storageService.UpdateFile(c, file, UserID, ClientID); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, NewErrorApiResult(501, "删除文件失败"))
		return
	}

}

func (fs *FSController) RenameFile(c *gin.Context) {
	var file services.File
	if err := c.ShouldBindQuery(&file); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, NewErrorApiResult(501, "解析请求数据失败"))
		return
	}

	ClientID := c.GetHeader("clientID")
	UserID := c.GetHeader("userID")

	if err := fs.storageService.RenameFile(c, file, UserID, ClientID); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, NewErrorApiResult(501, "删除文件失败"))
		return
	}

}

//folder

func (fs *FSController) ListFolder(c *gin.Context) {
	listFolder, err := fs.storageService.ListFolder(c, c.GetHeader("userID"))
	if err != nil {
		c.JSON(http.StatusOK, NewErrorApiResult(222, "获取文件夹列表失败"))
	}
	c.JSON(http.StatusOK, NewApiResult(200, "获取成功", listFolder))
}

func (fs *FSController) CreatFolder(c *gin.Context) {
	var folder services.Folder
	if err := c.ShouldBindQuery(&folder); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, NewErrorApiResult(501, "解析请求数据失败"))
		return
	}

	folderMeta, err := fs.storageService.CreatFolder(c, folder, c.GetHeader("userID"), c.GetHeader("clientID"))
	if err != nil {
		c.JSON(http.StatusOK, NewErrorApiResult(210, "创建文件夹失败"))
		return
	}
	c.JSON(http.StatusOK, NewApiResult(123, "创建文件夹成功", folderMeta))
}

func (fs *FSController) DeleteFolder(c *gin.Context) {
	folderID := c.Param("folder-id")

	err := fs.storageService.DeleteFolder(c, folderID, c.GetHeader("userID"), c.GetHeader("clientID"))
	if err != nil {
		c.JSON(http.StatusOK, NewErrorApiResult(210, "删除文件夹失败"))
		return
	}

	c.JSON(http.StatusOK, NewApiResult(20, "删除文件夹成功", nil))
}

func (fs *FSController) RenameFolder(c *gin.Context) {
	var folder services.Folder
	if err := c.ShouldBindQuery(&folder); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, NewErrorApiResult(501, "解析请求数据失败"))
		return
	}

	ClientID := c.GetHeader("clientID")
	UserID := c.GetHeader("userID")

	if err := fs.storageService.RenameFolder(c, folder, UserID, ClientID); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, NewErrorApiResult(501, "删除文件失败"))
		return
	}
	c.JSON(http.StatusOK, NewApiResult(123, "重命名成功", nil))
}
