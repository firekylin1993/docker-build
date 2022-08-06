package main

import (
	"flag"
	"fmt"
)

var (
	Name    string
	Version string
)

func init() {
	flag.StringVar(&Name, "name", "", "服务名")
}

func main() {
	flag.Parse()
	fmt.Printf("hello %s, this is %s\n", Name, Version)
}
