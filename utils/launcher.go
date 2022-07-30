package utils

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func CreateLinuxLauncher(path string, user User, code string) bool {
	shortcutText := "[Desktop Entry]" + "\n" +
		"Version=1.0" + "\n" +
		"Type=Application" + "\n" +
		"Name=" + code + "" + "\n" +
		"Icon=" + path + user.ProductFullName + "/bin/goland.png" + "\n" +
		"Exec=" + path + user.ProductFullName + "/bin/goland.sh" + "\n" +
		"Comment=" + code + " IDE" + "\n" +
		"Terminal=false" + "\n" +
		"StartupWMClass=jetbrains-goland"

	error := ioutil.WriteFile(os.ExpandEnv("$HOME/Desktop/"+code+".desktop"), ([]byte)(shortcutText), 0644)

	if error != nil {
		log.Fatalln("error happened when create launcher.")
		log.Fatalln(error)
		return false
	}
	return true
}

// CreateWindowsLauncher invoke powershell create shortcut
func CreateWindowsLauncher(path string, user User, code string) bool {
	sourceFilePath := "$SourceFilePath = \"D:\\PowerShell\\\""
	shortcutPath := "$ShortcutPath = \"C:\\Users\\admin\\Desktop\\powershell.lnk\""
	wscriptObj := "$WScriptObj = New-Object -ComObject (\"WScript.Shell\")"
	shortcut := "$shortcut = $WscriptObj.CreateShortcut($ShortcutPath)"
	shortcutTargetPath := "$shortcut.TargetPath = $SourceFilePath"
	shortcutWindowStyle := "$shortcut.WindowStyle = 1"
	shortcutSave := "$shortcut.Save()"

	command := sourceFilePath + shortcutPath + wscriptObj + shortcut + shortcutTargetPath + shortcutWindowStyle + shortcutSave

	_, error := exec.Command("powershell", "-Exec bypass", command).CombinedOutput()
	if error != nil {
		log.Fatalln("error happened when create launcher.")
		log.Fatalln(error)
		return false
	}
	return true
}
