package primeGeneration

import "fmt"

/**
 * 找出前N个质数, N > 3
 */
func primeGeneration(n int) int {
	prime := make([]int, n)
	var gap = 2
	var count = 3
	var maybePrime = 5
	var i int
	var isPrime bool

	/* 注意：2, 3, 5 是质数 */
	prime[0] = 2
	prime[1] = 3
	prime[2] = 5

	for count < n {
		maybePrime += gap
		gap = 6 - gap
		isPrime = true
		for i = 2; prime[i]*prime[i] <= maybePrime && isPrime; i++ {
			if maybePrime%prime[i] == 0 {
				isPrime = false
			}
			if isPrime {
				prime[count] = maybePrime
				count++
			}
		}
	}

	fmt.Printf("\nFirst %d Prime Numbers are :\n", count)

	for i = 0; i < count; i++ {
		if i%10 == 0 {
			fmt.Println()
		}
		fmt.Printf("%5d", prime[i])
	}
	fmt.Println()
	return 0
}
