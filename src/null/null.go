package NullString

import (
	MethodNotImplementedError "github.com/golang-cop/error/src/method_not_implemented"
	Null "github.com/golang-cop/null/src"
	Result "github.com/golang-cop/result/src"
	String "github.com/golang-cop/string/src"
)

type Interface interface {
	String.Interface
}

type data struct {
	value Null.Interface
}

func New() Interface {
	return &data{
		value: Null.New(),
	}
}

func (d data) ToGoString() string {
	return ``
}

func (d data) IsNull() bool {
	return true
}

func (d data) Split(separator string) Result.Interface {
	return Result.New(
		Result.WithPayload([]string{}),
	)
}

func (d data) Set(value string) Result.Interface {
	return Result.New(
		Result.WithError(
			MethodNotImplementedError.New(`Set`),
		),
	)
}
