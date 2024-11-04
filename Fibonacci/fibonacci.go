package Fibonacci

// fibonacci numbers  (0, 1, 1, 2, 3, 5, 8, 13, ...)

func Fibonacci() func() int {
	f1 := 0 // n-1
	f2 := 1 // n-2
	n := 0
	return func() int {
		result := 0

		if n > 1 {
			result = f1 + f2

			f1 = f2
			f2 = result
		} else {
			result = n
		}

		n++
		return result
	}
}

/*
	Correction

	func Fibonacci() func() int {
		f1, f2 := 0, 1
		return func() int {
			result := f1
			f1, f2 = f2, f1 + f2
			return result
		}
	}
*/
