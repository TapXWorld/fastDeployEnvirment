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

func validateInput(u utils.User) bool {
	for i := 0; i < len(u.ProductName); i++ {
		inputError := true

		for j := 0; j < len(meta.Software); j++ {
			if strings.ToLower(u.ProductName[i]) == strings.ToLower(meta.Software[j].ProductName) {
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
		fmt.Println("Which product are you want Download? ")
		fmt.Println("	1.IDEA")
		fmt.Println("	2.Goland")
		fmt.Println("	3.WebStorm")
		fmt.Println("	4.PhpStorm")
		fmt.Println("	5.PyCharm")
		fmt.Print("Input Product Name(Example: IDEA,Goland): ")

		str, _, _ := reader.ReadLine()

		product := strings.Replace(string(str), " ", "", -1)

		user.ProductName = strings.Split(product, ",")

		fmt.Println("Where to save? ")
		fmt.Print("----->  (default: \"D:\\software\\\"): ")

		savePath, _, _ := reader.ReadLine()

		user.DownloadPath = string(savePath)

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
			productMap["Go"] = goland
		case "PCP":
			pyCharm := &structs.PyCharm{}
			json.Unmarshal(strByte, &pyCharm)
			productMap["PCP"] = pyCharm
		case "PS":
			phpStorm := &structs.PhpStorm{}
			json.Unmarshal(strByte, &phpStorm)
			productMap["PS"] = phpStorm
		case "WS":
			webStorm := &structs.WebStorm{}
			json.Unmarshal(strByte, &webStorm)
			productMap["WS"] = webStorm
		default:
			break
		}
	}
}

var user = utils.User{}

func main() {
	loadProductJson()
	cacheVersionInfo()
	installOptions()

	if "windows" == runtime.GOOS {
		user.SystemType = 0
	} else {
		user.SystemType = 1
	}
	for _, name := range user.ProductName {
		for j := 0; j < len(meta.Software); j++ {
			if strings.EqualFold(name, meta.Software[j].ProductName) {
				productCode := meta.Software[j].ProductCode

				//install IDE
				utils.Install(productCode, productMap, &user)

				//create Launcher
				if user.SystemType == 0 {
					utils.CreateWindowsLauncher(user.DownloadPath, user, name)
				} else {
					utils.CreateLinuxLauncher(user.DownloadPath, user, name)
				}
				//create license
				utils.CrackLicense()
			}
		}
	}
	fmt.Println("All Install Completed.")
}
