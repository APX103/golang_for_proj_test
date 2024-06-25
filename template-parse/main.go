package main

import (
	"bytes"
	"fmt"
	"log"
	"text/template"
)

type Task struct {
	ID string
}

func main() {
	fmt.Println("6")
	tmpl, err := template.New("test").Parse("66 {{.ID}} 66")
	if err != nil {
		fmt.Println("憋屈")
	}
	task := Task{
		ID: "你好",
	}

	res := new(bytes.Buffer)
	err = tmpl.Execute(res, task)
	if err != nil {
		log.Fatalf("template execution: %s", err)
	}
	fmt.Println(res)
}
