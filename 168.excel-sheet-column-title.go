/*
 * @lc app=leetcode id=168 lang=golang
 *
 * [168] Excel Sheet Column Title
 */

package leetcode

func convertToTitle(n int) string {
	var title []byte
	for n > 0 {
		n--
		r := byte(n % 26)
		title = append(title, 'A'+r)
		n = n / 26
	}

	// reverse
	l := len(title)
	if l > 1 {
		for i := 0; i < len(title)/2; i++ {
			title[i], title[l-i-1] = title[l-i-1], title[i]
		}
	}
	return string(title)
}
