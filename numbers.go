package main

import (

	//"flag"

	//"bufio"
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

func generator(ch chan int) {
	defer wg.Done()
	//arr := make([]int, 0, 100)
	//fmt.Println(arr)
	var x int
	//var i int
	var min = 1
	var max = 101
	x = min + rand.Intn(max-min)
	ch <- x
	//close(ch)
	//close(chA)
}
func isUnique(ch chan int, arr []int) []int {
	// defer wg.Done()
	var temp int
	for a := range ch {
		temp = a
		//println(temp)
		var i int
		for i = 0; i < len(arr); i++ {
			if temp != arr[i] {
				break
			}
			//bufio.NewReader(os.Stdin).ReadBytes('\n')
		}
		if i == len(arr) {
			arr = append(arr, temp)
			//bufio.NewReader(os.Stdin).ReadBytes('\n')
			fmt.Println(arr)
		}
	}
	return arr
}
func main() {
	rand.Seed(time.Now().UnixNano())

	arr := make([]int, 0, 100)
	intCh := make(chan int, 100)
	go isUnique(intCh, arr)
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go generator(intCh)
		//time.Sleep(time.Millisecond*5000)
	}
	//sort.Ints(arr)
	//fmt.Println(arr)
	// close(intCh)
	wg.Wait()
	//close(intCh)
	//close(intCh)
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
