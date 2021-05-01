public class Solution {
  public boolean isSubsequence(String s, String t) {
    int slen = s.length();
    int tlen = t.length();
    int sind = 0, tind = 0;
    while (sind < slen && tind < tlen) {
      if (s.charAt(sind) == t.charAt(tind)) {
        sind++;
      }
      tind++;
    }
    return sind == slen;
  }
}
