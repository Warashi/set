package set

// Interface is string set interface
type Interface interface {
	Empty()
	Has(string) bool
	Len() int
	Items() <-chan string
	Add(string)
	Delete(string)
	Update(Interface)
}

func And(dst, a, b Interface) Interface {
	dst.Empty()
	for i := range a.Items() {
		if b.Has(i) {
			dst.Add(i)
		}
	}
	return dst
}

func Not(dst, a, b Interface) Interface {
	dst.Empty()
	for i := range a.Items() {
		if !b.Has(i) {
			dst.Add(i)
		}
	}
	return dst
}

func Or(dst, a, b Interface) Interface {
	dst.Empty()
	dst.Update(a)
	dst.Update(b)
	return dst
}

func Xor(dst, a, b Interface) Interface {
	tmp1, tmp2 := NewPatricia(), NewPatricia()
	return Or(dst, Not(tmp1, a, b), Not(tmp2, b, a))
}
