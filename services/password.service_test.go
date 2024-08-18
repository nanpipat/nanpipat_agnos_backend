package services

import (
	"testing"
)

func TestCalculateSteps(t *testing.T) {
	tests := []struct {
		name     string
		password string
		want     int
	}{
		{"Short password", "aA1", 3},
		{"Valid password", "1337C0d3", 0},
		{"Missing uppercase", "password123", 1},
		{"Missing lowercase", "PASSWORD123", 1},
		{"Missing digit", "Password", 1},
		{"Repeating characters", "AAAbc123", 1},
		{"Edge case - exactly 6 chars", "Aa1234", 0},
		{"Only special characters", ".!.!.!", 3},
		{"Mix of valid and invalid", "AAA111aaa", 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateSteps(tt.password); got != tt.want {
				t.Errorf("calculateSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdjustLength(t *testing.T) {
	tests := []struct {
		name     string
		password string
		want     int
	}{
		{"Too short", "abc", 3},
		{"Correct length", "password", 0},
		{"Exactly 6 chars", "sixchr", 0},
		{"Exactly 20 chars", "12345678901234567890", 0},
		{"Too long by 1", "123456789012345678901", 1},
		{"Too long by 3", "123456789012345678901234", 1},
		{"Very long", "ThisPasswordIsWayTooLongAndNeedsToBeShortened", 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := adjustLength(tt.password); got != tt.want {
				t.Errorf("adjustLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestEnsureCharacterTypes(t *testing.T) {
	tests := []struct {
		name     string
		password string
		want     int
	}{
		{"All types present", "Pass1", 0},
		{"Missing uppercase", "pass1", 1},
		{"Missing lowercase", "PASS1", 1},
		{"Missing digit", "Pass", 1},
		{"Missing all", ".....", 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ensureCharacterTypes(tt.password); got != tt.want {
				t.Errorf("ensureCharacterTypes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveRepeatingChars(t *testing.T) {
	tests := []struct {
		name     string
		password string
		want     int
	}{
		{"No repeats", "abcdef", 0},
		{"One repeat", "aaabcdef", 1},
		{"Multiple repeats", "aaabbbccc", 3},
		{"Edge case - entire string repeats", "aaaaaa", 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeRepeatingChars(tt.password); got != tt.want {
				t.Errorf("removeRepeatingChars() = %v, want %v", got, tt.want)
			}
		})
	}
}
