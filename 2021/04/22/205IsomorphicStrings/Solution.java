import java.util.HashMap;

public class Solution {
  public boolean isIsomorphic(String s, String t) {
    HashMap<Character, Character> ahashMap = new HashMap<>();
    HashMap<Character, Character> bhashMap = new HashMap<>();
    for (int i = 0; i < s.length(); i++) {
      if (ahashMap.get(s.charAt(i)) != null) {
        if (bhashMap.get(t.charAt(i)) != null) {
          if (bhashMap.get(t.charAt(i)) != s.charAt(i)) {
            return false;
          }
        } else {
          return false;
        }
      } else {
        if (bhashMap.get(t.charAt(i)) != null) {
          return false;
        }
        ahashMap.put(s.charAt(i), t.charAt(i));
        bhashMap.put(t.charAt(i), s.charAt(i));
      }
    }

    return true;
  }
}
