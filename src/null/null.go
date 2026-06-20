package NullString

import (
	MethodNotImplementedError "github.com/go-composites/error/src/method_not_implemented"
	Null "github.com/go-composites/null/src"
	Result "github.com/go-composites/result/src"
	String "github.com/go-composites/string/src"
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

func (d data) Length() int {
	return 0
}

func (d data) Concat(other String.Interface) Result.Interface {
	return Result.New(
		Result.WithPayload(d),
	)
}

func (d data) Contains(substr string) bool {
	return false
}

func (d data) Replace(old, new string) Result.Interface {
	return Result.New(
		Result.WithPayload(d),
	)
}

func (d data) Upper() Result.Interface {
	return Result.New(
		Result.WithPayload(d),
	)
}

func (d data) Lower() Result.Interface {
	return Result.New(
		Result.WithPayload(d),
	)
}

func (d data) Trim() Result.Interface {
	return Result.New(
		Result.WithPayload(d),
	)
}

func (d data) Equal(other String.Interface) bool {
	return false
}
