package main

import (
	"math/rand"
	"testing"
)

func Benchmark_FindPathIgnoreBlock(b *testing.B) {
	//rand.Rand.Seed(time.Now().Unix())
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			//FindPathIgnoreBlock([]int32{0, 0}, [][]int32{{6, 5}, {4, 3}})
			FindPathIgnoreBlock([]int32{rand.Int31n(12), rand.Int31n(12)},
				[][]int32{{rand.Int31n(12), rand.Int31n(12)}})
		}
	})
}
