public class Solution {
  public int removeDuplicates(int[] nums) {
    int index = 0;
    if (nums.length == 0) {
      return index;
    }
    int temp = nums[0];
    for (int i = 0; i < nums.length; i++) {
      if (nums[i] != temp) {
        index++;
        nums[index] = nums[i];
        temp = nums[i];
      }
    }
    return index + 1;
  }
  public static void main(String[] args) {

  }
}
