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

class Solution {
  int diameter = 0;
  public int diameterOfBinaryTree(TreeNode root) {
    postOrder(root);
    return diameter;
  }

  private int postOrder(TreeNode root){
    if(root == null) return 0;
    int leftSize = postOrder(root.left);
    int rightSize = postOrder(root.right);
    diameter = Math.max(diameter, leftSize + rightSize);
    return Math.max(leftSize, rightSize) + 1;
  }
}
