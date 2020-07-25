package main
import(
	"fmt"
)

func main(){
	var r Orderdict
	r.Objects = make([]string,0)
	r.hashi = make(map[string]int)
	r.add("kartik", 5)
	r.add("sam", 4)
	r.add("fiscal", 6)
	r.add("kartik",2)
	fmt.Println(r.get("kartik"))
	fmt.Println(r.deltop())
	f:=r.retrieve()
	fmt.Println(f)
}

type Orderdict struct{
	Objects []string
	hashi map[string]int
}

type tuple struct{
	key string
	value int
}

func (l *Orderdict) add(key string,value int){
	v, found := l.hashi[key]
	_=v
	if(found){
		l.hashi[key]=value
	}else{
	l.hashi[key]=value
	l.Objects = append(l.Objects, key)
	}
	
}

func (l *Orderdict) get(key string) int{
	return(l.hashi[key])
}
func (l *Orderdict) deltop() tuple{
	var s tuple
	s = tuple{l.Objects[0],l.hashi[l.Objects[0]]}
	delete(l.hashi,l.Objects[0])
	l.Objects = l.Objects[1:]
	return(s)
}
func (l *Orderdict) retrieve() []tuple{
	n:=len(l.Objects)
	tupp:=make([]tuple,n)
	count:=0
	for i:=0;i<n;i++{
		sam:=tuple{l.Objects[i],l.hashi[l.Objects[i]]}
		tupp[count]=sam
		count++
	}
	return(tupp)
}