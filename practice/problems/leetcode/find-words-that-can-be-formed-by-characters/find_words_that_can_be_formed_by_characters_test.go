package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestCountCharacters runs test cases for CountCharacters.
func TestCountCharacters(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(6, CountCharacters([]string{"cat", "bt", "hat", "tree"}, "atach"))
	assert.Equal(10, CountCharacters([]string{"hello", "world", "leetcode"}, "welldonehoneyr"))
	assert.Equal(0, CountCharacters(
		[]string{
			"dyiclysmffuhibgfvapygkorkqllqlvokosagyelotobicwcmebnpznjbirzrzsrtzjxhsfpiwyfhzyonmuabtlwin", 
			"ndqeyhhcquplmznwslewjzuyfgklssvkqxmqjpwhrshycmvrb", 
			"ulrrbpspyudncdlbkxkrqpivfftrggemkpyjl", 
			"boygirdlggnh",
			"xmqohbyqwagkjzpyawsydmdaattthmuvjbzwpyopyafphx",
			"nulvimegcsiwvhwuiyednoxpugfeimnnyeoczuzxgxbqjvegcxeqnjbwnbvowastqhojepisusvsidhqmszbrnynkyop",
			"hiefuovybkpgzygprmndrkyspoiyapdwkxebgsmodhzpx",
			"juldqdzeskpffaoqcyyxiqqowsalqumddcufhouhrskozhlmobiwzxnhdkidr",
			"lnnvsdcrvzfmrvurucrzlfyigcycffpiuoo",
			"oxgaskztzroxuntiwlfyufddl",
			"tfspedteabxatkaypitjfkhkkigdwdkctqbczcugripkgcyfezpuklfqfcsccboarbfbjfrkxp",
			"qnagrpfzlyrouolqquytwnwnsqnmuzphne",
			"eeilfdaookieawrrbvtnqfzcricvhpiv",
			"sisvsjzyrbdsjcwwygdnxcjhzhsxhpceqz",
			"yhouqhjevqxtecomahbwoptzlkyvjexhzcbccusbjjdgcfzlkoqwiwue",
			"hwxxighzvceaplsycajkhynkhzkwkouszwaiuzqcleyflqrxgjsvlegvupzqijbornbfwpefhxekgpuvgiyeudhncv",
			"cpwcjwgbcquirnsazumgjjcltitmeyfaudbnbqhflvecjsupjmgwfbjo",
			"teyygdmmyadppuopvqdodaczob",
			"qaeowuwqsqffvibrtxnjnzvzuuonrkwpysyxvkijemmpdmtnqxwekbpfzs",
			"qqxpxpmemkldghbmbyxpkwgkaykaerhmwwjonrhcsubchs",
		},
		"usdruypficfbpfbivlrhutcgvyjenlxzeovdyjtgvvfdjzcmikjraspdfp",
		),
	)
}