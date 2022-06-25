package main

import (
	"fmt"
	"github.com/TapXWorld/oneKeyDeveloment/linux"
	"github.com/TapXWorld/oneKeyDeveloment/windows"
	"runtime"
)

func main() {
	if "windows" == runtime.GOOS {

		wf := windows.WindowsPlatform{Path: "test"}

		fmt.Println(wf)

	} else {
		lp := linux.LinuxPlatform{Path: ""}
	}
}
