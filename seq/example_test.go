package seq_test

import (
	"fmt"

	"github.com/mokiat/gog/seq"
)

func ExampleTimes() {
	for v := range seq.Times(3) {
		fmt.Println(v)
	}

	// Output:
	// 0
	// 1
	// 2
}

func ExampleRange() {
	for v := range seq.Range(1, 3) {
		fmt.Println(v)
	}

	// Output:
	// 1
	// 2
	// 3
}

func ExampleMap() {
	source := seq.Times(4)
	target := seq.Map(source, func(v int) string {
		return fmt.Sprintf("item %d", v)
	})
	for v := range target {
		fmt.Println(v)
	}

	// Output:
	// item 0
	// item 1
	// item 2
	// item 3
}

func ExampleCollectCap() {
	source := seq.Times(4)
	target := seq.CollectCap(source, 12)
	for _, v := range target {
		fmt.Println(v)
	}
	fmt.Println()
	fmt.Println(cap(target))

	// Output:
	// 0
	// 1
	// 2
	// 3
	//
	// 12
}

func ExampleBatchSlice() {
	source := []string{
		"1Hello", "1World",
		"2This", "2Is", "2Longer",
		"3Yes",
	}
	eqFn := func(a, b string) bool {
		return a[0] == b[0]
	}

	for batch := range seq.BatchSlice(source, eqFn, 0) {
		fmt.Printf("%#v\n", batch)
	}

	// Output:
	// []string{"1Hello", "1World"}
	// []string{"2This", "2Is", "2Longer"}
	// []string{"3Yes"}
}
