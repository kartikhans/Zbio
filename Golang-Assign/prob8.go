package main
import(
	"fmt"
)

func main(){

}

type Orderdict struct{
	// arr:=make([]string,0)
	arr []string
	hashi map[string]int
}

type tuple struct{
	key string
	value int
}

func (l *Orderdict) init(){
	l.arr:=make([]string,0)
}
func (l *Orderdict) add(key string,value int){
	l.hashi[key]=value
	l.arr:=append(l.arr, key)
}

func (l *Orderdict) get(key string) int{
	return(l.hashi[key])
}

func (l *Orderdict) retrieve() []tuple{
	n:=len(l.arr)
	tupp:=make([]tuple,n)
	count:=0
	for i:=0;i<n;i++{
		sam:=tuple{l.arr[i],l.hashi[l.arr[i]]}
		tupp[count]=sam
		count++
	}
	return(tupp)
}