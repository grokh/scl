package main

import (
	"fmt"
	"flag"
)

var _ = fmt.Println

var port = flag.Int("port", 5443, "TLS listening port")

func main() {
	flag.Parse()

	// receive TLS
	// verify certificate
	// check OCSP
	// check AD // go get github.com/mavricknz/ldap
	// return TLS good/bad/unknown
}