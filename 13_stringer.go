package main

import (
	"fmt"
	"strconv"
)

type IPAddr [4]int

// TODO: Add a "String() string" method to IPAddr.
func (ip IPAddr)String() string{
	return strconv.Itoa(ip[0])+"."+strconv.Itoa(ip[0])+"."+strconv.Itoa(ip[0])+"."+strconv.Itoa(ip[0])
}
func main() {
	addrs := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for n, a := range addrs {
		fmt.Printf("%v: %v\n", n, a)
	}
}