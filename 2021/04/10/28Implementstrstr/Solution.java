public class Solution {
  public int strStr(String haystack, String needle) {
    if (needle.equals("")) {
      return 0;
    }
    if (haystack.equals("")) {
      return -1;
    }
    int needleLength = needle.length();
    int haystackLength = haystack.length();
    int index = 0;
    while (index <= haystackLength - needleLength) {
      for (int i = 0; i < needleLength; i++) {
        if (needle.charAt(i) != haystack.charAt(i+index)) {
          break;
        }
        if (i == needleLength-1) {
          return index;
        }
      }
      index++;
    }
    return -1;
  }
}
