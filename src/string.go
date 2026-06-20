package String

import (
	"fmt"
	"strings"
	"unicode/utf8"

	Array "github.com/go-composites/array/src"
	Result "github.com/go-composites/result/src"
)

type Interface interface {
	Set(string) Result.Interface
	ToGoString() string
	Split(string) Result.Interface
	Length() int
	Concat(Interface) Result.Interface
	Contains(substr string) bool
	Replace(old, new string) Result.Interface
	Upper() Result.Interface
	Lower() Result.Interface
	Trim() Result.Interface
	Equal(Interface) bool
	StartsWith(prefix string) bool
	EndsWith(suffix string) bool
	Format(args ...interface{}) Result.Interface
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

/*
Length returns the number of runes in the String (not the byte length), so
multibyte UTF-8 content is counted correctly.
*/
func (d data) Length() int {
	return utf8.RuneCountInString(d.value)
}

/*
Concat returns, in the Result payload, a new String equal to the receiver
followed by the argument.
*/
func (d data) Concat(other Interface) Result.Interface {
	return Result.New(
		Result.WithPayload(New(WithGoString(d.value + other.ToGoString()))),
	)
}

/*
Contains reports, as a Go bool, whether substr is within the String.
*/
func (d data) Contains(substr string) bool {
	return strings.Contains(d.value, substr)
}

/*
Replace returns, in the Result payload, a new String with all non-overlapping
instances of old replaced by new.
*/
func (d data) Replace(old, new string) Result.Interface {
	return Result.New(
		Result.WithPayload(New(WithGoString(strings.ReplaceAll(d.value, old, new)))),
	)
}

/*
Upper returns, in the Result payload, a new String with all characters mapped to
their upper case.
*/
func (d data) Upper() Result.Interface {
	return Result.New(
		Result.WithPayload(New(WithGoString(strings.ToUpper(d.value)))),
	)
}

/*
Lower returns, in the Result payload, a new String with all characters mapped to
their lower case.
*/
func (d data) Lower() Result.Interface {
	return Result.New(
		Result.WithPayload(New(WithGoString(strings.ToLower(d.value)))),
	)
}

/*
Trim returns, in the Result payload, a new String with leading and trailing
white space removed.
*/
func (d data) Trim() Result.Interface {
	return Result.New(
		Result.WithPayload(New(WithGoString(strings.TrimSpace(d.value)))),
	)
}

/*
Equal reports, as a Go bool, whether the receiver and the argument carry the same
underlying string.
*/
func (d data) Equal(other Interface) bool {
	return d.value == other.ToGoString()
}

/*
StartsWith reports, as a Go bool, whether the String begins with prefix.
*/
func (d data) StartsWith(prefix string) bool {
	return strings.HasPrefix(d.value, prefix)
}

/*
EndsWith reports, as a Go bool, whether the String ends with suffix.
*/
func (d data) EndsWith(suffix string) bool {
	return strings.HasSuffix(d.value, suffix)
}

/*
Format returns, in the Result payload, a new String equal to fmt.Sprintf applied
to the receiver's value (used as the format string) and args.
*/
func (d data) Format(args ...interface{}) Result.Interface {
	return Result.New(
		Result.WithPayload(New(WithGoString(fmt.Sprintf(d.value, args...)))),
	)
}

func (d data) IsNull() bool {
	return false
}

// null is the Null-Object variant of a String: an empty, immutable placeholder
// that honours the full Interface without ever being nil. New-value operations
// return a successful Result wrapping the null String (or, where chaining is
// meaningful, the receiver); predicates always report false; ToGoString is "".
type null struct{}

// Null returns the Null-Object String.
func Null() Interface {
	return &null{}
}

func (n *null) Set(value string) Result.Interface {
	return Result.New(
		Result.WithPayload(n),
	)
}

func (n *null) ToGoString() string { return `` }

func (n *null) Split(separator string) Result.Interface {
	return Result.New(
		Result.WithPayload(Array.New()),
	)
}

func (n *null) Length() int { return 0 }

func (n *null) Concat(other Interface) Result.Interface {
	return Result.New(
		Result.WithPayload(n),
	)
}

func (n *null) Contains(substr string) bool { return false }

func (n *null) Replace(old, new string) Result.Interface {
	return Result.New(
		Result.WithPayload(n),
	)
}

func (n *null) Upper() Result.Interface {
	return Result.New(
		Result.WithPayload(n),
	)
}

func (n *null) Lower() Result.Interface {
	return Result.New(
		Result.WithPayload(n),
	)
}

func (n *null) Trim() Result.Interface {
	return Result.New(
		Result.WithPayload(n),
	)
}

func (n *null) Equal(other Interface) bool { return false }

func (n *null) StartsWith(prefix string) bool { return false }

func (n *null) EndsWith(suffix string) bool { return false }

func (n *null) Format(args ...interface{}) Result.Interface {
	return Result.New(
		Result.WithPayload(n),
	)
}

// IsNull reports that this is the null String.
func (n *null) IsNull() bool { return true }
