package main

import (
	"fmt"
)



type Node struct {
	value int
	left *Node
	right *Node
}

func makebtree(inord []int, postord []int, n int) *Node{
	pindex:=make([]int,1)
	pindex[0]=n-1
	return(funcbtree(inord, postord, 0, n-1, pindex))
}

func funcbtree(inord []int, postord []int, strt int, end int, pindex []int) *Node{
	var s *Node
	if(strt>end){
		return(s)
	}
	var node *Node
	node = &Node{value: postord[pindex[0]], left: nil, right:nil}
	pindex[0]--
	if(strt==end){
		return(node)
	}
	iindex:= search(inord, strt, end, node.value)
	node.right = funcbtree(inord, postord, iindex+1, end, pindex)
	node.left = funcbtree(inord, postord, strt, iindex-1, pindex)

	return(node)
}

func search(arr []int, strt int, end int, value int) int{
	i:=0
	for i=strt;i<end+1;i++{
		if(arr[i]==value){
			break
		}
	}
	return(i)
}

func preorder(x *Node) {
	if(x==nil){
		return
	}
	fmt.Print(x.value," ")
	preorder(x.left)
	preorder(x.right)
}

func main(){
	inord:=[]int{9,3,15,20,7}
	postord:=[]int{9,15,7,20,3}
	leni:=len(inord)
	var x *Node
	x = makebtree(inord,postord,leni)
	preorder(x)
	fmt.Println()
	// fmt.Println(x.right.right.value)
}