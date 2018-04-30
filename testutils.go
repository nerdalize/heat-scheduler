package main

import k8sApi "k8s.io/kubernetes/pkg/api"

// this file contains functions that are shared among both test files.

// newNode returns a new k8sApi.Node given a name and joules.
func newNode(name string, joules string) k8sApi.Node {
	jmap := make(map[string]string)
	if joules != "" {
		jmap["nerdalize/temp"] = joules
	}
	return k8sApi.Node{
		ObjectMeta: k8sApi.ObjectMeta{
			Name:        name,
			Annotations: jmap,
		},
	}
}

// newNodeList returns a node list.
func newNodeList(nodes ...k8sApi.Node) k8sApi.NodeList {
	return k8sApi.NodeList{
		Items: nodes,
	}
}
