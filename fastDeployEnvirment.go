package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/TapXWorld/fastDeployEnvirment/utils"
	"io"
	"net/http"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
)

func installOptions() (string, string) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Which product are you want Download? ")
		fmt.Println("	1.IDEA(IIU)")
		fmt.Println("	2.Goland(Go)")
		fmt.Println("	3.WebStorm(WS)")
		fmt.Println("	4.PhpStorm(PS)")
		fmt.Println("	5.PyCharm(PCP)")
		fmt.Print("Input Product Name(Example: 1,2): ")
		str, _, _ := reader.ReadLine()
		products := strings.Replace(string(str), " ", "", -1)

		fmt.Println("Where to save? ")
		fmt.Print("----->  (default: \".\\software\"): ")
		savePath, _, _ := reader.ReadLine()
		if !utils.PathExists(string(savePath)) {
			os.MkdirAll(".\\software", os.ModePerm)
		}
		return string(savePath), products
	}
}

var codeMap = map[string]int{
	"IIU": 1,
	"Go":  2,
	"PCP": 3,
	"PS":  4,
	"WS":  5,
}

func main() {
	savePath, installs := installOptions()
	systemType := 0

	if "windows" != runtime.GOOS {
		systemType = 1
	}

	for _, selectedIndex := range strings.Split(installs, ",") {
		selectedIndex = strings.TrimSpace(selectedIndex)
		if selectedIndex == "" {
			continue
		}

		for productCode, productIndex := range codeMap {
			if selectedIndexInt, _ := strconv.Atoi(selectedIndex); selectedIndexInt == productIndex {
				url := fmt.Sprintf("https://data.services.jetbrains.com/products/releases?code=%s&latest=true&type=release", productCode)
				downloadURL, err := getDownloadURL(url, systemType)
				if err != nil {
					fmt.Printf("Get Download Link (%s): %v\n", productCode, err)
					continue
				}

				if len(savePath) == 0 {
					savePath = ".\\software"
				}

				err = downloadFile(downloadURL, savePath, productCode)
				if err != nil {
					fmt.Printf("File Download Failed (%s): %v\n", productCode, err)
				} else {
					fmt.Printf("File Download Success: %s\n", productCode)
				}
			}
		}
	}
}

func getDownloadURL(apiURL string, systemType int) (string, error) {
	resp, err := http.Get(apiURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP Request failed, Status Code: %d", resp.StatusCode)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	// 解析 JSON 数据
	for _, product := range result {
		if productArray, ok := product.([]interface{}); ok && len(productArray) > 0 {
			if productDetails, ok := productArray[0].(map[string]interface{}); ok {
				if downloads, ok := productDetails["downloads"].(map[string]interface{}); ok {
					var fileType string
					if systemType == 0 {
						fileType = "windows"
					} else {
						fileType = "linux"
					}

					if binary, ok := downloads[fileType].(map[string]interface{}); ok {
						if link, ok := binary["link"].(string); ok {
							return link, nil
						}
					}
				}
			}
		}
	}
	return "", fmt.Errorf("not found download link")
}

// 下载文件到指定目录
func downloadFile(downloadURL, savePath, productCode string) error {
	fmt.Println("Download URL: " + downloadURL)
	resp, err := http.Get(downloadURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("download Failed, Status Code: %d", resp.StatusCode)
	}
	filePath := fmt.Sprintf("%s/%s", savePath, path.Base(downloadURL))
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(file, resp.Body)
	return err
}
