package pkg

type Node struct {
	Children map[rune]*Node
	IsWord   bool
}

func NewNode() *Node {
	return &Node{
		Children: make(map[rune]*Node),
	}
}

type Trie struct {
	Root *Node
}

func NewTrie() *Trie {
	return &Trie{
		Root: NewNode(),
	}
}

func (t *Trie) Insert(word string) {

	// node := t.Root

	// for _, c := range word {
	// 	if _, ok := node.Children[c]; !ok {
	// 		node.Children[c] = NewNode()
	// 	}
	//
	// 	node = node.Children[c]
	// }
	//
	// node.IsWord = true
}

func (t *Trie) Search(word string) bool {
	node := t.Root

	for _, c := range word {
		if _, ok := node.Children[c]; !ok {
			return false
		}

		node = node.Children[c]
	}

	return node.IsWord
}

func (t *Trie) FindWordsWithPrefix(prefix string) []string {
	node := t.Root
	for _, c := range prefix {
		if child, ok := node.Children[c]; !ok {
			return []string{} // Prefix is not found in the trie.
		} else {
			node = child
		}
	}

	words := []string{}
	t.dfsWordsWithPrefix(node, prefix, &words)
	return words
}

func (t *Trie) dfsWordsWithPrefix(node *Node, currentPrefix string, words *[]string) {
	if node == nil {
		return
	}

	if node.IsWord {
		*words = append(*words, currentPrefix)
	}

	for char, childNode := range node.Children {
		t.dfsWordsWithPrefix(childNode, currentPrefix+string(char), words)
	}
}

func (t *Trie) Remove(word string) {
	t.remove(t.Root, word, 0)
}

func (t *Trie) remove(node *Node, word string, index int) bool {
	if node == nil {
		return false
	}

	if index == len(word) {
		if !node.IsWord {
			return false
		}

		node.IsWord = false
		return len(node.Children) == 0
	}

	char := rune(word[index])
	if childNode, exists := node.Children[char]; exists {
		shouldRemoveChild := t.remove(childNode, word, index+1)
		if shouldRemoveChild {
			delete(node.Children, char)
			return !node.IsWord && len(node.Children) == 0
		}
	}

	return false
}
