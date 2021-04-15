public class Solution {
  public int maxProfit(int[] prices) {
    if (prices.length == 0) {
      return 0;
    }

    int total = 0;
    int min = prices[0];

    for (int i = 0; i < prices.length; i++) {
      if (prices[i] > min) {
        total += prices[i] - min;
        min = prices[i];
      }
      if (prices[i] < min) {
        min = prices[i];
      }
    }

    return total;
  }
}
