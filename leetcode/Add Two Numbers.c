/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
struct ListNode* addTwoNumbers(struct ListNode* l1, struct ListNode* l2) {
    int flag = 0;
    int result = l1;
    for(;;) {
        l1->val = (l1->val + l2->val) + flag;
        flag = 0;
          if(l1->next == NULL && l2->next != NULL) {
                l1->next = (struct ListNode*)malloc(sizeof(struct ListNode));
                l1->next->val = 0;
                l1->next->next = NULL;
            }
            if(l2->next == NULL && l1->next != NULL) {
                l2->next = (struct ListNode*)malloc(sizeof(struct ListNode));
                l2->next->val = 0;
                l2->next->next = NULL;
            }
        if(l1->val >= 10) {
            flag = 1;
            l1->val -= 10;
            if(l1->next == NULL) {
                l1->next = (struct ListNode*)malloc(sizeof(struct ListNode));
                l1->next->val = 0;
                l1->next->next = NULL;
            }
            if(l2->next == NULL) {
                l2->next = (struct ListNode*)malloc(sizeof(struct ListNode));
                l2->next->val = 0;
                l2->next->next = NULL;
            }
        }
        if((l1 = l1->next) == NULL)
            break;
        if((l2 = l2->next) == NULL)
            break;
    }
    return result;
}
