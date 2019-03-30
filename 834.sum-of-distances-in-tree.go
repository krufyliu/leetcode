/*
 * @lc app=leetcode id=834 lang=golang
 *
 * [834] Sum of Distances in Tree
 *
 * https://leetcode.com/problems/sum-of-distances-in-tree/description/
 *
 * algorithms
 * Hard (38.29%)
 * Total Accepted:    6.8K
 * Total Submissions: 17.5K
 * Testcase Example:  '6\n[[0,1],[0,2],[2,3],[2,4],[2,5]]'
 *
 * An undirected, connected tree with N nodes labelled 0...N-1 and N-1 edges
 * are given.
 *
 * The ith edge connects nodes edges[i][0] and edges[i][1] together.
 *
 * Return a list ans, where ans[i] is the sum of the distances between node i
 * and all other nodes.
 *
 * Example 1:
 *
 *
 * Input: N = 6, edges = [[0,1],[0,2],[2,3],[2,4],[2,5]]
 * Output: [8,12,6,10,10,10]
 * Explanation:
 * Here is a diagram of the given tree:
 * ⁠ 0
 * ⁠/ \
 * 1   2
 * ⁠  /|\
 * ⁠ 3 4 5
 * We can see that dist(0,1) + dist(0,2) + dist(0,3) + dist(0,4) + dist(0,5)
 * equals 1 + 1 + 2 + 2 + 2 = 8.  Hence, answer[0] = 8, and so on.
 *
 *
 * Note: 1 <= N <= 10000
 *
 */
package leetcode

func sumOfDistancesInTree(N int, edges [][]int) []int {
	dis := make([]int, N)
	disSon := make([]int, N)
	nums := make([]int, N)
	trees := make([][]int, N)
	visited := make([]int, N)

	reset(nums, 1)

	for _, edge := range edges {
		trees[edge[0]] = append(trees[edge[0]], edge[1])
		trees[edge[1]] = append(trees[edge[1]], edge[0])
	}
	dfs(0, disSon, nums, trees, visited)
	dis[0] = disSon[0]
	reset(visited, 0)
	dfs2(0, trees, nums, visited, dis)
	return dis
}

func dfs(start int, disSon []int, nums []int, trees [][]int, visited []int) {
	visited[start] = 1
	vertex := trees[start]
	for i := 0; i < len(vertex); i++ {
		if visited[vertex[i]] == 1 {
			continue
		}
		dfs(vertex[i], disSon, nums, trees, visited)
		disSon[start] += nums[vertex[i]] + disSon[vertex[i]]
		nums[start] += nums[vertex[i]]
	}
}

func dfs2(start int, trees [][]int, nums []int, visited []int, dis []int) {
	visited[start] = 1
	vertex := trees[start]
	for i := 0; i < len(vertex); i++ {
		if visited[vertex[i]] == 1 {
			continue
		}
		dis[vertex[i]] = dis[start] + len(dis) - 2*nums[vertex[i]]
		dfs2(vertex[i], trees, nums, visited, dis)
	}
}

func reset(s []int, v int) {
	for i := 0; i < len(s); i++ {
		s[i] = v
	}
}
