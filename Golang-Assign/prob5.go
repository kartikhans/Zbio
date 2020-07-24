package main
import(
	"fmt"
	"errors"
)

func main(){
	var r lru
	r.Objects = make([]int,0)
	r.capacity = 2
	r.hashi = make(map[int]int)
	r.add(4)
	r.add(6)
	r.add(7)
	fmt.Println(r.show(4))
	fmt.Println(r.hashi)
}

type lru struct{
	Objects []int
	capacity int
	hashi map[int]int
}

func (l *lru) len() int{
	return(len(l.Objects))
}
func (l *lru) add(n int){
	if(l.hashi[n]!=0){
		sam:=l.hashi[n]
		// fmt.Println("n")
		delete(l.hashi, n)
		for i:=sam;i<l.len()-1;i++{
			l.hashi[l.Objects[i+1]]--
			l.Objects[i]=l.Objects[i+1]
		}
		l.Objects=l.Objects[:l.len()-1]
	}
	if(l.len()==l.capacity){
		sam:=l.Objects[0]
		delete(l.hashi, sam)
		l.Objects = l.Objects[1:]
		for i:=0;i<l.len();i++{
			l.hashi[l.Objects[i]]--
		}
		// for i:=0;i<l.len()-1;i++{
		// 	l.hashi[l.Objects[i+1]]--
		// 	l.Objects[i]=l.Objects[i+1]
		// }
		// l.Objects = l.Objects[:l.len()-1]
		delete(l.hashi, sam)
	}
	l.Objects = append(l.Objects, n)
	l.hashi[n]=l.len()-1
}

func (l *lru) get() (int,error){
	if(l.len()==0){
		return -1, errors.New("lru is empty")
	}
	sam:=l.Objects[0]
	l.Objects = l.Objects[1:]
	return sam, nil
}


func (l *lru) show(n int) bool{
	v, found := l.hashi[n]
	_ = v
	if(found){
		return(true)
	}
	return(false)
}
