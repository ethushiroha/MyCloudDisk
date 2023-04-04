package CryptoUtils

import (
	"MyCloudDisk/model/DBModels"
	"crypto/md5"
	"encoding/hex"
	"io"
	"strings"
)

func GetPostfix(filename string) string {
	index := strings.LastIndex(filename, ".")
	return strings.ToLower(filename[index+1:])
}

func FileHash(file io.Reader) string {
	md5h := md5.New()
	io.Copy(md5h, file)
	return hex.EncodeToString(md5h.Sum(nil))
}

func FileType(filename string) int {
	postFix := GetPostfix(filename)
	switch postFix {
	case "txt", "text", "bin":
		return DBModels.TEXT
	case "jpg", "png", "gif", "jpeg", "bmp", "svg":
		return DBModels.IMAGE

	default:
		return DBModels.OTHER
	}
}
