import java.util.HashMap;

class Solution {
  public boolean wordPattern(String pattern, String s) {
    String[] sArr = s.split(" ");
    if (pattern.length() != sArr.length) {
      return false;
    }
    HashMap<Character, String> pMap = new HashMap<>();
    HashMap<String, Character> sMap = new HashMap<>();
    for (int i = 0; i < sArr.length; i++) {
      String pval = pMap.get(pattern.charAt(i));
      Character sval = sMap.get(sArr[i]);
      if ((pval != null && sval == null) || (pval == null && sval != null)) {
        return false;
      }
      if (pval != null && sval != null) {
        if (!pval.equals(sArr[i]) || sval != pattern.charAt(i)) {
          return false;
        }
      } else {
        pMap.put(pattern.charAt(i), sArr[i]);
        sMap.put(sArr[i], pattern.charAt(i));
      }
    }

    return true;
  }
}
