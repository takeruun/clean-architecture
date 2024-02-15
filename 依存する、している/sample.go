package main

import "fmt"

type PCParts struct {
	Name string
}

func (p PCParts) Exec() {
	fmt.Printf("this parts is: %s\n", p.Name)
}

type PC struct {
	PCParts PCParts
}

func NewPC() *PC {
	pCParts := PCParts{Name: "CPU"}

	return &PC{
		// PC は PCParts に依存している
		PCParts: pCParts,
	}
}

func main() {
	pc := NewPC()
	pc.PCParts.Exec()
}
