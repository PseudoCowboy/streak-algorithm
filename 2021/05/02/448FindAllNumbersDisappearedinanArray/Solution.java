import java.util.ArrayList;

class Solution {
    public List<Integer> findDisappearedNumbers(int[] nums) {
        ArrayList<Integer> list = new ArrayList<>();
        int[] arr = new int[nums.length];
        for (int i = 0; i < nums.length; i++) {
            arr[nums[i]-1]++;
        }
        for (int i = 0; i < arr.length; i++) {
            if (arr[i] == 0) {
                list.add(i+1);
            }
        }
        return list;
    }
}