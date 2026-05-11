class Solution {
    public boolean isAnagram(String s, String t) {
        int[] a = new int[26];
        int[] b = new int[26];

        for (char c: s.toCharArray()) {
            int idx = (int) c - 97;
            a[idx]++;
        }

        for (char c: t.toCharArray()) {
            int idx = (int) c - 97;
            b[idx]++;
        }

        return Arrays.equals(a, b);
    }
}
