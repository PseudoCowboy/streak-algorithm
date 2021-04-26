public class Solution {
  public int addDigits(int num) {
    if (num < 10) {
      return num;
    }

    while (num > 9) {
      int cur = 0;
      while (num != 0) {
        cur += num%10;
        num /= 10;
      }
      num = cur;
    }

    return num;
  }
}
