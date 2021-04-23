package dynamic_programming

/*
最大子序和

输入: [-2,1,-3,4,-1,2,1,-5,4],
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
 */

func maxSumSubArray(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	dp := make([]int, len(nums))
	//设置初始化值
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		//处理 dp[i-1] < 0 的情况
		if dp[i-1] < 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = dp[i-1] + nums[i]
		}
	}
	result := -1 << 31
	for _, k := range dp {
		result = max(result, k)
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}