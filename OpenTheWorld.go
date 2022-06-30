package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/TapXWorld/fastDeployEnvirment/base/structs"
	"github.com/TapXWorld/fastDeployEnvirment/utils"
	"github.com/goccy/go-yaml"
	"io"
	"io/ioutil"
	"log"
	"os"
	"reflect"
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

func initStruct() {
	regStruct["IIU"] = structs.IIU{}
	regStruct["Go"] = structs.Goland{}
	regStruct["WS"] = structs.WebStorm{}
	regStruct["PS"] = structs.PhpStorm{}
	regStruct["PCP"] = structs.PyCharm{}
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

func main() {
	initStruct()

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
				code := meta.Software[j].ProductCode

				if regStruct[code] != nil {
					newObj := reflect.ValueOf(regStruct[code]).Type()

					b, _ := ioutil.ReadFile(".cache/" + code + ".cache")

					t := reflect.New(newObj).Elem()

					if "IIU" == code {
						obj := t.Interface().(structs.IIU)
						json.Unmarshal(b, &obj)

						go install(code, obj)
					} else if "Go" == code {
						obj := t.Interface().(structs.Goland)
						json.Unmarshal(b, &obj)

						go install(code, obj)
					} else if "WS" == code {
						obj := t.Interface().(structs.WebStorm)
						json.Unmarshal(b, &obj)

						go install(code, obj)
					} else if "PS" == code {
						obj := t.Interface().(structs.PhpStorm)
						json.Unmarshal(b, &obj)

						go install(code, obj)
					} else if "PCP" == code {
						obj := t.Interface().(structs.PyCharm)
						json.Unmarshal(b, &obj)

						go install(code, obj)
					} else {
						return
					}
					fmt.Println("ok")
				}
			}
		}
	}
	time.Sleep(time.Second * 5000)
	fmt.Println("All Completed.")
}

var regStruct = make(map[string]interface{})

type UnzipStruct struct {
	src  string
	desc string
}

func install(name string, s interface{}) {
	unzipFileArr := make([]UnzipStruct, 0)

	if "IIU" == name {
		url := s.(structs.IIU).Releases[0].Downloads.WindowsZip.Link
		urlArr := strings.Split(url, "/")

		filePath := user.downloadPath + "/" + urlArr[len(urlArr)-1]
		utils.HttpDownload(url, user.downloadPath, urlArr[len(urlArr)-1])

		unzipFileArr = append(unzipFileArr, UnzipStruct{src: filePath, desc: user.downloadPath + "/" + name})
	} else if "Go" == name {
		url := s.(structs.Goland).Releases[0].Downloads.WindowsZip.Link
		urlArr := strings.Split(url, "/")

		filePath := user.downloadPath + "/" + urlArr[len(urlArr)-1]
		utils.HttpDownload(url, user.downloadPath, urlArr[len(urlArr)-1])

		unzipFileArr = append(unzipFileArr, UnzipStruct{src: filePath, desc: user.downloadPath + "/" + name})
	} else if "WS" == name {
		url := s.(structs.WebStorm).Releases[0].Downloads.Windows.Link
		urlArr := strings.Split(url, "/")

		filePath := user.downloadPath + "/" + urlArr[len(urlArr)-1]
		utils.HttpDownload(url, user.downloadPath, urlArr[len(urlArr)-1])

		unzipFileArr = append(unzipFileArr, UnzipStruct{src: filePath, desc: user.downloadPath + "/" + name})
	} else if "PS" == name {
		url := s.(structs.PhpStorm).Releases[0].Downloads.Windows.Link
		urlArr := strings.Split(url, "/")

		filePath := user.downloadPath + "/" + urlArr[len(urlArr)-1]
		utils.HttpDownload(url, user.downloadPath, urlArr[len(urlArr)-1])

		unzipFileArr = append(unzipFileArr, UnzipStruct{src: filePath, desc: user.downloadPath + "/" + name})
	} else if "PCP" == name {
		url := s.(structs.PyCharm).Releases[0].Downloads.Windows.Link
		urlArr := strings.Split(url, "/")

		filePath := user.downloadPath + "/" + urlArr[len(urlArr)-1]
		utils.HttpDownload(url, user.downloadPath, urlArr[len(urlArr)-1])

		unzipFileArr = append(unzipFileArr, UnzipStruct{src: filePath, desc: user.downloadPath + "/" + name})
	}
}
