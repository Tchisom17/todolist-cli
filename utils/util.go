package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Todolist struct {
	Task 		string	`json:"task"`
	Completed	bool	`json:"completed"`
}

const doc = "todo.json"
var (
	todos = []Todolist{}
	Todo = Todolist{}
)

func init()  {
	Todo.Read()
}

func (t Todolist) Read() {
	file, err := ioutil.ReadFile(doc)
	if err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(file, &todos); err != nil {
		log.Fatal(err)
	}
}

func (t Todolist) Clear() {
	for i:=0; i<len(todos); i++ {
		if todos[i].Completed && i != len(todos)-1{
			todos = append(todos[:i], todos[i+1:]...)
			i--
		} else if todos[i].Completed && i == len(todos)-1 {
			todos = todos[:i]
		}
	}

	t.Write()
}
func (t Todolist) Add(title string) {
	if title == "" {
		fmt.Println("Task cannot be empty")
	}
	list := Todolist{Task: title}
	todos = append(todos, list)
	t.Write()
}

func (t Todolist) Write() {
	res, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	if err = ioutil.WriteFile(doc, res, 0666); err != nil {
		log.Fatal(err)
	}
}

func (t Todolist) Done(val int) {
	if len(todos) < val {
		fmt.Println("Index out of range")
	} else {
		todos[val - 1].Completed = true
		t.Write()
	}
}
func (t Todolist) Undone(val int) {
	if len(todos) < val {
		fmt.Println("Index out of range")
	} else {
		todos[val - 1].Completed = false
		t.Write()
	}
}
func (t Todolist) List() {
	for i, v := range todos {
		switch v.Completed {
		case true:
			fmt.Printf("%d(*)\t %v\n", i+1, v.Task)
		default:
			fmt.Printf("%d\t %v\n", i+1, v.Task)
		}
	}
}
func AddTodoWithTerm() string{
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}