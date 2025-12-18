package kata

// Given an integer n and two other values, build an array of size n filled with these two values alternating.

func Alternate(n int, firstValue string, secondValue string) []string {
	arr := make([]string, n)
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			arr[i] = firstValue
		} else {
			arr[i] = secondValue
		}
	}
	return arr
}
