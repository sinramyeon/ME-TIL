we have a double ended queue. or deque. additions and deletions can be made from both the rear and the front.

if you have a deque implementation, that can give you a queue or stack implementaiton easily. but in some cases stacks or even queues can have somewhat simpler underlying structure. starting with a linked list implementation, double linked list gives really easy way to implement a deque. insertLeft and deleteRight and vice versa. for a regular quque, is like singly linked list with a tail reference. for a deque, used double linked list. stack is even simpler with singly linked list, without a tail refernce. in collections for python, it has a deque.

deque is faster in terms of adding elements to the end, and the beginning of a list. so rather than using a list you want to use a deque, it's gonna be faster. 

