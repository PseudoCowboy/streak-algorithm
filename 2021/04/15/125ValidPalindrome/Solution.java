public class Solution {
  public boolean isPalindrome(String s) {
    int start = 0;
    int end = s.length() - 1;
    while (start < end) {
      if (isChar(s.charAt(start)) && isChar(s.charAt(end))) {
        if (Character.toLowerCase(s.charAt(start++)) != Character.toLowerCase(s.charAt(end--))) {
          return false;
        }
      } else {
        if (!isChar(s.charAt(start))) {
          start++;
        }
        if (!isChar(s.charAt(end))) {
          end--;
        }
      }
    }
    return true;
  }

  private boolean isChar(char c) {
    if (Character.isLetter(c) || Character.isDigit(c)) {
      return true;
    }
    return false;
  }
}
