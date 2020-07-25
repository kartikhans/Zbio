package main
import(
	"fmt"
	"strings"
)

func main(){
	x:= "ZBIO has come to campus for hiring today. ZBIO is a VC funded company in the enterprise infrastructure domain. This company is creating a platform to observe service to service communication"
	sam:=strings.Split(x," ")
	fmt.Println(uniquecount(sam))
}

func uniquecount(stri []string) int{
	mapi := make(map[string]int)
	i:=0
	counter:=0
	for i<len(stri){
		// var sam string
		if(stri[i][0]=='"' || stri[i][0]=='!'){
			stri[i]=stri[i][1:]
		}
		if(stri[i][len(stri[i])-1]=='.' || stri[i][len(stri[i])-1]=='"'){
			stri[i]=stri[i][:len(stri[i])-1]
		}
		elem, status := mapi[stri[i]]
		_ =elem
		if(!status){
			counter++
			mapi[stri[i]]++
		}else{
			if(elem==1){
				counter--
			}
			mapi[stri[i]]++
		}
		i++
	}
	return(counter)
}