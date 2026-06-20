package String

import (
	"testing"

	Array "github.com/go-composites/array/src"
)

func TestNewDefaultIsEmpty(t *testing.T) {
	s := New()
	if got := s.ToGoString(); got != `` {
		t.Fatalf("New() ToGoString = %q, want empty", got)
	}
}

func TestNewWithGoString(t *testing.T) {
	s := New(WithGoString(`Hello World!`))
	if got := s.ToGoString(); got != `Hello World!` {
		t.Fatalf("ToGoString = %q, want %q", got, `Hello World!`)
	}
}

func TestSet(t *testing.T) {
	s := New()
	r := s.Set(`changed`)
	payload, ok := r.Payload().(Interface)
	if !ok {
		t.Fatalf("Set payload is not String.Interface: %T", r.Payload())
	}
	if got := payload.ToGoString(); got != `changed` {
		t.Fatalf("after Set ToGoString = %q, want %q", got, `changed`)
	}
	if got := s.ToGoString(); got != `changed` {
		t.Fatalf("receiver mutated value = %q, want %q", got, `changed`)
	}
}

func TestIsNull(t *testing.T) {
	if New().IsNull() {
		t.Fatal("String.IsNull() = true, want false")
	}
}

func TestSplit(t *testing.T) {
	s := New(WithGoString(`a b c`))
	r := s.Split(` `)
	arr, ok := r.Payload().(Array.Interface)
	if !ok {
		t.Fatalf("Split payload is not Array.Interface: %T", r.Payload())
	}

	firstItem, ok := arr.First().Payload().(Interface)
	if !ok {
		t.Fatalf("first item is not String.Interface: %T", arr.First().Payload())
	}
	if got := firstItem.ToGoString(); got != `a` {
		t.Fatalf("first split field = %q, want %q", got, `a`)
	}

	thirdItem, ok := arr.Fetch(2).Payload().(Interface)
	if !ok {
		t.Fatalf("third item is not String.Interface: %T", arr.Fetch(2).Payload())
	}
	if got := thirdItem.ToGoString(); got != `c` {
		t.Fatalf("third split field = %q, want %q", got, `c`)
	}
}

// TestSplitUTF8 exercises byte-level multibyte UTF-8 handling so the same
// behaviour is asserted on every (little- and big-endian) architecture.
func TestSplitUTF8(t *testing.T) {
	s := New(WithGoString(`café→naïve→Ω`))
	arr := s.Split(`→`).Payload().(Array.Interface)

	first := arr.First().Payload().(Interface)
	if got := first.ToGoString(); got != `café` {
		t.Fatalf("first UTF-8 field = %q, want %q", got, `café`)
	}
	third := arr.Fetch(2).Payload().(Interface)
	if got := third.ToGoString(); got != `Ω` {
		t.Fatalf("third UTF-8 field = %q, want %q", got, `Ω`)
	}
}
