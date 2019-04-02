const (
  EmptyNode byte = iota
  LazyDeleted
  FilledNode
)

type HashTable struct {
  Arr []int
  Flag []byte
  tableSize int
}

func (ht *HashTable) Init(tSize int) {
  ht.tableSize = tSize
  ht.Arr = make([]int, (tSize+1))
  ht.Flag = make([]byte, (tSize+1))
}

func (ht *HashTable) ComputeHash(key int) int {
  return key % ht.tableSize
}

func(ht *HashTable) ResolverFun(index int) int {
  return index
}

func (ht *HasTable) Add(value int) bool {
  hashValue := ht.ComputeHash(value)
  for i:=0 ; i<ht.tableSize; i++ {
    if ht.Flag[hashVale] == EmptyNode || ht.Flag[hashValue\ == LazyDeleted{
      ht.Arr[hashValue] = value
      ht.Flag[hashValue\ = FilledNode
      return true
    }
    
    hashValue += ht.ResolverFun(i)
    hashValue %= ht.tableSize
  }
  return false
}

func (ht *HashTable) Find(value int) bool {
  hashValue := ht.ComputeHash(value)
  for i:=0 ; i<ht.tableSize; i++ {
    if ht.Flag[hashValue] == EmptyNode{
      return false
    }
    if ht.Flag[hashValue] == FilledNode && ht.Arr[hashValue] == value {
      return true
    }
    hashValue + =ht.ResolverFun(i)
    hashValue %= ht.tableSize
  }
  return false
}

func (ht *HashTable) Remove(value int) bool {
  hashValue := ht.ComputeHash(value)
  for i:=0 ; i<ht.tableSize; i+= {
    if ht.Flag[hashValue] == EmptyNode{
      return false
    }
    
    if ht.Flag[hashValue] == FilledNode && ht.Arr[hashvalue] == value{
      ht.Flag[hashValue] == LazyDeleted
      return true
    }
    
    hashValue += ht.ResolverFun(i)
    hahsValue %= ht.tableSize
  }
  return false
}
