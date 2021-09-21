import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

class Solution {
    public int countKDifference(int[] nums, int k) {
        int count = 0;
        for (int i = 0; i < nums.length; i++) {
            for (int j = i+1; j < nums.length; j++) {
                if (Math.abs(nums[i]- nums[j]) == k) {
                    count++;
                }
            }
        }
        return count;
    }

    public int[] findOriginalArray(int[] changed) {
        if (changed.length % 2 != 0) {
            return new int[0];
        }
        Map<Integer, Integer> m = new HashMap<>();
        for (int i = 0; i < changed.length; i++) {
            m.computeIfAbsent(changed[i], k -> 1);
            m.computeIfPresent(changed[i], (k, v) -> v+1);
        }
        Arrays.sort(changed);
        List<Integer> ans = new ArrayList<>();
        for (int i = changed.length-1; i >= 0 ; i--) {
            if (changed[i]==0) {
                if (m.get(0) > 1) {
                    m.computeIfPresent(0, (k,v) -> v-2);
                    ans.add(0);
                }
                continue;
            }
            if (changed[i]%2 == 0) {
                int v = changed[i] /2 ;
                if (m.get(changed[i]) > 0 && m.get(v) > 0) {
                    m.computeIfPresent(changed[i], (k, val) -> val-1);
                    m.computeIfPresent(v, (k, val) -> val-1);
                    ans.add(v);
                }
            }
        }
        return ans.size() == changed.length/2 ? (int[]) ans.stream().mapToInt(i->i).toArray() : new int[0];
    }
}