package sbtg

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
	buff *bytes.Buffer
}

func Constructor() Codec {
	return Codec{buff: new(bytes.Buffer)}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	this.buff.Reset()
	enc := gob.NewEncoder(this.buff)
	enc.Encode(*root)
	return hex.EncodeToString(this.buff.Bytes())
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	b, err := hex.DecodeString(data)
	if err != nil {
		panic(err)
	}

	t := new(TreeNode)
	dec := gob.NewDecoder(bytes.NewBuffer(b))
	dec.Decode(t)
	return t
}
