package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)
type Todolist struct {
	Task 		string
	Completed	bool
}

var (
	todos = []Todolist{}
	doc = "todo.json"
	Todo = Todolist{}
)

func init()  {
	Todo.Read()
}

func (t Todolist) Read() {
	doc, err := ioutil.ReadFile("todo.json")
	if err != nil {
		log.Fatal(err)
	}

	_ = json.Unmarshal(doc, &todos)
}

func (t *Todolist) Clear() {
	for i:=0; i<len(todos); i++ {
		if todos[i].Completed && i != len(todos)-1{
			todos = append(todos[:i], todos[i+1:]...)
			i--
		} else if todos[i].Completed && i == len(todos)-1 {
			todos = todos[:i]
		}
	}

	t.write()
}
func (t Todolist) Add(title string) {
	list := Todolist{Task: title}
	todos = append(todos, list)
	t.write()
}

func (t Todolist) write() {
	res, _ := json.MarshalIndent(todos, "", "  ")
	_ = ioutil.WriteFile(doc, res, 0666)
}

func (t Todolist) Done(val int) {
	if len(todos) < val {
		log.Println("Index out of range")
	} else {
		todos[val - 1].Completed = true
		t.write()
	}
}
func (t Todolist) Undone(val int) {
	if len(todos) < val {
		log.Println("Index out of range")
	} else {
		todos[val - 1].Completed = false
		t.write()
	}
}
func (t Todolist) List() {
	for i, v:= range todos {
		if v.Completed {
			continue
		}
		fmt.Printf("%d\t %v\n", i+1, v.Task)
	}
}

