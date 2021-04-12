import java.util.LinkedList;

import apple.laf.JRSUIUtils.Tree;

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
  public boolean isSymmetric(TreeNode root) {
    if (root == null) {
      return true;
    }
    LinkedList<TreeNode> list = new LinkedList<>();
    list.add(root.left);
    list.add(root.right);
    TreeNode left, right;
    while (list.size() > 0) {
      left = list.pop();
      right = list.pop();
      if (left == null && right == null) {
        continue;
      }
      if (left.val != right.val) {
        return false;
      }
      if (left == null || right == null) {
        return false;
      }
      list.add(left.left);
      list.add(right.right);
      list.add(left.right);
      list.add(right.left);
    }
    return true;
  }
}
