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

// New returns new set based on map[string]struct{}
func New(items ...string) Interface {
	return newPatricia(items...)
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
