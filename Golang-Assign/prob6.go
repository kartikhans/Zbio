package main

import (
	"fmt"
)

func main(){
	app:=[]string{"sonam","harish","kartik","rishi"}
	app2:=[]int{2,3,4,5}
	sam:=hashi(app,app2)
	fmt.Println(sam["kartik"])
}

func hashi(arr1 []string, arr2 []int) map[string]int{
	n:=len(arr1)

	mapi := make(map[string]int)

	for i:=0;i<n;i++{
		mapi[arr1[i]]=arr2[i]
	}
	return(mapi)
}