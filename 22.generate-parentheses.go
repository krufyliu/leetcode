/*
 * @lc app=leetcode id=22 lang=golang
 *
 * [22] Generate Parentheses
 */
package leetcode

func generateParenthesis(n int) []string {
	if n == 0 {
		return nil
	}

	var result []string
	set := make([]byte, 2*n)
	// 递归填充对应的括号
	var setFill func(leftParenthesesNum, rightParenthesesNum, curPos int, curSet []byte)
	setFill = func(lpn, rpn, curPos int, curSet []byte) {
		if curPos == 2*n {
			result = append(result, string(curSet))
			return
		}

		if lpn == rpn { // 左右括号数相等时当前位置只能填充'('
			curSet[curPos] = '('
			setFill(lpn+1, rpn, curPos+1, curSet)
		} else if lpn == n { // 左括号数已经达到最大值当前只能填充')'
			curSet[curPos] = ')'
			setFill(lpn, rpn+1, curPos+1, curSet)
		} else {
			curSet[curPos] = '('
			setFill(lpn+1, rpn, curPos+1, curSet)
			curSet[curPos] = ')'
			setFill(lpn, rpn+1, curPos+1, curSet)
		}
	}

	setFill(0, 0, 0, set)

	return result
}
