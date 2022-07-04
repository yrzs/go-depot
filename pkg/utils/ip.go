package utils

import (
	"net"
	"strings"
)

func GetLocalIP() (ip string) {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		return "0.0.0.0"
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	ip = strings.Split(localAddr.String(), ":")[0]
	return
}
