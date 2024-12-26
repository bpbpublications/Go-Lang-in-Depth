package main

import (
	"fmt"
)

type Node struct {
	property int
	nextNode *Node
}

type SinglyLinkedList struct {
	headNode *Node
}

func (singlyLinkedList *SinglyLinkedList) AddToTheRoot(property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil

	if singlyLinkedList.headNode != nil {
		node.nextNode = singlyLinkedList.headNode
	}

	singlyLinkedList.headNode = node

}

func (singlyLinkedList *SinglyLinkedList) CreateNode(property int) *Node {
	var node *Node
	var nodeWith *Node
	for node = singlyLinkedList.headNode; node != nil; node = node.nextNode {
		if node.property == property {
			nodeWith = node
			break
		}
	}
	return nodeWith
}

func (singlyLinkedList *SinglyLinkedList) AddAfterThisElement(nodeProperty int, property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil

	var nodeWith *Node

	nodeWith = singlyLinkedList.CreateNode(nodeProperty)
	if nodeWith != nil {
		node.nextNode = nodeWith.nextNode
		nodeWith.nextNode = node
	}

}

func (singlyLinkedList *SinglyLinkedList) GetLastNode() *Node {
	var node *Node
	var lastNode *Node
	for node = singlyLinkedList.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			lastNode = node
		}
	}
	return lastNode
}

func (singlyLinkedList *SinglyLinkedList) AddToLastNode(property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil

	var lastNode *Node

	lastNode = singlyLinkedList.GetLastNode()

	if lastNode != nil {
		lastNode.nextNode = node
	}

}

func (singlyLinkedList *SinglyLinkedList) ListIterator() {

	var node *Node
	for node = singlyLinkedList.headNode; node != nil; node = node.nextNode {
		fmt.Println(node.property)

	}
}

func main() {

	var singlyLinkedList SinglyLinkedList

	singlyLinkedList = SinglyLinkedList{}

	singlyLinkedList.AddToTheRoot(1)
	singlyLinkedList.AddToTheRoot(2)
	singlyLinkedList.AddToLastNode(4)
	singlyLinkedList.AddAfterThisElement(1, 6)

	singlyLinkedList.ListIterator()

}
