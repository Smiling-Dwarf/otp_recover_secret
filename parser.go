package main

import (
	"fmt"
	"os"
	"net/url"
	"log"
	"strings"
	"github.com/golang/protobuf/proto"
)

func main() {
	fmt.Println("Hello, world.")
	argsWithoutProg := os.Args[1:]
	fmt.Println(argsWithoutProg)
	decodedValue, err := url.QueryUnescape(strings.Join(argsWithoutProg, " "))
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(decodedValue)
}
