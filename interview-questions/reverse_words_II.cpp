#include <algorithm>
#include <iostream>
#include <string>
#include <vector>

using namespace std;

/* Given an input string, reverse the string word by word. A word is defined as
 * a sequence of non-space characters.
 *
 * The input string does not contain leading or trailing spaces and the words
 * are always separated by a single space.
 *
 * For example,
 * Given s = "the sky is blue",
 * return "blue is sky the".
 *
 * Could you do it in-place without allocating extra space?
 * */

class Solution {
public:
  void reverseWords(string &s) {
    // Reverse the entire string first.
    int size = s.size();
    for (int i = 0; i < size / 2; i++) {
      swap(s[i], s[size - i - 1]);
    };
    // Now, reverse each word using a sliding window.
    for (int l = 0, r = 0; r < size;) {
      if (s[l] == ' ') {
        l = ++r;
      } else if (r == size - 1 || s[r + 1] == ' ') {
        for (int d = 0; d <= (r - l) / 2; d++) {
          swap(s[l + d], s[r - d]);
        }
        l = ++r;
      } else {
        r++;
      }
    }
  };
};

int main() {
  vector<string> s = {"the sky is blue", "", "a", " ", "a spoonful of Sugar"};
  Solution *sol = new Solution();
  for (int i = 0; i < s.size(); i++) {
    sol->reverseWords(s[i]);
    cout << s[i] << endl;
  };
}
