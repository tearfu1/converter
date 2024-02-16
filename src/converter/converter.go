package main

import (
	"encoding/xml"
	"envelope"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	data, err := os.ReadFile("eurofxref-daily.xml")
	check(err)

	envlp := new(envelope.Envelope)
	err = xml.Unmarshal([]byte(data), envlp)
	check(err)

	fmt.Println(envlp)
}
