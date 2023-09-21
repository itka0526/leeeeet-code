package tries

import "fmt"

type WordDictionary struct {
	root     *WordDictionary
	isWord   bool
	children map[rune]*WordDictionary
}

func constructor() WordDictionary {
	return WordDictionary{root: &WordDictionary{}, isWord: false, children: make(map[rune]*WordDictionary)}
}

func (wd *WordDictionary) AddWord(word string) {
	curr := wd.root

	for _, c := range word {
		if curr.children == nil {
			curr.children = make(map[rune]*WordDictionary)
		}
		if _, ok := curr.children[c]; !ok {
			curr.children[c] = &WordDictionary{}
		}
		curr = curr.children[c]
	}
	curr.isWord = true
}

func (wd *WordDictionary) Search(word string) bool {
	var dfs func(int, *WordDictionary) bool

	dfs = func(start int, node *WordDictionary) bool {
		curr := node

		for i := start; i < len(word); i++ {
			c := []rune(word)[i]

			if c == '.' {
				for _, child := range curr.children {
					if dfs(i+1, child) {
						return true
					}
				}
				return false
			} else {
				if _, ok := curr.children[c]; !ok {
					return false
				}
				curr = curr.children[c]
			}
		}

		return curr.isWord
	}

	return dfs(0, wd.root)
}

func TestImplementWD() {
	word := "apple"
	obj := constructor()
	obj.AddWord(word)
	param_2 := obj.Search(".p.l.")

	fmt.Println(param_2)

	curr := obj.root

	for curr.children != nil {
		for k, v := range curr.children {
			fmt.Println("Letter: ", string(k), "Children: ", v)
			curr = v
		}
	}
}
