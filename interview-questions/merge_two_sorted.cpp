#include <cstddef>
#include <iostream>

using namespace std;

struct ListNode {
    int val;
    ListNode* next;
    ListNode(int x): val(x), next(NULL) {}
};

class Solution {
    public:
        ListNode* mergeTwoLists(ListNode* l1, ListNode* l2) {
            ListNode *head = nullptr, *cur = nullptr;
            while (l1 || l2) {
                if (!head) {
                    cur = head = new ListNode(0);
                } else {
                    cur = cur->next = new ListNode(0);
                }
                if ((l1 && !l2) || (l1 && l2 && l1->val <= l2->val)) {
                    cur->val = l1->val;
                    l1 = l1->next;
                    continue;
                }
                cur->val = l2->val;
                l2 = l2->next;
            }
            return head;
        }
};

int main() {
    Solution s;
    ListNode *l1 = new ListNode(1);
    l1->next = new ListNode(3);
    l1->next->next = new ListNode(8);
    ListNode *l2 = new ListNode(2);
    l2->next = new ListNode(15);
    l2->next->next = new ListNode(16);
    ListNode *res = s.mergeTwoLists(l1, l2);
    while (res) {
        cout << res->val << endl;
        res = res->next;
    }
}
