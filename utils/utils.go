package utils

import (
	"archive/zip"
	b64 "encoding/base64"
	"fmt"
	"github.com/schollz/progressbar/v3"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
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

	f, _ := os.OpenFile(path+"/"+saveName, os.O_CREATE|os.O_WRONLY, 0644)

	defer f.Close()

	fmt.Println(resp.ContentLength)

	bar := progressbar.DefaultBytes(
		resp.ContentLength,
		"downloading",
	)
	io.Copy(io.MultiWriter(f, bar), resp.Body)
	return true
}

func Unzip(src string, dest string) ([]string, error) {
	var filenames []string

	r, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	defer r.Close()

	for _, f := range r.File {
		path := filepath.Join(dest, f.Name)

		if !strings.HasPrefix(path, filepath.Clean(dest)+string(os.PathSeparator)) {
			return filenames, fmt.Errorf("%s: illegal file path", path)
		}

		filenames = append(filenames, path)

		if f.FileInfo().IsDir() {
			// Make Folder
			os.MkdirAll(path, os.ModePerm)
			continue
		}
		if err = os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
			return filenames, err
		}

		outFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return filenames, err
		}
		rc, _ := f.Open()

		io.Copy(outFile, rc)

		outFile.Close()
		rc.Close()
		if err != nil {
			return filenames, err
		}
	}
	return filenames, nil
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
