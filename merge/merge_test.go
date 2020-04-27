package merge

import (
	"fmt"
	"testing"
)

var funcs = []struct {
	name string
	f    func(...<-chan int) <-chan int
}{
	{"goroutines", MergeIteration},
	{"reflection", MergeReflection},
	{"recursion", MergeRecursion},
}

func TestMerge(t *testing.T) {
	for _, f := range funcs {
		t.Run(f.name, func(t *testing.T) {
			c := f.f(AsChan(1, 2, 3), AsChan(4, 5, 6), AsChan(7, 8, 9))
			seen := make(map[int]bool)
			for v := range c {
				if seen[v] {
					t.Errorf("saw %d at least twice", v)
				}
				seen[v] = true
			}
			for i := 1; i <= 9; i++ {
				if !seen[i] {
					t.Errorf("didn't see %d", i)
				}
			}
		})
	}
}

func BenchmarkMerge(b *testing.B) {
	for _, f := range funcs {
		for n := 1; n <= 1024; n *= 2 {
			chans := make([]<-chan int, n)
			b.Run(fmt.Sprintf("%s/%d", f.name, n), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					b.StopTimer()
					for i := range chans {
						chans[i] = AsChan(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
					}
					b.StartTimer()

					c := f.f(chans...)
					for range c {
					}
				}
			})
		}
	}
}
