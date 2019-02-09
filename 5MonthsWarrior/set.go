// set is used to store only unique elements
// set is implemented using a hash table so elemets are not stored in sequential order

func main() {
  s1 := set.New()
  s1.Insert()
  s1.Insert(2)
  fmt.Println(st.Has(1))
  
  fmt.Println(st.Has(3))
}

type set map[interface{}]bool

func (s *Set) Add (key interface{}){
  (*s)[key] = true
}

func (s *Set) Remove (key interface{}){
  delete((*s), key)
}

func (s *Set) Find(key interface{}){
  return (*s)[key]
}

func main(){
  mp := make(Set)
  mp.Add("A")
  mp.Add("B")
  mp.Add("A")
  fmt.Println(mp.FInd("A"))
  fmt.Println(mp.FInd("B"))
  fmt.Println(mp.FInd("C"))
}
