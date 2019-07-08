/*
 * @lc app=leetcode id=190 lang=golang
 *
 * [190] Reverse Bits
 */
package leetcode

func reverseBits(num uint32) uint32 {
	var res uint32
	count := 0

	for count < 32 {
		b := num & 1
		res <<= 1
		res |= b
		count++

		num >>= 1
	}

	return res
}
