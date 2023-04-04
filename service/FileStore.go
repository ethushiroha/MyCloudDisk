package service

import (
	"MyCloudDisk/global"
	"MyCloudDisk/model/DBModels"
)

func GetFileStoreByUserID(userID int) (*DBModels.FileStore, error) {
	sql := "select * from FileStore where UserID = ?"
	fileStore := &DBModels.FileStore{}
	err := global.DB.Get(fileStore, sql, userID)
	return fileStore, err
}

func GetFileStoreIDByUserID(userID int) (int, error) {
	sql := "SELECT StoreID FROM FileStore WHERE UserID = ?"
	var t int
	err := global.DB.Get(&t, sql, userID)
	return t, err
}

func GetUserIDByFileStoreID(fileStoreID int) (int, error) {
	sql := "select UserID from FileStore where StoreID = ?"
	var t int
	err := global.DB.Get(&t, sql, fileStoreID)
	return t, err
}

func GetFileStoreRootFolderID(fileStoreID int) (int, error) {
	sql := "SELECT RootFolderID FROM FileStore WHERE StoreID = ?"
	var t int
	err := global.DB.Get(&t, sql, fileStoreID)
	return t, err
}
