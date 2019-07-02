package net

import (
	"fmt"
	"net"
)

// LookupAddr is
func LookupAddr() {
	names, err := net.LookupAddr("35.190.247.0")
	fmt.Println(names)
	fmt.Println(err)

	name, err := net.LookupCNAME("golang.org")
	fmt.Println(name)
	fmt.Println(err)

	addrs, err := net.LookupHost("golang.org")
	fmt.Println(addrs)
	fmt.Println(err)

	port, err := net.LookupPort("tcp", "http")
	fmt.Println(port)
	fmt.Println(err)

	txts, err := net.LookupTXT("golang.org")
	fmt.Println(txts)
	fmt.Println(err)
}

// InterfaceAddrs is
func InterfaceAddrs() {
	addrs, err := net.InterfaceAddrs()
	fmt.Println(err)
	for _, addr := range addrs {
		fmt.Println("network:" + addr.Network() + ", address:" + addr.String())
	}
}
