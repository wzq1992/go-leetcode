package lessons

type Trie struct {
	root *Node
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		root: NewNode(),
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	this.solve(word, true, false)
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	return this.solve(word, false, false)
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	return this.solve(prefix, false, true)
}

func (this *Trie) solve(word string, insertIfNotExist, searchPrefix bool) bool {
	curr := this.root

	for _, ch := range word {
		if _, ok := curr.child[ch]; !ok {
			if insertIfNotExist {
				curr.child[ch] = NewNode()
			} else {
				return false
			}
		}

		curr = curr.child[ch]
	}

	if searchPrefix {
		return true
	}

	if insertIfNotExist {
		curr.count++
	}

	return curr.count > 0
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

type Node struct {
	count int
	child map[rune]*Node
}

func NewNode() *Node {
	return &Node{
		child: make(map[rune]*Node, 0),
	}
}
