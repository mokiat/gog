package gog_test

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/mokiat/gog"
)

func ExampleDedupe() {
	source := []int{1, 1, 2, 3, 3, 4}
	target := gog.Dedupe(source)
	fmt.Printf("%#v\n", target)

	// Output:
	// []int{1, 2, 3, 4}
}

func ExampleDerefElements() {
	first, second, third := "first", "second", "third"
	source := []*string{&first, &second, &third}
	target := gog.DerefElements(source)
	fmt.Printf("%#v\n", target)

	// Output:
	// []string{"first", "second", "third"}
}

func ExampleFindFunc() {
	source := []string{"user 01", "user 02", "user 03"}
	target, ok := gog.FindFunc(source, func(item string) bool {
		return strings.Contains(item, "02")
	})
	fmt.Println(ok)
	fmt.Println(target)

	// Output:
	// true
	// user 02
}

func ExampleFindFuncPtr() {
	source := []string{"user 01", "user 02", "user 03"}
	target := gog.FindFuncPtr(source, func(item string) bool {
		return strings.Contains(item, "02")
	})
	fmt.Println(target != nil)
	fmt.Println(*target)
	*target = "user 04"
	fmt.Printf("%#v\n", source)

	// Output:
	// true
	// user 02
	// []string{"user 01", "user 04", "user 03"}
}

func ExampleFlatten() {
	source := [][]int{
		{1, 2, 3},
		{4, 5, 6, 7},
	}
	target := gog.Flatten(source)
	fmt.Printf("%#v\n", target)

	// Output:
	// []int{1, 2, 3, 4, 5, 6, 7}
}

func ExampleMap() {
	source := []int{1, 2, 3}
	target := gog.Map(source, func(item int) string {
		return strconv.Itoa(item * 2)
	})
	fmt.Printf("%#v\n", target)

	// Output:
	// []string{"2", "4", "6"}
}

func ExampleMapping() {
	source := []int{1, 2, 3, 4, 5}
	target := gog.Mapping(source, func(item int) (bool, string) {
		isEven := item%2 == 0
		return isEven, strconv.Itoa(item)
	})
	fmt.Printf("even: %#v\n", target[true])
	fmt.Printf("odd: %#v\n", target[false])

	// Output:
	// even: []string{"2", "4"}
	// odd: []string{"1", "3", "5"}
}

func ExampleMutate() {
	source := []int{1, 2, 3, 4, 5}
	gog.Mutate(source, func(item *int) {
		*item = *item + 10
	})
	fmt.Printf("%#v\n", source)

	// Output:
	// []int{11, 12, 13, 14, 15}
}

func ExamplePartition() {
	source := []int{1, 2, 3, 4, 5}
	target := gog.Partition(source, func(item int) bool {
		isEven := item%2 == 0
		return isEven
	})
	fmt.Printf("even: %#v\n", target[true])
	fmt.Printf("odd: %#v\n", target[false])

	// Output:
	// even: []int{2, 4}
	// odd: []int{1, 3, 5}
}

func ExamplePtrOf() {
	ptr := gog.PtrOf("hello world") // since &"hello world" is not possible
	fmt.Println(ptr != nil)
	fmt.Println(*ptr)

	// Output:
	// true
	// hello world
}

func ExampleReduce() {
	source := []int{2, 30, 100}
	target := gog.Reduce(source, 0, func(accum, item int) int {
		return accum + item*2
	})
	fmt.Println(target)

	// Output:
	// 264
}

func ExampleRefElements() {
	source := []int{5, 7}
	target := gog.RefElements(source)
	*target[0] = 9
	*target[1] = 11
	fmt.Printf("%#v\n", source)

	// Output:
	// []int{9, 11}
}

func ExampleSelect() {
	source := []int{5, 8, 9, 10}
	target := gog.Select(source, func(item int) bool {
		isEven := item%2 == 0
		return isEven
	})
	fmt.Printf("%#v\n", target)

	// Output:
	// []int{8, 10}
}

func ExampleValueOf() {
	firstValue := "first"

	var firstPtr *string = &firstValue
	var secondPtr *string = nil

	first := gog.ValueOf(firstPtr, "default")
	fmt.Println(first)

	second := gog.ValueOf(secondPtr, "default")
	fmt.Println(second)

	// Output:
	// first
	// default
}

func ExampleEntries() {
	source := map[string][]int{
		"even": {2, 4, 6},
		"odd":  {1, 3, 5},
	}
	target := gog.Entries(source)
	slices.SortFunc(target, func(first, second gog.KV[string, []int]) int {
		return cmp.Compare(first.Key, second.Key)
	})
	fmt.Printf("%+v\n", target)

	// Output:
	// [{Key:even Value:[2 4 6]} {Key:odd Value:[1 3 5]}]
}

func ExampleConcat() {
	first := []int{1, 2}
	second := []int{3, 4}
	result := gog.Concat(first, second)
	fmt.Printf("%#v\n", result)

	// Output:
	// []int{1, 2, 3, 4}
}

func ExampleMerge() {
	first := map[int]string{
		1: "odd",
		2: "even",
	}
	second := map[int]string{
		1:     "small",
		10000: "large",
	}
	result := gog.Merge(first, second)

	fmt.Printf("1: %#v\n", result[1])
	fmt.Printf("2: %#v\n", result[2])
	fmt.Printf("10000: %#v\n", result[10000])

	// Output:
	// 1: "small"
	// 2: "even"
	// 10000: "large"
}
