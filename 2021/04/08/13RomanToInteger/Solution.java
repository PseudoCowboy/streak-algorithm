import java.util.HashMap;

public class Solution {
  public int romanToInt(String s) {
    HashMap<Character, Integer> map = new HashMap<>();
    map.put('I', 1);
    map.put('V', 5);
    map.put('X', 10);
    map.put('L', 50);
    map.put('C', 100);
    map.put('D', 500);
    map.put('M', 1000);
    char[] arr =  s.toCharArray();
    boolean skip = false;
    int result = 0;
    for (int i = 0; i < arr.length; i++) {
      if (skip == true) {
        skip = false;
        continue;
      }
      if (i+1 < arr.length && map.get(arr[i]) < map.get(arr[i+1])) {
        skip = true;
        result += map.get(arr[i+1]) - map.get(arr[i]);
      } else {
        result += map.get(arr[i]);
      }
    }
    return result;
  }

  public static void main(String[] args) {
    Solution solution = new Solution();
    solution.romanToInt("III");
  }
}
