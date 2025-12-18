package kata

import "strconv"

// Given a positive integer n: 0 < n < 1e6, remove the last digit until you're left with a number that is a multiple of three.

// Return n if the input is already a multiple of three, and if no such number exists, return null, a similar empty value, or -1.

func PrevMultOfThree(n int) interface{} {
	if n <= 2 {
		return nil
	}
	if n%3 == 0 {
		return n
	}
	strNum := strconv.Itoa(n)
	newNum, err := strconv.Atoi(strNum[:len(strNum)-1])
	if err != nil {
		return nil
	}
	return PrevMultOfThree(newNum)
}

// More idiomatic solution
func PrevMultOfThreeIdiomatic(n int) interface{} {

	for i := n; i > 0; i /= 10 {
		if i%3 == 0 {
			return i
		}
	}

	return nil
}
