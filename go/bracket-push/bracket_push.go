package brackets

const testVersion = 3

func Bracket(input string) (res bool, err error) {
	stack := &Stack{}
	for _, c := range input {
		r := rune(c)
		cb, ok := closeBracket(r)
		if ok {
			stack.Push(cb)
		}
		if isCloseBracket(r) {
			p := stack.Pop()
			if p == nil || r != p.(rune) {
				return false, nil
			}
		}
	}

	if stack.Len() == 0 {
		res = true
	}
	return
}
func isCloseBracket(r rune) bool {
	if r == '}' || r == ']' || r == ')' {
		return true
	}
	return false
}

func closeBracket(open rune) (close rune, ok bool) {
	switch open {
	case '(':
		close = ')'
		ok = true
	case '[':
		close = ']'
		ok = true
	case '{':
		close = '}'
		ok = true
	}
	return
}

type Stack struct {
	top  *Element
	size int
}

type Element struct {
	value interface{} // All types satisfy the empty interface, so we can store anything here.
	next  *Element
}

// Return the stack's length
func (s *Stack) Len() int {
	return s.size
}

// Push a new element onto the stack
func (s *Stack) Push(value interface{}) {
	s.top = &Element{value, s.top}
	s.size++
}

// Remove the top element from the stack and return it's value
// If the stack is empty, return nil
func (s *Stack) Pop() (value interface{}) {
	if s.size > 0 {
		value, s.top = s.top.value, s.top.next
		s.size--
		return
	}
	return nil
}
