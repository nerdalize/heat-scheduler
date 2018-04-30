package main

import (
	"testing"

	k8sApi "k8s.io/kubernetes/pkg/api"
)

// TestSelectNodeSuccess tests selectNode using different inputs.
func TestSelectNodeSuccess(t *testing.T) {
	testCases := map[string]struct {
		list     k8sApi.NodeList
		expected string
	}{
		"sorted": {
			list: newNodeList(
				newNode("node1", "50.5"),
				newNode("node2", "70.5"),
				newNode("node3", "80.5"),
			),
			expected: "node1",
		},
		"reverse sorted": {
			list: newNodeList(
				newNode("node1", "80.5"),
				newNode("node2", "70.5"),
				newNode("node3", "50.5"),
			),
			expected: "node3",
		},
		"mixed": {
			list: newNodeList(
				newNode("node1", "80.5"),
				newNode("node2", "50.5"),
				newNode("node3", "70.5"),
			),
			expected: "node2",
		},
		"illigal joules": {
			list: newNodeList(
				newNode("node1", "55.5"),
				newNode("node2", "65.5"),
				newNode("node3", "illigal"),
			),
			expected: "node1",
		},
		"no joules": {
			list: newNodeList(
				newNode("node1", "55.5"),
				newNode("node2", "65.5"),
				newNode("node3", ""),
			),
			expected: "node1",
		},
	}

	for desc, tc := range testCases {
		nodes, err := selectNode(&tc.list)
		if err != nil {
			t.Errorf("Error when testing case %v: %v", desc, err)
		} else {
			if nodes[0].Name != tc.expected {
				t.Errorf("Test case %v: expected %v but got %v", desc, tc.expected, nodes[0].Name)
			}
		}
	}
}

// TestSelectNodeFail tests the case when selecting a node fails.
func TestSelectNodeFail(t *testing.T) {
	list := newNodeList()
	_, err := selectNode(&list)
	if err == nil {
		t.Errorf("Expected error because list was empty")
	}
}
