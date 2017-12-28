package main

import "testing"

var number uint64 = 0x9876543210fedbca

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(number)
	}
}
