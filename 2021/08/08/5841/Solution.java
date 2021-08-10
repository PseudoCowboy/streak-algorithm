class Solution {
    public int lengthOfLIS(int[] nums) {
        int[] tails = new int[obstacles.length];
        int[] dp = new int[obstacles.length];
        int res = 0;
        for (int k = 0; k < obstacles.length; k++) {
            int i = 0, j = res;
            while(i < j) {
                int m = (i + j) / 2;
                if(tails[m] < obstacles[k]) i = m + 1;
                else j = m;
            }
            tails[i] = obstacles[k];
            if(res == j) res++;
            dp[k] = res;
        }
        return res;
    }
}
