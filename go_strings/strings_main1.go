package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strings"
)

func BytesToint(bys []byte) string {
	bytebuff := bytes.NewBuffer(bys)
	var data string
	binary.Read(bytebuff, binary.BigEndian, &data)
	return string(data)
}

type IPAddr [4]byte

func (s IPAddr) String() string {
	var strings_array []string
	for _, value := range s {
		strings_array = append(strings_array, BytesToint(value))
	}

	return strings.Join(strings_array, ",")
}

func main() {
	ipAddrs := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range ipAddrs {
		fmt.Printf("%v:%v\n", name, ip)
	}
}
