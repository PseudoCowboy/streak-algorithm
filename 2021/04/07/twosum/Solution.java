import java.util.Arrays;
import java.util.HashMap;


public class Solution {

  public int[] twoSum(int[] nums, int target) {
    int len = nums.length;
    HashMap<Integer, Integer> map = new HashMap<>();
    for (int i = 0; i < len; i++) {
      map.put(nums[i], i);

      final Integer value = map.get(target-nums[i]);
      if (value != null) {
        return new int[]{value, i};
      }
    }
    return null;
  }

  public static void main(String[] args) {
    Solution solution = new Solution();
    int[] nums = new int[]{2, 7, 11, 15};
    int target = 9;
    System.out.println(Arrays.toString(solution.twoSum(nums, target)));
  }
}
