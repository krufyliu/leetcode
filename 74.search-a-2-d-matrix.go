/*
 * @lc app=leetcode id=74 lang=golang
 *
 * [74] Search a 2D Matrix
 *
 * https://leetcode.com/problems/search-a-2d-matrix/description/
 *
 * algorithms
 * Medium (34.66%)
 * Total Accepted:    212.4K
 * Total Submissions: 611.9K
 * Testcase Example:  '[[1,3,5,7],[10,11,16,20],[23,30,34,50]]\n3'
 *
 * Write an efficient algorithm that searches for a value in an m x n matrix.
 * This matrix has the following properties:
 *
 *
 * Integers in each row are sorted from left to right.
 * The first integer of each row is greater than the last integer of the
 * previous row.
 *
 *
 * Example 1:
 *
 *
 * Input:
 * matrix = [
 * ⁠ [1,   3,  5,  7],
 * ⁠ [10, 11, 16, 20],
 * ⁠ [23, 30, 34, 50]
 * ]
 * target = 3
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input:
 * matrix = [
 * ⁠ [1,   3,  5,  7],
 * ⁠ [10, 11, 16, 20],
 * ⁠ [23, 30, 34, 50]
 * ]
 * target = 13
 * Output: false
 *
 */
package leetcode

func searchMatrix(matrix [][]int, target int) bool {
	rows := len(matrix)
	if rows == 0 {
		return false
	}
	cols := len(matrix[0])
	l := 0
	r := rows*cols - 1

	if r < 0 {
		return false
	}

	for l <= r {
		m := (l + r) / 2
		t := matrix[m/cols][m%cols]
		if t == target {
			return true
		} else if t < target {
			l = l + 1
		} else {
			r = r - 1
		}
	}

	return false
}
