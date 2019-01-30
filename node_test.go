// Copyright © 2019 Weald Technology Trading
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

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNodeString(t *testing.T) {
	n1 := &Node{
		Data: &testData{data: "n1"},
	}
	n2 := &Node{
		Data: &testData{data: "n2"},
	}
	n3 := &Node{
		Left:  n1,
		Right: n2,
	}
	assert.Equal(t, "n1", n1.String(), "incorrect value for data node")
	assert.Equal(t, "n2", n2.String(), "incorrect value for data node")
	assert.Equal(t, "n1/\\n2", n3.String(), "incorrect value for parent node")
}
