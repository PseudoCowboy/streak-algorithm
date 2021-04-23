import java.util.HashMap;

public class Solution {
  public boolean containsDuplicate(int[] nums) {
    HashMap<Integer, Boolean> hashMap = new HashMap<>();

    for (int i = 0; i < nums.length; i++) {
      if (hashMap.get(nums[i]) != null) {
        return true;
      } else {
        hashMap.put(nums[i], true);
      }
    }

    return false;
  }
}
