package main
import(
	"fmt"
)

func main(){
	var r lru
	r.Objects = make([]int,0)
	r.capacity = 10

}

type lru struct{
	Objects []int
	capacity int
	hashi map[int]int
}

func (l *lru) len(){
	return(len(l.Objects))
}
func (l *lru) add(n int){
	if(l.hashi[int]!=0){
		sam:=l.hashi[n]
		delete(l.hashi, int)

	}
	if(l.len()==l.capacity){
		l.Objects = l.Objects[1:]
	}
	l.Objects = append(l.Objects, n)
	l.hashi[n]=len(l.Objects)-1
}

func (l *lru) get()
