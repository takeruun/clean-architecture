package main

type PC struct {
	PCParts PCParts
}

func (p PC) Start() string {
	return "fanfanfan"
}

type PCParts struct{}
