package leetcode

func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}

	triangle := make([][]int, numRows)
	triangle[0] = []int{1}

	for i := 1; i < numRows; i++ {
		tmp := make([]int, i+1)
		triangle[i] = tmp
		tmp[0], tmp[i] = 1, 1
		l := len(triangle[i-1])
		for j := 0; j < l-1; j++ {
			tmp[j+1] = triangle[i-1][j] + triangle[i-1][j+1]
		}
	}

	return triangle
}
