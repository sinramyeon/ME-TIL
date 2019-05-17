// Counters are used to count the number of occurrence of values in a List

type Counter map[interface{}] int

func (* Counter) Add (key interface{}){
  (*s)[key] += 1
}

func (s *Counter) Find(key interface{}) bool {
  _, ok := (*s)[key]
  return ok
}

func (s *Counter) Get(key interface{}) (int, bool) {
  val, ok := (*s)[key]
  return val, ok
}
