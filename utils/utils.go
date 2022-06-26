package utils

import (
	b64 "encoding/base64"
	"log"
	"net/http"
	"os"
)

func HttpGet(url string) *http.Response {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	return resp
}

func decodeBase64(key string) string {
	uDec, _ := b64.URLEncoding.DecodeString(key)
	return string(uDec)
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
