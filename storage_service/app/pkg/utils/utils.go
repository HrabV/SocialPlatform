package utils

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/google/uuid"
	"net/http"
	"strings"
)

func GenerateFileObjectName(fileName string) string {
	ext := strings.Split(fileName, ".")[1]

	hash := md5.Sum([]byte(fileName))

	sb := strings.Builder{}

	sb.WriteString(hex.EncodeToString(hash[:]))
	sb.WriteString("__")
	sb.WriteString(uuid.New().String())
	sb.WriteString(".")
	sb.WriteString(ext)

	return sb.String()
}

func MinioLiveCheck(endpoint string, useSSL bool) bool {
	schema := "http://"

	if useSSL {
		schema = "https://"
	}

	_, err := http.Get(schema + endpoint + "/minio/health/live")
	if err != nil {
		return false
	}
	return true
}
