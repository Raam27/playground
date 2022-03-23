package main

import "time"

type Item struct {
	id       int
	action   string
	deadline time.Time
}

type Todos struct {
	items []Item
}

func (todos *Todos) Add(item Item) {
	//beginanswer
	todos.items = append(todos.items, item)
	//endanswer
}

func (todos *Todos) GetItems() []Item {
	//beginanswer
	return todos.items
	//endanswer
}

func NewItem(id int, action string, deadline time.Time) Item {
	return Item{id, action, deadline}
}

func NewTodos() Todos {
	//beginanswer
	return Todos{}
	//endanswer
}
