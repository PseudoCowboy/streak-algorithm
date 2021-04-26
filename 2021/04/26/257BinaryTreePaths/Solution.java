import java.util.ArrayList;
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
  public List<String> binaryTreePaths(TreeNode root) {
    List<String> list = new ArrayList<>();
    if (root==null) {
      return list;
    }

    if (root.left == null && root.right == null) {
      list.add(String.valueOf(root.val));
      return list;
    }

    List<String> left = binaryTreePaths(root.left);
    for (int i = 0; i < left.size(); i++) {
      list.add(String.valueOf(root.val) + "->" + left.get(i));
    }

    List<String> right = binaryTreePaths(root.right);
    for (int i = 0; i < right.size(); i++) {
      list.add(String.valueOf(root.val) + "->" + right.get(i));
    }

    return list;

  }
}
