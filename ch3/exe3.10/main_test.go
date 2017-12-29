package main

import "testing"

const (
	s = "98765432123456789987654321"
)

func BenchmarkComma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		comma(s)
	}
}

func BenchmarkComma1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		comma1(s)
	}
}

func BenchmarkComma2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		comma2(s)
	}
}
