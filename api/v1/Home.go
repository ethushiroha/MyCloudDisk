package v1

import (
	"MyCloudDisk/global"
	"MyCloudDisk/model/ApiModels"
	"MyCloudDisk/model/DBModels"
	"MyCloudDisk/service"
	"MyCloudDisk/utils/ApiUtils"
	"MyCloudDisk/utils/CryptoUtils"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"time"
)

func Index(c *gin.Context) {
	if t, exists := c.Get("UserID"); exists {
		userID := int(t.(uint))

		//获取用户信息
		user, _ := service.GetUserFromUserID(userID)
		//获取用户仓库信息
		userFileStore, _ := service.GetFileStoreByUserID(user.UserID)
		//获取用户文件数量
		fileCount, _ := service.GetStoreFileCount(user.FileStoreID)
		//获取用户文件夹数量
		fileFolderCount, _ := service.GetFolderFileCount(user.FileStoreID)
		//获取用户文件使用明细数量
		fileDetailUse := service.GetFileDetailUse(user.FileStoreID)
		data := map[string]interface{}{
			"user":            user,
			"currIndex":       "active",
			"userFileStore":   userFileStore,
			"fileCount":       fileCount,
			"fileFolderCount": fileFolderCount,
			"fileDetailUse":   fileDetailUse,
		}

		resp := &ApiModels.IndexRespBody{}
		resp.SetData(data).
			SetStatus("success").
			SetMessage("查询成功").
			Build()
		ApiUtils.SuccessResponse(c, resp)
	}
}

func Files(c *gin.Context) {
	if t, exists := c.Get("UserID"); exists {
		var body map[string]interface{}
		err := c.ShouldBindJSON(&body)
		if err != nil {
			ApiUtils.ErrorResponse(c, err)
			return
		}
		userID := int(t.(uint))
		var folderID int

		if value, ok := body["folderID"]; ok {
			folderID = int(value.(float64))
			if folderID == -1 {
				folderID, err = service.GetUserRootFolderID(userID)
				if err != nil {
					ApiUtils.ErrorResponse(c, err)
					return
				}
			}
		} else {
			folderID, err = service.GetUserRootFolderID(userID)
			if err != nil {
				ApiUtils.ErrorResponse(c, err)
				return
			}
		}

		//fId := c.DefaultQuery("fId", "0")
		//获取用户信息
		user, _ := service.GetUserFromUserID(userID)

		//获取当前目录所有文件
		files, _ := service.GetStoreFilesInFolder(folderID, user.FileStoreID)
		//获取当前目录所有文件夹
		fileFolders, _ := service.GetStoreFoldersInFolder(folderID, user.FileStoreID)

		//获取父级的文件夹信息
		parentFolder, _ := service.GetParentFolderByFolderID(folderID)

		//获取当前目录所有父级
		//currentAllParent := model.GetCurrentAllParent(parentFolder, make([]model.FileFolder, 0))

		allParentFolder, err := service.GetAllParentsFoldersByFolderID(folderID)
		if err != nil {
			ApiUtils.ErrorResponse(c, err)
			return
		}

		//获取当前目录信息
		currentFolder, _ := service.GetFolderByFolderID(folderID)

		//获取用户文件使用明细数量
		fileDetailUse := service.GetFileDetailUse(user.FileStoreID)

		data := map[string]interface{}{
			"currAll":       "active",
			"user":          user,
			"folderID":      folderID,
			"folderName":    currentFolder.FileFolderName,
			"files":         files,
			"folders":       fileFolders,
			"parentFolder":  parentFolder,
			"allParents":    allParentFolder,
			"fileDetailUse": fileDetailUse,
		}
		resp := &ApiModels.FilesRespBody{}
		resp.SetData(data).
			SetStatus("success").
			SetMessage("成功").
			Build()
		ApiUtils.SuccessResponse(c, resp)
	}
}

func Upload(c *gin.Context) {
	var userID int
	if t, exists := c.Get("UserID"); !exists {
		ApiUtils.UnauthorizedResponse(c)
		return
	} else {
		userID = int(t.(uint))
	}

	parentFolderIDStr := c.Query("folderID")
	if parentFolderIDStr == "" {
		ApiUtils.BadRequestResponse(c, "parentFolderID is required")
		return
	}
	parentFolderID, err := strconv.Atoi(parentFolderIDStr)
	if err != nil {
		ApiUtils.ErrorResponse(c, err)
		return
	}
	if parentFolderID == -1 {
		parentFolderID, _ = service.GetUserRootFolderID(userID)
	}

	if u, err := service.GetUserIDByFolderID(parentFolderID); err != nil {
		ApiUtils.ErrorResponse(c, err)
		return
	} else {
		if u != userID {
			ApiUtils.FailureResponse(c, "该文件夹不属于你")
			return
		}
	}

	storeID, err := service.GetFileStoreIDByFolderID(parentFolderID)
	if err != nil {
		ApiUtils.ErrorResponse(c, err)
		return
	}

	header, err := c.FormFile("file")
	if err != nil {
		ApiUtils.ErrorResponse(c, err)
		return
	}
	dst := global.CloudDiskBasicDir + header.Filename
	if err = c.SaveUploadedFile(header, dst); err != nil {
		ApiUtils.ErrorResponse(c, err)
		return
	}
	tmpFile, err := header.Open()
	if err != nil {
		ApiUtils.ErrorResponse(c, err)
		return
	}
	defer tmpFile.Close()

	file := &DBModels.MyFile{
		FileName:       header.Filename,
		Size:           header.Size,
		FileHash:       CryptoUtils.FileHash(tmpFile),
		Postfix:        CryptoUtils.GetPostfix(header.Filename),
		UpdateTime:     strconv.FormatInt(time.Now().Unix(), 10),
		SizeStr:        "b",
		ParentFolderID: parentFolderID,
		FileStoreID:    storeID,
		Type:           CryptoUtils.FileType(header.Filename),
	}

	err = service.InsertAFile(file)
	if err != nil {
		ApiUtils.ErrorResponse(c, err)
		return
	}
	resp := ApiModels.BasicRespBody{}
	r := resp.SetStatus("success").
		SetMessage("上传成功").
		Build()
	ApiUtils.SuccessResponse(c, r)
}

func AddFolder(c *gin.Context) {
	if t, exists := c.Get("UserID"); exists {
		resp := &ApiModels.BasicRespBody{}
		userID := int(t.(uint))
		user, _ := service.GetUserFromUserID(userID)
		var body map[string]interface{}
		err := c.ShouldBindJSON(&body)
		if err != nil {
			ApiUtils.ErrorResponse(c, err)
			return
		}
		var folderName string
		var parentFolderID int
		if name, ok := body["folderName"]; ok {
			folderName = name.(string)
		} else {
			resp.SetStatus("failed").SetMessage("缺少文件夹名")
			ApiUtils.SuccessResponse(c, resp)
			return
		}
		if parent, ok := body["parentFolderID"]; ok {
			parentFolderID = int(parent.(float64))
			if parentFolderID == -1 {
				parentFolderID, err = service.GetUserRootFolderID(userID)
				if err != nil {
					ApiUtils.ErrorResponse(c, err)
					return
				}
			}
		}

		ok, err := service.CreateFolder(folderName, parentFolderID, user.FileStoreID, userID)
		if err != nil {
			ApiUtils.ErrorResponse(c, err)
			return
		}

		if !ok {
			resp.SetStatus("failed").SetMessage("创建失败").Build()
		} else {
			resp.SetStatus("success").SetMessage("创建成功").Build()
		}
		ApiUtils.SuccessResponse(c, resp)
	}
}

func DownloadFile(c *gin.Context) {
	if t, exists := c.Get("UserID"); exists {
		userID := int(t.(uint))
		//user, _ := service.GetUserFromUserID(userID)
		var body map[string]interface{}
		err := c.ShouldBindJSON(&body)
		if err != nil {
			ApiUtils.ErrorResponse(c, err)
			return
		}
		var fileID int
		if fID, ok := body["fileID"]; ok {
			fileID = int(fID.(float64))
		} else {
			// todo:
			//ApiUtils.ErrorResponse()
			return
		}
		// 验证文件的归属
		u, err := service.GetFileUserID(fileID)
		if err != nil {
			ApiUtils.ErrorResponse(c, err)
			return
		}
		if u != userID {
			//
			return
		}
		file, err := service.GetFileByFileID(fileID)
		if err != nil {
			ApiUtils.ErrorResponse(c, err)
			return
		}

		filePath := global.CloudDiskBasicDir + file.FileName

		f, err := os.Open(filePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer func(f *os.File) {
			err = f.Close()
			if err != nil {
				ApiUtils.ErrorResponse(c, err)
				return
			}
		}(f)

		fileInfo, err := f.Stat()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.Header("Content-Type", "application/octet-stream")
		c.Header("Content-Disposition", "attachment; filename="+file.FileName)
		c.Header("Content-Length", strconv.FormatInt(fileInfo.Size(), 10))
		c.File(filePath)
	}
}
