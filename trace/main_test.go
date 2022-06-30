package main

import (
	"bytes"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("Return from New should not be nil")
	} else {
		tracer.Trace("Hello trace.")
		if buf.String() != "Hello trace.\n" {
			t.Errorf("Trace should not write '%s'.", buf.String())
		}
	}
}

func TestSkip(t *testing.T) {
	t.Skip("deneme")
}
func TestError(t *testing.T) {
	time.Sleep(1)
	t.Fatalf("denme")

}
