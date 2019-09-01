Insertion Sort

Do you ever played card games? You can remember you tried to sort a card with the order.
insertion sort is also same.

it has sorted, unsorted section and has also a pointer.
compare the current item with the first item, and do a swap, it is a basic idea of insertion sort.

lets say you have [7, 8, 5, 2 ... ] array.

compare 7 and 8, compare 8 and 5 next 7 and 5. the pointer separates the sorted section of the list on the left from the unsorted section on the right.
the pointer keeps going to the right till the last elements in the array.

to get familiar with this algorithm, imagine playing cards with laying random order.
here is the following pseudocode.

```
Set pointer to second item in the list
Repeat
 select first item in unsorted section and make it current
 Repeat
   Compare current item with next item in sorted section
   Move sorted item one place to the right if bigger than current item
 Until sorted item smaller than current item or all sorted item checked
 Insert current item into sorted section
 Advance the pointer
Until unsorted section is empty
```

to summarise, the insertion sort is like sorting a hand of cards.
it is particulary good for lists that are almost sorted. also good for inserting new values into a sorted list. Sorted and unsorted sections of the lis are separated by a pointer. left is sorted, right is unsorted.
The current item compared with each item in the sorted section. this is more efficient than buble sort.(little faster)

repeatedely scanning a list might not be a best solution but Insertion sort is still can be improved.

do you remember the pesudo code of bubble sort?

```
For i = 0 to length(arraytosort) - 2
  if arraytosort(i) > arraytosort(i+1) then
    temp = arraytosrot(i)
    arraytosort(i) = arraytosort(i+1)
    arraytosort(i+1) = temp
  End if
Next i
```

you can imagine this pseudocode, the inner loop needs to run several times so the outer loop controls the number of times the list is scanned that is the number of passes through the list.

 ```
 For ipass = ` to length(arraytosort) -1 
  For i = 0 to length(arraytosort) - 2
    if arraytosort(i) > arraytosort(i+1) then
     temp = arraytosrot(i)
      arraytosort(i) = arraytosort(i+1)
     arraytosort(i+1) = temp
     End if
  Next i
Next ipass
```

also, from bubble sort, it is not very hard to make a pesudocode for insertion sort.

```
For pointer = 1 to array.length -1
  current = array[pointer]
  position = pointer
  while position > 0 and array[position-1] > current
    array[position] = array[position-1]
    position = position -1
  end while
  array[position] = current
Next pointer
```

here we are using pointer not like bubble sort!
