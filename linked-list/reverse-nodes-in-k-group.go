/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {

    ListNode* reverse(ListNode* head){
        if(head==NULL || head->next==NULL)
            return head;
        
        ListNode* nh=reverse(head->next);
        head->next->next=head;
        head->next=NULL;
        return nh;
    }

public:
    ListNode* reverseKGroup(ListNode* head, int k) {
        ListNode* t=head;
        ListNode* t1=head;
        
        if(k==0 || k==1)
            return head;

        while(t!=NULL && t->next!=NULL){
            
            for(int i=0;i<k && t!=NULL;i++){
                
                if(i==(k-1)){
                    ListNode* nh=t->next;
                    t->next=NULL;
                    if(head==t1){
                        reverse(t1);
                        head=t;
                    }else{
                        reverse(t1->next);
                        t1->next=t;
                    }
                    while(t->next!=NULL)
                        t=t->next;
                    t1=t;
                    t->next=nh;            
                }
                    t=t->next;
                
            }
        }
        return head;        
    }
};