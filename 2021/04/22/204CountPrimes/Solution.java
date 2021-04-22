public class Solution {
  public int countPrimes(int n) {
    int[] isNotPrime = new int[n];
    for (int i = 2; i < n ; i++) {
      if (isNotPrime[i] == 1) {
        continue;
      }
      for (int j = i*i; j < n; j = j+i) {
        isNotPrime[j] = 1;
      }
    }

    int count = 0;
    for (int i = 2; i < n; i++) {
      if (isNotPrime[i] == 0) {
        count++;
      }
    }

    return count;
  }
}
