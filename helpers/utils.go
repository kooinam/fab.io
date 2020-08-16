package helpers

// H is alias for map[string]interface{}
type H map[string]interface{}

// Range used to create slice
func Range(min int, max int) []int {
	ls := []int{}

	for i := min; i < max; i++ {
		ls = append(ls, i)
	}

	return ls
}

// Times used to iterate n times with handler
func Times(n int, handler func(int) bool) {
	for i := range Range(0, n) {
		shouldContinue := handler(i)

		if !shouldContinue {
			break
		}
	}
}

// Find used to find an element in a slice
func Find(ls []interface{}, predicate func(interface{}) bool) interface{} {
	var found interface{}

	for el := range ls {
		if predicate(el) {
			found = el

			break
		}
	}

	return found
}
