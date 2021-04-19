public class Solution {
  public int titleToNumber(String columnTitle) {
    int val = 0;
    int res = 0;
    for (int i = 0; i < columnTitle.length(); i++) {
      val = (int)(columnTitle.charAt(i) - 'A' + 1);
      res = res * 26 + val;
    }

    return res;
  }
}
