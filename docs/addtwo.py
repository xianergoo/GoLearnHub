from typing import Optional

# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next
        
class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        carry = 0
        cur = dummy = ListNode()
        while l1 or l2 or carry:
            carry += (l1.val if l1 else 0) + (l2.val if l2.val else 0)
            cur.next = ListNode(carry % 10)
            carry = carry // 10
            cur = cur.next
            if l1: l1 = l1.next
            if l2: l2 = l2.next
        return dummy.next
    
if __name__ == '__main__':
    solution = Solution()
    l1 = ListNode()
    l2 = ListNode()
    solution.addTwoNumbers(l1, l2)
            
        


        