the queue a dynamic data structure in which items join at the back, and leave from the front.

think of a line of people waiting to be fed at a school canteen.
the person at the front of the queue, will be sered first, and the person at the back will be served last.

it's a dynamic data structrue because the amount of data it contians can be increase and decrease while it's in use.

a queue can be implemented as a linear queue. this could invole an array variable, more specifically a dynamic array, one with no fixed size.

this queue has some data in it, and we also hae a  system of pointers, a pointer to indicate the front of the queue, and another pointer to indicate the next free space.

taking data from the queue is known as dequeing. and this simply means taking data from the position given by the front pointer, and then incrementing the front pointer to indicate the new front of the queue.

if we dequeue again, its the same. take data from the position given by the front pointer, and then increment that pointer.

adding data to a queue is known as enqueueing. and this simply means putting data at the position given by next free, of course we then need to increment next free and again and queuing data at the position of next free, and incrementing next free.

implemening a linear queue with a dynamic array, has its advantages it's easy to do...howeer when we take data from the queue we are not actually remoing it fromt the array. all we are doing is incrementing the front pointer. so the memory is never actually freed up. now this is fine if you have got plenty of memory, but in a busy queue, it may end up working its way through the memory, and use up a lot without actually containing much data.

an alternative approach then is to use a static array, this time we have got a slightly different system of pointers we have one to indicat the front, and another to indicate the rear. we also have a variable to keep count of the number of items in the queue.

so when we dequeue this time, we take data from the position given by front, increment front like we did before, then update the value of the number in the queue. when we enqueue data, this means incrementing the value of the rear pointer and placing data at that position, and once again we update the number of items in the queue. if we enque again, its the same.

now there is going to come a point, where we can't add anymore at the back. we can continute to dequqe data but if we attempt to put some more at the back, we are not going to be able to do this, unless we first shuffle things along to create more space. so after shuffle, now this all well and good, but if you think about it with a big queue, that is a very expensive operation in terms of processing - not ideal.

To the rescue comes, the circular queue. the circurlar queue uses teh same system of pointers we have just seen. one for the front and one for the rear. we also have a variable to keep count of the number of items in the queue. when we want to include new data, we increment the rear pointer and we place the data there. deququeing data means taking it from the front, and then incrementing the front pointer. just as we've seen before. but this time if we want to include data and the rear pointer has reached its maximum, all we need to do is redefine the rear of the queue as the first element of the array. you can see then why it's called a circular queue.

theres going to come a point where the queue is full. and if we try to include more data, the number of items in the queue will match the number of items allowed by the size of the array. so we have to report that the queue is full. but of course we can stil continue to dequeue data. and its worth noticing what happends to the front pointer as we continue to do so. you can see now that the front pointer has reached max, we need to redefine the front pointer as one. so there you have it, the circular queue.


to summerize then, the queue is a dynamic data structure, its contents maybe changing all the time while its in use. new items are added to the rear we say that enqued, and items leave from the front we say there dequeued. a queue is a first in first out data structure. we say it is a FIFO. a queue can be linear or circular. it is an abstract data type(ADT). the data could be stored for example in a text file or database table, whats importatnt is that it is notionally a queue and items removed from the front and they join from the back. queues have all kinds of uses in the field of computing, the operating system makes use of queus in process schedulling. queus used as bufferes, for example for keyboard buffer when you are typing. also used to spool print output! 
