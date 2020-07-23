package main
import(
	"fmt"
)

func main(){
	nap lru("hi")
	fmt.println(nap.whatname())
}

type lru struct{
	name string
	func lru(n string){
		name = n
	}
	func whatname() string{
		return(name)
	}
}