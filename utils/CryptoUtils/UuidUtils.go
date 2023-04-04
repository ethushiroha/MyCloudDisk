package CryptoUtils

import (
	"github.com/google/uuid"
	"strings"
)

func GetUuid() string {
	uid := uuid.New().String()
	return strings.Replace(uid, "-", "", -1)
}
