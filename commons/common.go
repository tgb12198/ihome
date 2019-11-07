package commons

import (
	"encoding/base64"
)

//加密
func EncryptData(str string) string {
	strBytes := []byte(str)
	secretKey := base64.StdEncoding.EncodeToString(strBytes)
	return secretKey
}

//解密
func DecryptData(str string) string {
	decoded, err := base64.StdEncoding.DecodeString(str)
	if err == nil {
		return string(decoded)
	}
	return ""
}

const BASEPATH = "http://192.168.0.128/"
