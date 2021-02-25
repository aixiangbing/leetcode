package RandPrint

import (
	"fmt"
	"math/rand"
	"sync"
)

type RandNum struct {

}

func (this *RandNum) RandPrintNum() {
	out := make(chan int)

	wg := sync.WaitGroup{}

	wg.Add(2)

	go func() {
		defer wg.Done()

		for i := 0; i < 5 ; i++ {
			out <- rand.Int()
		}
		close(out)
	}()

	go func() {
		defer wg.Done()
		for i := range out {
			fmt.Println(i)
		}
	}()
	wg.Wait()
}