package helpers

import (
	"testing"
)

func TestCleanText(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Remove extra spaces and duplicate words", "This  is  is a  test", "This is a test"},
		{"Remove duplicate punctuation", "Yes!!", "Yes!"},
		{"Remove duplicate words and punctuation", "good good morning!!", "good morning!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CleanText(tt.input); got != tt.expected {
				t.Errorf("%s: CleanText() = %v, want %v", tt.name, got, tt.expected)
			}
		})
	}
}

func TestRemoveDuplicateWords(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Consecutive duplicate", "This is is a test", "This is a test"},
		{"No duplicate", "Hello world", "Hello world"},
		{"All duplicates", "Hey Hey Hey", "Hey"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveDuplicateWords(tt.input); got != tt.expected {
				t.Errorf("%s: RemoveDuplicateWords() = %v, want %v", tt.name, got, tt.expected)
			}
		})
	}
}

func TestRemoveDuplicatePunc(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Multiple exclamation marks", "Wow!!!", "Wow!"},
		{"Mixed punctuation", "Hello??!!", "Hello?!"},
		{"No duplicates", "Greetings!", "Greetings!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveDuplicatePunctuation(tt.input); got != tt.expected {
				t.Errorf("%s: RemoveDuplicatePunc() = %v, want %v", tt.name, got, tt.expected)
			}
		})
	}
}
