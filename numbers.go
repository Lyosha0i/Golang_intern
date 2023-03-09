package main

import (

	//"flag"

	"bufio"
	"fmt"

	//"io/ioutil"
	//"log"
	//"net/http"
	"math/rand"
	"os"
	"sync"
	"time"
)

// wg — глобальный sync.WaitGroup
var wg sync.WaitGroup

func generator(chA chan int, ch chan int, arr []int) {
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	//arr := make([]int, 0, 100)
	fmt.Println(arr)
	var x int
	//var i int
	var min = 1
	var max = 101
	x = -1
	if len(arr) < 100 {
		x = min + rand.Intn(max-min)
	} // for len(arr) < 100 {
	// x = min + rand.Intn(max-min)
	// //go isUnique(ch, arr)
	ch <- x
	for i := range arr {
		chA <- i
	}
	close(ch)
	close(chA)
	// //wg.Add(1) //увеличение счётчика wg
	// //bufio.NewReader(os.Stdin).ReadBytes('\n')
	// fmt.Println(arr)
	//}
}
func isUnique(chA chan int, ch chan int, arr []int) {
	defer wg.Done()
	var j int
	for i := range chA {
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		for j = 0; j < len(chA) && <-ch != i; j++ {
			bufio.NewReader(os.Stdin).ReadBytes('\n')
		}
	}
	if j == len(chA) {
		arr = append(arr, <-ch)
		//return
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		fmt.Println(arr)
	} else {
		return
	}
}
func main() {
	arr := make([]int, 0, 100)
	intsCh := make(chan int, 100)
	intCh := make(chan int, 1)
	wg.Add(2) //увеличение счётчика wg
	go generator(intsCh, intCh, arr)
	go isUnique(intsCh, intCh, arr)
	//sort.Ints(arr)
	//fmt.Println(arr)
	wg.Wait()
	os.Exit(0)
}

// func main() {
// 	arr := make([]int, 0)
// 	var x int
// 		var i int
// 		//var rSrc Source
// 	for len(arr) < 100 {
// 		var rSrc cryptoSource
//     rnd := rand.New(rSrc)

//     // действительно случайное число от 0 до 999
//     x=rnd.Intn(1000)
// 		//x = rand.New(rSrc)
// 		for i := 0; i < len(arr) && x == arr[i]; i++ {
// 			fmt.Print(fmt.Sprintf("%d ",x))
// 		}
// 		if i == len(arr) {
// 			arr = append(arr, x)
// 		} else {
// 			continue
// 		}
// 	}
// 	fmt.Println(arr)
// 	src := flag.String("src", "/", "") //флаги
// 	to := flag.String("to", "/", "")
// 	flag.Parse()
// 	fmt.Println(*src, *to)
// 	file, err := os.Open(*src)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		resp, err := http.Get("https://" + scanner.Text()) //GET
// 		if err != nil {
// 			fmt.Println(err)
// 			continue
// 		}
// 		defer resp.Body.Close()
// 		body, err := ioutil.ReadAll(resp.Body) //подготовка к записи тела в файл
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		wg.Add(1)                                  //увеличение счётчика wg
// 		go responseWrite(to, scanner.Text(), body) //запись в файл тело ответа на GET
// 		if err := scanner.Err(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}
// 	wg.Wait()
// }
