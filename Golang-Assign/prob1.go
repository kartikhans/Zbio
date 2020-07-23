package main
import (
    "fmt"  
)

func main() {
	arr:=[]int{0,1,0,3,12}
	arr = allzero(arr)
	fmt.Println(arr)
}

func allzero(kaim []int) []int{
	counter:=0
	for i:=0;i<len(kaim);i++{
		if(kaim[i]!=0){
			kaim[i],kaim[counter]=kaim[counter],kaim[i]
			counter++
		}
	}
	return(kaim)
}