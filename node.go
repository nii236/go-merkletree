// Copyright © 2018, 2019 Weald Technology Trading
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package merkletree

import "fmt"

// Node is a node in the merkle tree.  It keeps track of child relationships, along with the data if this is a leaf.
type Node struct {
	Left  *Node
	Right *Node
	Hash  []byte
	Data  NodeData
}

// IsLeaf returns true if the node is a leaf.  A leaf is defined as a node that has no children.
func (n *Node) IsLeaf() bool {
	//return n.Left == nil && n.Right == nil
	return n.Data != nil
}

// String implements the stringer interface
func (n *Node) String() string {
	if n.Data != nil {
		return fmt.Sprintf("%v", n.Data)
	}
	return fmt.Sprintf("%v/\\%v", n.Left, n.Right)
}
