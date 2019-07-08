/*
 * @lc app=leetcode id=371 lang=golang
 *
 * [371] Sum of Two Integers
 */

package leetcode

func getSum(a int, b int) int {
	var sum int
	var currentBit, sumBit, carryBit int
	currentBit = 1
	// var bits = 1
	var aBit, bBit int

	for currentBit > 0 {
		aBit = a & currentBit
		bBit = b & currentBit
		sumBit = aBit ^ bBit ^ carryBit
		carryBit = (aBit & bBit) | (aBit & carryBit) | (bBit & carryBit)
		carryBit <<= 1
		sum |= sumBit
		currentBit <<= 1
		// bits++
	}

	aBit = a & currentBit
	bBit = b & currentBit
	sumBit = aBit ^ bBit ^ carryBit
	sum |= sumBit

	return sum
}
