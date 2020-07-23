package main
import (
	"fmt"
)

func main(){
	arr1 := []int{1,3,4,5}
	arr2 := []int{2,4,6,8}
	arr3 := mergetwo(arr1,arr2)
	fmt.Println(arr3)
}

func mergetwo(arr1 []int, arr2 []int) []int{
	len1 := len(arr1)
	len2 := len(arr2)
	arr3 := make([]int, len1+len2)
	var i,j,k int
	i=0
	for i<len1 && j<len2{
		if(arr1[i]<arr2[j]){
			arr3[k]=arr1[i]
			k++
			i++
		} else{
			arr3[k]=arr2[j]
			k++
			j++
		}
	}
	for i<len1{
		arr3[k]=arr1[i]
		k++
		i++
	}
	for j<len2{
		arr3[k]=arr2[j]
		k++
		j++
	}
	return(arr3)
}