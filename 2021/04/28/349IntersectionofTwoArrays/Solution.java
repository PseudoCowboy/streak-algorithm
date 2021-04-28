import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;

public class Solution {
  public int[] intersection(int[] nums1, int[] nums2) {
    HashMap<Integer, Boolean> hashMap = new HashMap<>();
    for (int i = 0; i < nums1.length; i++) {
      hashMap.put(nums1[i], true);
    }
    List<Integer> list = new ArrayList<>();
    for (int i = 0; i < nums2.length; i++) {
      if (hashMap.get(nums2[i]) == true) {
        list.add(nums2[i]);
        hashMap.remove(nums2[i]);
      }
    }

    int[] res = new int[list.size()];
    for (int i = 0; i < res.length; i++) {
      res[i] = list.get(i);
    }
    return res;
  }
}
