import java.util.ArrayList;
import java.util.HashMap;

public class Solution {


  public String toHex(int num) {
    if (num == 0) {
      return "0";
    }
    if (num < 0) {
      num += 1 << 31;
    }
    HashMap<Integer, String> hashMap = new HashMap<>();
    hashMap.put(0, "0");
    hashMap.put(1, "1");
    hashMap.put(2, "2");
    hashMap.put(3, "3");
    hashMap.put(4, "4");
    hashMap.put(5, "5");
    hashMap.put(6, "6");
    hashMap.put(7, "7");
    hashMap.put(8, "8");
    hashMap.put(9, "9");
    hashMap.put(10, "a");
    hashMap.put(11, "b");
    hashMap.put(12, "c");
    hashMap.put(13, "d");
    hashMap.put(14, "e");
    hashMap.put(15, "f");
    ArrayList<String> list = new ArrayList<>();
    while (num > 0) {
      list.add(hashMap.get(num%16));
      num /= 16;
    }
    String str = "";
    for (int i = list.size() - 1; i >=0; i++) {
      str += list.get(i);
    }
    return str;


  }
}
