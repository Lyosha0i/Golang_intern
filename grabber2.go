package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
)

// wg — глобальный sync.WaitGroup
var wg sync.WaitGroup

// responseWrite — запись в файл ответа на GET запрос
func responseWrite(to *string, text string, body []byte) {
	{
		defer wg.Done()
		file, err := os.Create(fmt.Sprintf("%s%s.htm", *to,text)) // создаём файл
		if err != nil {                             // если возникла ошибка
			fmt.Println("Unable to create file:", err)
		}
		defer file.Close()
		file.Write([]byte(body))
		fmt.Println(text)
	}
}

func main() {
	src := flag.String("src", "/", "") //флаги
	to := flag.String("to", "/", "")
	flag.Parse()
	fmt.Println(*src, *to)
	file, err := os.Open(*src)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		resp, err := http.Get("https://" + scanner.Text()) //GET
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body) //подготовка к записи тела в файл
		if err != nil {
			fmt.Println(err)
		}
		wg.Add(1)//увеличение счётчика wg
		go responseWrite(to, scanner.Text(), body) //запись в файл тело ответа на GET
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
		}
	}
	wg.Wait()
}
