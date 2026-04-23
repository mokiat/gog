package ds_test

import (
	"bytes"
	"fmt"

	"github.com/mokiat/gog/ds"
)

func ExampleHeap() {
	heap := ds.EmptyHeap(func(a, b int) bool {
		return a < b
	})
	heap.Push(100)
	heap.Push(20)
	heap.Push(50)
	heap.Push(13)
	heap.Push(300)

	fmt.Println(heap.Pop())
	fmt.Println(heap.Pop())
	fmt.Println(heap.Pop())
	fmt.Println(heap.Pop())
	fmt.Println(heap.Pop())

	// Output:
	// 13
	// 20
	// 50
	// 100
	// 300
}

func ExampleList() {
	list := ds.EmptyList[string]()
	list.Add("first")
	list.Add("second")
	list.Add("third")
	list.Add("fourth")
	list.Remove("second")
	list.Remove("third")
	fmt.Printf("%#v\n", list.Items())

	// Output:
	// []string{"first", "fourth"}
}

func ExamplePool() {
	pool := ds.EmptyPool[bytes.Buffer]()

	buffer := pool.Fetch()
	buffer.Reset()
	buffer.WriteString("this buffer is mine")
	pool.Restore(buffer)

	newBuffer := pool.Fetch()
	// newBuffer.Reset() // commented out to show that the instance is reused
	fmt.Println(newBuffer.String())

	// Output:
	// this buffer is mine
}

func ExampleSet() {
	set := ds.EmptySet[int]()
	set.Add(2)
	set.Add(3)
	set.Add(5)
	set.Add(7)
	fmt.Println(set.Contains(2))
	fmt.Println(set.Contains(10))
	fmt.Println(set.Contains(5))
	fmt.Println(set.Contains(8))

	// Output:
	// true
	// false
	// true
	// false
}

func ExampleStack() {
	stack := ds.EmptyStack[string]()
	stack.Push("first")
	stack.Push("second")
	stack.Push("third")
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())

	// Output:
	// third
	// second
	// first
}
