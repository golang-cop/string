package String

import (
	"testing"

	Array "github.com/go-composites/array/src"
	Result "github.com/go-composites/result/src"
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

func TestLength(t *testing.T) {
	// café→Ω is 6 runes but more than 6 bytes; rune count is what we assert.
	s := New(WithGoString(`café→Ω`))
	if got := s.Length(); got != 6 {
		t.Fatalf("Length = %d, want 6", got)
	}
	if got := New().Length(); got != 0 {
		t.Fatalf("empty Length = %d, want 0", got)
	}
}

func TestConcat(t *testing.T) {
	s := New(WithGoString(`Hello, `))
	r := s.Concat(New(WithGoString(`World!`)))
	payload, ok := r.Payload().(Interface)
	if !ok {
		t.Fatalf("Concat payload is not String.Interface: %T", r.Payload())
	}
	if got := payload.ToGoString(); got != `Hello, World!` {
		t.Fatalf("Concat = %q, want %q", got, `Hello, World!`)
	}
	if got := s.ToGoString(); got != `Hello, ` {
		t.Fatalf("Concat mutated receiver = %q, want %q", got, `Hello, `)
	}
}

func TestContains(t *testing.T) {
	s := New(WithGoString(`Hello World`))
	if !s.Contains(`World`) {
		t.Fatal("Contains(World) = false, want true")
	}
	if s.Contains(`xyz`) {
		t.Fatal("Contains(xyz) = true, want false")
	}
}

func TestReplace(t *testing.T) {
	s := New(WithGoString(`a-b-c`))
	r := s.Replace(`-`, `_`)
	payload, ok := r.Payload().(Interface)
	if !ok {
		t.Fatalf("Replace payload is not String.Interface: %T", r.Payload())
	}
	if got := payload.ToGoString(); got != `a_b_c` {
		t.Fatalf("Replace = %q, want %q", got, `a_b_c`)
	}
}

func TestUpperLowerTrim(t *testing.T) {
	upper := New(WithGoString(`Hello`)).Upper().Payload().(Interface)
	if got := upper.ToGoString(); got != `HELLO` {
		t.Fatalf("Upper = %q, want %q", got, `HELLO`)
	}
	lower := New(WithGoString(`Hello`)).Lower().Payload().(Interface)
	if got := lower.ToGoString(); got != `hello` {
		t.Fatalf("Lower = %q, want %q", got, `hello`)
	}
	trim := New(WithGoString("  pad  ")).Trim().Payload().(Interface)
	if got := trim.ToGoString(); got != `pad` {
		t.Fatalf("Trim = %q, want %q", got, `pad`)
	}
}

func TestEqual(t *testing.T) {
	a := New(WithGoString(`same`))
	b := New(WithGoString(`same`))
	c := New(WithGoString(`other`))
	if !a.Equal(b) {
		t.Fatal("Equal(same) = false, want true")
	}
	if a.Equal(c) {
		t.Fatal("Equal(other) = true, want false")
	}
}

func TestNull(t *testing.T) {
	n := Null()
	if !n.IsNull() {
		t.Fatal("Null().IsNull() = false, want true")
	}
	if got := n.ToGoString(); got != `` {
		t.Fatalf("Null ToGoString = %q, want empty", got)
	}
	if got := n.Length(); got != 0 {
		t.Fatalf("Null Length = %d, want 0", got)
	}
	if n.Contains(`x`) {
		t.Fatal("Null Contains = true, want false")
	}
	if n.Equal(New(WithGoString(``))) {
		t.Fatal("Null Equal = true, want false")
	}

	// Set returns a successful Result wrapping the null String.
	setR := n.Set(`ignored`)
	if setR.HasError() {
		t.Fatalf("Null Set HasError = true, want false")
	}
	setPayload, ok := setR.Payload().(Interface)
	if !ok || !setPayload.IsNull() {
		t.Fatalf("Null Set payload not the null String: %T", setR.Payload())
	}

	// Split returns a successful Result wrapping an empty Array.
	splitR := n.Split(`,`)
	if splitR.HasError() {
		t.Fatal("Null Split HasError = true, want false")
	}
	if _, ok := splitR.Payload().(Array.Interface); !ok {
		t.Fatalf("Null Split payload not an Array: %T", splitR.Payload())
	}

	// Every new-value op returns a successful Result wrapping the null String.
	for name, r := range map[string]interface{ HasError() bool }{
		`Concat`:  n.Concat(New(WithGoString(`x`))),
		`Replace`: n.Replace(`a`, `b`),
		`Upper`:   n.Upper(),
		`Lower`:   n.Lower(),
		`Trim`:    n.Trim(),
	} {
		if r.HasError() {
			t.Fatalf("Null %s HasError = true, want false", name)
		}
	}

	concatPayload := n.Concat(New(WithGoString(`x`))).Payload().(Interface)
	if !concatPayload.IsNull() {
		t.Fatal("Null Concat payload is not null")
	}
	for _, r := range []Result.Interface{n.Replace(`a`, `b`), n.Upper(), n.Lower(), n.Trim()} {
		if p, ok := r.Payload().(Interface); !ok || !p.IsNull() {
			t.Fatalf("Null new-value payload is not null: %T", r.Payload())
		}
	}
}
