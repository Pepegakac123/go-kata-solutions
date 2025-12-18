package kata_test

import (
	. "go-kata-solutions/7kyu/clothes_size_number_converter"
	"testing"
)

func TestSizeToNumber(t *testing.T) {
	tests := []struct {
		name          string
		size          string
		expectedValue int
		expectedOk    bool
	}{
		// Basic tests
		{"Small size", "s", 36, true},
		{"Medium size", "m", 38, true},
		{"Large size", "l", 40, true},
		{"Extra large", "xl", 42, true},
		{"Extra small", "xs", 34, true},

		// Extra modifier tests
		{"3-extra small", "xxxs", 30, true},
		{"3-extra large", "xxxl", 46, true},

		// Invalid/wrong sizes/input
		{"Blank string is invalid", "", 0, false},
		{"Cannot apply modifier for medium size (xm)", "xm", 0, false},
		{"Cannot apply modifier for medium size (xxxm)", "xxxm", 0, false},
		{"No base size provided", "xxxx", 0, false},
		{"Only one base size is allowed", "ssss", 0, false},
		{"Any other text is invalid", "hello world", 0, false},
		{"Two bases (sm)", "sm", 0, false},
		{"Two bases (ml)", "ml", 0, false},
		{"Two bases (lm)", "lm", 0, false},
		{"Base should be last (lx)", "lx", 0, false},
	}

	for _, tt := range tests {
		// Używamy t.Run do stworzenia sub-testów
		t.Run(tt.name, func(t *testing.T) {
			// t.Parallel() pozwoli na równoległe uruchamianie testów
			t.Parallel()

			actualValue, actualOk := SizeToNumber(tt.size)

			if actualValue != tt.expectedValue {
				t.Errorf("SizeToNumber(%q) value = %d; want %d", tt.size, actualValue, tt.expectedValue)
			}

			if actualOk != tt.expectedOk {
				t.Errorf("SizeToNumber(%q) ok = %v; want %v", tt.size, actualOk, tt.expectedOk)
			}
		})
	}
}
