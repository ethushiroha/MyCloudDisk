package service

import (
	"MyCloudDisk/global"
	"MyCloudDisk/model/DBModels"
)

func GetStoreFileCount(UserID int) (int, error) {
	result := 0
	sql := "SELECT count(*) FROM Files WHERE FileStoreID = ?"
	err := global.DB.Get(&result, sql, UserID)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func GetStoreFilesInFolder(storeID int, folderID int) ([]DBModels.MyFile, error) {
	sql := "SELECT * FROM Files WHERE FileStoreID = ? AND ParentFolderID = ?"
	var t []DBModels.MyFile
	err := global.DB.Select(&t, sql, storeID, folderID)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func GetFileDetailUse(storeID int) map[string]interface{} {
	return map[string]interface{}{
		"docCount":   1,
		"imgCount":   3,
		"musicCount": 4,
		"otherCount": 6,
	}
}

func GetStoreIDByFileID(fileID int) (int, error) {
	sql := "SELECT FileStoreID FROM Files WHERE ID=?"
	var t int
	err := global.DB.Get(&t, sql, fileID)
	return t, err
}

func GetFileUserID(fileID int) (int, error) {
	fileStoreID, err := GetStoreIDByFileID(fileID)
	if err != nil {
		return 0, err
	}
	return GetUserIDByFileStoreID(fileStoreID)
}

func GetFileByFileID(fileID int) (*DBModels.MyFile, error) {
	file := new(DBModels.MyFile)
	sql := "SELECT * FROM Files WHERE ID=?"
	err := global.DB.Get(file, sql, fileID)
	return file, err
}

func GetParentFolderIDByFileID(fileID int) (int, error) {
	sql := "SELECT ParentFolderID FROM Files WHERE ID=?"
	var t int
	err := global.DB.Get(&t, sql, fileID)
	return t, err
}

func GetAllParentFoldersNameByFileID(fileID int) ([]string, error) {
	folderID, err := GetParentFolderIDByFileID(fileID)
	if err != nil {
		return nil, err
	}
	return GetAllParentsFoldersNameByFolderID(folderID)
}

func InsertAFile(file *DBModels.MyFile) error {
	sql := "INSERT INTO Files VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_, err := global.DB.Exec(sql, nil, file.FileName, file.FileHash, file.FileStoreID, file.FileName, 0, file.UpdateTime, file.ParentFolderID, file.Size, file.SizeStr, file.Type, file.Postfix)
	if err != nil {
		return err
	}
	return err
}
