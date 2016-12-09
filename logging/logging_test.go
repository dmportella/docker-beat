package logging

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

var (
	actualTrace   string
	actualInfo    string
	actualWarning string
	actualError   string

	traceBuffer   *bytes.Buffer
	infoBuffer    *bytes.Buffer
	warningBuffer *bytes.Buffer
	errorBuffer   *bytes.Buffer
)

func TestMain(m *testing.M) {
	traceBuffer = bytes.NewBufferString(actualTrace)
	infoBuffer = bytes.NewBufferString(actualInfo)
	warningBuffer = bytes.NewBufferString(actualWarning)
	errorBuffer = bytes.NewBufferString(actualError)

	retCode := m.Run()

	os.Exit(retCode)
}

func TestLoggingInit(t *testing.T) {
	Init(traceBuffer, infoBuffer, warningBuffer, errorBuffer)

	if Trace == nil || Info == nil || Warning == nil || Error == nil {
		t.Error("Logging init failed to set the loggers")
	}
}

func TestLoggingTrace(t *testing.T) {
	Trace.Println("Hello")

	if !strings.Contains(traceBuffer.String(), "Hello") {
		t.Error("Failed logging message to trace output")
	}
}

func TestLoggingInfo(t *testing.T) {
	Info.Println("Hello")

	if !strings.Contains(infoBuffer.String(), "Hello") {
		t.Error("Failed logging message to info output")
	}
}

func TestLoggingWarning(t *testing.T) {
	Warning.Println("Hello")

	if !strings.Contains(warningBuffer.String(), "Hello") {
		t.Error("Failed logging message to warning output")
	}
}

func TestLoggingError(t *testing.T) {
	Error.Println("Hello")

	if !strings.Contains(errorBuffer.String(), "Hello") {
		t.Error("Failed logging message to error output")
	}
}
