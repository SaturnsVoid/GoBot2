package components

import (
	"net/http"
	"os"
)

func startServer() {
	if isHosting {
		if isAdmin { //Check for Admin
			isopen, _ := openPort(80)
			if isopen { //Try to open port 80
				err := os.MkdirAll(tmpPath+"srv\\", os.FileMode(544)) //Make folder
				if err != nil {
				}
				n_html, _ := os.Create(tmpPath + "srv\\" + "index.html") //Make defult index
				n_html.WriteString(rawHTMLPage)
				n_html.Close()
				go srvHandle() //start webserver
			}
		}
	}
}

func editPage(name string, html string) {
	err := deleteFile(tmpPath + "srv\\" + name) //Delete old
	if err != nil {
	}
	n_html, _ := os.Create(tmpPath + "srv\\" + name) //write new
	n_html.WriteString(base64Decode(html))
	n_html.Close()
}

func srvHandle() {
	NewDebugUpdate("Hosting Webserver.")
	http.ListenAndServe(":80", http.FileServer(http.Dir(tmpPath+"srv/")))
}
