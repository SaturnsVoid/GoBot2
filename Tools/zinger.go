package main

import(
	"fmt"
	"io/ioutil"
	"net/http"
	"crypto/tls"
)

func main(){
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	//client := &http.Client{}
			req, _ := http.NewRequest("GET", "https://192.168.0.10:7777/command", nil)
			req.Header.Set("User-Agent", "d5900619da0c8a72e569e88027cd9490")
			q := req.URL.Query()
			q.Add("0", "86b4f9e6-366b-47b0-ab4e-15c6cd2f7074")
			req.URL.RawQuery = q.Encode()
			fmt.Println(req.URL.String())
			resp, err := client.Do(req)
			defer resp.Body.Close()
			resp_body, _ := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("HTTP Client Error")
			}
			if resp.StatusCode == 200 {
					if string(resp_body) == "0" {
						fmt.Println("Needs to Register")
						//break
					} else {
						fmt.Println(string(resp_body))
						//break
					}
			}
}