package leetcode

// // LeetCode #207.
// Link: https://leetcode.com/problems/course-schedule/

// There are a total of numCourses courses you have to take, labeled from 0 to
// numCourses - 1. You are given an array prerequisites where
// prerequisites[i] = [ai, bi] indicates that you must take course bi first if
// you want to take course ai.
// - For example, the pair [0, 1], indicates that to take course 0 you have to
// 	 first take course 1.

// Return true if you can finish all courses. Otherwise, return false.

// Example 1.
// Input: numCourses = 2, prerequisites = [[1,0]]
// Output: true
// Explanation: There are a total of 2 courses to take.
// To take course 1 you should have finished course 0. So it is possible.

// Example 2.
// Input: numCourses = 2, prerequisites = [[1,0],[0,1]]
// Output: false
// Explanation: There are a total of 2 courses to take.
// To take course 1 you should have finished course 0, and to take course 0 you
// should also have finished course 1. So it is impossible.

// Constraints:
// - 1 <= numCourses <= 2000
// - 0 <= prerequisites.length <= 5000
// - prerequisites[i].length == 2
// - 0 <= ai, bi < numCourses
// - All the pairs prerequisites[i] are unique.

// CanFinish is a solution to the course schedule problem
func CanFinish(numCourses int, prerequisites [][]int) bool {
	return canFinish(numCourses, prerequisites)
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	// 1. Build an adjacency list based on the prerequisite pairs (edges)
	adjList := buildAdjacencyList(numCourses, prerequisites)

	// 2. DFS the graph, looking for a cycle. If we find no cycles then we
	// know that the graph _could_ be top-sorted which is enough to know
	// that all courses could be taken.
	visited := make(map[int]struct{})
	for course := range adjList {
		path := make(map[int]struct{})
		if !dfs(course, visited, path, adjList) {
			return false
		}
	}
	return true
}

func buildAdjacencyList(numCourses int, prerequisites [][]int) [][]int {
	adjList := make([][]int, numCourses)
	for _, pair := range prerequisites {
		course, prereq := pair[0], pair[1]
		if adjList[course] == nil {
			adjList[course] = []int{prereq}
		} else {
			adjList[course] = append(adjList[course], prereq)
		}
	}
	return adjList
}

func dfs(course int, visited, path map[int]struct{}, adjList [][]int) bool {
	if _, ok := visited[course]; ok {
		return true
	}
	if _, ok := path[course]; ok {
		return false
	}
	path[course] = struct{}{}
	for _, prereq := range adjList[course] {
		if !dfs(prereq, visited, path, adjList) {
			return false
		}
	}
	visited[course] = struct{}{}
	delete(path, course)
	return true
}
