package service

import (
	"MyCloudDisk/global"
	"MyCloudDisk/model/DBModels"
	"crypto/md5"
)

func generateUserID(username string) string {
	sum := md5.Sum([]byte(username))
	return string(sum[:])
}

func GetUserFromUserID(userID int) (*DBModels.User, error) {
	u := new(DBModels.User)
	sql := "select * from User where UserID = ?"
	err := global.DB.Get(u, sql, userID)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func JudgeUser(Username, Password string) *DBModels.User {
	u := new(DBModels.User)
	sql := "select * from User where Username = ? and Password = ?"
	err := global.DB.Get(u, sql, Username, Password)
	if err != nil {
		return nil
	}
	return u
}

func GetUser(username string, password string) *DBModels.User {
	return JudgeUser(username, password)
}

func NewUser(username string, password string) (bool, error) {
	ok := GetUser(username, password)
	if ok != nil {
		return false, nil
	}
	sql := "insert into Users values (?, ?, ?)"
	userID := generateUserID(username)
	_, err := global.DB.Exec(sql, userID, username, password)
	if err != nil {
		return false, err
	}

	return true, nil
}

func GetUserRootFolderID(userID int) (int, error) {
	storeID, err := GetFileStoreIDByUserID(userID)
	if err != nil {
		return 0, err
	}
	return GetFileStoreRootFolderID(storeID)
}
