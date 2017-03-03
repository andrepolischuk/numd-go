package numd

import "strconv"

// Genetive plural test
func isGenitivePlural(num int) bool {
	nn := num % 10
	return (num > 10 && ((num % 100) - ((num % 100) % 10)) / 10 == 1) || (nn == 0 || nn >= 5)
}

// Genetive singular test
func isGenitiveSingular(num int) bool {
	return num % 10 >= 2
}

// Numeral decline
func Decline(num int, args ...string) string {
	length := len(args)

	if length < 2 {
		return ""
	}

	value := strconv.Itoa(num)

	if num < 0 {
		num = -num
	}

	switch {
		case num != 1 && length == 2:
			return value + " " + args[1]
		case isGenitivePlural(num):
			return value + " " + args[2]
		case isGenitiveSingular(num):
			return value + " " + args[1]
		default:
			return value + " " + args[0]
	}
}
