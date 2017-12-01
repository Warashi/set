package set

type S map[string]struct{}

func New(items ...string) *S {
	s := make(S)
	for _, item := range items {
		s[item] = struct{}{}
	}
	return &s
}

func (s *S) Has(i string) bool {
	_, ok := (*s)[i]
	return ok
}

func (s *S) Add(i string) {
	(*s)[i] = struct{}{}
}

func (s *S) Update(u *S) {
	for i := range *u {
		s.Add(i)
	}
}

func And(a, b *S) *S {
	s := New()
	for i := range *a {
		if _, ok := (*b)[i]; ok {
			s.Add(i)
		}
	}
	return s
}

func Not(a, b *S) *S {
	s := New()
	for i := range *a {
		if _, ok := (*b)[i]; !ok {
			s.Add(i)
		}
	}
	return s
}

func Or(a, b *S) *S {
	s := New()
	s.Update(a)
	s.Update(b)
	return s
}

func Xor(a, b *S) *S {
	return Or(Not(a, b), Not(b, a))
}
