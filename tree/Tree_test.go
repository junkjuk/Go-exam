package tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTree_count(t *testing.T) {
	type testCase struct {
		name           string
		input          Tree
		expectedResult int
	}
	testTable := []testCase{
		{
			name:           "test b-node traverse with count comparing on 5 nodes",
			input:          GetNTree5Nodes(),
			expectedResult: 5,
		},
		{
			name:           "test b-node traverse with count comparing on 4 nodes",
			input:          GetNTree4Nodes(),
			expectedResult: 4,
		},
		{
			name:           "test b-node traverse with count comparing on 3 nodes",
			input:          GetNTree3Nodes(),
			expectedResult: 3,
		},
	}

	for _, test := range testTable {
		count := 0
		test.input.Traverse(func(node INode) {
			count++
		})
		assert.Equal(t, test.expectedResult, count, test.name)
	}
}

func TestTree_sum(t *testing.T) {
	type testCase struct {
		name           string
		input          Tree
		expectedResult int
	}
	testTable := []testCase{
		{
			name:           "test n-node sum on in 1-5 nodes",
			input:          GetNTree5Nodes(),
			expectedResult: 15,
		},
		{
			name:           "test n-node sum on in 1-4 nodes",
			input:          GetNTree4Nodes(),
			expectedResult: 10,
		},
		{
			name:           "test n-node sum on in 1-3 nodes",
			input:          GetNTree3Nodes(),
			expectedResult: 6,
		},
	}

	for _, test := range testTable {
		sum := SumTree(&test.input)
		assert.Equal(t, test.expectedResult, sum.GetData(), test.name)
	}
}

func GetNTree5Nodes() Tree {
	var n1, n2, n3, n4, n5 Node

	tree := Tree{}
	n2 = Node{Data: 2}
	n4 = Node{Data: 4}
	n5 = Node{Data: 5}
	n3 = Node{Data: 3, Children: []*Node{&n4, &n5}}
	n1 = Node{Data: 1, Children: []*Node{&n2, &n3}}
	tree.Root = n1
	return tree
}

func GetNTree4Nodes() Tree {
	var n1, n2, n3, n4 Node

	tree := Tree{}
	n2 = Node{Data: 2}
	n4 = Node{Data: 4}
	n3 = Node{Data: 3}
	n1 = Node{Data: 1, Children: []*Node{&n2, &n3, &n4}}
	tree.Root = n1
	return tree
}

func GetNTree3Nodes() Tree {
	var n1, n2, n3 Node

	tree := Tree{}
	n2 = Node{Data: 2}
	n3 = Node{Data: 3}
	n1 = Node{Data: 1, Children: []*Node{&n2, &n3}}
	tree.Root = n1
	return tree
}
