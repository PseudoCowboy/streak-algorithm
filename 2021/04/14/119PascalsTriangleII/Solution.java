import java.util.ArrayList;
import java.util.List;

public class Solution {
  public static void main(String[] args) {
    List<Integer> list = new ArrayList<>(6);
    System.out.println(list);
        // list.set(0, 1);
        // list.set(rowIndex, 1);
  }

  public List<Integer> getRow(int rowIndex) {
    List<Integer> res = new ArrayList<>();
    for (int i = 0; i <= rowIndex; ++i) {
        res.add(1);
        for (int j = i - 1; j > 0; --j) {
            res.set(j, res.get(j - 1) + res.get(j));
        }
    }
    return res;
  }
}
