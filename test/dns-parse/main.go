package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	dns := "test-2-37.apulis.com.cn"

	ns, err := net.LookupHost(dns)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Err: %s", err.Error())
		return
	}
	fmt.Println(ns)
}
