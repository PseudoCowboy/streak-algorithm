class Solution {

	public int minSpaceWastedKResizing(int[] nums, int k) {
		int[][] dp = new int[k + 1][nums.length];
		for (int i = 0; i <= k; i++) {
			for (int j = 0; j < nums.length; j++) {
				dp[i][j] = Integer.MAX_VALUE;
				for (int l = j, sum = 0, max = 0; l >= i; l--) {
					sum += nums[l];
					max = Math.max(max, nums[l]);
					dp[i][j] = Math.min(dp[i][j], i == 0 && l > 0 ? Integer.MAX_VALUE
							: (l > 0 ? dp[i - 1][l - 1] : 0) + (max * (j - l + 1) - sum));
				}
			}
		}
		return dp[k][nums.length - 1];
	}
}