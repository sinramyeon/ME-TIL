[stack.py](https://github.com/sinramyeon/ME-TIL/blob/master/Beakjoon/stack.py) 에 쓰인 스택 자료구조에 대해 설명해 보자.

a stack is a dynamic data structure with a variety of applications in the field of computing.

it is known as a dynamic data structre because the amount of data it conitains can increase and decrease while it is in use.

a stack is known as the last in first out = it is sometimes referred to as a LIFO data structre. because new items are added to the top of the stack, and items can only be removed from the top of the stack.

it is kind of a pile of dirty dishes that needs washing.

to implement a stack, we could use an array variable. we we would normally also include variable to indicate how many items the stack can contain, and a pointer variable to indicate the top of the stack, with an empty stack the top pointer will be 0.

i used len as a length of the array which is stack in here.

now lets say we want to put some data on to the stack, this is known as a push operation. so we're going to push the data on to this stack. this involves incrementing the value of top, and then putting that data into the array at that position.

another push we increment the value of top, and add the data. and so on.

become a point, where we try to push another item of data, but the value of the top is the same as the value of maximum size, in which case the stack is full, and we can't add anymore.

taking data from a stack, is known as popping data, so when we perform a pop operation, we copy data from the position denoted by the top, and then decrement top. it's worth noticing as well that when we take data from the stack, we don't necessarliy have to remove it from the array that contains the data, all we need to do is redefine the top.

but in my code, i made `del` for get rid of the variable from the stack.

so, when we push new data on to the stack, it is at this stage we redefine top and if necessary overwrite data which was previously popped.

