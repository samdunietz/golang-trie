// trietest
package trie

import (
	// "fmt"
	goset "github.com/deckarep/golang-set"
	"github.com/stretchr/testify/assert"
	"testing"
	"trie"
)

var L0 []string = []string{"at"}
var L1 []string = []string{"at", "the", "then"}
var L2 []string = []string{"at", "the", "then", "them", "the", "thematic"}

var T0 trie.Trie = trie.New(L0)
var T1 trie.Trie = trie.New(L1)
var T2 trie.Trie = trie.New(L2)

func setWithSlice(words []string) goset.Set {
	var s = goset.NewSet()
	for _, v := range words {
		s.Add(v)
	}

	return s
}

func TestWords(t *testing.T) {
	assert.True(t, setWithSlice(T0.Words()).Equal(setWithSlice(L0)))
	assert.True(t, setWithSlice(T1.Words()).Equal(setWithSlice(L1)))
	assert.True(t, setWithSlice(T2.Words()).Equal(setWithSlice(L2)))
}

func TestUnicode(t *testing.T) {
	unicodeWords := []string{"אב", "אבא"}
	unicodeTrie := trie.New(unicodeWords)

	assert.Equal(t, setWithSlice(unicodeTrie.WordsWithPrefix("א")), setWithSlice(unicodeWords))
	assert.True(t, unicodeTrie.ContainsWord("אבא"))
	assert.False(t, unicodeTrie.ContainsWord("asdf"))
}

func TestWordsWithPrefix(t *testing.T) {
	assert.True(t, setWithSlice(T0.WordsWithPrefix("")).Equal(setWithSlice(L0)))
	assert.True(t, setWithSlice(T0.WordsWithPrefix("at")).Equal(setWithSlice([]string{"at"})))
	assert.True(t, setWithSlice(T0.WordsWithPrefix("t")).Equal(goset.NewSet()))

	assert.True(t, setWithSlice(T2.WordsWithPrefix("them")).Equal(setWithSlice([]string{"them", "thematic"})))
}

func TestAdd(t *testing.T) {
	temp := trie.New(L0)
	assert.False(t, temp.Add("at"))
	assert.True(t, temp.Add("a"))
	assert.True(t, temp.Add("ate"))
	assert.True(t, temp.Add("the"))
	assert.Equal(t, setWithSlice(temp.Words()), setWithSlice([]string{"a", "ate", "at", "the"}))
}

func TestAddAll(t *testing.T) {
	temp := trie.New(L1)
	assert.True(t, temp.AddAll(L2))
	assert.False(t, temp.AddAll(L1))
	assert.Equal(t, setWithSlice(temp.Words()), setWithSlice(T2.Words()))
}

func TestRemove(t *testing.T) {
	temp := trie.New(L1)
	assert.True(t, temp.Remove("the"))
	assert.Equal(t, setWithSlice(temp.Words()), setWithSlice([]string{"at", "then"}))
	assert.False(t, temp.Remove("a"))
	assert.False(t, temp.Remove("asdfa"))
	assert.Equal(t, setWithSlice(temp.Words()), setWithSlice([]string{"at", "then"}))
	assert.True(t, temp.Remove("then"))
	assert.Equal(t, setWithSlice(temp.Words()), setWithSlice([]string{"at"}))
	assert.True(t, temp.Remove("at"))
	assert.True(t, temp.IsEmpty())
}

func TestRemoveAndAdd(t *testing.T) {
	temp := trie.New(L1)
	assert.True(t, temp.RemoveAll([]string{"at", "then", "the"}))
	assert.True(t, temp.IsEmpty())
	assert.True(t, temp.Add("at"))
	assert.True(t, temp.Add("the"))
	assert.True(t, temp.Add("then"))
}

func TestIsEmpty(t *testing.T) {
	t1 := trie.New([]string{})
	assert.True(t, t1.IsEmpty())
	t2 := trie.New(nil)
	assert.True(t, t2.IsEmpty())
	t3 := trie.New([]string{""})
	assert.False(t, t3.IsEmpty())
	assert.False(t, T0.IsEmpty())
}

func ContainsWord(t *testing.T) {
	temp := trie.New(L2)
	assert.False(t, temp.ContainsWord("a"))
	assert.False(t, temp.ContainsWord("th"))
	assert.False(t, temp.ContainsWord("ate"))

	assert.True(t, temp.ContainsWord("at"))
	assert.True(t, temp.ContainsWord("thematic"))

	temp.Remove("them")
	assert.False(t, temp.ContainsWord("them"))
}





