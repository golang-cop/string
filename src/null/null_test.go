package NullString

import (
	"testing"

	String "github.com/go-composites/string/src"
)

func TestNewIsNull(t *testing.T) {
	s := New()
	if !s.IsNull() {
		t.Fatal("NullString.IsNull() = false, want true")
	}
}

func TestToGoStringEmpty(t *testing.T) {
	if got := New().ToGoString(); got != `` {
		t.Fatalf("NullString.ToGoString() = %q, want empty", got)
	}
}

func TestSplitEmpty(t *testing.T) {
	r := New().Split(` `)
	payload, ok := r.Payload().([]string)
	if !ok {
		t.Fatalf("Split payload is not []string: %T", r.Payload())
	}
	if len(payload) != 0 {
		t.Fatalf("Split payload len = %d, want 0", len(payload))
	}
}

func TestSetNotImplemented(t *testing.T) {
	r := New().Set(`x`)
	if got := r.Error().Message(); got == `` {
		t.Fatal("Set on NullString should attach a non-empty error message")
	}
}

func TestLengthZero(t *testing.T) {
	if got := New().Length(); got != 0 {
		t.Fatalf("NullString.Length() = %d, want 0", got)
	}
}

func TestContainsFalse(t *testing.T) {
	if New().Contains(`x`) {
		t.Fatal("NullString.Contains() = true, want false")
	}
}

func TestEqualFalse(t *testing.T) {
	if New().Equal(String.New()) {
		t.Fatal("NullString.Equal() = true, want false")
	}
}

func TestStartsWithFalse(t *testing.T) {
	if New().StartsWith(`x`) {
		t.Fatal("NullString.StartsWith() = true, want false")
	}
}

func TestEndsWithFalse(t *testing.T) {
	if New().EndsWith(`x`) {
		t.Fatal("NullString.EndsWith() = true, want false")
	}
}

func TestFormatReturnsNull(t *testing.T) {
	r := New().Format(`x`)
	if r.HasError() {
		t.Fatal("NullString.Format HasError = true, want false")
	}
	payload, ok := r.Payload().(String.Interface)
	if !ok || !payload.IsNull() {
		t.Fatalf("NullString.Format payload is not the null String: %T", r.Payload())
	}
}

func TestNewValueOpsReturnNull(t *testing.T) {
	n := New()
	results := map[string]struct {
		r String.Interface
	}{}
	for name, r := range map[string]interface {
		HasError() bool
		Payload() interface{}
	}{
		`Concat`:  n.Concat(String.New(String.WithGoString(`x`))),
		`Replace`: n.Replace(`a`, `b`),
		`Upper`:   n.Upper(),
		`Lower`:   n.Lower(),
		`Trim`:    n.Trim(),
	} {
		if r.HasError() {
			t.Fatalf("NullString.%s HasError = true, want false", name)
		}
		payload, ok := r.Payload().(String.Interface)
		if !ok || !payload.IsNull() {
			t.Fatalf("NullString.%s payload is not the null String: %T", name, r.Payload())
		}
		results[name] = struct{ r String.Interface }{payload}
	}
	if len(results) != 5 {
		t.Fatalf("covered %d new-value ops, want 5", len(results))
	}
}
