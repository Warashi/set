package set

// Interface is string set interface
type Interface interface {
	Has(string) bool
	Len() int
	Items() <-chan string
	Add(string)
	Delete(string)
	Update(Interface)
}

type set map[string]struct{}

// New returns new set based on map[string]struct{}
func New(items ...string) Interface {
	s := make(set)
	for _, item := range items {
		s[item] = struct{}{}
	}
	return &s
}

func (s *set) Has(i string) bool {
	_, ok := (*s)[i]
	return ok
}

func (s *set) Len() int {
	return len(*s)
}

func (s *set) Items() <-chan string {
	ch := make(chan string)
	go func() {
		for i := range *s {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func (s *set) Add(i string) {
	(*s)[i] = struct{}{}
}

func (s *set) Delete(i string) {
	delete(*s, i)
}

func (s *set) Update(u Interface) {
	for i := range u.Items() {
		s.Add(i)
	}
}

// And returns set which contains items in both a and b
func And(a, b Interface) Interface {
	s := New()
	for i := range a.Items() {
		if b.Has(i) {
			s.Add(i)
		}
	}
	return s
}

// Not returns set which contains items not in b but a
func Not(a, b Interface) Interface {
	s := New()
	for i := range a.Items() {
		if !b.Has(i) {
			s.Add(i)
		}
	}
	return s
}

// Or returns set which contains items in either a or b
func Or(a, b Interface) Interface {
	s := New()
	s.Update(a)
	s.Update(b)
	return s
}

// Xor returns set which contains items in either (not in b but a) or (not in a but b)
func Xor(a, b Interface) Interface {
	return Or(Not(a, b), Not(b, a))
}
