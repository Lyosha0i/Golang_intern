package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

// wg — глобальный sync.WaitGroup
var wg sync.WaitGroup
//generator — генерация случайного числа в случае недостаточного количества чисел
func generator(ch chan int, arr []int, length int) {
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	var x int
	var min = 1
	var max = 101
	x = min + rand.Intn(max-min)
	if len(arr) < length {
		ch <- x
	}
}
//isUnique — проверка сгенерированного числа на уникальность
func isUnique(ch chan int, arr []int, length int) {
	defer wg.Done()
	for len(arr) < length{
		x := <-ch
		var i int
		for i = 0; i < len(arr); i++ {
			if x == arr[i] {
				break//если число неуникально
			}
		}
		if i == len(arr) {//если уникально
			arr = append(arr, x)
			fmt.Print(fmt.Sprintf("%d ",x))
		} else {
			wg.Add(1)
			go generator(ch, arr, length)
		}
	}
}
func main() {
	amount := flag.Int("amount", 1, "")
	flag.Parse()
	defer wg.Done()
	arr := make([]int, 0, *amount)
	ch := make(chan int)
	for i := 0; i < *amount; i++ {
		wg.Add(1)
		go generator(ch, arr, *amount)
	}
	wg.Add(1)
	go isUnique(ch, arr, *amount)
	wg.Wait()
	os.Exit(0)
}