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
  public boolean hasPathSum(TreeNode root, int targetSum) {
    if (root == null) {
      return false;
    }

    boolean left = false;
    boolean right = false;
    if (root.right != null) {
      right = hasPathSum(root.right, targetSum-root.val);
    }
    if (root.left != null) {
      left = hasPathSum(root.left, targetSum-root.val);
    }

    if (root.right == null && root.left == null) {
      return targetSum == root.val;
    }

    return left || right;
  }
}
