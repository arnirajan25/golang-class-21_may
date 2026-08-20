package main
import"fmt"
import"sync"
import "time"

func hello(){
fmt.Println("Hello mahim")
}

func suhag(wg *sync.WaitGroup){
p:=10
t:=2
r:=5
si:=(p*t*r)/100
fmt.Println(si)
wg.Done()
}

func mahim(wg *sync.WaitGroup){
num:=12
for i:=1;i<=10;i++ {
fmt.Printf("%d*%d=%d\n",num,i,num*i)
}
wg.Done()
}

func main(){
var wg sync.WaitGroup
start:=time.Now()
wg.Add(2)
go hello()
go suhag(&wg)
go mahim(&wg)
fmt.Println(time.Since(start))
wg.Wait()
}
