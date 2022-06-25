package main

import (
	"fmt"
	"github.com/TapXWorld/oneKey_develoment/linux"
	"github.com/TapXWorld/oneKey_develoment/windows"
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
