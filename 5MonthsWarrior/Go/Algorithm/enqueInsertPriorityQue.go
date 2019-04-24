func HeapSort(arrInput []int) {
   hp := NewHeap(arrInput, true)
   n := len(arrInput)
   for i:= 0; i<n; i++{
    arrInput[i] = hp.Remove()
   }
}
