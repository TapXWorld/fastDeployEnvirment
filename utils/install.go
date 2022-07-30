package utils

import (
	"github.com/TapXWorld/fastDeployEnvirment/base/structs"
	"log"
	"os"
	"strings"
)

type User struct {
	ProductName     []string
	DownloadPath    string
	SystemType      int //0 window 1 linux
	ProductFullName string
}

func Install(code string, productMap map[string]interface{}, user *User) {
	var url string
	var obj = productMap[code]

	switch obj.(type) {
	case *structs.IIU:
		if user.SystemType == 0 {
			url = obj.(*structs.IIU).Releases[0].Downloads.WindowsZip.Link
		} else {
			url = obj.(*structs.IIU).Releases[0].Downloads.Linux.Link
		}
	case *structs.Goland:
		if user.SystemType == 0 {
			url = obj.(*structs.Goland).Releases[0].Downloads.WindowsZip.Link
		} else {
			url = obj.(*structs.Goland).Releases[0].Downloads.Linux.Link
		}
	case *structs.PyCharm:
		if user.SystemType == 0 {
			url = obj.(*structs.PyCharm).Releases[0].Downloads.Windows.Link
		} else {
			url = obj.(*structs.PyCharm).Releases[0].Downloads.Linux.Link
		}
	case *structs.PhpStorm:
		if user.SystemType == 0 {
			url = obj.(*structs.PhpStorm).Releases[0].Downloads.Windows.Link
		} else {
			url = obj.(*structs.PhpStorm).Releases[0].Downloads.Linux.Link
		}
	case *structs.WebStorm:
		if user.SystemType == 0 {
			url = obj.(*structs.WebStorm).Releases[0].Downloads.Windows.Link
		} else {
			url = obj.(*structs.WebStorm).Releases[0].Downloads.Linux.Link
		}
	}
	urlArr := strings.Split(url, "/")

	filePath := user.DownloadPath + urlArr[len(urlArr)-1]

	if _, err := os.Stat(filePath); err != nil {
		HttpDownload(url, user.DownloadPath, urlArr[len(urlArr)-1])
	}

	nameArr := strings.Split(strings.Replace(urlArr[len(urlArr)-1], ".", "_", 0), ".")
	fullName := strings.Join(nameArr[0:len(nameArr)-2], ".")
	user.ProductFullName = fullName

	log.Println(fullName + " extracting ...")
	DecompressFile(filePath, user.DownloadPath)

	log.Println(fullName + " extract ok.")
}
