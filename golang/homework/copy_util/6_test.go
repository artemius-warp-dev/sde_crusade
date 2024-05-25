package copy_util

import (
	"os"
	"testing"
	"fmt"
)

func TestCopy(t *testing.T) {

	fromFile, err := os.CreateTemp("", "source")
	if err != nil {
		t.Fatalf("Failed to create temperary source file")
	}
	defer os.Remove(fromFile.Name())
	defer fromFile.Close()

	content := []byte("Hello world")

	if _, err := fromFile.Write(content); err != nil {
		t.Fatalf("Failed to write content to source file")
	}

	toFile, err := os.CreateTemp("", "dest")
	if err != nil {
		t.Fatalf("Failed to create temperary source file")
	}
	defer os.Remove(toFile.Name())
	defer fromFile.Close()
	fmt.Println(fromFile.Name())
	err = Copy(fromFile.Name(), toFile.Name(), 0, 0)
	
	if err != nil {
		t.Fatalf("Copy failed: %v", err)
	}

	copiedContent, err := os.ReadFile(toFile.Name())

	if err != nil {
		t.Fatalf("Failed read from copied file: %v", err)
	}

	if string(copiedContent) != string(content) {
		t.Errorf("Copied content does not match source content")
	}
}
