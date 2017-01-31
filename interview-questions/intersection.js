function ListNode(val) {
    this.val = val;
    this.next = null;
}

var getLength = function(node) {
    var length = 0;
    while (node != null) {
        length += 1;
        node = node.next;
    }
    return length;
};

var getIntersectionNode = function(headA, headB) {
    if (headA == null || headB == null) {
        return null;
    }
    var aLen = getLength(headA);
    var bLen = getLength(headB);
    var i = 0;
    if (aLen > bLen) {
        for (i = 0; i < (aLen - bLen); i += 1) {
            headA = headA.next;
        }
    } else if (bLen > aLen) {
        for (i = 0; i < (bLen - aLen); i += 1) {
            headB = headB.next;
        }
    }
    console.log(headA.val, headB.val);
    while (headA != null && headB != null) {
        if (headA == headB) {
            return headA;
        }
        headA = headA.next;
        headB = headB.next;
    }
    return null
};

var a1 = new ListNode("a1");
var a2 = new ListNode("a2");
var a3 = new ListNode("a3");
var a4 = new ListNode("a4");
var c1 = new ListNode("c1");
var c2 = new ListNode("c2");
var c3 = new ListNode("c3");
var b1 = new ListNode("b1");
var b2 = new ListNode("b2");
var b3 = new ListNode("b3");
var x = new ListNode("");

console.log(getIntersectionNode(a1, b1));
a1.next = a2;
a2.next = a3;
a3.next = a4;
a4.next = c1;
c1.next = c2;
c2.next = c3;
b1.next = b2;
b2.next = b3;
b3.next = c1;
console.log(getIntersectionNode(a1, x));
console.log(getIntersectionNode(null, a1));
console.log(getIntersectionNode(a1, c1));
console.log(getIntersectionNode(a1, b1));
console.log(getIntersectionNode(a1, b2));
console.log(getIntersectionNode(a2, b1));
