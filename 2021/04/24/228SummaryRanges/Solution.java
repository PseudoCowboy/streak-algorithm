import java.util.ArrayList;
import java.util.List;

class Solution {
    public List<String> summaryRanges(int[] nums) {
        List<String> list = new ArrayList<>();
        for (int i = 0, n = nums.length; i < n; ) {
            int left = i;
            for (i++; i < n && nums[i-1]+1 == nums[i]; i++) {
            }
            String s = String.valueOf(nums[left]);
            if (left != i-1) {
                s += "->" + String.valueOf(nums[i-1]);
            }
            list.add(s);
        }

        return list;

    }
}