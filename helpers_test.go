package walker_test

import (
	"fmt"
	"testing"
	"testing/fstest"
)

// createTestFs creates a test filesystem for the tests.
func createTestFs() fstest.MapFS {
	return fstest.MapFS{
		"myRoot/bob.md":         {},
		"myRoot/ted.md":         {},
		"myRoot/zoe.txt":        {},
		"myRoot/_tea.txt":       {},
		"myRoot/folder/lea.txt": {},
		"myRoot/folder/_tea.md": {},
		"myRoot/folder/kim.jpg": {},
	}
}

// assertLen asserts the length of the items.
func assertLen[K any](t *testing.T, items []K, expected int) {
	if len(items) != expected {
		t.Fatalf("expected %d items, got %d", expected, len(items))
	}
}

// assertStringType asserts the type of the string via Sprintf.
func assertStringType[K any](t *testing.T, got K, expected string) {
	if fmt.Sprintf("%T", got) != expected {
		t.Errorf("expected %s, got %T", expected, got)
	}
}

// assertNoError asserts the error is nil.
func assertNoError(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

// assertEqual asserts the equality of the values.
func assertEqual[K comparable](t *testing.T, got, expected K) {
	if got != expected {
		t.Fatalf("expected %v, got %v", expected, got)
	}
}
