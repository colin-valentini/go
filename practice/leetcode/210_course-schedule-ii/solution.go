package leetcode

// LeetCode #210.
// Link: https://leetcode.com/problems/course-schedule-ii/

// There are a total of numCourses courses you have to take, labeled from 0 to
// numCourses - 1. You are given an array prerequisites where
// prerequisites[i] = [ai, bi] indicates that you must take course bi first if
// you want to take course ai.

// For example, the pair [0, 1], indicates that to take course 0 you have to
// first take course 1. Return the ordering of courses you should take to finish
// all courses. If there are many valid answers, return any of them. If it is
// impossible to finish all courses, return an empty array.

// Example 1.
// Input: numCourses = 2, prerequisites = [[1,0]]
// Output: [0,1]
// Explanation: There are a total of 2 courses to take. To take course 1 you
// should have finished course 0. So the correct course order is [0,1].

// Example 2.
// Input: numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
// Output: [0,2,1,3]
// Explanation: There are a total of 4 courses to take. To take course 3 you
// should have finished both courses 1 and 2. Both courses 1 and 2 should be
// taken after you finished course 0.
// So one correct course order is [0,1,2,3]. Another correct ordering is [0,2,1,3].

// Example 3.
// Input: numCourses = 1, prerequisites = []
// Output: [0]

// Constraints:
// - 1 <= numCourses <= 2000
// - 0 <= prerequisites.length <= numCourses * (numCourses - 1)
// - prerequisites[i].length == 2
// - 0 <= ai, bi < numCourses
// - ai != bi
// - All the pairs [ai, bi] are distinct.

// FindOrder is a solution to the course schedule ii problem.
func FindOrder(numCourses int, prerequisites [][]int) []int {
	return findOrder(numCourses, prerequisites)
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	// Build adjacency list from prerequisites
	adjList := buildAdjacencyList(numCourses, prerequisites)

	// DFS from each node:
	//  - Maintain a set that indicates which nodes are in our current DFS path
	//  - If we revisit a node in our current path, then a cycle exists and we
	//    know that there is no solution (early return)
	//  - When we get to the end of the path, we put that final node into our
	//    "visited" set, and append it to our course order slice, then return
	//  - DFS function calls higher in the stack will do the same as we return
	//    upwards.
	order := make([]int, 0, numCourses)
	visited := make(map[int]struct{}, numCourses)
	for node := range adjList {
		path := make(map[int]struct{})
		if !dfs(node, adjList, path, visited, &order) {
			return []int{}
		}
	}
	return order
}

func buildAdjacencyList(numCourses int, prerequisites [][]int) [][]int {
	adjList := make([][]int, numCourses)
	for _, pair := range prerequisites {
		node, prereq := pair[0], pair[1]
		if adjList[node] == nil {
			adjList[node] = []int{}
		}
		adjList[node] = append(adjList[node], prereq)
	}
	return adjList
}

func dfs(node int, adjList [][]int, path, visited map[int]struct{}, order *[]int) bool {
	// If we revisit a node in our current path, then we know there's a cycle
	// and we can return false for "no solution"
	if _, ok := path[node]; ok {
		return false
	}
	// If we've already visited this node, then we know it's been added to our
	// order list and we can avoid re-searching it's sub-paths.
	if _, ok := visited[node]; ok {
		return true
	}
	path[node] = struct{}{}
	for _, prereq := range adjList[node] {
		if !dfs(prereq, adjList, path, visited, order) {
			return false
		}
	}
	visited[node] = struct{}{}
	delete(path, node)
	*order = append(*order, node)
	return true
}
