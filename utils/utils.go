package utils

import (
	"bytes"
	"context"
	"github.com/codeclysm/extract/v3"
	"github.com/schollz/progressbar/v3"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func HttpGet(url string) *http.Response {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	return resp
}

func HttpDownload(url string, path string, saveName string) bool {
	req, _ := http.NewRequest("GET", url, nil)
	resp, _ := http.DefaultClient.Do(req)

	defer resp.Body.Close()

	f, _ := os.OpenFile(path+"/"+saveName, os.O_CREATE|os.O_WRONLY, 0644)

	defer f.Close()

	bar := progressbar.DefaultBytes(
		resp.ContentLength,
		saveName+"  -  downloading",
	)
	io.Copy(io.MultiWriter(f, bar), resp.Body)
	return true
}

func DecompressFile(filePath string, extractPath string) {
	data, _ := ioutil.ReadFile(filePath)
	buffer := bytes.NewBuffer(data)

	type Files map[string]string

	var shift = func(path string) string {
		return strings.ToLower(path)
	}
	err := extract.Archive(context.Background(), buffer, extractPath, shift)
	if err != nil {
		log.Fatalln("err happened when extract file.", err)
	}
}

func createFile(name string) (*os.File, error) {
	err := os.MkdirAll(string([]rune(name)[0:strings.LastIndex(name, "/")]), 0755)
	if err != nil {
		return nil, err
	}
	return os.Create(name)
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
