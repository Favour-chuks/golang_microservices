package main

import "fmt"


func main(){
	var myString = []rune( "resume")
	var indexed = myString[0]
	fmt.Printf("%v, %T\n", indexed, indexed)
	
	for i, v := range myString{
		fmt.Println(i,v)
	}
	// fmt.Print(len(myString))
}
