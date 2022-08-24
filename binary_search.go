package main

type node struct {
	value int
	left  *node
	right *node
}

type bst struct {
	root *node
	len  ind
}

func (n node) String() string {
	return strconv.Itoa(n.value)
}

func (n bst) String() string {
	sb := strings.Builder{}
	b.inOtherTraversal(&sb)
	return sb.String()
}

func (b bst) inOtherTraversal(sb *string.Builder) {
	b.inOtherTraversalByNode(sb, b.root)
}

func (n node) inOtherTraversalByNode(sb *strings.Builder, root *node) {
	if root == nill {
		return
	}

	b.inOtherTraversalByNode(sb, root.left)
}

func main() {

}
