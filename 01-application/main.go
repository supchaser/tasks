package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
)

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "127.0.0.1"
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return "127.0.0.1"
}

func handler(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	ip := getLocalIP()
	author := os.Getenv("AUTHOR")

	fmt.Fprintf(w, `
	<h1>Hello from Cloud.ru Camp!</h1>
	<p><b>Hostname:</b> %s</p>
	<p><b>IP Address:</b> %s</p>
	<p><b>Author:</b> %s</p>
	`, hostname, ip, author)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server running on :8000")
	http.ListenAndServe(":8000", nil)
}
