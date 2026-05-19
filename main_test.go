package main

import (
	"os/exec"
	"testing"
)

// helper function to run your program
func runApp(input string) (string, error) {
	cmd := exec.Command("go", "run", ".", input)
	out, err := cmd.CombinedOutput()
	return string(out), err
}

// ✅ 1. Basic test
func TestHello(t *testing.T) {
	out, err := runApp("Hello")
	if err != nil || len(out) == 0 {
		t.Error("Failed Hello test")
	}
}

// ✅ 2. Mixed case (lower + upper)
func TestMixedCase(t *testing.T) {
	out, err := runApp("abCDxyZ")
	if err != nil || len(out) == 0 {
		t.Error("Failed mixed case test")
	}
}

// ✅ 3. Numbers + space
func TestNumbersAndSpace(t *testing.T) {
	out, err := runApp("hello 42")
	if err != nil || len(out) == 0 {
		t.Error("Failed numbers + space test")
	}
}

// ✅ 4. Special characters
func TestSpecialChars(t *testing.T) {
	out, err := runApp("A@#B")
	if err != nil || len(out) == 0 {
		t.Error("Failed special characters test")
	}
}

// ✅ 5. Complex input (spaces + numbers + symbols + upper/lower)
func TestComplexInput(t *testing.T) {
	out, err := runApp("ab  3@CDX")
	if err != nil || len(out) == 0 {
		t.Error("Failed complex input test")
	}
}

// ✅ 6. Newline handling
func TestNewline(t *testing.T) {
	out, err := runApp("Hello\\nWorld")
	if err != nil || len(out) == 0 {
		t.Error("Failed newline test")
	}
}

// ✅ 7. Double newline
func TestDoubleNewline(t *testing.T) {
	out, err := runApp("Hello\\n\\nWorld")
	if err != nil || len(out) == 0 {
		t.Error("Failed double newline test")
	}
}

// ✅ 8. Empty input (should print a newline)
func TestEmptyInput(t *testing.T) {
	out, err := runApp("")
	if err != nil {
		t.Error("Error on empty input")
	}
	if out != "\n" {
		t.Errorf("Expected newline for empty input, got: %q", out)
	}
}

// ✅ 9. Only newline input
func TestOnlyNewline(t *testing.T) {
	out, err := runApp("\\n")
	if err != nil || len(out) == 0 {
		t.Error("Failed only newline test")
	}
}