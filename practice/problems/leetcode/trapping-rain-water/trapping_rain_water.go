package leetcode

// LeetCode #42.
// Link: https://leetcode.com/problems/trapping-rain-water/

// Given n non-negative integers representing an elevation map where the width of each bar is 1,
// compute how much water it can trap after raining.

// Example 1:
// Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
// Output: 6
// Explanation: The above elevation map (black section) is represented by array
// [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are
// being trapped.

// 3              █
// 2      █ • • • █ █ • █
// 1  █ • █ █ • █ █ █ █ █ █
// 0 ‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾

// Example 2:
// Input: height = [4,2,0,3,2,5]
// Output: 9

// 5            █
// 4  █ • • • • █
// 3  █ • • █ • █
// 2  █ █ • █ █ █
// 1  █ █ • █ █ █
// 0 ‾‾‾‾‾‾‾‾‾‾‾‾

// Example 3:
// Input: height = [2,0,0,1,0,1,4,0,1,0,1,0,1,0,1,0,3]
// Output: 31

// 4              █
// 3              █ • • • • • • • • • █
// 2  █ • • • • • █ • • • • • • • • • █
// 1  █ • • █ • █ █ • █ • █ • █ • █ • █
// 0 ‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾

// Example 4:
// Input: height = [0,0,0,1,0,1,4,0,1,0,1,0,1,0,1,0,3]
// Output: 24

// 4              █
// 3              █
// 2              █
// 1        █ • █ █ • █ • █ • █ • █ • █
// 0 ‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾

// Example 5:
// Input: height = [2,0,0,1,0,1,4,0,1,0,3,0,1,0,2,0,1]
// Output: 22

// 4              █
// 3              █ • • • █   
// 2  █ • • • • • █ • • • █ • • • █  
// 1  █ • • █ • █ █ • █ • █ • █ • █ • █
// 0 ‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾

// Example 6:
// Input: height = [1,2,3,4,5,6]
// Output: 0

// 6            █
// 5          █ █
// 4        █ █ █
// 3      █ █ █ █
// 2    █ █ █ █ █
// 1  █ █ █ █ █ █
// 0 ‾‾‾‾‾‾‾‾‾‾‾‾

// Example 7:
// Input: height = [6,5,4,3,2,1]
// Output: 0

// 6  █          
// 5  █ █        
// 4  █ █ █       
// 3  █ █ █ █    
// 2  █ █ █ █ █  
// 1  █ █ █ █ █ █
// 0 ‾‾‾‾‾‾‾‾‾‾‾‾

// Solution outline:
// If, for each position in the height slice, we knew the tallest height
// to the left, and the tallest height to the right then we could easily
// calculate the water column above any given position.
//
// The amount of water above each position would just be the min of the
// tallest to the left, and the tallest to the right less the height
// at the current position. Only add this value if it's positive.
//
// Time: O(N) | Space: O(N)
func TrappingRainWater(height []int) int {
	return trappingRainWater(height)
}

func trappingRainWater(height []int) int {
	// Need at least three elements to form a basin.
	if len(height) < 3 {
		return 0
	}
	
	// One pass in reverse order to get the "tallest to the right" slice.
	// We could also generate the "tallest to the left" slice but since
	// that requires forward traversal, and we are going to forward traverse
	// anyway, we don't actually need to store it ahead of time.
	maxRight := newTallestToTheRight(height)

	// One more pass to calculate the water area.
	water, maxLeft := 0, 0
	for i, h := range height {
		water += max(min(maxLeft, maxRight[i]) - h, 0)
		if h > maxLeft {
			maxLeft = h
		}
	}
	return water
}

func newTallestToTheRight(height []int) ([]int) {
	right := make([]int, len(height))
	maxRight := 0
	for i := len(height)-1; i >= 0; i-- {
		right[i] = maxRight
		if height[i] > maxRight {
			maxRight = height[i]
		}
	}
	return right
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}