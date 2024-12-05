package main

import "slices"

type rule struct {
	first  int
	second int
}

func (r rule) isInOrder(list []int) bool {
	if !r.containsNumbers(list) {
		return true
	}

	firstLocated := false
	for _, nbr := range list {
		if nbr == r.first {
			firstLocated = true
		}
		if nbr == r.second {
			if firstLocated {
				return true
			} else {
				return false
			}
		}
	}
	return true
}

func (r rule) containsNumbers(numbers []int) bool {
	return slices.Contains(numbers, r.first) && slices.Contains(numbers, r.second)
}
