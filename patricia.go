package set

import (
	"bytes"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getCommonLength(a, b string) int {
	var i int
	for i = 0; i < min(len(a), len(b)); i++ {
		if a[i] != b[i] {
			break
		}
	}
	return i
}

type patricia struct {
	path     string
	parent   *patricia
	children []*patricia
	flag     bool
}

type stack []*patricia

func (s stack) Len() int {
	return len(s)
}
func (s *stack) Push(p *patricia) {
	*s = append(*s, p)
}
func (s *stack) Pop() *patricia {
	p := (*s)[s.Len()-1]
	*s = (*s)[:s.Len()-1]
	return p
}

func newPatricia(paths ...string) *patricia {
	s := &patricia{}
	for _, p := range paths {
		s.Add(p)
	}
	return s
}

func (s *patricia) fullPATH() string {
	plen := 0
	var list []string
	for {
		plen += len(s.path)
		list = append(list, s.path)
		if s.parent == nil {
			break
		}
		s = s.parent
	}

	buf := bytes.NewBuffer(make([]byte, 0, plen))
	for i := len(list) - 1; i >= 0; i-- {
		buf.WriteString(list[i])
	}
	return buf.String()
}

func (s *patricia) Empty() {
	*s = patricia{}
}

func (s *patricia) Has(path string) bool {
	if path == s.path {
		return s.flag
	}
	if len(path) < len(s.path) {
		return false
	}
	idx := len(s.path)
	if path[:idx] != s.path {
		return false
	}

	path = path[idx:]
	for _, c := range s.children {
		if path[0] == c.path[0] {
			return c.Has(path)
		}
	}
	return false
}

func (s *patricia) Len() int {
	l := 0
	var st stack
	st.Push(s)
	for st.Len() > 0 {
		n := st.Pop()
		if n.flag {
			l++
		}
		for _, c := range n.children {
			st.Push(c)
		}
	}
	return l
}

func (s *patricia) Items() <-chan string {
	ch := make(chan string)
	go func() {
		var st stack
		st.Push(s)
		for st.Len() > 0 {
			n := st.Pop()
			if n.flag {
				ch <- n.fullPATH()
			}
			for _, c := range n.children {
				st.Push(c)
			}
		}
		close(ch)
	}()
	return ch
}

func (s *patricia) Add(path string) {
walk:
	for {
		idx := getCommonLength(path, s.path)
		switch {
		case idx == len(s.path) && idx == len(path):
			s.flag = true
			return
		case idx < len(s.path): // split node
			c := &patricia{
				path:     s.path[idx:],
				parent:   s,
				children: s.children,
				flag:     s.flag,
			}
			s.path = s.path[:idx]
			s.children = []*patricia{c}
			s.flag = idx == len(path)
			return

		case idx < len(path):
			path = path[idx:]
			if idx == len(s.path) { // hit middle node, continue to search children
				for _, c := range s.children {
					if c.path[0] == path[0] {
						s = c
						continue walk
					}
				}
			}

			c := &patricia{
				path:   path,
				parent: s,
				flag:   true,
			}
			s.children = append(s.children, c)
			return
		}

	}
}

func (s *patricia) mergeChild() {
	if len(s.children) != 1 { // 子供がひとつの場合にだけマージ
		return
	}
	if s.flag { // 親がエントリを表さない場合のみマージ
		return
	}

	c := s.children[0]
	s.path += c.path
	s.children = c.children
	s.flag = c.flag
}

func (s *patricia) delete() {
	s.flag = false
	if len(s.children) > 0 {
		return
	}
	if s.parent == nil {
		return
	}

	p := s.parent
	children := make([]*patricia, 0, len(p.children)-1)
	for _, c := range p.children {
		if c != s {
			children = append(children, c)
		}
	}
	p.children = children
	p.mergeChild()
}

func (s *patricia) Delete(path string) {
	if path == s.path {
		s.delete()
		return
	}
	if len(path) < len(s.path) {
		return
	}
	idx := len(s.path)
	if path[:idx] != s.path {
		return
	}

	path = path[idx:]
	for _, c := range s.children {
		if path[0] == c.path[0] {
			c.Delete(path)
		}
	}
}

func (s *patricia) Update(u Interface) {
	for i := range u.Items() {
		s.Add(i)
	}
}
