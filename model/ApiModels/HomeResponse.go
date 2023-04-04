package ApiModels

import "MyCloudDisk/model/DBModels"

type indexData struct {
	User            *DBModels.User         `json:"user"`
	CurrentIndex    string                 `json:"currIndex"`
	UserFileStore   *DBModels.FileStore    `json:"userFileStore"`
	FileCount       int                    `json:"fileCount"`
	FileFolderCount int                    `json:"fileFolderCount"`
	FileDetailUse   map[string]interface{} `json:"fileDetailUse"`
}

type IndexRespBody struct {
	BasicRespBody
	indexData `json:"data"`
}

func (index *IndexRespBody) SetData(data map[string]interface{}) RespBody {
	index.User = data["user"].(*DBModels.User)
	index.CurrentIndex = data["currIndex"].(string)
	index.UserFileStore = data["userFileStore"].(*DBModels.FileStore)
	index.FileCount = data["fileCount"].(int)
	index.FileFolderCount = data["fileFolderCount"].(int)
	index.FileDetailUse = data["fileDetailUse"].(map[string]interface{})
	return index
}

type filesData struct {
	CurrAll       string                 `json:"currAll"`
	User          *DBModels.User         `json:"user"`
	FolderID      int                    `json:"folderID"`
	FolderName    string                 `json:"folder"`
	Files         []DBModels.MyFile      `json:"files"`
	FileFolders   []DBModels.FileFolder  `json:"fileFolder"`
	ParentFolder  *DBModels.FileFolder   `json:"parentFolder"`
	FileDetailUse map[string]interface{} `json:"fileDetailUse"`
	AllParents    []*DBModels.FileFolder `json:"allParents"`
	//currentAllParents
}

type FilesRespBody struct {
	BasicRespBody
	filesData `json:"data"`
}

func (f *FilesRespBody) SetData(data map[string]interface{}) RespBody {
	f.CurrAll = data["currAll"].(string)
	f.User = data["user"].(*DBModels.User)
	f.FolderID = data["folderID"].(int)
	f.FolderName = data["folderName"].(string)
	f.Files = data["files"].([]DBModels.MyFile)
	f.FileFolders = data["folders"].([]DBModels.FileFolder)
	if c, ok := data["parent"]; ok && c != nil {
		f.ParentFolder = c.(*DBModels.FileFolder)
	}

	f.FileDetailUse = data["fileDetailUse"].(map[string]interface{})
	f.AllParents = data["allParents"].([]*DBModels.FileFolder)
	return f
}
