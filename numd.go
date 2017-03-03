package numd

import "strconv"

// Genetive plural test
func isGenitivePlural(num int) bool {
	nn := num % 10
	return (num > 10 && ((num%100)-((num%100)%10))/10 == 1) || (nn == 0 || nn >= 5)
}

// Genetive singular test
func isGenitiveSingular(num int) bool {
	return num%10 >= 2
}

// Numeral variant index
func numeralIndexOf(num int) int {
	if num < 0 {
		num = -num
	}
	switch {
		case isGenitivePlural(num):
			return 2
		case isGenitiveSingular(num):
			return 1
		default:
			return 0
	}
}

// Numeral decline
func Decline(num int, args ...string) string {
	if len(args) != 3 {
		return ""
	}
	return strconv.Itoa(num) + " " + args[numeralIndexOf(num)]
}
