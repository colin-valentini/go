package algoexpert

// See https://www.algoexpert.io/questions/knapsack-problem

// You're given an array of arrays where each subarray holds two integer values
// and represents an item; the first integer is the item's value, and the second
// integer is the item's weight. You're also given an integer representing the
// maximum capacity of a knapsack that you have.

// Your goal is to fit items in your knapsack without having the sum of their
// weights exceed the knapsack's capacity, all the while maximizing their
// combined value. Note that you only have one of each item at your disposal.

// Write a function that returns the maximized combined value of the items that
// you should pick as well as an array of the indices of each item picked.

// If there are multiple combinations of items that maximize the total value in
// the knapsack, your function can return any of them.

// Sample Input
// items = [[1, 2], [4, 3], [5, 6], [6, 7]]
// capacity = 10

// Sample Output
// [10, [1, 3]] // items [4, 3] and [6, 7]

// KnapsackProblem is a solution to the Knapsack problem
func KnapsackProblem(items [][]int, capacity int) []any {
	return newKnapsackProblemSolver(items, capacity).Solve()
}

type knapsackProblemSolver struct {
	items    [][]int
	capacity int
}

func newKnapsackProblemSolver(items [][]int, capacity int) *knapsackProblemSolver {
	return &knapsackProblemSolver{items, capacity}
}

func (s *knapsackProblemSolver) Solve() []any {
	// TODO

	// 1. Create a matrix for subsolutions
	//  - Each row represents an item, with the first row being an empty item
	//  - Each column represents a capacity, with the last column being the
	//    specific capacity for this problem
	subSolutions := make([][]int, len(s.items)+1)

	for i := range subSolutions {
		subSolutions[i] = make([]int, s.capacity+1)

		// Can skip the empty set row, leave filled as zero
		if i == 0 {
			continue
		}
		itemIdx := i - 1
		value, weight := s.items[itemIdx][0], s.items[itemIdx][1]
		for subCapacity := range subSolutions[i] {
			// This item's weight doesn't fit, so just take the solution
			// from above which we know doesn't include the current item.
			solutionAbove := subSolutions[i-1][subCapacity]
			if weight > subCapacity {
				subSolutions[i][subCapacity] = solutionAbove
				continue
			}
			// This item's weight DOES fit, so we consider the better of what
			// we'd get with this item, versus what we had without it.
			solutionWithoutThisItem := subSolutions[i-1][subCapacity-weight]
			if solutionWithoutThisItem+value <= solutionAbove {
				subSolutions[i][subCapacity] = solutionAbove
				continue
			} else {
				subSolutions[i][subCapacity] = solutionWithoutThisItem + value
			}
		}
	}
	totalValue := subSolutions[len(subSolutions)-1][s.capacity]

	// Backtrack to find which items were included
	v := totalValue
	itemIdx, subCapacity := len(s.items), s.capacity
	itemIndices := make([]int, 0, len(s.items))
	for v > 0 {
		// If value in this cell is greater than the value in the cell
		// above, we know that this item was definitely included.
		if v > subSolutions[itemIdx-1][subCapacity] {
			itemIndices = append(itemIndices, itemIdx-1)
			itemWeight := s.items[itemIdx-1][1]
			subCapacity -= itemWeight
		}
		itemIdx--
		v = subSolutions[itemIdx][subCapacity]
	}

	// AlgoExpert requires this untyped return, which is disappointing.
	return []any{
		totalValue,
		itemIndices,
	}
}
