package ch03

// Sort stack in ascending order (min values on top)
// Use temporary stack to sort values in descending
// order. Then reverse it into original stack.
func Sort(stack *Stack) {
	sorted := new(Stack)

	// fill in sorted stack in descending order
	for !stack.IsEmpty() {
		tmp, _ := stack.Pop()

		for !sorted.IsEmpty() {
			top, _ := sorted.Peek()
			if top < tmp {
				break
			}

			top, _ = sorted.Pop()
			stack.Push(top)
		}

		sorted.Push(tmp)
	}

	// reverse sorted stack
	for !sorted.IsEmpty() {
		top, _ := sorted.Pop()
		stack.Push(top)
	}
}
