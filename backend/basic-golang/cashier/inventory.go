package main

type Item struct {
	id    int
	name  string
	price float64
}

type Inventory struct {
	items []Item
}

func NewItem(id int, name string, price float64) Item {
	return Item{id, name, price}
}

func (i *Inventory) Add(item Item) {
<<<<<<< HEAD
	//beginanswer
	i.items = append(i.items, item)
	//endanswer
=======
	// TODO: answer here
>>>>>>> 6a266f35b3e5d854980b80d4d6208d897f7008b9
}

func (i Inventory) Items() []Item {
	return i.items
}
