/*
 * Given two words (beginWord and endWord), and a dictionary's word list, find
 * all shortest transformation sequence(s) from beginWord to endWord, such
 * that:
 *
 * Only one letter can be changed at a time
 * Each transformed word must exist in the word list. Note that beginWord is
 * not a transformed word.
 * For example,
 *
 * Given:
 * beginWord = "hit"
 * endWord = "cog"
 * wordList = ["hot","dot","dog","lot","log","cog"]
 * Return
 *   [
 *       ["hit","hot","dot","dog","cog"],
 *       ["hit","hot","lot","log","cog"]
 *   ]
 *             Note:
 *             Return an empty list if there is no such transformation
 *             sequence.
 *             All words have the same length.
 *             All words contain only lowercase alphabetic characters.
 *             You may assume no duplicates in the word list.
 *             You may assume beginWord and endWord are non-empty and are not
 *             the same.
 *             UPDATE (2017/1/20):
 *             The wordList parameter had been changed to a list of strings
 *             (instead of a set of strings). Please reload the code definition
 *             to get the latest changes.
 */

#include <iostream>
#include <string>
#include <unordered_set>
#include <vector>

using namespace std;

class Solution {
public:
  vector<vector<string>> findLadders(string beginWord, string endWord,
                                     vector<string> &wordList) {
    unordered_set<string> dict;
    vector<string> transform = {beginWord};
    vector<vector<string>> results;
    for (auto w : wordList) {
        if (w == beginWord) {
            continue;
        }
        dict.insert(w);
    }
    if (dict.find(endWord) == dict.end()) {
      return results;
    }
    // Try all possible combinations of the current word
    fillResults(endWord, transform, dict, results);
  }

private:
  void fillResults(const string end, vector<string> &transform,
                   unordered_set<string> &dict,
                   vector<vector<string>> &results) {
    if (dict.size() == 0) {
      return;
    }
    // Backtrack by going through all the positions in the
    // last string. If the position is end, then add that one
    // and return.
    auto last = transform.back();
    for (int i = 0; i < last.size(); i++) {
      auto old = last[i];
      for (int d = 0; d < 26; d++) {
        last[i] = 'a' + d;
        transform.push_back(last);
        if (last == end) {
            //cout << last << endl;
          results.push_back(transform);
        } else if (dict.find(last) != dict.end()) {
          dict.erase(last);
            //cout << last << endl;
          fillResults(end, transform, dict, results);
          dict.insert(last);
        }
        transform.pop_back();
      }
      last[i] = old;
    }
  }
};

struct test {
  string begin;
  string end;
  vector<string> words;
  test(string b, string e, vector<string> w) : begin(b), end(e), words(w) {}
};

int main() {
  Solution s;
  vector<test> ts = {
      test("hit", "cog",
           vector<string>({"hot", "dot", "dog", "lot", "log", "cog"})),
      //test("hit", "pog",
           //vector<string>({"hot", "dot", "dog", "lot", "log", "cog"})),
  };
  for (auto t : ts) {
    auto res = s.findLadders(t.begin, t.end, t.words);
    for (auto i : res) {
      for (auto j : i) {
        cout << j << " ";
      }
      cout << endl;
    }
  }
}
