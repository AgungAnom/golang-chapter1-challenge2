package main

import (
	"fmt"
)


func main() {

for i :=0; i<5; i++{
	fmt.Println("Nilai i = ", i)
	if i == 4{
		for j:=0; j<11;j++{
			
			if j == 5{
				for index, value := range "САШАРВО"{
					fmt.Printf("character %U '%c' starts at byte position %d\n", value,value,index)
				}
			} else {
				fmt.Println("Nilai j = ", j)
			}
		}
	}
}
}