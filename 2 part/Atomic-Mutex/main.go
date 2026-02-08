// Гонка данных (data race) представляет ситуацию,
// когда два или несколько потоков одновременно обращаются к одному и тому же участку памяти
// и выполняют по крайней мере одну операцию записи

package main

import (
	"fmt"
	"sync"
	//"sync/atomic"
)

// var number int = 0
// var number atomic.Int64
var slice []int

var mtx sync.Mutex

func increase(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 1000; i++ {
		//mtx.Lock()
		slice = append(slice, i)
		//number.Add(1)
		//number++
		//mtx.Unlock()
	}
}

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(10)
	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)

	wg.Wait()
	fmt.Println(len(slice))
}
