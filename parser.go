package main

import (
	"fmt"
	"os"
	"reflect"
	"net/url"
	"encoding/base32"
	"encoding/base64"
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
	sDec, err := base64.StdEncoding.DecodeString(decodedValue)
	if err != nil {
		fmt.Printf("Error decoding string: %s ", err.Error())
		return
	}

	otp_keys := &MigrationPayload{}
	err = proto.Unmarshal([]byte(sDec), otp_keys)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	msg := otp_keys.GetOtpParameters()[0]
	fmt.Println(msg)
	fmt.Println(msg.GetSecret())
	str := base32.StdEncoding.EncodeToString(msg.GetSecret())
	fmt.Println(str)
	// for i, msg := msgs {
	// 	fmt.Println(msg)
	// 	//fmt.Println(msg.GetSecret())
	// }
	fmt.Println("=============================")
	fmt.Println(reflect.TypeOf(msg))
	fmt.Println("=============================")
}
