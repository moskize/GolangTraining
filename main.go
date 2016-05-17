package main

import "fmt"

func main() {
	list := NewLinkedList()
	list.Add("a")
	list.Add("b")
	list.Add("c")
	list.Add("d")

	print(list)

	fmt.Println("remove element : " + list.Remove(2))

	print(list)

	fmt.Println("replace element" + list.Set(1, "123"))

	print(list)
}

func print(list *linkedList) {
	fmt.Println("----")
	for i := 0; i < list.Length(); i++ {
		fmt.Println(list.Get(i))
	}
}
