
# Bubble Sort



Bubble sort is comparing the next pair of indexes and swapping if necessary.



if the numbers are in the completely wrong order, the bubble sort scanning all list, and pick nearest two indexes. it is a bit like bubbles rising up through the water. it is also sometimes called a ripple sort too.



you can see how much longer it takes to sort a list of completely wrong order. the list needs to be scanned len(n)-1 times, consider if we one item, we will search one and the next one. obvious.



when it's written down in words, it could be like this below.



>  Repeat number of items in the list -1 time

> For each adjacent pair of items in the list

>   if first of the pair > second of the pair then

>     swap positions

>   end if

>  next pair of items

>  end repeat



the bubble sort scans a list, comparing pairs of values and swapping them if necessary. For n time,s the list is scanned n-1 times. This sort is good for small lists of data. The bubble sort is not just for sorting numbers, can sort text too.



So how can we make enhancements for the bubble sort to make it more efficient?



For swapping, the bubble sort uses the temporary variable. it copies the first element to a temporary variable, and swapping  - the overwrite.



let's see the pseudocode.



```

FOR i = 0 to Length(ArrayToSort) -2

  IF ArrayToSort(i) > ArrayToSort(i+1) THEN

    Temp = ArrayToSort(i)

    ArrayToSort(i) = ArrayToSort(i+1)

    ArrayToSort(i+1) = Temp

  END IF

NEXT i

```



this could be an inner loop, and the outer loop should be like below.



```

For iPass = 1 to Length(ArrayToSort) -1



FOR i = 0 to Length(ArrayToSort) -2

  IF ArrayToSort(i) > ArrayToSort(i+1) THEN

    Temp = ArrayToSort(i)

    ArrayToSort(i) = ArrayToSort(i+1)

    ArrayToSort(i+1) = Temp

  END IF

NEXT i



Next iPass

```



but using a nested loop is not the best idea. now we should have noticed that the largest item is at the end of the list after the first pass. and also second largest is in the correct position after the second pass. what does it mean? the inner loop can run one less time with each pass.



```

count = 0

For iPass = 1 to Length(ArrayToSort) -1



FOR i = 0 to Length(ArrayToSort) -2 - count

  IF ArrayToSort(i) > ArrayToSort(i+1) THEN

    Temp = ArrayToSort(i)

    ArrayToSort(i) = ArrayToSort(i+1)

    ArrayToSort(i+1) = Temp

  END IF

NEXT i

count = count+1

Next iPass

```



if we use a value of count, inner loops run less. it could be the same code as



```

For iPass = 1 to Length(ArrayToSort) -1



FOR i = 0 to Length(ArrayToSort) -1 - iPass

  IF ArrayToSort(i) > ArrayToSort(i+1) THEN

    Temp = ArrayToSort(i)

    ArrayToSort(i) = ArrayToSort(i+1)

    ArrayToSort(i+1) = Temp

  END IF

NEXT i

Next iPass

```



it doesn't look like much but if we think about it this enhanced version of the bubble sort is doing half as much work as the original and comfortable.



there is another way for an alternative. if there have been no swaps, then all of the data must be in the correct order. next enhancement causes early finish ones that data is sorted.



```

REPEAT

  Swapped = False

   FOR i = 0 to Length(ArrayToSort) -2

  IF ArrayToSort(i) > ArrayToSort(i+1) THEN

    Temp = ArrayToSort(i)

    ArrayToSort(i) = ArrayToSort(i+1)

    ArrayToSort(i+1) = Temp

    Swapped = True

  END IF

NEXT i



UNTIL Swapped = False

```



this can be pseudocode for skipping.

Free grammar check
