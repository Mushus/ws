package marker

type runeNode struct {
	rune rune
	next *runeNode
}
type runeStack struct {
	start *runeNode
	count int64
}

func (r *runeStack) pushEmpty() {
	rn := &runeNode{
		rune: empty,
		next: r.start,
	}
	r.start = rn
}

func (r *runeStack) push(c rune) {
	rn := &runeNode{
		rune: c,
		next: r.start,
	}
	r.start = rn
	r.count++
}

func (r *runeStack) pop() rune {
	if r.start == nil {
		return eof
	}
	c := r.start.rune
	r.start = r.start.next
	if c != empty {
		r.count--
	}
	return c
}

func (r *runeStack) clear() {
	r.start = nil
	r.count = 0
}

func (r *runeStack) takeout() []rune {
	list := make([]rune, r.count)
	for {
		c := r.pop()
		if c == eof {
			break
		}
		if c == empty {
			continue
		}
		list[r.count] = c
	}
	return list
}
