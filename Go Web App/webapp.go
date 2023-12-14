package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type TodoList struct {
	TodoCount int
	Todos     []string
}

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getStrings(filename string) []string {
	lines := []string{}

	file, err := os.Open(filename)
	if os.IsNotExist(err) {
		return nil
	}
	errorCheck(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	errorCheck(scanner.Err())

	return lines
}

func write(wr http.ResponseWriter, msg string) {
	_, err := wr.Write([]byte(msg))
	errorCheck(err)

}

func englishHandler(wr http.ResponseWriter, req *http.Request) {
	write(wr, "Hello, World!")
}

func spanishHandler(wr http.ResponseWriter, req *http.Request) {
	write(wr, "Hola, Internet!")
}

func interactHandler(wr http.ResponseWriter, req *http.Request) {
	todoValues := getStrings("totos.txt")
	fmt.Printf("%#v\n", todoValues)

	tmpl, err := template.ParseFiles("view.html")
	errorCheck(err)
	todos := TodoList{
		Todos:     todoValues,
		TodoCount: len(todoValues),
	}

	err = tmpl.Execute(wr, todos)
	errorCheck(err)
}

func newHandler(wr http.ResponseWriter, req *http.Request) {

}
