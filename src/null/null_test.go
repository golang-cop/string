package NullString

import (
	"testing"
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
