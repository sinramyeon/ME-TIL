package BubbleSort

  func BubbleSort(array []int) {

    swapCount := 1

      for swapCount > 0 {

        swapCount = 0

          for itemIndex := 0; itemIndex < len(array)-1; itemIndex++ {

            if array[itemIndex] > array[itemIndex+1] { // 지금게 다음거보다 크면

               array[itemIndex], array[itemIndex+1] = array[itemIndex+1], array[itemIndex] // 바꾼다

               swapCount += 1

             }

           }

      }

}
