import java.util.HashMap;

public class Solution {
  public boolean containsNearbyDuplicate(int[] nums, int k) {
    HashMap<Integer, Integer> hashMap = new HashMap<>();

    for (int i = 0; i < nums.length; i++) {
      if (hashMap.get(nums[i]) != null) {
        if (i-hashMap.get(nums[i]) <= k) {
          return true;
        }
      }

      hashMap.put(nums[i], i);
    }
    return false;
  }

}
