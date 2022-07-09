package main

import (
	"fmt"
	"log"

	m "main.go"
	"main.go/pkg/handler"
	"main.go/pkg/repository"
	"main.go/pkg/service"
	/* 	"bufio"
	   	"bytes"
	   	"fmt"
	   	"os"
	   	"strings" */)

func main() {
	repos := repository.NewRepository("./data/file.htm")
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	fmt.Print(handlers)

	srv := new(m.Server)
	if err := srv.Run("8008", handlers.InitRoutes()); err != nil {
		log.Fatalf("error ocurated while running http server %s", err.Error())
	}

}

/*
func main() {
	fmt.Print("\n\n\nStart Psrsing... \n\n\n")

	file, err := os.Open("./data/file.htm")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	wr := bytes.Buffer{}
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		wr.WriteString(sc.Text())
	}

	res := FindWord("атайын кӕнын", wr.String())
	for i := 0; i < len(res); i++ {
		fmt.Print("\n" + res[i] + "\n\n\n")
	}

}

func SplitTextToArray(text string) []string {
	bodyText := strings.Split(text, "body")[1]
	splitText := strings.Split(bodyText, "<p>")

	slice := []string{"start"}
	index := 0

	for i := 0; i < len(splitText); i++ {
		if strings.Contains(splitText[i], "style=\"font-weight:bold;\"") {
			index++
			slice = append(slice, splitText[i])
			continue
		}
		slice[index] += splitText[i]
	}
	return slice
}

func GetSplitTranslate(str string) []string {
	res := strings.Split(str, "</p>")

	return res
}

func FindWord(word string, str string) []string {
	slice := SplitTextToArray(str)

	res := make([]string, 0)
	for i := 0; i < len(slice); i++ {
		split := GetSplitTranslate(slice[i])
		if strings.Contains(split[0], word) {
			res = append(res, strings.ReplaceAll(slice[i], "</p>", ""))
		}
	}

	return res
}
*/
