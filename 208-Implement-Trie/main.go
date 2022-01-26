// time complexity for inserting will take O(m) where the length of the word is m
// space complexity for inserting will take O(m) where size of tree is also dependent on the length of the word
// for searching of word and prfix:
// the time complexity is O(m)
// the space complexity is O(1) since we don't take extra space at all after the trie(try) has been created

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {

	if len(word) == 0 {
		return
	}

	cur := this // this gives us a reference to this and allows us to move through the trie without affecting root node
	for i := 0; i < len(word); i++ {

		pos := helperFunction(word[i])
		if cur.children[pos] == nil { // this means that we don't have the word stored
			cur.children[pos] = &Trie{}
		}
		cur = cur.children[pos] // this shifts cur to the next one, allow us to traverse the Trie Node
	}
	cur.isEnd = true // this ensures that we mark the end of the word with true

}

func (this *Trie) Search(word string) bool {
	if len(word) == 0 {
		return false
	}

	cur := this // this gives us a reference to this and allows us to move through the trie without affecting root node
	for i := 0; i < len(word); i++ {
		pos := helperFunction(word[i])
		if cur.children[pos] == nil { // this means the alphabet doesn't exist
			return false
		}
		cur = cur.children[pos] // this shifts cur to the next one, allow us to traverse the Trie Node
	}

	return cur.isEnd // this will return true if it's the end of the word else false

}

func (this *Trie) StartsWith(prefix string) bool {
	if len(prefix) == 0 {
		return false
	}

	cur := this // this gives us a reference to this and allows us to move through the trie without affecting root node

	for i := 0; i < len(prefix); i++ {
		pos := helperFunction(prefix[i])
		if cur.children[pos] == nil { // this means that the alphabet doesn't exist at all
			return false
		}
		cur = cur.children[pos] // this shifts cur to the next one, allow us to traverse the Trie Node
	}
	return true // this returns true as long as we reach the end of the word
}

func helperFunction(c uint8) uint8 {
	return c - 'a' // this will return the index of the word with reference to the alphabet count. i.e. a = 0 while b = 1
}
