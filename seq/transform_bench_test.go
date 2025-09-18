package seq_test

import (
	"testing"

	"github.com/mokiat/gog/seq"
)

type largeStruct struct {
	key  int
	data [1024]byte
}

func BenchmarkBatchSlice(b *testing.B) {
	source := make([]largeStruct, 1024)
	for i := range source {
		source[i].key = i / 32
		source[i].data[0] = byte(i)
		source[i].data[1023] = byte(i)
	}

	b.ResetTimer()

	for b.Loop() {
		iter := seq.BatchSlice(source, func(a, b largeStruct) bool {
			return a.key == b.key
		}, 0)
		count := 0
		for range iter {
			count++
		}
		if count != 1024/32 {
			b.Fatalf("unexpected count: %d", count)
		}
	}
}

func BenchmarkBatchSliceFast(b *testing.B) {
	source := make([]largeStruct, 1024)
	for i := range source {
		source[i].key = i / 32
		source[i].data[0] = byte(i)
		source[i].data[1023] = byte(i)
	}

	b.ResetTimer()

	for b.Loop() {
		iter := seq.BatchSliceFast(source, func(items []largeStruct, i, j int) bool {
			a := &items[i]
			b := &items[j]
			return a.key == b.key
		}, 0)
		count := 0
		for range iter {
			count++
		}
		if count != 1024/32 {
			b.Fatalf("unexpected count: %d", count)
		}
	}
}
