/*
 * @lc app=leetcode id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 */
package leetcode

var digitAlphaMap = map[byte][]byte{
	'2': []byte{'a', 'b', 'c'},
	'3': []byte{'d', 'e', 'f'},
	'4': []byte{'g', 'h', 'i'},
	'5': []byte{'j', 'k', 'l'},
	'6': []byte{'m', 'n', 'o'},
	'7': []byte{'p', 'q', 'r', 's'},
	'8': []byte{'t', 'u', 'v'},
	'9': []byte{'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}

	var result []string

	var solution func(string, int, []byte)
	solution = func(digits string, curPos int, curSet []byte) {
		if len(digits) == 0 {
			result = append(result, string(curSet))
			return
		}

		letters := digitAlphaMap[digits[0]]
		for _, c := range letters {
			curSet[curPos] = c
			solution(digits[1:], curPos+1, curSet)
		}
	}

	set := make([]byte, len(digits))
	solution(digits, 0, set)

	return result
}
