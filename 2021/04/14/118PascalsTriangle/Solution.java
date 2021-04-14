public class Solution {
  public List<List<Integer>> generate(int numRows) {
    List<List<Integer>> list = new ArrayList<>();
    for (int i = 0; i < numRows; ++i) {
        List<Integer> sub = new ArrayList<>();
        for (int j = 0; j <= i; ++j) {
            if (j == 0 || j == i) {
                sub.add(1);
            } else {
                List<Integer> upSub = list.get(i - 1);
                sub.add(upSub.get(j - 1) + upSub.get(j));
            }
        }
        list.add(sub);
    }
    return list;
  }
}
