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
	return &knapsackProblemSolver{items: items, capacity: capacity}
}

// Solve solves the knapsack problem and returns the solution.
func (s *knapsackProblemSolver) Solve() []any {
	solution := s.solve()
	// AlgoExpert (unfortunately) requires an untype return
	return []any{
		solution.totalValue,
		solution.itemIndices,
	}
}

type knapsackProblemSolution struct {
	totalValue  int
	itemIndices []int
}

func (s *knapsackProblemSolver) solve() knapsackProblemSolution {
	return s.solution(s.subSolutions())
}

func (s *knapsackProblemSolver) subSolutions() [][]int {
	sub := make([][]int, len(s.items)+1)
	for i := range sub {
		sub[i] = make([]int, s.capacity+1)

		// First row represents an empty item, which we know has no real
		// solution to any capacity (item of weight zero, and value zero).
		if i == 0 {
			continue
		}
		value, weight := s.items[i-1][0], s.items[i-1][1]
		for subCapacity := range sub[i] {
			// If the current item is too heavy to fit in the subproblem
			// capacity, then take the subsolution above.
			solutionAbove := sub[i-1][subCapacity]
			if weight > subCapacity {
				sub[i][subCapacity] = solutionAbove
				continue
			}
			// Solution at this cell is the better of what we'd get
			// with this item versus what we had without it.
			solutionWithItem := sub[i-1][subCapacity-weight] + value
			if solutionWithItem > solutionAbove {
				sub[i][subCapacity] = solutionWithItem
			} else {
				sub[i][subCapacity] = solutionAbove
			}
		}
	}
	return sub
}

func (s *knapsackProblemSolver) solution(sub [][]int) knapsackProblemSolution {
	solution := knapsackProblemSolution{}
	solution.totalValue = sub[len(s.items)][s.capacity]

	// Backtrack to find which items were included
	v := solution.totalValue
	itemIdx, subCapacity := len(s.items), s.capacity
	solution.itemIndices = make([]int, 0, len(s.items))
	for v > 0 {
		// If value in this cell is greater than the value in the cell
		// above, we know that this item was included in the subsolution.
		if v > sub[itemIdx-1][subCapacity] {
			solution.itemIndices = append(solution.itemIndices, itemIdx-1)
			itemWeight := s.items[itemIdx-1][1]
			subCapacity -= itemWeight
		}
		itemIdx--
		v = sub[itemIdx][subCapacity]
	}

	return solution
}
