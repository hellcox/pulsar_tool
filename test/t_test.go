package test

import (
	"fmt"
	"math/big"
	"net"
	"testing"
)

func Test_a(t *testing.T) {
	fmt.Println(Ip2Int(net.ParseIP("2001:0db8:85a3:0000:0000:8a2e:0370:7334")).String())
}

func Ip2Int(ip net.IP) *big.Int {
	i := big.NewInt(0)
	i.SetBytes(ip)
	return i
}

func Test_b(t *testing.T) {
	// fe80:0000:0000:0000:d376:5dc5:ba34:47e2
	// 0:0:0:0:0:FFFF:6787:66A0
	// 0xff
	ClientIP := []byte{0xfe, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xd3, 0xd3, 0xd3, 0xd3, 0xd3, 0xd3, 0xd3, 0xd3}
	ClientIP = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xff, 0xff, 0x67, 0x87, 0x66, 0xA0}
	h := fmt.Sprintf("%02x", ClientIP)
	ip := fmt.Sprintf("%s:%s:%s:%s:%s:%s:%s:%s", h[0:4], h[4:8], h[8:12], h[12:16], h[16:20], h[20:24], h[24:28], h[28:32])
	a := net.ParseIP(ip)
	fmt.Println(ip, a)
}
