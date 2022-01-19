package main

import (
	"flag"
	"fmt"
)

func main() {
	str := flag.String("foo", "bar", "just testing foo!")
	flag.Parse()
	fmt.Println("foo:", *str)
}
