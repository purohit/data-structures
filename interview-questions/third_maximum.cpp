#include <algorithm>
#include <iostream>
#include <unordered_set>
#include <vector>
#include <queue>

using namespace std;

/* Given a non-empty array of integers, return the third maximum distinct number
 * in this
 * array. If it does not exist, return the maximum number. The time complexity
 * must be in O(n).
 *
 * Input: [2, 2, 3, 1]
 * Output: 1
 * Explanation: Note that the third maximum here means the third maximum
 * distinct number.
 * Both numbers with value 2 are both considered as second maximum.
 */

class Solution {
public:
  int thirdMax(vector<int> &nums) {
    if (nums.size() == 0) {
      return 0;
    }
    priority_queue<int, vector<int>, greater<int>> h; // min queue
    unordered_set<int> seen;
    for (const auto n : nums) {
      if (seen.find(n) != seen.end()) {
        continue;
      }
      if (h.size() < 3 || n > h.top()) {
        if (h.size() == 3) {
          h.pop();
        }
        h.push(n);
        seen.insert(n);
      }
    }
    if (h.size() == 3) {
      return h.top();
    }
    int biggest = h.top();
    h.pop();
    if (!h.empty()) {
        biggest = max(biggest, h.top());
    }
    return biggest;
  }
};

int main() {
  vector<vector<int>> cases = {{3, 2, 1},
                               {1, 2},
                               {2, 2, 3, 1},
                               {2, 2, 2},
                               {2, 2, 1},
                               {1, -214777, 2},
                               {1, 4, 5, 7, 10, -3, 15, 2, 0}};
  Solution sol;
  for (auto &c : cases) {
    cout << "RESULT " << sol.thirdMax(c) << endl;
  };
}
