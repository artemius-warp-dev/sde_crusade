package main

import (
	"os"
	"testing"
)

func TestReadDir(t *testing.T) {
	testDir := "testdata"
	if err := os.Mkdir(testDir, 0755); err != nil {
		t.Fatal("Error creating test directory: ", err)
	}
	defer os.RemoveAll(testDir)

	testFiles := map[string]string{
		"A_ENV": "123",
		"B_ENV": "456",
	}

	for fileName, content := range testFiles {
		filePath := testDir + "/" + fileName
		if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
			t.Fatalf("Error creatring test file %s %v", filePath, err)
		}
	}

	envs, err := ReadDir(testDir)

	if err != nil {
		t.Fatalf("ReadDir return error: %v", err)
	}

	expectedEnvs := map[string]string{
		"A_ENV": "123",
		"B_ENV": "456",
	}
	for k, v := range expectedEnvs {
		if envs[k] != v {
			t.Errorf("Expected %s=%s, got %s=%s", k, v, k, envs[k])
		}
	}
}

func TestRunCmd(t *testing.T) {
	cmd := []string{"./hook/hook"}
	env := map[string]string{"ENV_VAR": "value"}

	exitCode := RunCmd(cmd, env)

	if exitCode != 0 {
		t.Errorf("Expected exit code 0, got %d", exitCode)
	}
}

func TestRunCmdWithError(t *testing.T) {
	cmd := []string{"non_existing_command"}
	env := map[string]string{"ENV_VAR": "value"}

	exitCode := RunCmd(cmd, env)
	if exitCode == 0 {
		t.Errorf("Expected non-zero exit code, got %d", exitCode)
	}
}
