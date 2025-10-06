package main

import (
	"fmt"
)

/**
Given a list of words from an alien language dictionary, where the words are sorted lexicographically according to this alien language, derive the order of the characters in the alien language.
Example:
Input:
words = ["wrt", "wrf", "er", "ett", "rftt"]
Output: "wertf"
**/

type node struct {
	Data  string
	Child string // would have to handle it better if I used *Node here while traversing the linked list
}

type alienLanguageConfig struct {
	hashMap map[string]*node
	arr     []string
}

func NewAlienLanugageConfig(arr []string) *alienLanguageConfig {
	return &alienLanguageConfig{
		hashMap: make(map[string]*node),
		arr:     arr,
	}
}

func (c *alienLanguageConfig) generateOrderedLetters() {

	alienLanguageConfig := &alienLanguageConfig{
		hashMap: make(map[string]*node),
		arr:     c.arr,
	}

	for i := 0; i < len(c.arr)-1; i++ {
		word1 := c.arr[i]
		word2 := c.arr[i+1]
		// minLength := min(len(word1), len(word2))

		alienLanguageConfig.compareAndUpdateHashMap(word1, word2)
	}

	alienLanguageConfig.printHashMap()
}

func (c *alienLanguageConfig) compareAndUpdateHashMap(word1, word2 string) {

	if word1[0] == word2[0] && (len(word1) > 1 || len(word2) > 1) {
		c.compareAndUpdateHashMap(word1[1:], word2[1:])
		return
	} else {
		if c.hashMap[word2[0:1]] == nil {
			c.hashMap[word2[0:1]] = &node{Data: word2[0:1]}
		}

		if c.hashMap[word1[0:1]] == nil {
			c.hashMap[word1[0:1]] = &node{Data: word1[0:1], Child: word2[0:1]}
		} else {
			c.hashMap[word1[0:1]].Child = word2[0:1]
		}
	}
}

func (c *alienLanguageConfig) printHashMap() {
	localNode := c.hashMap[c.arr[0][0:1]]
	for localNode != nil {
		fmt.Printf("Node: %s\n", localNode.Data)
		localNode = c.hashMap[localNode.Child]
	}
}
