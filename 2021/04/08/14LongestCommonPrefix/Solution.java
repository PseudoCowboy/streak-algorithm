public class Solution {
  public String longestCommonPrefix(String[] strs) {
    int index = 0;
    String prefix = "";
    if (strs.length == 0) {
      return prefix;
    }
    String s = "";
    boolean flag = true;
    while (flag) {
      for (int i = 0; i < strs.length; i++) {
        if (index > strs[i].length()-1) {
          flag = false;
          break;
        }
        if (i == 0) {
          s = String.valueOf(strs[i].charAt(index));
        }
        if (!String.valueOf(strs[i].charAt(index)).equals(String.valueOf(s))) {
          flag = false;
          break;
        }
        if (i == strs.length - 1) {
          prefix += s;
          s = "";
        }
      }
      index++;
    }
    return prefix;
  }
  public static void main(String[] args) {
    Solution solution = new Solution();
    String[] strs = {"flower","flow","flight"};

    System.out.println(solution.longestCommonPrefix(strs));
  }
}
