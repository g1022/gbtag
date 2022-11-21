package gbtag

import (
	"encoding/base64"
	"fmt"
)

var Base64Data = "" //以Base64编码的数据，一串的json

func DecodeBase64() string {
	if len(Base64Data) == 0 {
		return ""
	}
	jsonByte, err := base64.StdEncoding.DecodeString(Base64Data)
	if err != nil {
		fmt.Println("error: ", err)
	}
	return string(jsonByte)
}

func ToString() string {
	return DecodeBase64()
}
