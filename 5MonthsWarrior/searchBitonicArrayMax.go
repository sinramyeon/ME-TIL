// a bitonic list comprises of an increasing sequence of integers
// immediately followed by a decreasing sequence of integers.
//https://en.wikipedia.org/wiki/Bitonic_sorter
// 먼솔야...?

func SearchBitonicArrayMax(data []int) (int, bool) {
  size := len(data)
  start := 0
  end := size - 1
  mid := (start+end) / 2
  
  maximaFound := 0
  if size < 3{
    return 0, false
  }
  
  for start <= end {
    mid := (start+end)/2
    
    if data[mid-1] < data[mid] && data[mid+1] < data[mid]{
      maximaFound = 1
      break
    } else if data[mid-1] < data[mid] && data[mid] < data[mid+1] {
      start = mid+1
    } else if data[mid-1] > data[mid] && data[mid] > data[mid+1] {
      end = mid-1 
    } else {
      break
    }
  }
  
  if maximaFound == 0 {
    return 0, false
  }
  
  return mid, true
}

