package main

import "time"

type Item struct {
	Title    string
	Deadline time.Time
}

type Todos struct {
	items []Item
}

func (todos *Todos) Add(item Item) {
<<<<<<< HEAD
	//beginanswer
	todos.items = append(todos.items, item)
	//endanswer
}

func (todos *Todos) GetAll() []Item {
	//beginanswer
	return todos.items
	//endanswer return []Item{}
}

func (todos *Todos) GetUpcoming() []Item {
	//beginanswer
	var upcoming []Item
	for _, item := range todos.items {
		if item.Deadline.After(time.Now()) {
			upcoming = append(upcoming, item)
		}
	}
	return upcoming
	//endanswer return []Item{}
=======
	// TODO: answer here
}

func (todos *Todos) GetAll() []Item {
	return []Item{} // TODO: replace this
}

func (todos *Todos) GetUpcoming() []Item {
	return []Item{} // TODO: replace this
>>>>>>> 6a266f35b3e5d854980b80d4d6208d897f7008b9
}

func NewItem(title string, deadline time.Time) Item {
	return Item{title, deadline}
}

func NewTodos() Todos {
	return Todos{}
}
