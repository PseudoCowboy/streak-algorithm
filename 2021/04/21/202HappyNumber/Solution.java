import java.util.HashMap;

public class Solution {
  public boolean isHappy(int n) {
    if (n == 0) {
      return false;
    }

    int res = 0;
    int num = n;
    HashMap<Integer, Integer> hashMap = new HashMap<>();
    while (true) {
      while (num != 0) {
        res += (num%10) * (num%10);
        num /= 10;
      }
      if (hashMap.get(res) == null) {
        if (res == 1) {
          return true;
        }
        num = res;
        hashMap.put(res, res);
        res = 0;
        continue;
      } else {
        return false;
      }
    }
  }
}
