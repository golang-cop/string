package main

import (
	"fmt"

	String "github.com/golang-oop/string/src"
)

func main() {
	s := String.New(
		String.WithGoString(`Hello World!`),
	)
	fmt.Println(s.ToGoString())
	if r := s.Split(` `); r.HasError() {
		fmt.Printf(
			"%s\n",
			r.Error().Message(),
		)
	} else {
		fmt.Printf(
			"%+v",
			r.Payload(),
		)
	}
}
