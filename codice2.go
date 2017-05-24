package main

import (
    "fmt"
    "os"
)

func main(){
    if len(os.Args)!=2 {
	fmt.Println("Mi aspettavo un solo argomento")
	os.Exit(1)
    }

    fmt.Println("Ho ricevuto questo argomento:",os.Args[1])
}
