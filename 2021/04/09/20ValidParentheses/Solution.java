import java.util.HashMap;
import java.util.Stack;

public class Solution {
  public boolean isValid(String s) {
    if (s.length()%2 != 0) {
      return false;
    }

    HashMap<Character, Character> map = new HashMap<>();
    map.put('(', ')');
    map.put('{', '}');
    map.put('[', ']');
    Stack stack = new Stack<Character>();

    for (int i = 0; i < s.length(); i++) {
      if (map.get(s.charAt(i)) != null) {
        stack.push(s.charAt(i));
      } else {
        if (stack.empty()) {
          return false;
        }
        if (map.get(stack.peek()) != s.charAt(i)) {
          return false;
        } else {
          stack.pop();
        }
      }
    }
    return stack.empty();
  }
  public static void main(String[] args) {
    String s = "()";
    Solution solution = new Solution();
    System.out.println(solution.isValid(s));
  }
}
