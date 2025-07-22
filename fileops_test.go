package main

import (
	"os"
	"testing"
)

func TestSaveDataAtomic_Success(t *testing.T) {
	testFile := "test_atomic.txt"
	testData := []byte("Hello, atomic world!")

	defer os.Remove(testFile)
	os.Remove(testFile)

	err := SaveDataAtomic(testFile, testData)
	if err != nil {
		t.Fatalf("SaveDataAtomic failed: %v", err)
	}

	readData, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	if string(readData) != string(testData) {
		t.Errorf("File content mismatch. Expected %s, Got %s", testData, readData)
	}
}

func TestSaveDataAtomic_OverwriteExisting(t *testing.T) {
	testFile := "test_overwrite.txt"
	defer os.Remove(testFile)

	initialData := []byte("initial data")
	err := os.WriteFile(testFile, initialData, 0644)
	if err != nil {
		t.Fatalf("Failed to create initial file: %v", err)
	}

	newData := []byte("new atomic data")
	err = SaveDataAtomic(testFile, newData)
	if err != nil {
		t.Fatalf("SaveDataAtomic failed: %v", err)
	}

	readData, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	if string(readData) != string(newData) {
		t.Errorf("Overwrite failed. Expected: %s, Got: %s", newData, readData)
	}
}

func TestSaveDataAtomic_InvalidPath(t *testing.T) {
	invalidPath := "this/directory/does/not/exist/file.txt"
	testData := []byte("test data")

	err := SaveDataAtomic(invalidPath, testData)
	if err == nil {
		t.Error("Expected error for non-existent directory path, got nil")
	}

	t.Logf("Got expected error: %v", err)
}

func TestRandomInt(t *testing.T) {
	val1 := randomInt()
	val2 := randomInt()

	if val1 < 0 || val1 >= 999999 {
		t.Errorf("randomInt returned value out of range: %d", val1)
	}

	if val2 < 0 || val2 >= 999999 {
		t.Errorf("randomInt returned value out of range: %d", val2)
	}

	t.Logf("Generated values: %d, %d", val1, val2)
}
