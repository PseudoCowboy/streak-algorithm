class Solution {
    public int longestPalindrome(String s) {
        int[] chArr = new int[52];
        for (int i = 0; i < s.length(); i++) {
            chArr[s.charAt(i)-'A']++;
        }
        int sum = 0;
        for (int o: chArr) {
            sum += (o/2)*2;
        }
        if (sum < s.length()) {
            sum++;
        }
        return sum;
    }
}