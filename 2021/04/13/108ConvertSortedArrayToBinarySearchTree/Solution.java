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
  public TreeNode sortedArrayToBST(int[] nums) {
    if (nums == null || nums.length == 0) return null;
    return helper(nums, 0, nums.length - 1);
}

private TreeNode helper(int[] nums, int left, int right) {
    if (left > right) return null;
    int mid = (left + right) >> 1;
    TreeNode node = new TreeNode(nums[mid]);
    node.left = helper(nums, left, mid - 1);
    node.right = helper(nums, mid + 1, right);
    return node;
}
  public static void main(String[] args) {
    System.out.println(1>>>1);
    System.out.println(2>>>1);
    System.out.println(3>>>1);
  }
}
