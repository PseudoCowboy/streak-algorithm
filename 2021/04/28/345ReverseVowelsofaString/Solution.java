public class Solution {
  public String reverseVowels(String s) {
		char[] array = s.toCharArray();
		int left = 0, right = s.length() - 1;

		while (left <= right) {
			if (isVowel(array[left]) && isVowel(array[right])) {
				swap(array, left, right);
				left++;
				right--;
			} else if (!isVowel(array[left])) {
				left++;
			} else if (!isVowel(array[right])) {
				right--;
			}
		}

		return new String(array);
	}

    public void swap(char[] arr, int i, int j) {
		char temp = arr[i];
		arr[i] = arr[j];
		arr[j] = temp;
	}

	public boolean isVowel(Character c) {
		return c == 'a' || c == 'A' || c == 'e' || c == 'E' || c == 'i' || c == 'I'
				|| c == 'o' || c == 'O' || c == 'u' || c == 'U';
	}
}
