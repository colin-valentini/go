package binarysearch

const notFound = -1

// BinarySearch is a simple implementation of a binary search algorithm.
// The input slice `in`, should be sorted, otherwise the behavior is undefined.
// If the target is found (using the given comparator function), the index of
// it's first occurrence is returned. If target is not found, the function
// returns -1.
func BinarySearch(in []any, compare func(i int) int) int {
	left, right := 0, len(in)-1
	if compare(left) < 0 || compare(right) > 0 {
		return notFound // Target is completely out of range.
	}
	idx := notFound
	for left <= right {
		mid := left + (right-left)/2
		if res := compare(mid); res == 0 {
			idx = mid
			right = mid - 1 // Continue searching the lower window.
		} else if res < 0 {
			right = mid - 1
		} else /* if res > 0 */ {
			left = mid + 1
		}
	}
	return idx
}
