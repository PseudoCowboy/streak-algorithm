public class Solution {
  public int removeElement(int[] nums, int val) {
    if (nums.length == 0) {
      return 0;
    }

    int outer = nums.length-1;
    for (int i = outer; i >= 0; i--) {
      if (nums[i] == val) {
        nums[i] = nums[outer];
        outer--;
      }
    }
    return outer+1;
  }
}
