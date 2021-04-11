public class Solution {
  public String addBinary(String a, String b) {
    StringBuilder sb = new StringBuilder();
    int carry = 0, p1 = a.length() - 1, p2 = b.length() - 1;
    while (p1 >= 0 && p2 >= 0) {
        carry += a.charAt(p1--) - '0';
        carry += b.charAt(p2--) - '0';
        sb.insert(0, (char) (carry % 2 + '0'));
        carry >>= 1;
    }
    while (p1 >= 0) {
        carry += a.charAt(p1--) - '0';
        sb.insert(0, (char) (carry % 2 + '0'));
        carry >>= 1;
    }
    while (p2 >= 0) {
        carry += b.charAt(p2--) - '0';
        sb.insert(0, (char) (carry % 2 + '0'));
        carry >>= 1;
    }
    if (carry == 1) {
        sb.insert(0, '1');
    }
    return sb.toString();
  }

  public static void main(String[] args) {
    // int carry = 3;
    // System.out.println(carry >>= 1);
    // System.out.println(carry);
    // System.out.println(carry <<= 1);
    // System.out.println(carry >>= 1);
    int[] digits = {1,2,3};
    digits = new int[digits.length + 1];
    digits[0] = 1;
    System.out.println(digits[0]);
    System.out.println(digits[1]);
    System.out.println(digits[2]);
    System.out.println(digits[3]);

  }
}
