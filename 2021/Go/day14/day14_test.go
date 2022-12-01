package main

import (
	"testing"
)

func BenchmarkComplete(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main()
	}
}
