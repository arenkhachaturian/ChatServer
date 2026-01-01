package trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)

	if tracer == nil {
		t.Fatal("return from New should not be nil")
	}

	traceMsg := "Hello trace package."
	expected := traceMsg + "\n"

	tracer.Trace(traceMsg)

	if got := buf.String(); got != expected {
		t.Errorf("Trace wrote: %q, expected: %q", got, expected)
	}
}

func TestOff(t *testing.T) {
	var silentTracer Tracer = Off()
	silentTracer.Trace("something")
}
