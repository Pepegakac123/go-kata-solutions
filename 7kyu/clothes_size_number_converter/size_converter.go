// Description
// You have clothes international size as text (xs, s, xxl).
// Your goal is to return European number size as a number from that size.

// Notice that there is arbitrary amount of modifiers (x), excluding m size, and input can be any string.

// Linearity
// Base size for medium (m) is 38.
// (then, small (s) is 36, and large (l) is 40)

// The scale is linear; modifier x continues that by adding or subtracting 2 from resulting size.

// (For sizes where x modifier makes total size negative, treat that as OK, and return negative size)

// Invalid cases
// Function should handle wrong/invalid sizes.

// Valid input has any base size (s/m/l) and any amount of modifiers (x) before base size (like xxl).
// Notice that you cannot apply modifier for m size, consider that case as invalid!
// Anything else is disallowed and should be considered as invalid (None for Python, 0, false for Go, null for JavaScript).

// Invalid examples: xxx, sss, xm, other string

// Valid Examples
// xs: 34
// s: 36
// m: 38
// l: 40
// xl: 42
package kata

import "strings"

var sizeMappings = map[string]int{
	"xs": 34,
	"s":  36,
	"m":  38,
	"l":  40,
	"xl": 42,
}

func SizeToNumber(size string) (int, bool) {
	trimmedSize := strings.TrimSpace(size)
	if len(trimmedSize) <= 0 {
		return 0, false
	}
	if trimmedSize == "" {
		return 0, false
	}
	if val, ok := sizeMappings[trimmedSize]; ok {
		return val, true
	}
	var total int = 0
	var valToAddOrSubtract int = 0
	var hasModifier bool = false
	var hasBaseSize bool = false
	for _, char := range trimmedSize {
		switch char {
		case 'x':
			valToAddOrSubtract += 2
			hasModifier = true
		case 's':
			if !hasModifier {
				return 0, false
			}
			total += sizeMappings["s"] - valToAddOrSubtract
			hasBaseSize = true
		case 'm':
			return 0, false
		case 'l':
			if !hasModifier {
				return 0, false
			}
			total += sizeMappings["l"] + valToAddOrSubtract
			hasBaseSize = true
		default:
			return 0, false
		}
	}
	if !hasBaseSize {
		return 0, false
	}
	if total < 0 && hasBaseSize {
		return total, true
	}
	return total, true
}
