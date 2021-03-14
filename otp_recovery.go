package main

import (
	"fmt"
	"os"
	"net/url"
	"encoding/base32"
	"encoding/base64"
	"log"
	"strings"
	"github.com/golang/protobuf/proto"
)

func main() {
	argsWithoutProg := os.Args[1:]
	decodedValue, err := url.QueryUnescape(strings.Join(argsWithoutProg, " "))
	if err != nil {
		log.Fatal(err)
		return
	}
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
	for i, msg := range otp_keys.GetOtpParameters() {
		fmt.Printf("Record %d : %s\n", i, msg)
		readableSecret := base32.StdEncoding.EncodeToString(msg.GetSecret())
		fmt.Printf("%s\n\n", readableSecret)
	}
}
