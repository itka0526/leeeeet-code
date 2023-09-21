// package tries

// type Trie struct {
// 	children [26]*Trie
// 	isEnd    bool
// 	root     *Trie
// }

// func Constructor() Trie {
// 	return Trie{
// 		children: [26]*Trie{},
// 		isEnd:    false,
// 		root:     &Trie{},
// 	}
// }

// func (t *Trie) Insert(word string) {
// 	curr := t.root

// 	for _, c := range word {
// 		index := c - 'a'
// 		if curr.children[index] == nil {
// 			curr.children[index] = &Trie{}
// 		}
// 		curr = curr.children[index]
// 	}

// 	curr.isEnd = true
// }

// func (t *Trie) Search(word string) bool {
// 	curr := t.root

// 	for _, c := range word {
// 		index := c - 'a'
// 		if curr.children[index] == nil {
// 			return false
// 		}
// 		curr = curr.children[index]
// 	}

// 	return curr.isEnd
// }

// func (t *Trie) StartsWith(prefix string) bool {
// 	curr := t.root

// 	for _, c := range prefix {
// 		index := c - 'a'
// 		if curr.children[index] == nil {
// 			return false
// 		}
// 		curr = curr.children[index]
// 	}
// 	return true
// }

package tries

import "fmt"

type Trie struct {
	root     *Trie
	isEnd    bool
	children map[rune]*Trie
}

func Constructor() Trie {
	return Trie{
		root:     &Trie{},
		isEnd:    false,
		children: map[rune]*Trie{},
	}
}

func (t *Trie) Insert(word string) {
	curr := t.root

	for _, c := range word {
		if curr.children == nil {
			curr.children = make(map[rune]*Trie)
		}
		if _, ok := curr.children[c]; !ok {
			curr.children[c] = &Trie{}
		}
		curr = curr.children[c]
	}
	// forgot this crap!
	curr.isEnd = true
}

func (t *Trie) Search(word string) bool {
	curr := t.root
	for _, c := range word {
		if _, ok := curr.children[c]; !ok {
			return false
		}
		curr = curr.children[c]
	}
	return curr.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	curr := t.root
	for _, c := range prefix {
		if _, ok := curr.children[c]; !ok {
			return false
		}
		curr = curr.children[c]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

func TestImplementTrie() {
	word := "apple"
	obj := Constructor()
	obj.Insert(word)
	param_2 := obj.Search(word)
	param_3 := obj.StartsWith("app")

	fmt.Println(param_2, param_3)

	curr := obj.root

	for curr.children != nil {
		for k, v := range curr.children {
			fmt.Println("Letter: ", string(k), "Children: ", v)
			curr = v
		}
	}
}
