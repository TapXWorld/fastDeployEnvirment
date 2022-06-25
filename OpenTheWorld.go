package main

import (
	"runtime"
)

func main() {
	if "windows" == runtime.GOOS {

		wf := windowsPlatform{}

	} else {
		lp := linuxPlatform{}
	}
}
