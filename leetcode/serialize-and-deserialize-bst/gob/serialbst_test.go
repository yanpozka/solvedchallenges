package serialbst

import "testing"

func TestSerializeDeserialize(t *testing.T) {
	coder := Constructor()
	tree := &TreeNode{
		Val:  5,
		Left: &TreeNode{Val: 2},
		Right: &TreeNode{
			Val: 7,
			Right: &TreeNode{
				Val: 8,
			},
		},
	}

	serialStr1 := coder.serialize(tree)
	t.Log(serialStr1)

	decTree := coder.deserialize(serialStr1)
	serialStr2 := coder.serialize(decTree)
	t.Log(serialStr2)
	t.Log(decTree.Val, decTree.Left.Val, decTree.Left.Left, decTree.Left.Right, decTree.Right.Val, decTree.Right.Left, decTree.Right.Right.Val)

	if serialStr1 != serialStr2 {
		t.Fatalf("%q != %q", serialStr1, serialStr2)
	}
}

func TestSerializeDeserializeEmpty(t *testing.T) {
	coder := Constructor()
	var tree *TreeNode

	serialStr1 := coder.serialize(tree)
	t.Log(serialStr1)
	decTree := coder.deserialize(serialStr1)
	serialStr2 := coder.serialize(decTree)
	t.Log(serialStr2)

	if serialStr1 != serialStr2 {
		t.Fatalf("%q != %q", serialStr1, serialStr2)
	}
}
