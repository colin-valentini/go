package leetcode

// LeetCode #127.
// Link: https://leetcode.com/problems/word-ladder/

// A transformation sequence from word beginWord to word endWord using a
// dictionary wordList is a sequence of words beginWord -> s1 -> s2 -> ... -> sk
// such that:

// - Every adjacent pair of words differs by a single letter.
// - Every si for 1 <= i <= k is in wordList. Note that beginWord does not need
//   to be in wordList.
// - sk == endWord

// Given two words, beginWord and endWord, and a dictionary wordList, return the
// number of words in the shortest transformation sequence from beginWord to
// endWord, or 0 if no such sequence exists.

// Example 1.
// Input:
// 	beginWord = "hit", endWord = "cog",
// 	wordList = ["hot","dot","dog","lot","log","cog"]
// Output: 5
// Explanation: One shortest transformation sequence is
// "hit" -> "hot" -> "dot" -> "dog" -> cog", which is 5 words long.

// Example 2.
// Input:
// 	beginWord = "hit", endWord = "cog",
// 	wordList = ["hot","dot","dog","lot","log"]
// Output: 0
// Explanation: The endWord "cog" is not in wordList, therefore there is no
// valid transformation sequence.

// Constraints:
// - 1 <= beginWord.length <= 10
// - endWord.length == beginWord.length
// - 1 <= wordList.length <= 5000
// - wordList[i].length == beginWord.length
// - beginWord, endWord, and wordList[i] consist of lowercase English letters.
// - beginWord != endWord
// - All the words in wordList are unique.

// LadderLength is a solution to the word ladder problem
func LadderLength(beginWord string, endWord string, wordList []string) int {
	return ladderLength(beginWord, endWord, wordList)
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	// 1. Build an adjacency list of words in the word list that
	//    can be transformed between each other, i.e. an "edge" represents
	//    a valid transformation between two words.
	//  NOTE: If we notice that the endword is not in the wordlist, then we
	//  automatically know that no possible transformation exists.
	adjList := buildAdjacencyList(beginWord, wordList)
	if _, ok := adjList[endWord]; !ok {
		return 0
	}

	// 2. Use the graph (expressed as an adjacency list) to find the shortest
	//    path between the starting word and the ending word, and return the
	//    number of nodes in that path.
	return bfs(beginWord, endWord, adjList)
}

func buildAdjacencyList(beginWord string, wordList []string) map[string][]string {
	adjList := make(map[string][]string)
	// TODO: Can probably speed this up if we take advantage of the fact that
	// edges are bi-directional so knowing that word_i -> word_j implies
	// that word_j -> word_i
	for i, word := range append(wordList, beginWord) {
		// Skip words we've already processed (which is only going to be
		// `beginWord` if it's in the `wordList`` already)
		if _, ok := adjList[word]; ok {
			continue
		}
		adjList[word] = []string{}
		for j, otherWord := range wordList {
			// Skip self-edges
			if i == j {
				continue
			}
			// Can only do this naive character indexing because every word is
			// ascii, and each word is the exact same length.
			numDiffs := 0
			for k, char := range word {
				if byte(char) != otherWord[k] {
					numDiffs++
				}
				if numDiffs > 1 {
					break
				}
			}
			if numDiffs == 1 {
				adjList[word] = append(adjList[word], otherWord)
			}
		}
	}
	return adjList
}

func bfs(beginWord, endWord string, adjList map[string][]string) int {
	visited := make(map[string]struct{})
	queue := []string{beginWord}
	numWords := 1
	for len(queue) > 0 {
		for range queue {
			word := queue[0]
			queue = queue[1:]
			if word == endWord {
				return numWords
			}
			visited[word] = struct{}{}
			for _, neighbor := range adjList[word] {
				if _, ok := visited[neighbor]; ok {
					continue
				}
				queue = append(queue, neighbor)
			}
		}
		numWords++
	}
	// If we got here, we never arrived at endWord via graph traversal
	return 0
}
