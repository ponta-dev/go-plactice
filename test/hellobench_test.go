package test

import (
	"fmt"
	"testing"
)

func BenchmarkSampleEchoHello(b *testing.B) {
	b.ResetTimer()
	result := 0
	for i := 0; i < b.N; i++ {
		result += i
	}
	fmt.Printf("b.N count  : %d \n", b.N)
	fmt.Printf("b.N result : %d \n", result)
}
