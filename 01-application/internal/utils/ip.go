package util

import (
	"net"

	"github.com/sirupsen/logrus"
)

func GetLocalIP(log *logrus.Logger) string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Errorf("cannot list network interfaces: %v", err)
		return "unknown"
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ip4 := ipnet.IP.To4(); ip4 != nil {
				return ip4.String()
			}
		}
	}

	return "unknown"
}
