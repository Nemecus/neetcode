package main

import (
	"fmt"

	"github.com/Nemecus/neetcode/clear"
	"github.com/Nemecus/neetcode/dynamicarray"
	"github.com/Nemecus/neetcode/factorymethod"
	"github.com/Nemecus/neetcode/insertionsort"
	"github.com/Nemecus/neetcode/singleton"
	"github.com/Nemecus/neetcode/singlylinkedlist"
	"github.com/charmbracelet/lipgloss"
)

var header = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#04B575"))

var title = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#0000FF"))

func main() {
	clear.CallClear()
	fmt.Println(title.Render("Starting Dynamic Array!"))
	runDynamicArray()
	fmt.Println(title.Render("Starting Singly Linked List!"))
	runSinglyLinkedList()
	fmt.Println(title.Render("Starting Insertion Sort!"))
	runInsertionSort()
	fmt.Println(title.Render("Starting Factory Method!"))
	runFactoryMethod()
	fmt.Println(title.Render("Starting Singleton Pattern!"))
	runSingletonPattern()
}

func runDynamicArray() {
	// Test 1
	fmt.Println(header.Render("Test 1"))
	example1 := dynamicarray.SetupArray(1)
	fmt.Println("Current Size: ", example1.GetSize())
	fmt.Println("Current Capacity: ", example1.GetCapacity())

	// Test 2
	fmt.Println(header.Render("Test 2"))
	example2 := dynamicarray.SetupArray(1)
	fmt.Println("Running Pushback")
	example2.Pushback(1)
	fmt.Println("Current Capacity: ", example2.GetCapacity())
	fmt.Println("Running Pushback")
	example2.Pushback(2)
	fmt.Println("Current Capacity: ", example2.GetCapacity())

	// Test 3
	fmt.Println(header.Render("Test 3"))
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
	fmt.Println(header.Render("Test 1"))
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

	fmt.Println(header.Render("Test 2"))
	var example2 singlylinkedlist.LinkedList
	fmt.Println("Inserting Head value 1")
	example2.InsertHead(1)
	fmt.Println("Inserting Head value 2")
	example2.InsertHead(2)
	fmt.Println("Getting value 5: ", example2.Get(5))
}

func runInsertionSort() {
	pairs := map[int]string{5: "apple", 2: "banana", 9: "cherry"}

	fmt.Println(header.Render("Test 1"))
	keys := make([]int, 0, len(pairs))

	for k := range pairs {
		keys = append(keys, k)
	}
	fmt.Println("Unsorted Pairs: ")
	for _, k := range keys {
		fmt.Println(k, pairs[k])
	}

	sortedPairs := insertionsort.InsertionSort(pairs)
	newKeys := make([]int, 0, len(sortedPairs))
	for nk := range sortedPairs {
		newKeys = append(newKeys, nk)
	}
	fmt.Println("Sorted Pairs: ")
	for _, nk := range newKeys {
		fmt.Println(nk, sortedPairs[nk])
	}

	/*
		 	maps don't allow for duplicate keys

			pairs2 := map[int]string{3: "cat", 3: "bird", 2: "dog"}

			fmt.Println("Test 1")
			keys2 := make([]int, 0, len(pairs2))

			for k := range pairs {
				keys2 = append(keys2, k)
			}
			fmt.Println("Unsorted Pairs: ")
			for _, k := range keys2 {
				fmt.Println(k, pairs[k])
			}

			sortedPairs2 := insertionsort.InsertionSort(pairs2)
			newKeys2 := make([]int, 0, len(sortedPairs2))
			for nk := range sortedPairs2 {
				newKeys2 = append(newKeys2, nk)
			}
			fmt.Println("Sorted Pairs: ")
			for _, nk := range newKeys2 {
				fmt.Println(nk, sortedPairs2[nk])
			}
	*/
}

func runFactoryMethod() {
	fmt.Println(header.Render("Testing Vehicles"))
	carFactory := &factorymethod.CarFactory{}
	car := carFactory.CreateVehicle()
	fmt.Println("Vehicle is a ", car.GetType())

	truckFactory := &factorymethod.TruckFactory{}
	truck := truckFactory.CreateVehicle()
	fmt.Println("Vehicle is a ", truck.GetType())

	bikeFactory := &factorymethod.BikeFactory{}
	bike := bikeFactory.CreateVehicle()
	fmt.Println("Vehicle is a ", bike.GetType())
}

func runSingletonPattern() {
	fmt.Println(header.Render("Test Singleton"))
	fmt.Println("Creating Singleton Instance")

	var singleton singleton.ValueServiceSingleton

	s1 := singleton.GetService()
	fmt.Println("Getting Value: ", s1.GetValue())
	fmt.Println("Setting Value to 'a value string'")
	s1.SetValue("a value string")
	fmt.Println("Getting Value: ", s1.GetValue())

	fmt.Println("Creating Second Singleton Instance")
	s2 := singleton.GetService()
	fmt.Println("Getting Value from second instance: ", s2.GetValue())
}
