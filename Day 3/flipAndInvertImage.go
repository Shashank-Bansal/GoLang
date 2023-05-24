package main

func flipAndInvertImage(image [][]int) [][]int {
	n := len(image)

	for _, arr := range image {
		for i := 0; i < (n+1)/2; i++ {
			arr[i], arr[n-1-i] = invert(arr[n-1-i]), invert(arr[i])
		}
	}

	return image
}

func invert(n int) int {
	if n == 1 {
		return 0
	}

	return 1
}
