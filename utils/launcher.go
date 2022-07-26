package utils

import (
	"os"
	"os/user"
)

func CreateLinuxLauncher(path string, code string) {
	user, _ := user.Current()

	shortcutText := "[Desktop Entry]" + "\n" +
		"Version=1.0" + "\n" +
		"Type=Application" + "\n" +
		"Name=" + code + "" + "\n" +
		"Icon=" + path + "bin/goland.png" + "\n" +
		"Exec='" + path + "/bin/goland.sh %f'" + "\n" +
		"Comment=" + code + " IDE" + "\n" +
		"Categories=Development;IDE;" + "\n" +
		"Terminal=false" + "\n" +
		"StartupWMClass=jetbrains-gogland"

	username := user.Username

	shortcutLink := "/home/" + username + "/Desktop/" + code + ".desktop"

	file, _ := os.Create(shortcutLink)

	file.Write([]byte(shortcutText))
}

func CreateWindowsLauncher(path string, code string) {
	//get current user home

	//
}
