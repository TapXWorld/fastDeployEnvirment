package ut

import (
	"github.com/TapXWorld/fastDeployEnvirment/utils"
	"testing"
)

var code = "goland"

var user = utils.User{ProductFullName: "goland-222.3153.7"}

func TestCreateLinuxLauncher(t *testing.T) {
	t.Run("CreateLinuxLauncher", func(t *testing.T) {
		utils.CreateLinuxLauncher("/opt/tools/GoLand-2021.3.3/", user, code)
	})
}

func TestCreateWindowsLauncher(t *testing.T) {
	t.Run("CreateWindowsLauncher", func(t *testing.T) {
		utils.CreateWindowsLauncher("d:\test", user, code)
	})
}
