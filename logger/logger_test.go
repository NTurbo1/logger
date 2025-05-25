package logger

import (
	"bytes"
	"regexp"
	"testing"
	"strings"
)

const timestampFormatRegexp string = `^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(?:\.\d+)?(Z|[\+\-]\d{2}:\d{2})$`

func TestLogger_Info(t *testing.T) {
	resetSingleLogger() // Logger is a singleton

	var buf bytes.Buffer
	logMessage := "Info log message is logged!"
	logger := GetLoggerInstance(&buf);

	logger.Info(logMessage)

	actual := buf.String()

	testTimestampFormat(t, actual)
	expected := INFO.String() + ": " + logMessage

	if !strings.Contains(actual, expected) {
		t.Errorf("Expected log message: %s\nNot found in %s", expected, actual)
	}
}

func TestLogger_Debug(t *testing.T) {
	resetSingleLogger() // Logger is a singleton

	var buf bytes.Buffer
	logMessage := "Debug log message is logged!"
	logger := GetLoggerInstance(&buf);

	logger.Debug(logMessage)

	actual := buf.String()

	testTimestampFormat(t, actual)
	expected := DEBUG.String() + ": " + logMessage

	if !strings.Contains(actual, expected) {
		t.Errorf("Expected log message: %s\nNot found in %s", expected, actual)
	}
}

func TestLogger_Error(t *testing.T) {
	resetSingleLogger() // Logger is a singleton

	var buf bytes.Buffer
	logMessage := "Error log message is logged!"
	logger := GetLoggerInstance(&buf);

	logger.Error(logMessage)

	actual := buf.String()

	testTimestampFormat(t, actual)
	expected := ERROR.String() + ": " + logMessage

	if !strings.Contains(actual, expected) {
		t.Errorf("Expected log message: %s\nNot found in %s", expected, actual)
	}
}

func TestLogger_Warn(t *testing.T) {
	resetSingleLogger() // Logger is a singleton

	var buf bytes.Buffer
	logMessage := "Warn log message is logged!"
	logger := GetLoggerInstance(&buf);

	logger.Warn(logMessage)

	actual := buf.String()

	testTimestampFormat(t, actual)
	expected := WARN.String() + ": " + logMessage

	if !strings.Contains(actual, expected) {
		t.Errorf("Expected log message: %s\nNot found in %s", expected, actual)
	}
}

func testTimestampFormat(t *testing.T, output string) {
	// Extract timestamp
    start := strings.Index(output, "[")
    end := strings.Index(output, "]")
    if start == -1 || end == -1 || start >= end {
        t.Fatalf("Failed to extract timestamp from output: %q", output)
    }

    timestamp := output[start+1 : end]

    rfc3339 := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(?:\.\d+)?(Z|[\+\-]\d{2}:\d{2})$`)

    if !rfc3339.MatchString(timestamp) {
        t.Errorf("Timestamp doesn't match RFC3339:\nGot: %q\nFull line: %s", timestamp, output)
    }
}
