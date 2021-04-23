package dynamic_programming

/*
求斐波拉契数列
 */

// 自顶向下的备忘录法
func FibonacciUp2Down(n int) int {
	if n <= 0 {
		return n
	}
	Memo := make([]int, n+1)
	return fib(n, Memo)
}

func fib(n int, Memo []int) int {
	//如果已经求出了fib（n）的值直接返回
	if Memo[n] > 0 {
		return Memo[n]
	}
	//否则将求出的值保存在Memo备忘录中。
	if n <= 2 {
		Memo[n] = 1
	} else {
		Memo[n] = fib(n-1, Memo) + fib(n-2, Memo)
	}
	return Memo[n]
}

// 自底向上的动态规划
func FibonacciDown2Up(n int) int {
	if n <= 0 {
		return n
	}
	Memo := make([]int, n+1)
	Memo[0] = 0
	Memo[1] = 1
	for i := 2; i <= n; i++ {
		Memo[i] = Memo[i-1] + Memo[i-2]
	}
	return Memo[n]
}