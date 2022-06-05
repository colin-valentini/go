package leetcode

// LeetCode #269.
// https://leetcode.com/problems/alien-dictionary/

// There is a new alien language which uses the latin alphabet. However, the
// order among letters are unknown to you. You receive a list of non-empty words
// from the dictionary, where words are sorted lexicographically by the rules of
// this new language. Derive the order of letters in this language.

// Example 1:
// Input: ["wrt", "wrf", "er", "ett", "rftt"]
// Output: "wertf"

// Example 2:
// Input: ["z", "x"]
// Output: "zx"

// Example 3:
// Input: ["z", "x", "z"]
// Output: ""
// Explanation: The order is invalid, so return "".

// Constraints:
// - All letters are in English, lowercase.
// - If a is a prefix of b, then a must appear before b in the given dictionary.
// - If the order is invalid, return an empty string.
// - If there may be multiple valid order of letters, return any one of them.

func AlienOrder(words []string) string {
	return alienOrder(words)
}

func alienOrder(words []string) string {
	edges := discoverEdges(words)
	adjList := buildAdjacencyList(edges)
	order := []byte{}
	visited := make(map[byte]struct{})
	for char := range adjList {
		path := make(map[byte]struct{})
		if !dfs(char, adjList, path, visited, &order) {
			// Return empty string if a cycle is detected
			return ""
		}
	}
	return string(order)
}

func dfs(char byte, adjList map[byte][]byte, path, visited map[byte]struct{}, order *[]byte) bool {
	if _, ok := path[char]; ok {
		return false
	}
	if _, ok := visited[char]; ok {
		return true
	}
	path[char] = struct{}{}
	for _, lesserChar := range adjList[char] {
		if !dfs(lesserChar, adjList, path, visited, order) {
			return false
		}
	}
	visited[char] = struct{}{}
	delete(path, char)
	*order = append(*order, char)
	return true
}

func discoverEdges(words []string) [][]byte {
	edges := [][]byte{}
	for i := 1; i < len(words); i++ {
		a, b := words[i-1], words[i]
		for j := 0; j < min(len(a), len(b)); j++ {
			// This only works because we're using ascii (otherwise you'd need
			// to handle utf-8 multi-byte characters)
			if a[j] != b[j] {
				edges = append(edges, []byte{b[j], a[j]})
				break
			}
		}
		// We get here if one string is a prefix of other, which does not
		// give us any new information about the order of characters
	}
	return edges
}

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

func buildAdjacencyList(edges [][]byte) map[byte][]byte {
	adjList := make(map[byte][]byte)
	for _, pair := range edges {
		greaterChar := pair[0]
		if _, ok := adjList[greaterChar]; !ok {
			adjList[greaterChar] = []byte{}
		}
		adjList[greaterChar] = append(adjList[greaterChar], pair[1])
	}
	return adjList
}
