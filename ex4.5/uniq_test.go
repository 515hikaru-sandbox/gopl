package uniq

import (
	"testing"
)

func TestAlreadyUniqStrings(t *testing.T) {
	a := []string{"foo", "bar"}
	result := Uniq(a)
	expected := []string{"foo", "bar"}
	if len(result) != len(expected) {
		t.Errorf("Expected slice length=%d, got=%d", len(expected), len(result))
	}
	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("Expected value is %s, but got=%s", expected[i], result[i])
		}
	}
}

func TestNotUniqStrings(t *testing.T) {
	a := []string{"foo", "bar", "bar"}
	result := Uniq(a)
	expected := []string{"foo", "bar"}
	if len(result) != len(expected) {
		t.Errorf("Expected slice length=%d, got=%d", len(expected), len(result))
	}
	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("Expected value is %s, but got=%s", expected[i], result[i])
		}
	}
}

func TestNotNearStrings(t *testing.T) {
	a := []string{"foo", "bar", "foo", "bar"}
	result := Uniq(a)
	expected := []string{"foo", "bar", "foo", "bar"}
	if len(result) != len(expected) {
		t.Errorf("Expected slice length=%d, got=%d", len(expected), len(result))
	}
	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("Expected value is %s, but got=%s", expected[i], result[i])
		}
	}
}

func TestThreeUniqStrings(t *testing.T) {
	a := []string{"foo", "bar", "bar", "bar", "foo"}
	result := Uniq(a)
	expected := []string{"foo", "bar", "foo"}
	if len(result) != len(expected) {
		t.Errorf("Expected slice length=%d, got=%d", len(expected), len(result))
	}
	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("Expected value is %s, but got=%s", expected[i], result[i])
		}
	}
}
