package merge

import "reflect"

func MergeReflection(cs ...<-chan int) <-chan int {

	out := make(chan int)

	go func() {
		defer close(out)
		var cases []reflect.SelectCase
		for _, c := range cs {
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(c),
			})
		}

		for len(cases) > 0 {
			i, v, ok := reflect.Select(cases)
			if !ok {
				cases = append(cases[:i], cases[i+1:]...)
				continue
			}
			out <- v.Interface().(int)
		}
	}()

	return out
}
