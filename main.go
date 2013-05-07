package main

import (
	"fmt"
	"flag"
)

var port = flag.Int("port", 5443, "Listen port")

func main() {
	flag.Parse()
	fmt.Println("Hello, World!")

	// receive TLS
	// verify certificate
	// check OCSP
	// check AD
	// return TLS good/bad/unknown
}