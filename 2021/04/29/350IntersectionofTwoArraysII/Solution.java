class Solution {
  public int[] intersect(int[] nums1, int[] nums2) {
    List<Integer> arrList = new ArrayList<>();
    Map<Integer, Integer> map = new HashMap<>();

    for (int num : nums1) {
        map.put(num, map.getOrDefault(num, 0) + 1);
    }

    for (int num : nums2) {
        if (map.containsKey(num) && map.get(num) > 0) {
            arrList.add(num);
            map.put(num, map.get(num)-1);
        }
    }

    int[] arr = new int[arrList.size()];
    int index = 0;

    for (Integer value : arrList) {
         arr[index++] = value;
      }

    return arr;
  }
}
