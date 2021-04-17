import java.util.Stack;

public class Solution {

}

class MinStack {

  /** initialize your data structure here. */
  private Stack<Integer> stack;
  public MinStack() {
    this.stack = new Stack<>();
  }

  public void push(int val) {
    if (this.stack.empty()) {
      this.stack.push(val);
      this.stack.push(val);
    } else {
      int min = this.getMin();
      this.stack.push(val);
      if (min < val) {
        this.stack.push(min);
      } else {
        this.stack.push(val);
      }
    }
  }

  public void pop() {
    this.stack.pop();
    this.stack.pop();
  }

  public int top() {
    return this.stack.get(this.stack.size()-2);
  }

  public int getMin() {
    return this.stack.peek();
  }
}
