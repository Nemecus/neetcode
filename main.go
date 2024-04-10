package main

import (
	"fmt"

	"github.com/Nemecus/neetcode/dynamicarray"
	"github.com/Nemecus/neetcode/singlylinkedlist"
)

func main() {
	//runDynamicArray()
	runSinglyLinkedList()
}

func runDynamicArray() {
	// Test 1
	fmt.Println("Test 1")
	example1 := dynamicarray.SetupArray(1)
	fmt.Println("Current Size: ", example1.GetSize())
	fmt.Println("Current Capacity: ", example1.GetCapacity())

	// Test 2
	fmt.Println("Test 2")
	example2 := dynamicarray.SetupArray(1)
	fmt.Println("Running Pushback")
	example2.Pushback(1)
	fmt.Println("Current Capacity: ", example2.GetCapacity())
	fmt.Println("Running Pushback")
	example2.Pushback(2)
	fmt.Println("Current Capacity: ", example2.GetCapacity())

	// Test 3
	fmt.Println("Test 3")
	example3 := dynamicarray.SetupArray(1)
	fmt.Println("Current Size: ", example3.GetSize())
	fmt.Println("Current Capacity: ", example3.GetCapacity())
	fmt.Println("Running Pushback")
	example3.Pushback(1)
	fmt.Println("Current Size: ", example3.GetSize())
	fmt.Println("Current Capacity: ", example3.GetCapacity())
	fmt.Println("Running Pushback")
	example3.Pushback(2)
	fmt.Println("Current Size: ", example3.GetSize())
	fmt.Println("Current Capacity: ", example3.GetCapacity())
	fmt.Println("Running get: ", example3.Get(1))
	fmt.Println("Running set")
	example3.Set(1, 3)
	fmt.Println("Running get: ", example3.Get(1))
	fmt.Println("Running popback: ", example3.Popback())
	fmt.Println("Current Size: ", example3.GetSize())
	fmt.Println("Current Capacity: ", example3.GetCapacity())
}

func runSinglyLinkedList() {
	fmt.Println("Test 1")
	var example1 singlylinkedlist.LinkedList
	fmt.Println("Inserting Head value 1")
	example1.InsertHead(1)
	fmt.Println("Inserting Tail value 2")
	example1.InsertTail(2)
	fmt.Println("Inserting Head value 0")
	example1.InsertHead(0)
	fmt.Println("Removing 1")
	example1.Remove(1)
	fmt.Println("Full Values List: ", example1.GetValues())

	fmt.Println("Test 2")
	var example2 singlylinkedlist.LinkedList
	fmt.Println("Inserting Head value 1")
	example2.InsertHead(1)
	fmt.Println("Inserting Head value 2")
	example2.InsertHead(2)
	fmt.Println("Getting value 5: ", example2.Get(5))
}
