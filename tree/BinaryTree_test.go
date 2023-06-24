package tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinaryNode_count(t *testing.T) {
	type testCase struct {
		name           string
		input          BinaryTree
		expectedResult int
	}
	testTable := []testCase{
		{
			name:           "test b-node traverse with count comparing on 5 nodes",
			input:          GetBTree5Nodes(),
			expectedResult: 5,
		},
		{
			name:           "test b-node traverse with count comparing on 4 nodes",
			input:          GetBTree4Nodes(),
			expectedResult: 4,
		},
		{
			name:           "test b-node traverse with count comparing on 3 nodes",
			input:          GetBTree3Nodes(),
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

func TestBinaryNode_Sum(t *testing.T) {
	type testCase struct {
		name           string
		input          BinaryTree
		expectedResult int
	}
	testTable := []testCase{
		{
			name:           "test b-node sum on in 1-5 nodes",
			input:          GetBTree5Nodes(),
			expectedResult: 15,
		},
		{
			name:           "test b-node sum on in 1-4 nodes",
			input:          GetBTree4Nodes(),
			expectedResult: 10,
		},
		{
			name:           "test b-node sum on in 1-3 nodes",
			input:          GetBTree3Nodes(),
			expectedResult: 6,
		},
	}

	for _, test := range testTable {
		sum := SumTree(&test.input)
		assert.Equal(t, test.expectedResult, sum.GetData(), test.name)
	}
}

func GetBTree5Nodes() BinaryTree {
	var n1, n2, n3, n4, n5 BinaryNode

	tree := BinaryTree{}
	n2 = BinaryNode{Data: 2}
	n4 = BinaryNode{Data: 4}
	n5 = BinaryNode{Data: 5}
	n3 = BinaryNode{Data: 3, Left: &n4, Right: &n5}
	n1 = BinaryNode{Data: 1, Left: &n2, Right: &n3}
	tree.Root = &n1
	return tree
}

func GetBTree4Nodes() BinaryTree {
	var n1, n2, n3, n4 BinaryNode

	tree := BinaryTree{}
	n2 = BinaryNode{Data: 2}
	n4 = BinaryNode{Data: 4}
	n3 = BinaryNode{Data: 3, Left: &n4}
	n1 = BinaryNode{Data: 1, Left: &n2, Right: &n3}
	tree.Root = &n1
	return tree
}

func GetBTree3Nodes() BinaryTree {
	var n1, n2, n3 BinaryNode
	tree := BinaryTree{}
	n2 = BinaryNode{Data: 2}
	n3 = BinaryNode{Data: 3}
	n1 = BinaryNode{Data: 1, Left: &n2, Right: &n3}
	tree.Root = &n1
	return tree
}
