package main

import (
	"runtime"
		"github.com/TapXWorld/oneKey_develoment/linux"
	"github.com/TapXWorld/oneKey_develoment/windows"
)

func main() {
	if "windows" == runtime.GOOS {

		wf := windowsPlatform{}

	} else {
		lp := linuxPlatform{}
	}
}
