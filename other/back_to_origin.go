package other

// backToOrigin 圆环上有10个点，编号为0~9。从0点出发，每次可以逆时针和顺时针走一步，
// 问走n步回到0点共有多少种走法。
// dp[i][j] 代表从经过i步到j位置的走法数量
func backToOrigin(n int) int {
	length := 10
	dp := make([][]int, n+1)
	dp[0] = make([]int, length)
	dp[0][0] = 1
	for i := 1; i < n+1; i++ {
		dp[i] = make([]int, length)
		for j := 0; j < length; j++ {
			// (j-1+length)%length: 如果j-1==-1的话，可以退回到9
			// (j+1)%length: 如果j+1=10的话，可以回到原点0
			dp[i][j] = dp[i-1][(j-1+length)%length] + dp[i-1][(j+1)%length]
		}
	}

	// dp[n][0]代表走n步到0的走法数量
	return dp[n][0]
}
