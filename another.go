package main
import ("fmt"
"time")
func gole(){
fmt.Println("Horaa")
}
func pbag(){
fmt.Println("Honi")
}
func chup_hacker(){
fmt.Println("Kati Boleko")
}
func test_time(){
fmt.Println("Hello World!!!")
}
func main(){
fmt.Println(time.Minute)
fmt.Println(time.Nanosecond)
fmt.Println(time.Microsecond)
fmt.Println(time.Second)
fmt.Println(time.Millisecond)
//start:= time.Now()
//t:= time.Now()
//elapsed:= t.Sub(start)
//fmt.Println(elapsed)
time.Sleep(5*time.Second)
test_time()
fmt.Println("Write Now Before")
go chup_hacker()
fmt.Println("And Then ")
time.Sleep(7*time.Second)
//fmt.Println(time.Sleep)
start:=time.Now()
go gole()
time.Sleep(8*time.Second)
go pbag()
time.Sleep(10*time.Second)
end:=time.Since(start)
fmt.Println(start)
fmt.Println(end)
}
