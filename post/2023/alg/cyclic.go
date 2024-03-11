package alg

type Node struct {
	name     string
	Children []*Node
	visited  bool
	recStack bool
}

func NewNode(name string) *Node {
	return &Node{name: name}
}

func (n *Node) AddEdge(to *Node) {
	n.Children = append(n.Children, to)
}

func isCyclicUtil(node *Node) bool {
		if !node.visited {
			node.visited = true
			node.recStack = true

			for _, child := range node.Children {
				if !child.visited && isCyclicUtil(child) {
					return true
				} else if child.recStack {
					return true
				}
			}
		}
		node.recStack = false
		return false
}

func isCyclic(nodes []*Node) bool {
	for _, node := range nodes {
		if isCyclicUtil(node) {
			return true
		}
	}
	return false
}
