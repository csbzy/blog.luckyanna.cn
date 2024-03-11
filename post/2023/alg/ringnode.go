package alg

func backToOrigin(n int) int {
	length := 10
	// 走n步到0的方案数=走n-1步到1的方案数+走n-1步到9的方案数。
	// 因此，若设dp[i][j]为从0点出发走i步到j点的方案数，则递推式为：
	// dp[i][j]=dp[i-1][(j-1+length)%10]+dp[i-1][(j+1)%10]
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, length)
	}
	dp[0][0] = 1

	for i := 1; i < n+1; i++ {
		for j := 0; j < length; j++ {
			dp[i][j] = dp[i-1][(j-1+length)%length] + dp[i-1][(j+1)%length]
		}
	}
	return dp[n][0]
}
