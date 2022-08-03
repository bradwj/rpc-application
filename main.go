package main

import (
	"fmt"
)

type Item struct {
	title	string
	body	string
}

var database []Item

func GetByName(title string) Item {
	var getItem Item

	for _, val := range database {
		if val.title == title {
			getItem = val
		}
	}

	return getItem
}

func AddItem(item Item) Item {
	database = append(database, item)
	return item
}

func EditItem(title string, edit Item) Item {
	var changed Item

	for idx, val := range database {
		if val.title == title {
			database[idx] = edit
			changed = edit
		}
	}

	return changed
}

func DeleteItem(item Item) Item {
	var deleted Item

	for idx, val := range database {
		if val.title == item.title && val.body == item.body {
			database = append(database[:idx], database[idx+1:]...)
			deleted = item
			break
		}
	}

	return deleted
}

func main() {
	fmt.Println("initial database: ", database)
	a := Item{"first", "the first item"}
	b := Item{"second", "the second item"}
	c := Item{"third", "the third item"}

	AddItem(a)
	AddItem(b)
	AddItem(c)
	fmt.Println("database after adding items: ", database)

	DeleteItem(b)
	fmt.Println("database after deleting second item: ", database)

	EditItem("third", Item{"fourth", "a new item"})
	fmt.Println("database after editing third item: ", database)

	x := GetByName("fourth")
	y := GetByName("first")
	fmt.Println(x, y)
}