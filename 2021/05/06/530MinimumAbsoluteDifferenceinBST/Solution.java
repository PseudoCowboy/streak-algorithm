import java.util.List;

public class TreeNode {
  int val;
  TreeNode left;
  TreeNode right;
  TreeNode() {}
  TreeNode(int val) { this.val = val; }
  TreeNode(int val, TreeNode left, TreeNode right) {
    this.val = val;
    this.left = left;
    this.right = right;
  }
}

public class Solution {
  public int getMinimumDifference(TreeNode root) {
    Stack<TreeNode> stack = new Stack<>();
    List<Integer> inorder = new ArrayList<>();

    TreeNode current = root;

    while (!stack.isEmpty() || current != null) {
        if (current != null) {
            stack.add(current);
            current = current.left;
        } else {
            current = stack.pop();
            inorder.add(current.val);
            current = current.right;
        }
    }
    int diff = 10000;
    for (int i = 1; i < inorder.size(); i++) {
     diff = Math.min(Math.abs(inorder.get(i)-inorder.get(i-1)), diff);
    }
    return diff;
  }

  public int getMinimumDifference1(TreeNode root) {
    return traverse(root, Integer.MAX_VALUE)[0];
  }

  private int[] traverse(TreeNode root, int parent){
      int[] left = root.left == null ? new int[]{Integer.MAX_VALUE, root.val} : traverse(root.left, root.val);
      int[] right = root.right == null ? new int[]{Integer.MAX_VALUE, parent} : traverse(root.right, parent);

      int distance = Math.min(right[1] - root.val, left[0]);

      return new int[]{Math.min(distance, right[0]), left[1]};
  }
}
