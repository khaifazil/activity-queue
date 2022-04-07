package main

import (
	"errors"
	"fmt"
)

type node struct {
	priority int
	item     string
	next     *node
}

type queue struct {
	front *node
	back  *node
	size  int
}

func (p *queue) enqueue(name string, priorityNumber int) error {
	newNode := &node{
		priority: priorityNumber,
		item: name,
		next: nil,
	}

	if p.front == nil { //check if empty queue
		p.front = newNode
		p.back = newNode
	} else { // check if check is not empty

		if p.front.priority > priorityNumber { // check to see if newnode priorityNumber is lower.
			newNode.next = p.front // new node becomes front
			p.front = newNode

		} else {

			currentNode := p.front //init traversal node

			for currentNode.next != nil && currentNode.next.priority <= priorityNumber { //iterate and check for priority. will stop right before a higher priority number or equals to
				currentNode = currentNode.next
			}

			if currentNode.next == nil { //currentNode == p.back //if newNode has biggest priority number
				p.back = newNode
			}
			newNode.next = currentNode.next // connections to link
			currentNode.next = newNode
		}
	}
	p.size++ //add length
	return nil
}

func (p *queue) dequeue() (string, error) {
	var item string

	if p.front == nil { //check if empty queue
		return "", errors.New("empty queue")
	}

	item = p.front.item
	if p.size == 1 {
		p.front = nil
		p.back = nil
	} else {
		p.front = p.front.next
	}
	p.size--
	return item, nil
}

func (p *queue) printAllNodes() (int, error) {
	var count int = 0
	currentNode := p.front
	if currentNode == nil {
		return count, errors.New("queue is empty")
	}

	fmt.Printf("%+v\n", currentNode.item)
	for currentNode.next != nil {
		currentNode = currentNode.next
		fmt.Printf("%+v\n", currentNode.item)
		count++
	}
	return count, nil
}

func (p *queue) getSize() int {
	return p.size
}

func main() {
	myQueue := &queue{nil, nil, 0}

	fmt.Println("Enqueing...")
	myQueue.enqueue("jane", 3)
	myQueue.enqueue("mary", 15)
	myQueue.enqueue("peter", 8)
	myQueue.enqueue("fred", 2)

	// fmt.Println(myQueue.getSize())

	fmt.Println("Showing all nodes in queue..")
	myQueue.printAllNodes()

	myQueue.enqueue("johnnathan", 1)
	// myQueue.dequeue()

	fmt.Println("Showing all nodes in queue..")
	myQueue.printAllNodes()
}
