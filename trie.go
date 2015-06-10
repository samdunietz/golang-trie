// A simple, unicode-compliant Trie. Supports insertion and removal, although 
// removal only changes the isWord flag on the relevant node, rather than 
// deleting the entire unnecessary subtrie from memory. 
package trie

import (
	// "fmt"
)

type Trie struct {
	root trieNode
}

// Initializes a new Trie from a slice of strings.
// Supports unicode characters.
func New(words []string) Trie {
	t := Trie{initNode()}
	t.AddAll(words)
	return t
}

// Returns true if the trie stores no words.
// Note that currently, this does not necessarily mean that the trie has no nodes,
// as words that were removed remain as nodes.
func (t *Trie) IsEmpty() bool {
	return len(t.Words()) == 0
}

// Adds a new word to the trie.
func (t *Trie) Add(word string) bool {
	return t.root.add([]rune(word))
}

// Adds all of a slice of words to the trie.
func (t *Trie) AddAll(words []string) bool {
	hasChanged := false
	for _, v := range words {
		if t.Add(v) {
			hasChanged = true
		}
	}

	return hasChanged
}

// Removes a word from the trie by changing the isWord flag on the relevant node.
func (t *Trie) Remove(word string) bool {
	return t.root.remove([]rune(word))
}


// Removes all from a slice of words.
func (t *Trie) RemoveAll(words []string) bool {
	hasChanged := false
	for _, v := range words {
		if t.Remove(v) {
			hasChanged = true
		}
	}

	return hasChanged
}

// Gets all the words stored in the trie.
func (t *Trie) Words() []string {
	return t.WordsWithPrefix("")
}

// Gets all the words in the trie that start with prefix.
func (t *Trie) WordsWithPrefix(prefix string) []string {
	return t.root.WordsWithPrefix([]rune(prefix), []rune(prefix))
}

// Returns true if a word exists in the trie.
func (t *Trie) ContainsWord(word string) bool {
	return t.root.ContainsWord([]rune(word))
}

type trieNode struct {
	isWord   bool
	children map[rune]trieNode
}

func initNode() trieNode {
	return trieNode{false, make(map[rune]trieNode)}
}

func (node *trieNode) add(str []rune) bool {
	if len(str) == 0 {
		isNewWord := !node.isWord
		node.isWord = true
		return isNewWord
	} else {
		firstLetter := str[0]
		rest := str[1:]
		var isNewWord bool
		if child, ok := node.children[firstLetter]; ok {
			isNewWord = child.add(rest)
			node.children[firstLetter] = child
			return isNewWord
		} else {
			subnode := initNode()
			isNewWord = subnode.add(rest)
			node.children[firstLetter] = subnode
		}

		return isNewWord
	}
}

func (node *trieNode) remove(str []rune) bool {
	if len(str) == 0 {
		wasWord := node.isWord
		node.isWord = false
		return wasWord
	} else {
		firstLetter := str[0]
		rest := str[1:]
		var wasWord bool
		if child, ok := node.children[firstLetter]; ok {
			wasWord = child.remove(rest)
			node.children[firstLetter] = child
			return wasWord
		} else {
			subnode := initNode()
			wasWord = subnode.remove(rest)
			node.children[firstLetter] = subnode
		}

		return wasWord
	}
}

func (node *trieNode) WordsWithPrefix(initPrefix, currPrefix []rune) []string {
	if len(currPrefix) == 0 {
		return node.validWordsHelper(make([]string, 0), initPrefix)
	} else {
		currPrefixFirst := currPrefix[0]
		if child, ok := node.children[currPrefixFirst]; ok {
			return child.WordsWithPrefix(initPrefix, currPrefix[1:])
		} else {
			return []string{}
		}
	}
}

func (node *trieNode) validWordsHelper(words []string, prefix []rune) []string {
	if node.isWord {
		words = append(words, string(prefix))
	}

	for k, v := range node.children {
		words = v.validWordsHelper(words, append(prefix, k))
	}

	return words
}

func (node *trieNode) ContainsWord(currPrefix []rune) bool {
	if len(currPrefix) == 0 {
		return node.isWord
	} else {
		first := currPrefix[0]
		if child, ok := node.children[first]; ok {
			return child.ContainsWord(currPrefix[1:])
		} else {
			return false
		}
	}
}
