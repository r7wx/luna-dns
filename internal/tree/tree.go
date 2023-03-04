package tree

// Tree - DNS tree struct
type Tree struct {
	tlds map[string]*node
}
type node struct {
	childrens map[string]*node
	ip        string
}

// NewTree - Create a new DNS tree
func NewTree() *Tree {
	return &Tree{
		tlds: map[string]*node{},
	}
}
