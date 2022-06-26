package main

import (
	"bufio"
	"fmt"
	"github.com/TapXWorld/fastDeployEnvirment/utils"
	"github.com/goccy/go-yaml"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Meta struct {
	Url      string
	Software []struct {
		ProductName string `yaml:"product_name"`
		ProductCode string `yaml:"product_code"`
		Parameters  string `yaml:"parameters"`
	}
}

type User struct {
	productName  string
	downloadPath string
	systemType   int //0 window 1 linux
}

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
	productName := strings.Split(u.productName, ",")

	for i := 0; i < len(productName); i++ {
		inputError := true

		for j := 0; j < len(meta.Software); j++ {

			if strings.ToLower(productName[i]) == strings.ToLower(meta.Software[j].ProductName) {
				inputError = false
			}
		}
		if inputError {
			return true
		}
	}
	return false
}

func installOptions() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("--- Which product of you want download( IDEA,Goland,WebStorm,PhpStorm,PyCharm )? ")
		fmt.Print("  (default(IDEA,Goland): ): ")

		str, _, _ := reader.ReadLine()
		user.productName = string(str)

		fmt.Println("--- Where to save? ")
		fmt.Print("  (default: \"D:\\software\\\"): ")

		savePath, _, _ := reader.ReadLine()

		user.downloadPath = string(savePath)

		if validateInput(user) {
			fmt.Println("validate error,　input again.")
		} else {
			break
		}
	}
}

var user = User{}

var meta = Meta{}

func main() {
	cacheVersionInfo()
	installOptions()

	fmt.Println(user)
	//time.Sleep(time.Second * 10)
}
