package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}
type Template struct {
	eat   string
	loco  string
	noise string
}

func (t Template) Eat() {
	fmt.Println(t.eat)
}
func (t Template) Move() {
	fmt.Println(t.loco)
}

func (t Template) Speak() {
	fmt.Println(t.noise)
}
func Choices() {

	collection := make(map[string]Template)
	fmt.Println("Input Example: \nFor Creation: \"newanimal Tom Cow \"\nFor Query:\"query Tom eat\" ")

	input := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">")
		ch, _, _ := input.ReadLine()
		req := strings.Split(string(ch), " ")
		for k, _ := range req {
			req[k] = strings.ToLower(req[k])
		}

		switch req[0] {
		case "newanimal":
			switch req[2] {
			case "cow":
				collection[req[1]] = Template{"grass", "walk", "moo"}
				fmt.Println("Created it!")
			case "snake":
				collection[req[1]] = Template{"mice", "slither", "hsss"}
				fmt.Println("Created it!")
			case "bird":
				collection[req[1]] = Template{"worms", "fly", "peep"}
				fmt.Println("Created it!")
			default:
				fmt.Println("Invaild Animal Type")
			}
		case "query":
			switch req[2] {
			case "eat":
				s, err := collection[req[1]]
				if err != true {
					fmt.Println("Animal not Present")
					break
				}
				s.Eat()
			case "move":
				s, err := collection[req[1]]
				if err != true {
					fmt.Println("Animal not Present")
					break
				}
				s.Move()
			case "speak":
				s, err := collection[req[1]]
				if err != true {
					fmt.Println("Animal not Present")
					break
				}
				s.Speak()
			}
		default:
			fmt.Println("Invaild Command")
		}
	}

}
func main() {
	Choices()
}
