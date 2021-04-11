public class Solution {
  public int[] plusOne(int[] digits) {
    boolean extendFlag = false;
    int index = 0;

    while (digits.length-1-index >= 0) {
      if (digits[digits.length-1-index]+1 > 9) {
        digits[digits.length-1-index] = 0;
      } else {
        digits[digits.length-1-index] = digits[digits.length-1-index] + 1;
        break;
      }

      if (digits.length-1-index == 0) {
        extendFlag = true;
      }
      index++;
    }

    if (extendFlag == true) {
      digits = new int[digits.length + 1];
      digits[0] = 1;
    }

    return digits;
  }
}
