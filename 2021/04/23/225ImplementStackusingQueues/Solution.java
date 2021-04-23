import java.util.Collection;
import java.util.Collections;
import java.util.LinkedList;

public class Solution {

}

class MyStack {

  LinkedList<Integer> linkedList = new LinkedList<>();
  /** Initialize your data structure here. */
  public MyStack() {
  }

  /** Push element x onto stack. */
  public void push(int x) {
    Collections.reverse(linkedList);
    linkedList.add(x);
    Collections.reverse(linkedList);
  }

  /** Removes the element on top of the stack and returns that element. */
  public int pop() {
    return linkedList.pop();
  }

  /** Get the top element. */
  public int top() {
    return linkedList.peek();
  }

  /** Returns whether the stack is empty. */
  public boolean empty() {
    return linkedList.isEmpty();
  }
}
