public class Solution {
  public boolean isPalindrome(int x) {
    if (x < 0) {
      return false;
    }
    int temp = x;
    long result = 0;
    for (; temp != 0; temp/=10) {
      result = temp % 10 + result*10;
    }
    return result == x;
  }

  public static void main(String[] args) {
    Solution solution = new Solution();
    int input = 121;

    System.out.println(solution.isPalindrome(input));
  }
}
