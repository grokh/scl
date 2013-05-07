package main

import (
	"fmt"
	"flag"
)

var port = flag.Int("port", 5443, "Which port to listen on for requests. Default: 5443")

func main() {
	fmt.Println("Hello, World!")
}