type Trie struct {
	Next map[rune]*Trie
	End bool
}


/** Initialize your data structure here. */
func Constructor() Trie {
	root:=new(Trie)
	root.Next=make(map[rune]*Trie,26)
	root.End=false
	return *root
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	for _,v:=range word{
		if this.Next[v]==nil {
			node:=new(Trie)
			node.Next=make(map[rune]*Trie,26)
			node.End=false
			this.Next[v]=node
		}
		this=this.Next[v]
	}
	this.End=true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	for _,v:=range word{
		if this.Next[v]==nil {
			return false
		}
		this=this.Next[v]
	}
	return this.End
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	for _,v:=range prefix{
		if this.Next[v] == nil {
			return false
		}
		this=this.Next[v]
	}
	return true
}