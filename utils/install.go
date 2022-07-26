package utils

import (
	"github.com/TapXWorld/fastDeployEnvirment/base/structs"
	"log"
	"os"
	"strings"
)

type User struct {
	ProductName  []string
	DownloadPath string
	SystemType   int //0 window 1 linux
}

func Install(code string, productMap map[string]interface{}, user User) {
	var url string
	var obj = productMap[code]

	switch obj.(type) {
	case *structs.IIU:
		url = obj.(*structs.IIU).Releases[0].Downloads.WindowsZip.Link
	case *structs.Goland:
		url = obj.(*structs.Goland).Releases[0].Downloads.WindowsZip.Link
	case *structs.PyCharm:
		url = obj.(*structs.PyCharm).Releases[0].Downloads.Windows.Link
	case *structs.PhpStorm:
		url = obj.(*structs.PhpStorm).Releases[0].Downloads.Windows.Link
	case *structs.WebStorm:
		url = obj.(*structs.WebStorm).Releases[0].Downloads.Windows.Link
	}
	urlArr := strings.Split(url, "/")
	filePath := user.DownloadPath + "/" + urlArr[len(urlArr)-1]

	if _, err := os.Stat(filePath); err != nil {
		HttpDownload(url, user.DownloadPath, urlArr[len(urlArr)-1])
	}

	nameArr := strings.Split(strings.Replace(urlArr[len(urlArr)-1], ".", "_", 0), ".")
	fullName := strings.Join(nameArr[0:len(nameArr)-2], ".")

	log.Println(fullName + " extracting ...")
	Unzip(filePath, user.DownloadPath+"/"+fullName)
	log.Println(fullName + " extract ok.")
}
