package String

import (
	"strings"

	Array "github.com/golang-oop/array/src"
	Result "github.com/golang-oop/result/src"
)

type Interface interface {
	Set(string) Result.Interface
	ToGoString() string
	Split(string) Result.Interface
	IsNull() bool
}

type data struct {
	value string
}

type Option func(*data)

/*
Functional parameter to set a String value field when calling the New constructor.

# Usage

	s := String.New(String.WithGoString(`Hello World!`))
*/
func WithGoString(value string) Option {
	return func(d *data) {
		d.value = value
	}
}

/*
String constructor.

Called without the functionnal parameter String.WithGoString(), set the field value with an empty Go string.

# Usage

		s1 := String.New()
	    s2 := String.New(String.WithGoString(`Hello World!`))
*/
func New(options ...Option) Interface {
	d := &data{
		value: ``,
	}
	for _, opt := range options {
		opt(d)
	}
	return d
}

/*
Set the String value field
Return the modified String.Interface in the Result payload.
*/
func (d *data) Set(value string) Result.Interface {
	d.value = value
	return Result.New(
		Result.WithPayload(d),
	)
}

/*
Return the String value field
*/
func (d data) ToGoString() string {
	return d.value
}

/*
Return an Array of the splitted fields in the Result payload.

# Usage

	    s := String.New(String.WithGoString(`Hello World!`))
	    if r := s.Split(` `); r.HasError {
			fmt.Println(r.Error().Message())
		} else {
			r.Payload().(Array.Interface).First()
		}
*/
func (d data) Split(separator string) Result.Interface {
	src := strings.Split(d.value, separator)
	dst := Array.New()
	for _, item := range src {
		dst.Push(New(WithGoString(item)))
	}
	return Result.New(
		Result.WithPayload(dst),
	)
}

func (d data) IsNull() bool {
	return false
}
