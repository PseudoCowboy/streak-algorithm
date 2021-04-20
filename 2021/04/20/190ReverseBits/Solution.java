public class Solution {
  public int reverseBits(int n) {
      var rev = 0;
      for (int i = 0; i < 32; i++) {
        rev = ( rev << 1 ) | ( n & 1 );
        n >>= 1;
      }
      return rev;
  }
}
