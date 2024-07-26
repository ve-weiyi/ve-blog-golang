package slicex

import (
	"reflect"
	"testing"
)

func TestAddIds(t *testing.T) {
	// Test adding a new key to an empty slice
	ids := []string{}
	key := "123"
	expected := []string{"123"}
	result := AppendSlice(ids, key)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

	// Test adding a key that already exists in the slice
	ids = []string{"123", "456"}
	key = "123"
	expected = []string{"123", "456"}
	result = AppendSlice(ids, key)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

	// Test adding a new key to a non-empty slice
	ids = []string{"123", "456"}
	key = "789"
	expected = []string{"123", "456", "789"}
	result = AppendSlice(ids, key)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestIsExistSlice(t *testing.T) {
	// Test a key that exists in the slice
	ids := []string{"123", "456"}
	key := "123"
	expected := true
	result := IsExistSlice(ids, key)
	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

	// Test a key that does not exist in the slice
	ids = []string{"123", "456"}
	key = "789"
	expected = false
	result = IsExistSlice(ids, key)
	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
