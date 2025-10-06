package logger

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

// captureOutput redirects stdout to a buffer, executes the provided function,
// and returns the captured output as a string.
func captureOutput(f func()) string {
	originalStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = originalStdout

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

func TestInit(t *testing.T) {
	// Ensure logger is nil before Init() to test initialization.
	logger = nil
	Init()
	if logger == nil {
		t.Fatal("logger is nil after Init()")
	}
	if logger.level != InfoLevel {
		t.Errorf("expected default level to be %d (InfoLevel), but got %d", InfoLevel, logger.level)
	}
	if logger.infoLogger == nil || logger.warnLogger == nil || logger.errorLogger == nil {
		t.Error("one or more of the specific loggers were not initialized")
	}
}

func TestSetLevel(t *testing.T) {
	Init() // Guarantees logger is not nil

	testCases := []struct {
		level    LogLevel
		expected LogLevel
	}{
		{WarnLevel, WarnLevel},
		{ErrorLevel, ErrorLevel},
		{InfoLevel, InfoLevel},
	}

	for _, tc := range testCases {
		SetLevel(tc.level)
		if logger.level != tc.expected {
			t.Errorf("for level %d, expected logger.level to be %d, but got %d", tc.level, tc.expected, logger.level)
		}
	}
}

func TestLoggingExecution(t *testing.T) {
	// This test assumes the bug in Warn() is fixed.
	testCases := []struct {
		name         string
		levelToSet   LogLevel
		logFunc      func(string)
		message      string
		shouldLog    bool
		expectedText string // Prefix to check for, e.g., "INFO:"
	}{
		// Current Level: Info -> should log everything
		{"InfoLevel_Logs_Info", InfoLevel, Info, "info message", true, "INFO:"},
		{"InfoLevel_Logs_Warn", InfoLevel, Warn, "warn message", true, "WARN:"},
		{"InfoLevel_Logs_Error", InfoLevel, Error, "error message", true, "ERROR:"},

		// Current Level: Warn -> should log Warn and Error
		{"WarnLevel_Skips_Info", WarnLevel, Info, "info message", false, ""},
		{"WarnLevel_Logs_Warn", WarnLevel, Warn, "warn message", true, "WARN:"},
		{"WarnLevel_Logs_Error", WarnLevel, Error, "error message", true, "ERROR:"},

		// Current Level: Error -> should only log Error
		{"ErrorLevel_Skips_Info", ErrorLevel, Info, "info message", false, ""},
		{"ErrorLevel_Skips_Warn", ErrorLevel, Warn, "warn message", false, ""},
		{"ErrorLevel_Logs_Error", ErrorLevel, Error, "error message", true, "ERROR:"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			Init() // Reset logger state for each test case
			SetLevel(tc.levelToSet)

			output := captureOutput(func() {
				tc.logFunc(tc.message)
			})

			hasLogged := output != ""
			if hasLogged != tc.shouldLog {
				t.Fatalf("expected shouldLog=%v, but it was %v. Output: %q", tc.shouldLog, hasLogged, output)
			}

			if tc.shouldLog {
				if !strings.Contains(output, tc.expectedText) {
					t.Errorf("expected output to contain prefix %q, but got %q", tc.expectedText, output)
				}
				if !strings.Contains(output, tc.message) {
					t.Errorf("expected output to contain message %q, but got %q", tc.message, output)
				}
			}
		})
	}
}

func TestUninitializedLogger(t *testing.T) {
	// Set the global logger to nil to simulate an uninitialized state
	logger = nil

	// Capture output from all log functions
	output := captureOutput(func() {
		Info("this should not appear")
		Warn("this should not appear")
		Error("this should not appear")
	})

	// Expect no output because of the nil check
	if output != "" {
		t.Errorf("expected no output from uninitialized logger, but got: %q", output)
	}
}
