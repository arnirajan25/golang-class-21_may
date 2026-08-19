
package main
import "fmt"

type Rectangle struct{
length int
breadth int
}
func (rect*Rectangle) Perimeter(){
return 2* (rect.length+rect.breadth)
}
func (rect*Rectangle) Area(){
return rect.length * rect.breadth
}
//var num int = 10
//var decimal float64 = 10.10
//var True bool = true
//var suhag string
//var mahim string = "ncmt_student"
//var safal string = "ncmt_bcs_student"
func main(){
//suhag="python guy"
//var name string= "nirajan"
//temp :=30.85
//var temp_two float64= 32.32
//var age int= 20
//fmt.Println("temp:%.2f\n",temp)
//fmt.Println("name:%s\n",name)
//fmt.Println("age:%d\n",age)
//fmt.Println("temp:%.2f\n",temp_two)
//fmt.Println(suhag,safal,mahim,num,decimal,True)
var rect Rectangle
fmt.Println("Enter the length and breadth of a rectangle:")
fmt.Scanln(&rect.length,&rect.breadth)
area := rect.length * rect.breadth
perimeter := 2*(rect.length + rect.breadth)
fmt.Println("\nThe area of rectangle = %d",area)
fmt.Println("\nThe perimeter of rectangle = %d",perimeter)
fmt.Println("Perimeter =",rect.Perimeter())
fmt.Println("Area =",rect.Area())
}
