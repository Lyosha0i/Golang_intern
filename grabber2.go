package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// responseWrite — запись в файл ответа на GET запрос
func responseWrite(to *string, text string, body []byte) {
	{
		file, err := os.Create(*to + text + ".htm") // создаём файл
		if err != nil {                             // если возникла ошибка
			fmt.Println("Unable to create file:", err)
			os.Exit(1) // выходим из программы
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
			log.Fatalln(err)
		}
		defer file.Close()
		body, err := ioutil.ReadAll(resp.Body) //подготовка к записи тела в файл
		if err != nil {
			log.Fatalln(err)
		}
		go responseWrite(to, scanner.Text(), body) //запись в файл тело ответа на GET
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
	file.Close()
}
