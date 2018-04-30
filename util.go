package main

import (
	"fmt"
	"math"
	"strconv"

	k8sApi "k8s.io/kubernetes/pkg/api"
)

// logNodes prints a line for every node.
func logNodes(nodes *k8sApi.NodeList) {
	for _, n := range nodes.Items {
		fmt.Printf("Received node %v with joules %v\n", n.Name, n.Labels["joules"])
	}
}

// selectNode returns the one node with the lowest joules label value out
// of a list of nodes.
func selectNode(nodes *k8sApi.NodeList) ([]k8sApi.Node, error) {
	if len(nodes.Items) == 0 {
		return nil, fmt.Errorf("No nodes were provided")
	}

	// find min joules value
	min := math.MaxFloat64
	for _, node := range nodes.Items {
		min = math.Min(min, jouleFromLabels(&node))
	}

	// find node belonging to min joules value
	for _, node := range nodes.Items {
		if min == jouleFromLabels(&node) {
			return []k8sApi.Node{node}, nil
		}
	}

	return nil, fmt.Errorf("No suitable nodes found.")
}

// jouleFromLabels parses the joules from a node's label or returns
// the max float value if the label doesn't exist.
func jouleFromLabels(node *k8sApi.Node) float64 {
	jouleString, exists := node.Annotations["nerdalize/temp"]
	if exists {
		joule, err := strconv.ParseFloat(jouleString, 32)
		if err == nil {
			return joule
		}
	}
	return math.MaxFloat64
}
