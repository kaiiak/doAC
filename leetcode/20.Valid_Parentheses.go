// https://leetcode.com/problems/valid-parentheses/

func isValid(s string) bool {
	stack := NewStack()
	for i := 0; i < len(s); i++ {
		last := stack.Last()
		match := false
		switch s[i] {
		case ')':
			if last == '(' {
				match = true
			}
		case '}':
			if last == '{' {
				match = true
			}

		case ']':
			if last == '[' {
				match = true
			}
		}
		if match {
			stack.Pop()
		} else {
			stack.Push(s[i])
		}
		// fmt.Println(string(stack))
	}
	return stack.IsEmpty()
}

type Stack []byte

func NewStack() Stack {
	return []byte{}
}

func (s *Stack) Push(b byte) {
	*s = append(*s, b)
}

func (s *Stack) Pop() (b byte) {
	if s.IsEmpty() {
		return ' '
	}
	b = s.Last()
	*s = (*s)[:len(*s)-1]
	return
}

func (s Stack) Last() byte {
	if s.IsEmpty() {
		return byte(' ')
	}
	return s[len(s)-1]
}

func (s Stack) IsEmpty() bool {
	return len(s) == 0
}
