package main

func Pic(dx, dy int) [][]uint8 {
	array := make([][]uint8, dy)

	for i := range array {
		slice := make([]uint8, dx)

		for j := range slice {
			dataImage := setImage(i, j)
			slice[j] = dataImage
		}

		array[i] = slice
	}

	return array
}

func setImage(x, y int) uint8 {
	// return uint8(x * y)
	return uint8(x ^ y)
}
