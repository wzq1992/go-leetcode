package lessons

import (
	"strconv"
	"strings"
)

type Codec struct {
	seq []string
}

func Constructor() Codec {
	return Codec{make([]string, 0)}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	this.traverse(root)
	return strings.Join(this.seq, " ")
}

func (this *Codec) traverse(root *TreeNode) {
	if root == nil {
		this.seq = append(this.seq, "null")
		return
	}

	this.seq = append(this.seq, strconv.Itoa(root.Val))
	this.traverse(root.Left)
	this.traverse(root.Right)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	this.seq = strings.Split(data, " ")
	return this.buildNode()
}

func (this *Codec) buildNode() *TreeNode {
	nodeVal := this.seq[0]
	this.seq = this.seq[1:]

	if nodeVal == "null" {
		return nil
	}
	nodeValInt, _ := strconv.Atoi(nodeVal)
	root := &TreeNode{
		Val: nodeValInt,
	}

	root.Left = this.buildNode()
	root.Right = this.buildNode()

	return root
}
