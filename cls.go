package main
import "fmt"
func main() {
	var copies [7]string = [7]string{"hello","it's","just","me","and","my","workspace"}
	var num [3]int = [3]int{9,10,11}
	var notes [7]string
	mySlice := myArray[1:3]
	myslice :=[]string{"Hachu","YoYo"}
	notes[0] = "do"
	notes[1] = "re"
	notes[2] = "mi"
	fmt.Println(myslice[0],myslice[1])
	fmt.Println(notes[0])
	fmt.Println(notes[1])
	fmt.Println(num[0],num[1],num[2])
	fmt.Println(copies[0],copies[5],copies[6])
	for i := 0; i <= 6; i++ {
		fmt.Println(copies[i])
}
	fmt.Println(len(copies))
}
