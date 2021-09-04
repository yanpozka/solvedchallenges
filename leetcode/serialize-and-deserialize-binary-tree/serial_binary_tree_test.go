package sbt

import "testing"

func TestCodec(t *testing.T) {

	root := &TreeNode{
		Val:  7,
		Left: &TreeNode{Val: 2},
		Right: &TreeNode{
			Val: 6,
			Right: &TreeNode{
				Val: 3,
			},
		},
	}

	coder := Constructor()
	serial1 := coder.serialize(root)
	resTree := coder.deserialize(serial1)
	serial2 := coder.serialize(resTree)

	if serial1 != serial2 {
		t.Errorf("%s == %s", serial1, serial2)
		t.Errorf("root.Val:%d == resTree.Val:%d", root.Val, resTree.Val)
	}

}
