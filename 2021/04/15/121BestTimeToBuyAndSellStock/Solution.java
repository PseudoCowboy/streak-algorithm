public class Solution {
  public int maxProfit(int[] prices) {
    if (prices.length == 0) {
      return 0;
    }

    int max = 0;
    int minPrice = prices[0];

    for (int i = 0; i < prices.length; i++) {
      if (prices[i]<minPrice) {
        minPrice = prices[i];
      }
      if (prices[i] - minPrice > max) {
        max = prices[i] - minPrice;
      }
    }

    return max;
  }
}
