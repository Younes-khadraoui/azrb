package cmd

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"azrb/config"
)

func TestViewCommand(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "azrb_test")
	if err != nil {
		t.Fatal("Error creating temp dir:", err)
	}
	defer os.RemoveAll(tmpDir)
	config.StorageDir = tmpDir

	cheatName := "testsheet"
	filePath := filepath.Join(tmpDir, cheatName+".md")

	if err := os.WriteFile(filePath, []byte("## Test Cheat Sheet\nThis is a test."), 0644); err != nil {
		t.Fatal("Error writing file:", err)
	}

	expectedContent := "## Test Cheat Sheet\nThis is a test.\n"

	oldStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal("Error creating pipe:", err)
	}
	os.Stdout = w

	rootCmd.SetArgs([]string{cheatName})
	if err := rootCmd.Execute(); err != nil {
		t.Fatal("Error executing command:", err)
	}

	w.Close()
	outputBytes, err := io.ReadAll(r)
	if err != nil {
		t.Fatal("Error reading output:", err)
	}
	os.Stdout = oldStdout

	output := strings.TrimSpace(string(outputBytes))
	expected := strings.TrimSpace(expectedContent)
	if output != expected {
		t.Errorf("Expected output:\n%s\nGot:\n%s", expected, output)
	}
}
