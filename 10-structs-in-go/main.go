package main

import (
	"fmt"
)

type Engineer struct {
	Name    string
	Age     int
	Project Project
}

type Project struct {
	Name         string
	Priority     string
	Technologies []string
}

func main() {
	engineer := Engineer{
		Name: "Jesse",
		Age:  39,
		Project: Project{
			Name:         "Beginner's Guide to Go",
			Priority:     "High",
			Technologies: []string{"Go"},
		},
	}
	fmt.Printf("%+v\n", engineer)
	fmt.Println(engineer.Name)
	fmt.Println(engineer.Project.Name)

}
