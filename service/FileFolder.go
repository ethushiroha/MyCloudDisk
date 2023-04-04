package service

import (
	"MyCloudDisk/global"
	"MyCloudDisk/model/DBModels"
	"time"
)

func GetFolderFileCount(FolderID int) (int, error) {
	result := 0
	sql := "SELECT count(*) FROM FileFolder WHERE FileStoreID = ?"
	err := global.DB.Get(&result, sql, FolderID)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func GetStoreFoldersInFolder(folderID, storeID int) ([]DBModels.FileFolder, error) {
	sql := "SELECT * FROM FileFolder WHERE FileStoreID = ? AND ParentFolderID = ?"
	var t []DBModels.FileFolder
	err := global.DB.Select(&t, sql, storeID, folderID)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func GetParentFolderByFolderID(folderID int) (*DBModels.FileFolder, error) {
	sql := "SELECT ParentFolderID FROM FileFolder WHERE ID = ?"
	var t int
	err := global.DB.Get(&t, sql, folderID)
	if err != nil {
		return nil, err
	}
	sql = "SELECT * FROM FileFolder WHERE ID = ?"
	folder := new(DBModels.FileFolder)
	err = global.DB.Get(folder, sql, t)
	if err != nil {
		return nil, err
	}
	return folder, nil
}

func GetParentFolderIDByFolderID(folderID int) (int, error) {
	sql := "SELECT ParentFolderID FROM FileFolder WHERE ID = ?"
	var t int
	err := global.DB.Get(&t, sql, folderID)
	return t, err
}

func GetAllParentsFoldersNameByFolderID(folderID int) ([]string, error) {
	var result []string
	for folderID != -1 {
		name, err := GetFolderNameByFolderID(folderID)
		if err != nil {
			return nil, err
		}
		result = append(result, name)
		folderID, err = GetParentFolderIDByFolderID(folderID)
		if err != nil {
			return nil, err
		}
	}
	// reverse list
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result, nil
}

func GetAllParentsFoldersByFolderID(folderID int) ([]*DBModels.FileFolder, error) {
	var result []*DBModels.FileFolder
	for folderID != -1 {
		folder, err := GetFolderByFolderID(folderID)
		if err != nil {
			return nil, err
		}
		result = append(result, folder)
		folderID = folder.ParentFolderID
	}
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result, nil
}

func GetFolderNameByFolderID(folderID int) (string, error) {
	sql := "SELECT FileFolderName FROM FileFolder WHERE ID = ?"
	var t string
	err := global.DB.Get(&t, sql, folderID)
	return t, err
}

func GetFolderByFolderID(folderID int) (*DBModels.FileFolder, error) {
	sql := "SELECT * FROM FileFolder WHERE ID = ?"
	t := new(DBModels.FileFolder)
	err := global.DB.Get(t, sql, folderID)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func GetFileStoreIDByFolderID(folderID int) (int, error) {
	sql := "SELECT FileStoreID FROM FileFolder WHERE ID = ?"
	var t int
	err := global.DB.Get(&t, sql, folderID)
	return t, err
}

func CreateFolder(folderName string, parentFolderID, fileStoreID, userID int) (bool, error) {
	// 验证 该仓库是否属于该用户
	u, err := GetUserIDByFileStoreID(fileStoreID)
	if (err != nil) || (u != userID) {
		return false, err
	}

	// 验证 上级文件夹是否属于该仓库
	s, err := GetFileStoreIDByFolderID(parentFolderID)
	if (err != nil) || (s != fileStoreID) {
		return false, err
	}
	// 添加记录
	sql := "INSERT INTO FileFolder values (?, ?, ?, ?,?)"
	_, err = global.DB.Exec(sql, nil, folderName, parentFolderID, fileStoreID, time.Now().Unix())
	if err != nil {
		return false, err
	}
	return true, nil
}

func GetUserIDByFolderID(folderID int) (int, error) {
	sql := "SELECT FileStoreID FROM FileFolder WHERE ID = ?"
	var t int
	err := global.DB.Get(&t, sql, folderID)
	if err != nil {
		return t, err
	}
	return GetUserIDByFileStoreID(t)
}
