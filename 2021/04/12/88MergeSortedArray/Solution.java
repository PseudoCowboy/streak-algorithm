public class Solution {
  public void merge(int[] nums1, int m, int[] nums2, int n) {
    while (m+n-1 >= 0) {
      if (n==0) {
        return;
      }
      if (m==0) {
        nums1[m+n-1] = nums2[n-1];
			  n--;
        continue;
      }
      if (nums1[m-1] > nums2[n-1]) {
        nums1[m+n-1] = nums1[m-1];
        m--;
      } else {
        nums1[m+n-1] = nums2[n-1];
        n--;
      }
    }
  }
}
