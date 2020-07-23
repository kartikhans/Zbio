package main
import(
	"fmt"
)

func main(){
	kaim:=88
	fmt.Println(primerecur(kaim,2))
	fmt.Println(primeiter(kaim))
	fmt.Println(primeseive(kaim))
}


//iterative way of checking for prime
func primeiter(n int) bool{
	if(n<=2){
		if(n==2){
			return(true)
		}
		return(false)
	}
	for i:=3;i<n;i++{
		if(n%i==0){
			return(false)
		}
	}
	return(true)
}

//recursive way of checking for prime
func primerecur(n int, x int) bool{

	if(n<=2){
		if(n==2){
			return(true)
		}
		return(false)
	}
	if(n%x==0){
		return(false)
	}
	if(x*x>n){
		return(true)
	}
	return(primerecur(n,x+1))
}

//seive algorithm, the most efficient algorithm to check for a number is prime or not.
func primeseive(n int) bool{
	arr := make([]bool, n+1)
	for i:=0;i<n+1;i++{
		arr[i]=true
	}
	p:=2

	for (p*p<=n){
		if(arr[p]==true){
			for i:=p*p;i<n+1;i=i+p{
				arr[i]=false
			}
		}
		p++
	}
	if(arr[n]){
		return(true)
	}
	return(false)
}