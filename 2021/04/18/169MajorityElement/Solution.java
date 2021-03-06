public class Solution {
  public int majorityElement(int[] nums) {
    int res = nums[0];
    int count = 0;
    for (int i = 0; i < nums.length; i++) {
      if (count == 0) {
        res = nums[i];
        count = 1;
      } else {
        if (nums[i] == res) {
          count++;
        } else {
          count--;
        }
      }
    }
    return res;
  }
}
