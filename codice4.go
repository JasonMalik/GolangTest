package main

import(
    "fmt"
     "os"
     "strconv"
)

func safeDivision(num int, den int) (division float32,safe bool) {
    if den==0 {
	return 0,false
    }

    return float32(num)/float32(den),true
}

func main(){
    if len(os.Args)!=3 {
	fmt.Println("Mi aspettavo due argomenti")
	os.Exit(1)
    }
    num1,_ := strconv.Atoi(os.Args[1])
    num2,_ := strconv.Atoi(os.Args[2])

    res,safe := safeDivision(num1,num2)

    if !safe{
	fmt.Println("Hai diviso per zero")
        os.Exit(2)
    }

    fmt.Println(os.Args[1],"/",os.Args[2],"=",res)

}
