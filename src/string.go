package String

import (
	"strings"

	Array "github.com/golang-oop/array/src"
	Result "github.com/golang-oop/result/src"
)

type Interface interface {
	ToGoString() string
	Split(string) Result.Interface
	IsNull() bool
}

type data struct {
	value string
}

func New(value string) Interface {
	return &data{
		value: value,
	}
}

func (d data) ToGoString() string {
	return d.value
}

func (d data) Split(separator string) Result.Interface {
	src := strings.Split(d.value, separator)
	dst := Array.New()
	for _, item := range src {
		dst.Push(New(item))
	}
	return Result.New(
		Result.WithPayload(dst),
	)
}

func (d data) IsNull() bool {
	return false
}
