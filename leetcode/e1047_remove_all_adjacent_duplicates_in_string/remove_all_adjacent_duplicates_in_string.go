package removealladjacentduplicatesinstring

type myStack struct {
	stack []rune
}

func constructor() myStack {
	return myStack{
		stack: make([]rune, 0),
	}
}

func (s myStack) top() rune {
	if len(s.stack) > 0 {
		return s.stack[len(s.stack)-1]
	}

	return 0
}

func (s *myStack) push(r rune) {
	if (s != nil) && (s.stack != nil) {
		s.stack = append(s.stack, r)
	}
}

func (s *myStack) pop() rune {
	if len(s.stack) > 0 {
		top := s.top()
		s.stack = s.stack[:len(s.stack)-1]
		return top
	}

	return 0
}

func (s myStack) isEmpty() bool {
	return len(s.stack) == 0
}

// Time complexity: O(n)
// Space complexity: O(n)
func removeDuplicates(s string) string {
	stack := constructor()
	for _, r := range s {
		if stack.isEmpty() || (stack.top() != r) {
			stack.push(r)
		} else {
			stack.pop()

		}
	}

	result := []rune{}
	for !stack.isEmpty() {
		result = append(result, stack.pop())
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}
