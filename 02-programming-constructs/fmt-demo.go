package main

import "fmt"

func main() {
	//name := "Magesh"
	var name string
	fmt.Println("Enter the name : ")
	fmt.Scanf("%s", &name)
	//fmt.Println("Name is ", name)
	//fmt.Printf("Name is %v (of type %T)\n", name, name)
	s := fmt.Sprintf("Name is %v (of type %T)\n", name, name)
	fmt.Println(s)
}
