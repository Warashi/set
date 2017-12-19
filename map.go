package set

type Map map[string]struct{}

func NewMap(items ...string) *Map {
	s := make(Map, len(items)*2)
	for _, i := range items {
		s.Add(i)
	}
	return &s
}

func (s *Map) Empty() {
	*s = make(map[string]struct{})
}

func (s *Map) Has(i string) bool {
	_, ok := (*s)[i]
	return ok
}

func (s *Map) Len() int {
	return len(*s)
}

func (s *Map) Items() <-chan string {
	ch := make(chan string)
	go func() {
		for i := range *s {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func (s *Map) Add(i string) {
	(*s)[i] = struct{}{}
}

func (s *Map) Delete(i string) {
	delete(*s, i)
}

func (s *Map) Update(a Interface) {
	for i := range a.Items() {
		s.Add(i)
	}
}
