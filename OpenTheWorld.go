package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/TapXWorld/fastDeployEnvirment/base/structs"
	"github.com/TapXWorld/fastDeployEnvirment/utils"
	yaml "github.com/goccy/go-yaml"
	"io"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func cacheVersionInfo() {
	data, _ := os.ReadFile("meta.yml")

	if err := yaml.Unmarshal(data, &meta); err != nil {
		log.Println(err)
	}
	if utils.PathExists("./.cache/") {
		return
	} else {
		for i := 0; i < len(meta.Software); i++ {
			replaceParameter := strings.ReplaceAll(meta.Software[i].Parameters, "{timestamp}", strconv.FormatInt(time.Now().UnixMilli(), 10))

			res := utils.HttpGet(meta.Url + replaceParameter)
			response, _ := io.ReadAll(res.Body)

			os.Mkdir(".cache", os.ModePerm)
			ioutil.WriteFile(".cache/"+meta.Software[i].ProductCode+".cache", response, 0644)
		}

	}
}

func validateInput(u User) bool {
	for i := 0; i < len(u.productName); i++ {
		inputError := true

		for j := 0; j < len(meta.Software); j++ {
			if strings.ToLower(u.productName[i]) == strings.ToLower(meta.Software[j].ProductName) {
				inputError = false
			}
		}
		if inputError {
			fmt.Println("---> product name error. please input again.")
			return true
		}
	}
	return false
}

func installOptions() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Which product of you want download( IDEA,Goland,WebStorm,PhpStorm,PyCharm )? ")
		fmt.Print("----->  (default(IDEA,Goland): ): ")

		str, _, _ := reader.ReadLine()

		product := strings.Replace(string(str), " ", "", -1)

		user.productName = strings.Split(product, ",")

		fmt.Println("Where to save? ")
		fmt.Print("----->  (default: \"D:\\software\\\"): ")

		savePath, _, _ := reader.ReadLine()

		user.downloadPath = string(savePath)

		if validateInput(user) {
			fmt.Println("validate error,　input again.")
		} else {
			break
		}
	}
}

type Meta struct {
	Url      string
	Software []struct {
		ProductName string `yaml:"product_name"`
		ProductCode string `yaml:"product_code"`
		Parameters  string `yaml:"parameters"`
	}
}

type User struct {
	productName  []string
	downloadPath string
	systemType   int //0 window 1 linux
}

var user = User{}

var meta = Meta{}

var productCode = []string{"IIU", "Go", "PCP", "PS", "WS"}

var productMap = make(map[string]interface{})

func loadProductJson() {
	for _, p := range productCode {
		strByte, _ := ioutil.ReadFile(".cache/" + p + ".cache")
		switch p {
		case "IIU":
			iiu := &structs.IIU{}
			json.Unmarshal(strByte, &iiu)
			productMap["IIU"] = iiu
		case "Go":
			goland := &structs.Goland{}
			json.Unmarshal(strByte, &goland)
			productMap["Go"] = json.Unmarshal(strByte, &structs.Goland{})
		case "PCP":
			pyCharm := &structs.PyCharm{}
			json.Unmarshal(strByte, &pyCharm)
			productMap["PCP"] = json.Unmarshal(strByte, &structs.PyCharm{})
		case "PS":
			phpStorm := &structs.PhpStorm{}
			json.Unmarshal(strByte, &phpStorm)
			productMap["PS"] = json.Unmarshal(strByte, &structs.PhpStorm{})
		case "WS":
			webStorm := &structs.WebStorm{}
			json.Unmarshal(strByte, &webStorm)
			productMap["WS"] = json.Unmarshal(strByte, &structs.WebStorm{})
		default:
			break
		}
	}
}

func main() {
	loadProductJson()
	cacheVersionInfo()
	installOptions()

	if "windows" == runtime.GOOS {
		user.systemType = 0
	} else {
		user.systemType = 1
	}
	for _, name := range user.productName {
		for j := 0; j < len(meta.Software); j++ {
			if strings.EqualFold(name, meta.Software[j].ProductName) {
				go install(meta.Software[j].ProductCode)
			}
		}
	}
	time.Sleep(time.Second * 5000)
	fmt.Println("All Completed.")
}

func install(code string) {
	var url string
	var obj = productMap[code]

	switch obj.(type) {
	case *structs.IIU:
		url = obj.(*structs.IIU).Releases[0].Downloads.WindowsZip.Link
	case structs.Goland:
		url = obj.(*structs.Goland).Releases[0].Downloads.WindowsZip.Link
	case structs.PyCharm:
		url = obj.(*structs.PyCharm).Releases[0].Downloads.Windows.Link
	case structs.PhpStorm:
		url = obj.(*structs.PhpStorm).Releases[0].Downloads.Windows.Link
	case structs.WebStorm:
		url = obj.(*structs.WebStorm).Releases[0].Downloads.Windows.Link
	}
	urlArr := strings.Split(url, "/")
	filePath := user.downloadPath + "/" + urlArr[len(urlArr)-1]

	if _, err := os.Stat(filePath); err != nil {
		utils.HttpDownload(url, user.downloadPath, urlArr[len(urlArr)-1])
	}

	nameArr := strings.Split(strings.Replace(urlArr[len(urlArr)-1], ".", "_", 0), ".")
	fullName := strings.Join(nameArr[0:len(nameArr)-2], ".")

	utils.Unzip(filePath, user.downloadPath+"/"+fullName)
}
