package main

import (
	"fmt"
	"github.com/TapXWorld/fastDeployEnvirment/linux"
	"github.com/TapXWorld/fastDeployEnvirment/windows"
	"runtime"
)

func main() {
	if "windows" == runtime.GOOS {

		wf := windows.WindowsPlatform{path: "test"}

		fmt.Println(wf)

	} else {
		lp := linux.LinuxPlatform{path: ""}
	}
}
