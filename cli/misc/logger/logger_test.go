package logger

import (
	"bytes"
	"errors"
	"log/slog"
	"testing"

	"github.com/fatih/color"
	"github.com/stretchr/testify/assert"
)

// Logs at the level of a released bbe, which leaves out debug messages, into a buffer
func captureReleaseLogs(t *testing.T) *bytes.Buffer {
	output := &bytes.Buffer{}
	previous := slog.Default()
	noColor := color.NoColor
	color.NoColor = true
	slog.SetDefault(slog.New(&plainTextHandler{writer: output, level: slog.LevelInfo}))
	t.Cleanup(func() {
		slog.SetDefault(previous)
		color.NoColor = noColor
	})
	return output
}

func Test_Error_ShowsTheError_WithoutAMessage(t *testing.T) {
	output := captureReleaseLogs(t)

	Error("", errors.New("UPGRADE FAILED: post-upgrade hooks failed"))

	assert.Equal(t, "UPGRADE FAILED: post-upgrade hooks failed\n", output.String())
}

func Test_Error_ShowsTheError_AfterTheMessage(t *testing.T) {
	output := captureReleaseLogs(t)

	Error("Failed to update BBE packages", errors.New("permission denied"))

	assert.Equal(t, "Failed to update BBE packages: permission denied\n", output.String())
}

func Test_Error_ShowsTheMessage_WithoutAnError(t *testing.T) {
	output := captureReleaseLogs(t)

	Error("Error while determining IP address", nil)

	assert.Equal(t, "Error while determining IP address\n", output.String())
}
