package leetcode

// LeetCode #11.
// Link: https://leetcode.com/problems/container-with-most-water/

// You are given an integer array height of length n. There are n vertical lines drawn such that
// the two endpoints of the ith line are (i, 0) and (i, height[i]).

// Find two lines that together with the x-axis form a container, such that the container contains
// the most water.

// Return the maximum amount of water a container can store.

// Notice that you may not slant the container.

// Example 1:
// Input: height = [1,8,6,2,5,4,8,3,7]
// Output: 49
// Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case,
// the max area of water (blue section) the container can contain is 49.

// Example 2:
// Input: height = [1,1]
// Output: 1

// ContainerWithMostWater is a solution the container with most water problem.
func ContainerWithMostWater(height []int) int {
	return twoPointerContainerWithMostWater(height)
}

// twoPointerContainerWithMostWater is a two pointer solution to the container with
// most water problem. It runs in O(N) time, and O(1) space.
func twoPointerContainerWithMostWater(height []int) int {
	maxArea := 0
	i, j := 0, len(height)-1
	for i <= j {
		area := waterArea(height, i, j)
		if area > maxArea {
			maxArea = area
		}
		if height[i] <= height[j] {
			i++
		} else {
			j--
		}
	}
	return maxArea
}

// waterArea returns the water area contained within the bars defined by
// positions `i` and `j` and the bars they represent in `heights``.
func waterArea(height []int, i, j int) int {
    w := j - i
	h := height[i]
    if height[j] < height[i] {
        h = height[j]
    }
    return w * h
}