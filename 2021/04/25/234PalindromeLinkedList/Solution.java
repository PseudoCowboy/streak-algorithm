import java.util.ArrayList;
import java.util.List;

public class ListNode {
    int val;
    ListNode next;
    ListNode() {}
    ListNode(int val) { this.val = val; }
    ListNode(int val, ListNode next) { this.val = val; this.next = next; }
}

class Solution {
  public boolean isPalindrome(ListNode head) {
    List<Integer> list = new ArrayList<>();

    while (head != null) {
      list.add(head.val);
      head = head.next;
    }

    for (int start = 0 ,end = list.size()-1; start < end; ){
      if (list.get(start) == list.get(end)) {
        start++;
        end--;
        continue;
      }
      return false;
    }

    return true;
  }
}
