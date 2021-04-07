public class Solution {
  public int reverse(int x) {
    long result = 0;
    for (; x != 0; x/=10) {
      result = x % 10 + result * 10;
    }
    return result > Integer.MAX_VALUE || result < Integer.MIN_VALUE ? 0 : (int)result;
  }

  public static void main(String[] args) {
    Solution solution = new Solution();
    int input = 123;
    System.out.println(solution.reverse(input));

  }
}
