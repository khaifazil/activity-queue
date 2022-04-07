package main

import (
	"errors"
	"fmt"
)

type node struct {
	item string
	next *node
}

type queue struct {
	front *node
	back  *node
	size  int
}

func (p *queue) enqueue(name string) error {
	newNode := &node{name, nil}

	if p.front == nil { //check if empty queue
		p.front = newNode
	} else { // check if check is not empty
		p.back.next = newNode
	}
	p.back = newNode
	p.size++ //add length
	return nil
}

func (p *queue) dequeue() (string, error)  {
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
	myQueue.enqueue("jane")
	myQueue.enqueue("mary")
	myQueue.enqueue("peter")
	myQueue.enqueue("fred")

	fmt.Println(myQueue.getSize())

	fmt.Println("Showing all nodes in queue..")
	myQueue.printAllNodes()

	myQueue.dequeue()

	fmt.Println("Showing all nodes in queue..")
	myQueue.printAllNodes()
}
