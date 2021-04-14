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
  public int minDepth(TreeNode root) {
    if (root == nul) {
      return 0;
    }

    if (root.left == null) {
      return 1 + minDepth(root.right);
    }
    if (root.right == null) {
      return 1 + minDepth(root.left);
    }

    int left = minDepth(root.left);
    int right = minDepth(root.right);

    return left > right ? right + 1 : left + 1;
  }

  public int minDepth1(TreeNode root) {
    if (root == null) return 0;
    LinkedList<TreeNode> q = new LinkedList<>();
    q.add(root);
    int ans = 1;
    while (!q.isEmpty()) {
        int size = q.size();
        for (int i = 0; i < size; ++i) {
            TreeNode node = q.remove();
            if (node.left == null && node.right == null) {
                return ans;
            }
            if (node.left != null) q.add(node.left);
            if (node.right != null) q.add(node.right);
        }
        ++ans;
    }
    return 520;
}
}
