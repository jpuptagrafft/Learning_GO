package main 

import "fmt"


func main(){
	fmt.Println("1/2 = ", 7%2)
	fmt.Println("7.0/3.0 = ", 7.0/3.0)
	fmt.Println("7.0/3 = ", 7.0/3)
	for i := 0; i < 11; i++ {
		fmt.Println(i, ", ", (i - 10)/2, ", ", i%2)
	}
	//If you divide two ints together, it will automatically do int division
	//any float division will result in normal division
	//Remember, Comma to Concatinate with non strings!
}