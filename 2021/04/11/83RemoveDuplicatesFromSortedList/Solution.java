public class ListNode {
  int val;
  ListNode next;
  ListNode() {}
  ListNode(int val) { this.val = val; }
  ListNode(int val, ListNode next) { this.val = val; this.next = next; }
}

public class Solution {
  public ListNode deleteDuplicates(ListNode head) {
    if (head == null) {
      return null;
    }

    ListNode result = head;
    while (head.next != null) {
      if (head.val == head.next.val) {
        head.next = head.next.next;
        continue;
      }
      head = head.next;
    }

    return result;
  }
}
