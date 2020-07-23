package main

import (
	"fmt"
)



type Node struc {
	value int
	left Node
	right Node
}

func makebtree(inord []int, postord []int, n int[]) Node{
	pindex:=[1]int
	pindex[0]=n-1
	return(funcbtree(inord, postord, 0, n-1, pindex))
}

func funcbtree(inord []int, postord []int, strt int, end int, pindex []int) Node{
	if(strt>end){
		return(nil)
	}
	node:=Node{value: postord[pindex[0]], left: nil, right:nil}
	pindex[0]--
	if(strt==end){
		return(node)
	}
	iindex:= search(inord, strt, end, node.value)
	node.right = funcbtree(inord, postord, iindex+1, end, pindex)
	node.left = funcbtree(inord, postord, strt, iindex-1, pindex)

	return(node)
}

func search(arr []int, strt int, end int, value int){
	i:=0
	for i=strt;i<end+1;i++{
		if(arr[i]==value){
			break
		}
	}
	return(i)
}
func main(){
	inord:=[]int{9,3,15,20,7}
	postord:=[]int{9,15,7,20,3}
	leni:=len(inord)
	var x Node
	x = makebtree(inord,postord,leni)
	fmt.Prinltn(x.data)
}