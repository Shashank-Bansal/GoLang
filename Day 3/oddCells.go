package main

func oddCells(m int, n int, indices [][]int) int {
	count := 0
	arr := make([][]int, m, n)

	for i := 0; i < len(indices); i++ {
		for j := 0; j < n; j++ {
			arr[indices[i][0]][j]++
		}

		for j := 0; j < m; j++ {
			arr[j][indices[i][1]]++
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if arr[i][j]%2 != 0 {
				count++
			}
		}
	}

	return count
}
